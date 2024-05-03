package sales

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type SalesReportDateStruct struct {
	FromDate string `json:"date1"`
	ToDate   string `json:"date2"`
}

type SalesReportDetailsStruct struct {
	BillNo       int     `json:"billno"`
	BillDate     string  `json:"billdate"`
	MedicineName string  `json:"medname"`
	Quantity     int     `json:"quantity"`
	Amount       float64 `json:"amount"`
}

type SalesReportResp struct {
	SalesReportList []SalesReportDetailsStruct `json:"salesarr"`
	Status          string                     `json:"status"`
	ErrMsg          string                     `json:"errmsg"`
}

func SalesReport(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Sales Report Bill")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Credentials", "false")
	(w).Header().Set("Access-Control-Allow-Methods", "PUT,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "User,Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")
	if r.Method == "PUT" {
		var SRDateRec SalesReportDateStruct
		var resp SalesReportResp

		resp.Status = "S"

		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println("SR001", err)
			resp.ErrMsg = err.Error()
			resp.Status = "E"
		} else {
			err := json.Unmarshal(body, &SRDateRec)
			fmt.Println(SRDateRec)
			if err != nil {
				resp.ErrMsg = err.Error()
				fmt.Println("SR002", err)
				resp.Status = "E"
			} else {
				lsalesArr, err := GetSalesReportDetails(SRDateRec)
				if err != nil {
					resp.ErrMsg = err.Error()
					fmt.Println("SR003", err)
					resp.Status = "E"
				} else {
					data, err := json.Marshal(lsalesArr)
					if err != nil {
						resp.ErrMsg = err.Error()
						fmt.Println("SR004", err)
						resp.Status = "E"
					} else {
						fmt.Println("SalesReport Items")
						fmt.Fprintln(w, string(data))
						fmt.Println(string(data))
						fmt.Println("done")
					}
				}
			}
		}

	}
}

func GetSalesReportDetails(SRDateRec SalesReportDateStruct) (SalesReportResp, error) {
	var resp SalesReportResp
	var SRDetailsRec SalesReportDetailsStruct
	resp.Status = "S"

	db, err := LocalDBConnect()
	if err != nil {
		resp.ErrMsg = err.Error()
		fmt.Println("GSR001", err)
		resp.Status = "E"
	} else {
		defer db.Close()

		lSalesList := `select mbm.bill_no ,mbm.bill_date,mbd.medicine_name,mbd.quantity ,mbd.amount  
		from  loganathan.medapp_bill_master mbm
		join loganathan.medapp_bill_details mbd
		on mbm.bill_no = mbd.bill_no
		where mbm.bill_date between ? and ? ;`

		rows, err := db.Query(lSalesList, &SRDateRec.FromDate, &SRDateRec.ToDate)
		if err != nil {
			resp.ErrMsg = err.Error()
			fmt.Println("GSR002", err)
			resp.Status = "E"
		} else {
			for rows.Next() {
				err := rows.Scan(&SRDetailsRec.BillNo, &SRDetailsRec.BillDate, &SRDetailsRec.MedicineName, &SRDetailsRec.Quantity, &SRDetailsRec.Amount)
				if err != nil {
					resp.ErrMsg = err.Error()
					fmt.Println("GSR003", err)
					resp.Status = "E"
				} else {
					resp.SalesReportList = append(resp.SalesReportList, SRDetailsRec)
				}
			}
		}
	}
	return resp, err

}
