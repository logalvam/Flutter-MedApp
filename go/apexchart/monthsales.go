package apexchart

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type MothSales struct {
	Month string  `json:"month"`
	Sales float64 `json:"sales"`
}

type MonthSalesStruct struct {
	MonthList []string  `json:"monthList"`
	SalesList []float64 `json:"salesList"`
	Status    string    `json:"status"`
	ErrMsg    string    `json:"errmsg"`
}

func MonthSales(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Month sales (+)")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Credentials", "false")
	(w).Header().Set("Access-Control-Allow-Methods", "PUT,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "DATE,Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")

	if r.Method == "PUT" {
		var resp MonthSalesStruct
		var MonthRec MothSales
		resp.Status = "S"

		Msales, err := GetMonthSales(MonthRec)
		if err != nil {
			resp.Status = "E"
			resp.ErrMsg = err.Error()
		} else {
			data, err := json.Marshal(Msales)
			if err != nil {
				resp.ErrMsg = err.Error()
				resp.Status = "E"
			} else {
				fmt.Println(string(data))
				fmt.Fprintln(w, string(data))
			}
		}

	}
}

func GetMonthSales(MonthRec MothSales) (MonthSalesStruct, error) {
	var resp MonthSalesStruct
	resp.Status = "S"

	db, err := LocalDBConnect()
	if err != nil {
		resp.Status = "S"
		resp.ErrMsg = err.Error()

	} else {
		defer db.Close()

		monthsales := `SELECT 
		DATE_FORMAT(bill_date , '%b') AS month,
		SUM(bill_amount) AS total_amount
	FROM 
		loganathan.medapp_bill_master mbm 
	where
		date_format(bill_date, '%Y') = '2024'  
		GROUP BY 
		DATE_FORMAT(bill_date, '%Y-%m')
	ORDER BY 
	month(bill_date );`

		rows, err := db.Query(monthsales)
		if err != nil {
			resp.Status = "E"
			resp.ErrMsg = err.Error()
		} else {
			for rows.Next() {
				err := rows.Scan(&MonthRec.Month, &MonthRec.Sales)
				if err != nil {
					resp.Status = "E"
					resp.ErrMsg = err.Error()
				} else {
					resp.Status = "S"
					resp.ErrMsg = "Success"
					fmt.Println(resp)
					resp.MonthList = append(resp.MonthList, MonthRec.Month)
					resp.SalesList = append(resp.SalesList, MonthRec.Sales)

				}
			}
		}
	}
	return resp, err
}
