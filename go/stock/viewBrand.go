package stock

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type MedBrandStockStruct struct {
	MedicineName string `json:"medname"`
	Brand        string `json:"brand"`
}

type MedBrandStockResp struct {
	MedicineList []MedBrandStockStruct `json:"medlist"`
	Status       string                `json:"status"`
	ErrMsg       string                `json:"errmsg"`
}

func ViewBrand(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Stock View")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Credentials", "false")
	(w).Header().Set("Access-Control-Allow-Methods", "GET,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "User,Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")

	if r.Method == "GET" {
		fmt.Println("Stock View")

		var lMedicineRec MedBrandStockStruct
		var resp MedBrandStockResp
		resp.Status = "S"
		Result, err := BStockList(lMedicineRec)
		if err != nil {
			resp.ErrMsg = err.Error()
			resp.Status = "E"
			fmt.Println("VB001", err)
		} else {
			data, err := json.Marshal(Result)
			if err != nil {
				resp.Status = "E"
				fmt.Println("VB002", err)
				resp.ErrMsg = err.Error()
			} else {
				fmt.Println("MedStock list Geted")
				fmt.Fprintln(w, string(data))
			}
		}
	}

}

func BStockList(lMedicineRec MedBrandStockStruct) (MedBrandStockResp, error) {
	var resp MedBrandStockResp
	resp.Status = "S"

	db, err := LocalDBConnect()
	if err != nil {
		fmt.Println("BSL001", err)
		resp.ErrMsg = err.Error()
		resp.Status = "E"
	} else {
		defer db.Close()

		lMedList := `select mmm.medicine_name ,mmm.brand   from loganathan.medapp_medicine_master mmm`
		rows, err := db.Query(lMedList)
		if err != nil {
			resp.Status = "E"
			fmt.Println("BSL002", err)
			resp.ErrMsg = err.Error()
		} else {
			for rows.Next() {
				err := rows.Scan(&lMedicineRec.MedicineName, &lMedicineRec.Brand)
				if err != nil {
					resp.Status = "E"
					fmt.Println("BSL003", err)
					resp.ErrMsg = err.Error()
				} else {
					resp.MedicineList = append(resp.MedicineList, lMedicineRec)
				}
			}

		}
	}
	return resp, err
}
