package sales

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type TodaySalesStruct struct {
	Sales float64 `json:"sales"`
	Date  string  `json:"date"`
}
type TSalesResp struct {
	Status string `json:"status"`
	ErrMsg string `json:"errmsg"`
}

func TodaSales(w http.ResponseWriter, r *http.Request) {
	fmt.Println("today sales (+)")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Credentials", "false")
	(w).Header().Set("Access-Control-Allow-Methods", "PUT,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "USER,Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")

	UserId := r.Header.Get("USER")
	fmt.Println(UserId)
	if r.Method == "PUT" {
		var SalesRec TodaySalesStruct
		var resp TSalesResp
		resp.Status = "S"

		lSales, err := GetTodaySales(SalesRec, UserId)
		if err != nil {
			resp.Status = "E"
			fmt.Println("TS001", err)
			resp.ErrMsg = err.Error()
			fmt.Fprintln(w, resp)
			fmt.Println(resp)
		} else {
			data, err := json.Marshal(lSales)
			if err != nil {
				fmt.Println("TS002", err)
				resp.Status = "E"
				resp.ErrMsg = err.Error()
				fmt.Fprintln(w, resp)
				fmt.Println(resp)
			} else {
				fmt.Println("today sales (-)")

				fmt.Fprintln(w, string(data))
			}
		}

	}
}

func GetTodaySales(SalesRec TodaySalesStruct, UserId string) (TodaySalesStruct, error) {
	var resp TSalesResp
	resp.Status = "S"

	db, err := LocalDBConnect()
	if err != nil {
		resp.ErrMsg = err.Error()
		fmt.Println("GTS001", err)
		resp.Status = "E"
	} else {
		defer db.Close()

		lsalestoday := ` select ifnull(sum(bill_amount) ,0) from  loganathan.medapp_bill_master 
		where bill_date = curdate() and user_id = ?`
		rows, err := db.Query(lsalestoday, UserId)
		fmt.Println(SalesRec.Date)
		if err != nil {
			resp.ErrMsg = err.Error()
			fmt.Println("GTS002", err)
			resp.Status = "E"
		} else {
			for rows.Next() {
				err := rows.Scan(&SalesRec.Sales)
				if err != nil {
					resp.ErrMsg = err.Error()
					fmt.Println("GTS003", err)
					resp.Status = "E"
				} else {
					resp.ErrMsg = "Success"
					resp.Status = "S"
				}
			}
		}

	}

	return SalesRec, err
}

func ManagerSales(w http.ResponseWriter, r *http.Request) {
	fmt.Println("manager sales")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Credentials", "false")
	(w).Header().Set("Access-Control-Allow-Methods", "PUT,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "USER,Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")

	if r.Method == "PUT" {
		var SalesRec TodaySalesStruct
		var resp TSalesResp
		resp.Status = "S"

		lSales, err := GetManagerTodaySales(SalesRec)
		if err != nil {
			resp.Status = "E"
			fmt.Println("TS001", err)
			resp.ErrMsg = err.Error()
			fmt.Fprintln(w, resp)
			fmt.Println(resp)
		} else {
			data, err := json.Marshal(lSales)
			if err != nil {
				fmt.Println("TS002", err)
				resp.Status = "E"
				resp.ErrMsg = err.Error()
				fmt.Fprintln(w, resp)
				fmt.Println(resp)
			} else {

				fmt.Fprintln(w, string(data))
			}
		}

	}

}

func GetManagerTodaySales(SalesRec TodaySalesStruct) (TodaySalesStruct, error) {
	var resp TSalesResp
	resp.Status = "S"

	db, err := LocalDBConnect()
	if err != nil {
		resp.ErrMsg = err.Error()
		fmt.Println("GTS001", err)
		resp.Status = "E"
	} else {
		defer db.Close()

		lsalestoday := ` select ifnull(sum(bill_amount) ,0) from  loganathan.medapp_bill_master 
		where bill_date = curdate() `
		rows, err := db.Query(lsalestoday)
		fmt.Println(SalesRec.Date)
		if err != nil {
			resp.ErrMsg = err.Error()
			fmt.Println("GTS002", err)
			resp.Status = "E"
		} else {
			for rows.Next() {
				err := rows.Scan(&SalesRec.Sales)
				if err != nil {
					resp.ErrMsg = err.Error()
					fmt.Println("GTS003", err)
					resp.Status = "E"
				} else {
					resp.ErrMsg = "Success"
					resp.Status = "S"
				}
			}
		}

	}

	return SalesRec, err
}
