package bills

// Bill master insert Api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type BillMasterResp struct {
	Status string `json:"status"`
	ErrMsg string `json:"errMeg"`
}

type BillReq struct {
	Billno   int     `json:"billno"`
	Amount   float64 `json:"amount"`
	BillGst  float64 `json:"billgst"`
	NetPrice float64 `json:"netprice"`
	UserId   string  `json:"userId"`
}

func BillMasterDetails(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Purchased Done")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Credentials", "false")
	(w).Header().Set("Access-Control-Allow-Methods", "PUT,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "User,Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")
	var Items []BillReq
	var Mresp BillMasterResp
	// var BillRec BillReq
	Mresp.Status = "S"
	if r.Method == "PUT" {
		fmt.Println("Purchased Done")

		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println("BMD001", err)
			Mresp.ErrMsg = err.Error()
			Mresp.Status = "E"
		} else {
			err := json.Unmarshal(body, &Items)
			fmt.Println(Items)

			if err != nil {
				fmt.Println("BMD002", err)

				Mresp.ErrMsg = err.Error()
				Mresp.Status = "E"
			} else {
				BillMaster, err := InsertBillMaster(Items)
				if err != nil {
					fmt.Println("BMD003", err)

					Mresp.ErrMsg = err.Error()
					Mresp.Status = "E"
				} else {
					data, err := json.Marshal(BillMaster)
					if err != nil {
						fmt.Println("BMD004", err)

						Mresp.ErrMsg = err.Error()
						Mresp.Status = "E"
					} else {
						Mresp.ErrMsg = "Scucess"
						Mresp.Status = "S"
						fmt.Fprintln(w, string(data))
					}
				}
			}
		}
	}

}

func InsertBillMaster(Items []BillReq) (BillMasterResp, error) {
	// var BillRec BillReq
	var Mresp BillMasterResp
	Mresp.Status = "S"

	db, err := LocalDBConnect()
	if err != nil {
		Mresp.ErrMsg = err.Error()
		Mresp.Status = "E"
		fmt.Println("IBM001", err)

	} else {
		defer db.Close()
		for i := 0; i < len(Items); i++ {
			fmt.Println(Items[i])
		}
		for _, item := range Items {
			lBillListItem := `insert  into  loganathan.medapp_bill_master 
			( bill_no ,
				bill_date ,
				user_id,
				bill_gst ,
				net_price,
				created_by,
				created_date,
				updated_by,
				updated_date,
				bill_amount  )
			values(?,curdate(),?,?,?,?,curdate(),?,curdate(),?)`

			_, err := db.Exec(lBillListItem, item.Billno, item.UserId, item.BillGst, item.NetPrice, item.UserId, item.UserId, item.Amount)
			if err != nil {
				log.Println(err)
				Mresp.ErrMsg = err.Error()
				Mresp.Status = "E"
				fmt.Println("IBM002", err)

			} else {

				log.Println("updated sucessfully")
			}

		}

	}

	return Mresp, err

}
