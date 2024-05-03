package sales

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type CompareSalesStruct struct {
	Amount    float64 `json:"amount"`
	UserId    string  `json:"userid"`
	Yesterday string  `json:"yesterday"`
}
type CompareResp struct {
	Status string `json:"status"`
	ErrMsg string `json:"errmsg"`
}

func CompareSales(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Compare sales")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Credentials", "false")
	(w).Header().Set("Access-Control-Allow-Methods", "PUT,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "DATE,Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")

	if r.Method == "PUT" {
		var CompareRec CompareSalesStruct
		var resp CompareResp
		resp.Status = "S"
		body, err := io.ReadAll(r.Body)
		if err != nil {
			resp.Status = "E"
			resp.ErrMsg = err.Error()
			fmt.Fprintln(w, resp)
			fmt.Println(resp)
		} else {
			err := json.Unmarshal(body, &CompareRec)
			fmt.Println(CompareRec)
			if err != nil {
				resp.Status = "E"
				resp.ErrMsg = err.Error()
				fmt.Fprintln(w, resp)
				fmt.Println(resp)
			} else {

				lCSales, err := GetCompareSales(CompareRec)
				if err != nil {
					resp.Status = "E"
					resp.ErrMsg = err.Error()
					fmt.Fprintln(w, resp)
					fmt.Println(resp)
				} else {
					data, err := json.Marshal(lCSales)
					if err != nil {
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

	}
}

func GetCompareSales(CompareRec CompareSalesStruct) (CompareSalesStruct, error) {
	var resp CompareResp
	resp.Status = "S"

	db, err := LocalDBConnect()
	if err != nil {
		resp.ErrMsg = err.Error()
		resp.Status = "E"
	} else {
		defer db.Close()

		lYSales := `select sum(bill_amount) from  loganathan.medapp_bill_master 
				where DATE_FORMAT(bill_date , '%d/%m/%Y') =? and user_id =?`
		rows, err := db.Query(lYSales, &CompareRec.Yesterday, &CompareRec.UserId)
		if err != nil {
			resp.ErrMsg = err.Error()
			resp.Status = "E"
		} else {
			for rows.Next() {
				err := rows.Scan(&CompareRec.Amount)
				if err != nil {
					resp.ErrMsg = err.Error()
					resp.Status = "E"
				} else {
					resp.ErrMsg = "Success"
					resp.Status = "S"
				}
			}
		}
	}
	return CompareRec, err
}
