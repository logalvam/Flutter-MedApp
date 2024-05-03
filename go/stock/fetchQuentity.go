package stock

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type MedicineQuantityStruct struct {
	Quantity  int64   `json:"quantity"`
	UnitPrice float64 `json:"unitprice"`
	Status    string  `json:"status"`
	ErrMsg    string  `json:"errmsg"`
}

func FetchMedQuantity(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Quantity ")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Credentials", "false")
	(w).Header().Set("Access-Control-Allow-Methods", "PUT,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "MEDNAME,Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")

	medName := r.Header.Get("MEDNAME")

	if r.Method == "PUT" {
		var quantity MedicineQuantityStruct
		quantity.Status = "S"

		lQty, err := GetQuantity(quantity, medName)
		if err != nil {
			quantity.Status = "E"
			fmt.Println("TS001", err)
			quantity.ErrMsg = err.Error()
			fmt.Fprintln(w, quantity)
			fmt.Println(quantity)
		} else {
			data, err := json.Marshal(lQty)
			if err != nil {
				fmt.Println("TS002", err)
				quantity.Status = "E"
				quantity.ErrMsg = err.Error()
				fmt.Fprintln(w, quantity)
				fmt.Println(quantity)
			} else {

				fmt.Fprintln(w, string(data))
			}
		}
	}
}

func GetQuantity(quantity MedicineQuantityStruct, medName string) (MedicineQuantityStruct, error) {
	quantity.Status = "S"

	db, err := LocalDBConnect()
	if err != nil {
		quantity.ErrMsg = err.Error()
		fmt.Println("GTS001", err)
		quantity.Status = "E"
	} else {
		defer db.Close()

		Quantity := ` select quantity,unit_price from  loganathan.medapp_stock 
		where medicine_name =?`
		rows, err := db.Query(Quantity, medName)
		if err != nil {
			quantity.ErrMsg = err.Error()
			fmt.Println("GTS002", err)
			quantity.Status = "E"
		} else {
			for rows.Next() {
				err := rows.Scan(&quantity.Quantity, &quantity.UnitPrice)
				if err != nil {
					quantity.ErrMsg = err.Error()
					fmt.Println("GTS003", err)
					quantity.Status = "E"
				} else {
					quantity.ErrMsg = "Success"
					quantity.Status = "S"
				}
			}
		}

	}

	return quantity, err
}
