package bills

// Bill master insert Api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type BillDetailsResp struct {
	Status string `json:"status"`
	ErrMsg string `json:"errMeg"`
}

type BillDReq struct {
	Billno       int     `json:"billno"`
	MedicineName string  `json:"medname"`
	Quantity     int     `json:"quantity"`
	UniPrice     int     `json:"unitprice"`
	Amount       float64 `json:"amount"`
	UserId       string  `json:"userid"`
}

func BillDetails(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Purchased Done")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Credentials", "false")
	(w).Header().Set("Access-Control-Allow-Methods", "PUT,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "User,Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")
	var Items []BillDReq
	var Dresp BillDetailsResp
	// var BillRec BillDReq
	Dresp.Status = "S"
	if r.Method == "PUT" {
		fmt.Println("Purchased Done (+)")

		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println("BMD001", err)
			Dresp.ErrMsg = err.Error()
			Dresp.Status = "E"
		} else {
			err := json.Unmarshal(body, &Items)
			fmt.Println("Items", Items)

			if err != nil {
				fmt.Println("BMD002", err)

				Dresp.ErrMsg = err.Error()
				Dresp.Status = "E"
			} else {
				BillMaster, err := InserBillDetails(Items)
				if err != nil {
					fmt.Println("BMD003", err)

					Dresp.ErrMsg = err.Error()
					Dresp.Status = "E"
				} else {
					data, err := json.Marshal(BillMaster)
					if err != nil {
						fmt.Println("BMD004", err)

						Dresp.ErrMsg = err.Error()
						Dresp.Status = "E"
					} else {
						Dresp.ErrMsg = "Scucess"
						Dresp.Status = "S"
						fmt.Fprintln(w, string(data))
					}
				}
				ReduceStock, err := ReduceStockDetails(Items)
				if err != nil {
					fmt.Println("BMD005", err)

					Dresp.ErrMsg = err.Error()
					Dresp.Status = "E"
				} else {
					data, err := json.Marshal(ReduceStock)
					if err != nil {
						fmt.Println("BMD006", err)

						Dresp.ErrMsg = err.Error()
						Dresp.Status = "E"
					} else {
						Dresp.ErrMsg = "Scucess"
						Dresp.Status = "S"
						fmt.Fprintln(w, string(data))
					}
				}
			}
			fmt.Println("Purchased Done (-)")

		}
	}

}

func InserBillDetails(Items []BillDReq) (BillDetailsResp, error) {
	// var BillRec BillDReq
	var Dresp BillDetailsResp
	Dresp.Status = "S"

	db, err := LocalDBConnect()
	if err != nil {
		Dresp.ErrMsg = err.Error()
		Dresp.Status = "E"
		fmt.Println("IBD001", err)

	} else {
		defer db.Close()
		for i := 0; i < len(Items); i++ {
			fmt.Println(Items[i])
		}
		for _, item := range Items {
			lBillListItem := `insert  into  loganathan.medapp_bill_details
			( bill_no ,medicine_name ,quantity,unit_price,amount,created_by,created_date,updated_by,
				updated_date )
			values(?,?,?,?,?,?,curdate(),?,curdate())`
			item.UniPrice = 20
			_, err := db.Exec(lBillListItem, item.Billno, item.MedicineName, item.Quantity, item.UniPrice, item.Amount, item.UserId, item.UserId)
			if err != nil {
				log.Println(err)
				Dresp.ErrMsg = err.Error()
				fmt.Println("IBD002", err)
				Dresp.Status = "E"
			} else {

				log.Println("updated sucessfully")
			}

		}

	}

	return Dresp, err

}

func ReduceStockDetails(Items []BillDReq) (BillDetailsResp, error) {
	// var BillRec BillDReq
	var Dresp BillDetailsResp
	Dresp.Status = "S"

	db, err := LocalDBConnect()
	if err != nil {
		Dresp.ErrMsg = err.Error()
		Dresp.Status = "E"
		fmt.Println("RSD001", err)

	} else {
		defer db.Close()
		for i := 0; i < len(Items); i++ {
			fmt.Println(Items[i])
		}
		for _, item := range Items {
			ReduceStockQuery := `update loganathan.medapp_stock 
				set quantity = (select quantity from loganathan.medapp_stock where medicine_name = ?)-?
				where medicine_name  = ?`
			_, err := db.Exec(ReduceStockQuery, item.MedicineName, item.Quantity, item.MedicineName)
			if err != nil {
				log.Println(err)
				Dresp.ErrMsg = err.Error()
				fmt.Println("RSD002", err)
				Dresp.Status = "E"
			} else {

				log.Println("updated sucessfully")
			}

		}

	}

	return Dresp, err

}
