package stock

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type MedicineStruct struct {
	Medicinename string  `json:"medname"`
	Quantity     int     `json:"quantity"`
	Unitprice    float64 `json:"unitprice"`
	CreatedBy    string  `json:"createdBy"`
	CreatedDate  string  `json:"createdDate"`
	UpdatedBy    string  `json:"updatedBy"`
	UpdatedDate  string  `json:"updatedDate"`
}

type MedicineResp struct {
	// MedicineArr []MedicineStruct `json:"medicineArr"`
	Status string `json:"status"`
	ErrMsg string `json:"errMeg"`
}

func AddMedicine(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Insert Medicine Name")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Credentials", "false")
	(w).Header().Set("Access-Control-Allow-Methods", "PUT,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "USER,Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")

	if r.Method == "PUT" {
		var MedicineRec MedicineStruct
		var resp MedicineResp
		resp.Status = "S"

		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println("AD001", err)
			resp.Status = "E"
			resp.ErrMsg = err.Error()
		} else {
			err := json.Unmarshal(body, &MedicineRec)
			fmt.Println(MedicineRec)
			if err != nil {
				fmt.Println("AD002", err)
				resp.Status = "E"
				resp.ErrMsg = err.Error()
			}
			MedicineName, err := InsertMedicine(MedicineRec)
			if err != nil {
				fmt.Println("AD003", err)

				resp.Status = "E"
				resp.ErrMsg = err.Error()
				fmt.Println(resp)
				fmt.Fprintln(w, resp)
			} else {
				data, err := json.Marshal(MedicineName)
				if err != nil {
					fmt.Println("", err)
					resp.Status = "E"
					resp.ErrMsg = err.Error()
					fmt.Println(resp)
					fmt.Fprintln(w, resp)
				} else {
					fmt.Println("Stock is inserted")
					fmt.Fprintln(w, string(data))
				}

			}

		}
	}
}

func InsertMedicine(MedicineRec MedicineStruct) (MedicineResp, error) {
	var resp MedicineResp
	resp.Status = "S"
	db, err := LocalDBConnect()
	if err != nil {
		log.Println(err)
		resp.ErrMsg = err.Error()
		resp.Status = "E"
		fmt.Println("IM001", err)

	} else {
		defer db.Close()

		medicineInsert := `
		update loganathan.medapp_stock 
		set quantity = ? +(select quantity from loganathan.medapp_stock where medicine_name = ?) , unit_price =?
		where medicine_name  = ?
		`
		_, err := db.Exec(medicineInsert, &MedicineRec.Quantity, &MedicineRec.Medicinename, &MedicineRec.Unitprice, &MedicineRec.Medicinename)
		if err != nil {
			log.Println(err)
			resp.ErrMsg = err.Error()
			fmt.Println("IM002", err)
			resp.Status = "E"
		} else {
			// resp.MedicineArr = append(resp.MedicineArr, MedicineRec)
			resp.ErrMsg = "Success"
			resp.Status = "S"
		}

	}
	return resp, err
}
