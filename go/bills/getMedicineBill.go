package bills

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type MedicineInputStrct struct {
	MedicineName string `json:"medname"`
	Quantity     int    `json:"quantity"`
}

type MedicineInputDetails struct {
	MedicineName string `json:"medname"`
	Brand        string `json:"brand"`
	Quantity     int    `json:"quantity"`
	Amount       int    `json:"amount"`
}

type MedicineResponse struct {
	MedicineList []MedicineInputDetails `json:"medlist"`
	Status       string                 `json:"status"`
	ErrMsg       string                 `json:"errmsg"`
}

func GetMedicineInput(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Single medicine(+) ")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	(w).Header().Set("Access-Control-Allow-Methods", "PUT,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "User,Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")
	var resp MedicineResponse
	var InputRec MedicineInputStrct
	resp.Status = "S"
	body, err := io.ReadAll(r.Body)
	if err != nil {
		resp.Status = "E"
		fmt.Println("GMI001", err)
		resp.ErrMsg = err.Error()
	} else {
		err = json.Unmarshal(body, &InputRec)
		if err != nil {
			fmt.Println("GMI002", err)
			resp.Status = "E"
			resp.ErrMsg = err.Error()
		} else {
			MedicneDetils, err := GetMedicineDetails(InputRec)
			if err != nil {
				fmt.Println("GMI003", err)
				resp.Status = "E"
				resp.ErrMsg = err.Error()
			} else {
				data, err := json.Marshal(MedicneDetils)
				if err != nil {
					fmt.Println("GMI004", err)
					resp.Status = "E"
					resp.ErrMsg = err.Error()
				} else {
					fmt.Fprintln(w, string(data))
					// fmt.Println(resp)

				}
			}

		}
		fmt.Println("Single medicine (-) ")
	}
}

func GetMedicineDetails(InputRec MedicineInputStrct) (MedicineResponse, error) {

	var resp MedicineResponse
	var MedicineDRec MedicineInputDetails
	resp.Status = "S"

	db, err := LocalDBConnect()
	if err != nil {
		resp.Status = "E"
		fmt.Println("GMD001", err)

		resp.ErrMsg = err.Error()
	} else {
		defer db.Close()

		lBrand := ` select brand from loganathan.medapp_medicine_master
					where medicine_name = ?`
		rows, err := db.Query(lBrand, &InputRec.MedicineName)
		if err != nil {
			resp.ErrMsg = err.Error()
			fmt.Println("GMD002", err)
			resp.Status = "E"
		} else {
			for rows.Next() {
				err := rows.Scan(&MedicineDRec.Brand)
				if err != nil {
					resp.ErrMsg = err.Error()
					resp.Status = "E"
					fmt.Println("GMD003", err)

				} else {
					resp.ErrMsg = "Success"
					resp.Status = "S"
				}

			}

		}

		lAmount := `select unit_price from loganathan.medapp_stock 
					where medicine_name = ?  `
		arows, err := db.Query(lAmount, &InputRec.MedicineName)
		if err != nil {
			resp.ErrMsg = err.Error()
			fmt.Println("GMD004", err)
			resp.Status = "E"
		} else {
			qty := InputRec.Quantity
			MedicineDRec.Quantity = InputRec.Quantity
			MedicineDRec.MedicineName = InputRec.MedicineName
			for arows.Next() {
				err := arows.Scan(&MedicineDRec.Amount)
				MedicineDRec.Amount = qty * MedicineDRec.Amount
				fmt.Println(qty, MedicineDRec.Amount)
				if err != nil {
					resp.ErrMsg = err.Error()
					fmt.Println("GMD005", err)
					resp.Status = "E"
				} else {
					resp.ErrMsg = "Success"
					resp.Status = "S"
					resp.MedicineList = append(resp.MedicineList, MedicineDRec)
				}
			}
		}
	}
	return resp, err
}
