package download

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jung-kurt/gofpdf"
)

func BillPdf(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Download PDF(+)")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Credentials", "false")
	(w).Header().Set("Access-Control-Allow-Methods", "PUT,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "BILLNO,Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")

	if r.Method == "PUT" {
		var billpdf []Record
		var resp DownloadResp
		resp.Status = "S"
		billno := r.Header.Get("BILLNO")
		fmt.Println(billno, "bill")
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			resp.Status = "E"
			resp.ErrMsg = err.Error()
			fmt.Fprintln(w, resp)
		} else {
			err = json.Unmarshal(body, &billpdf)
			if err != nil {
				resp.ErrMsg = err.Error()
				resp.Status = "E"
				fmt.Fprintln(w, resp)
			} else {
				pdf := gofpdf.New("P", "mm", "A4", "")
				pdf.AddPage()

				pdf.SetFont("Arial", "B", 16)
				y := 10
				pdf.SetX(10)
				pdf.CellFormat(0, 10, "Medical Shop Name", "", 0, "CM", false, 0, "")
				pdf.Ln(-1)
				pdf.CellFormat(0, 10, "Bill No:"+billno, "", 0, "CM", false, 0, "")
				pdf.Ln(-1)

				y += 10

				pdf.SetX(10)
				pdf.CellFormat(80, 10, "Medicine Name", "1", 0, "", false, 0, "")
				pdf.CellFormat(30, 10, "Qty", "1", 0, "", false, 0, "")
				pdf.CellFormat(40, 10, "Amount", "1", 0, "", false, 0, "")
				pdf.Ln(-1)
				y += 10

				var total float64

				for _, item := range billpdf {
					pdf.SetX(10)
					pdf.CellFormat(80, 10, item.MedName, "", 0, "", false, 0, "")
					pdf.CellFormat(30, 10, fmt.Sprintf("%d", item.Quantity), "", 0, "", false, 0, "")
					pdf.CellFormat(40, 10, fmt.Sprintf("%.2f", item.Amount), "", 0, "", false, 0, "")
					pdf.Ln(-1)
					y += 10
					total += item.Amount
				}

				pdf.SetX(10)
				pdf.CellFormat(0, 10, fmt.Sprintf("Total: %.2f", total), "0", 0, "R", false, 0, "")
				pdf.Ln(-1)
				y += 10

				gst := total * 0.18
				pdf.SetX(10)
				pdf.CellFormat(0, 10, fmt.Sprintf("GST: %.2f", gst), "0", 0, "R", false, 0, "")
				pdf.Ln(-1)
				y += 10

				netPrice := total + gst
				pdf.SetX(10)
				pdf.CellFormat(0, 10, fmt.Sprintf("NetPrice: %.2f", netPrice), "0", 0, "R", false, 0, "")
				pdf.Ln(-1)
				y += 10

				filename := "Bill_no" + billno + ".pdf"
				err := pdf.OutputFileAndClose(filename)
				if err != nil {
					panic(err)
				}
			}
		}
		data, err := json.Marshal(resp)
		if err != nil {
			resp.Status = "E"
			resp.ErrMsg = err.Error()
			fmt.Fprintln(w, resp)
			fmt.Println(resp)
		} else {
			resp.Status = "S"
			fmt.Fprintln(w, string(data))
		}
	}

}
