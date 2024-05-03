package stock

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type MedStockStruct struct {
	MedicineName string  `json:"medname"`
	Brand        string  `json:"brand"`
	Quantity     int     `json:"quantity"`
	UnitPrice    float64 `json:"unitprice"`
}

type MedStockResp struct {
	MedicineList []MedStockStruct `json:"medlist"`
	Status       string           `json:"status"`
	ErrMsg       string           `json:"errmsg"`
}

func ViewMedicine(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Stock View")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Credentials", "false")
	(w).Header().Set("Access-Control-Allow-Methods", "GET,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "User,Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")

	if r.Method == "GET" {
		fmt.Println("Stock View")

		var lMedicineRec MedStockStruct
		var resp MedStockResp
		resp.Status = "S"
		Result, err := StockList(lMedicineRec)
		if err != nil {
			fmt.Println("VM001", err)
			resp.ErrMsg = err.Error()
			resp.Status = "E"
		} else {
			data, err := json.Marshal(Result)
			if err != nil {
				fmt.Println("2", err)
				fmt.Println("VM002", err)

				resp.Status = "E"
				resp.ErrMsg = err.Error()
			} else {
				fmt.Println("MedStock list Geted")
				fmt.Fprintln(w, string(data))
			}
		}
	}

}

func StockList(lMedicineRec MedStockStruct) (MedStockResp, error) {
	var resp MedStockResp
	resp.Status = "S"

	db, err := LocalDBConnect()
	if err != nil {
		fmt.Println("SL001", err)

		resp.ErrMsg = err.Error()
		resp.Status = "E"
	} else {
		defer db.Close()

		lMedList := `select ms.medicine_name ,mmm.brand ,ms.quantity ,ms.unit_price  from loganathan.medapp_stock ms join loganathan.medapp_medicine_master mmm on ms.medicine_name  = mmm.medicine_name`
		rows, err := db.Query(lMedList)
		if err != nil {
			fmt.Println("SL002", err)

			resp.Status = "E"
			resp.ErrMsg = err.Error()
		} else {
			for rows.Next() {
				err := rows.Scan(&lMedicineRec.MedicineName, &lMedicineRec.Brand, &lMedicineRec.Quantity, &lMedicineRec.UnitPrice)
				if err != nil {
					resp.Status = "E"
					fmt.Println("SL003", err)
					resp.ErrMsg = err.Error()
				} else {
					resp.MedicineList = append(resp.MedicineList, lMedicineRec)
				}
			}
		}
	}
	return resp, err
}
