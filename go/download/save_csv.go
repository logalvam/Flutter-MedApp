package download

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Record struct {
	MedName  string  `json:"medname"`
	Brand    string  `json:"brand"`
	Quantity int     `json:"quantity"`
	Amount   float64 `json:"amount"`
	BillGST  float64 `json:"billgst"`
	NetPrice float64 `json:"netprice"`
	BillNo   int     `json:"billno"`
}

type DownloadResp struct {
	Status string `json:"status"`
	ErrMsg string `json:"errmsg"`
}

func CsvFileWrite(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Download CSV(+)")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Credentials", "false")
	(w).Header().Set("Access-Control-Allow-Methods", "PUT,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "BILLNO,Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")

	if r.Method == "PUT" {

		var records []Record
		var resp DownloadResp
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
			resp.ErrMsg = err.Error()
			resp.Status = "E"
		} else {
			err := json.Unmarshal(body, &records)
			if err != nil {
				fmt.Println(err)
				resp.ErrMsg = err.Error()
				resp.Status = "E"
			} else {

				filename := fmt.Sprintf("BillNo_%d", records[0].BillNo)

				file, err := os.Create(filename)
				if err != nil {
					fmt.Println(err)
					resp.ErrMsg = err.Error()
					resp.Status = "E"
				} else {
					defer file.Close()
					writer := csv.NewWriter(file)
					writer.Write([]string{"MedName", "Brand", "Quantity", "Amount", "BillGST", "NetPrice", "BillNo"})
					for _, record := range records {
						writer.Write([]string{record.MedName, record.Brand, fmt.Sprintf("%d", record.Quantity), fmt.Sprintf("%f", record.Amount), fmt.Sprintf("%f", record.BillGST), fmt.Sprintf("%f", record.NetPrice), fmt.Sprintf("%d", record.BillNo)})
					}
					writer.Flush()
					resp.Status = "S"
					resp.ErrMsg = "Success"
					// fmt.Fprintln(w, resp)
					fmt.Println(resp)
					fmt.Println("Download CSV(-)")

					data, err := json.Marshal(resp)
					if err != nil {
						resp.ErrMsg = err.Error()
						resp.Status = "E"
					} else {
						resp.Status = "S"
						fmt.Fprintln(w, string(data))
					}
				}
			}
		}

	}
}
