package stock

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func LocalDBConnect() (*sql.DB, error) {
	log.Println("Local DBConnect +")
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", "root", "root", "192.168.2.5", 3306, "loganathan")
	db, err := sql.Open("mysql", connString)

	if err != nil {
		log.Println("open connection failed:", err.Error())
		return db, err
	}
	log.Println("local DBConnect -")
	return db, nil
}

type BrandStruct struct {
	Medicinename string `json:"medname"`
	Brand        string `json:"medbrand"`
	CreatedBy    string `json:"createdBy"`
	CreatedDate  string `json:"createdDate"`
	UpdatedBy    string `json:"updatedBy"`
	UpdatedDate  string `json:"updatedDate"`
}

type BrandResp struct {
	Status string `json:"status"`
	ErrMsg string `json:"errMeg"`
}

func InsertBrand(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Insert Medicine Name")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Credentials", "false")
	(w).Header().Set("Access-Control-Allow-Methods", "PUT,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "User,Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")
	if r.Method == "PUT" {
		var BrandRec BrandStruct
		var resp BrandResp
		resp.Status = "S"

		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println("IB001", err)
			resp.Status = "E"
			resp.ErrMsg = err.Error()
		} else {
			err := json.Unmarshal(body, &BrandRec)
			if err != nil {
				resp.Status = "E"
				fmt.Println("IB001", err)

				resp.ErrMsg = err.Error()
			}
			fmt.Println(BrandRec)
			brand, err := MedicineBrand(BrandRec)
			if err != nil {
				resp.Status = "E"
				fmt.Println("IB001", err)

				resp.ErrMsg = err.Error()
				fmt.Println(resp)
				fmt.Fprintln(w, resp)
			}
			data, err := json.Marshal(brand)
			if err != nil {
				resp.Status = "E"
				fmt.Println("IB001", err)

				resp.ErrMsg = err.Error()
				fmt.Println(resp)
				fmt.Fprintln(w, resp)
			} else {
				fmt.Print("insert brand")
				fmt.Fprintln(w, string(data))
			}
		}
	}
}

func MedicineBrand(BrandRec BrandStruct) (BrandResp, error) {
	var resp BrandResp
	resp.Status = "S"
	db, err := LocalDBConnect()
	if err != nil {
		log.Println(err)
		resp.ErrMsg = err.Error()
		resp.Status = "E"
	} else {
		defer db.Close()
		brandinsert := `insert into loganathan.medapp_medicine_master(medicine_name,brand,created_by,created_date,updated_by,updated_date)
		select ?, ?,"admin", CURDATE(), "admin", CURDATE() from dual
		where not exists(select 1 from loganathan.medapp_stock 
			where medicine_name = ?) `
		_, err := db.Exec(brandinsert, &BrandRec.Medicinename, &BrandRec.Brand, &BrandRec.Medicinename)
		if err != nil {
			log.Println(err)
			fmt.Println("MB001", err)

			resp.ErrMsg = err.Error()
			resp.Status = "E"
		} else {
			resp.ErrMsg = "Success"
			resp.Status = "S"
		}
		inserInStock := `insert  into  loganathan.medapp_stock (medicine_name,quantity,unit_price,created_by,created_date,updated_by,updated_date)
		values(?,0,0,"manager",curdate(),"manager",curdate())
		 `
		_, err1 := db.Exec(inserInStock, BrandRec.Medicinename)
		if err1 != nil {
			fmt.Println("MB002", err1)
			log.Println(err1)
			resp.ErrMsg = err1.Error()
			resp.Status = "E"
		} else {
			resp.ErrMsg = "Success"
			resp.Status = "S"
		}
	}
	return resp, err
}

func MedicineName(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Insert Medicine Name")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Credentials", "false")
	(w).Header().Set("Access-Control-Allow-Methods", "PUT,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "MEDNAME,Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")
	if r.Method == "PUT" {
		// var BrandRec BrandStruct
		var resp BrandResp
		resp.Status = "S"

		medicineName := r.Header.Get("MEDNAME")
		brand, err := FetchMedicineName(medicineName)
		if err != nil {
			resp.Status = "E"
			fmt.Println("MN001", err)

			resp.ErrMsg = err.Error()
			fmt.Println(resp)
			fmt.Fprintln(w, resp)
		}
		data, err := json.Marshal(brand)
		if err != nil {
			fmt.Println("MN002", err)
			resp.Status = "E"
			resp.ErrMsg = err.Error()
			fmt.Println(resp)
			fmt.Fprintln(w, resp)
		} else {
			fmt.Print("insert brand")
			fmt.Fprintln(w, string(data))
		}
		// }
	}
}
func FetchMedicineName(medicineName string) (string, error) {
	var brand string
	var resp BrandResp
	resp.Status = "S"
	db, err := LocalDBConnect()
	if err != nil {
		log.Println(err)
		resp.ErrMsg = err.Error()
		resp.Status = "E"
	} else {
		defer db.Close()

		getbrand := ` select  brand  from loganathan.medapp_medicine_master mmm
		where  medicine_name  =?`
		rows, err := db.Query(getbrand, medicineName)
		if err != nil {
			log.Println(err)
			resp.ErrMsg = err.Error()
			resp.Status = "E"
			fmt.Println("FMN001", err)

		} else {
			for rows.Next() {
				err := rows.Scan(&brand)
				if err != nil {
					log.Println(err)
					resp.ErrMsg = err.Error()
					resp.Status = "E"
					fmt.Println("FMN002", err)

				} else {
					resp.ErrMsg = "Success"
					resp.Status = "S"
				}
			}
		}

	}
	fmt.Println(brand)

	return brand, err
}
