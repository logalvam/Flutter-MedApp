package apexchart

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type BillerSale struct {
	Biller string  `json:"biller"`
	Sales  float64 `json:"sales"`
}

type BillerSalesStruct struct {
	BillerList []string  `json:"billerlist"`
	AmountList []float64 `json:"amountlist"`
	Status     string    `json:"status"`
	ErrMsg     string    `json:"errmsg"`
}

func BillerSales(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Month sales (+)")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Credentials", "false")
	(w).Header().Set("Access-Control-Allow-Methods", "PUT,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "DATE,Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")

	if r.Method == "PUT" {
		var resp BillerSalesStruct
		var BillerRec BillerSale
		resp.Status = "S"

		Bsales, err := GetBillerSales(BillerRec)
		if err != nil {
			resp.Status = "E"
			resp.ErrMsg = err.Error()
		} else {
			data, err := json.Marshal(Bsales)
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

func GetBillerSales(BillerRec BillerSale) (BillerSalesStruct, error) {
	var resp BillerSalesStruct
	resp.Status = "S"

	db, err := LocalDBConnect()
	if err != nil {
		resp.Status = "S"
		resp.ErrMsg = err.Error()

	} else {
		defer db.Close()

		billersold := `select sum(bill_amount),user_id 
		from loganathan.medapp_bill_master mbm 
		group by user_id `

		rows, err := db.Query(billersold)
		if err != nil {
			resp.Status = "E"
			resp.ErrMsg = err.Error()
		} else {
			for rows.Next() {
				err := rows.Scan(&BillerRec.Sales, &BillerRec.Biller)
				if err != nil {
					resp.Status = "E"
					resp.ErrMsg = err.Error()
				} else {
					resp.Status = "S"
					resp.ErrMsg = "Success"
					resp.BillerList = append(resp.BillerList, BillerRec.Biller)
					resp.AmountList = append(resp.AmountList, BillerRec.Sales)

				}
			}
		}
	}
	return resp, err
}
