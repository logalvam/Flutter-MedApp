package sales

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type StockPriceStruct struct {
	Total    float64 `json:total"`
	Amount   float64 `json:"amount"`
	Quantity float64 `json:"quantity"`
}

type StockResp struct {
	Status string `json:"status"`
	ErrMsg string `json:"errmsg"`
}

func StockPrices(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Credentials", "false")
	(w).Header().Set("Access-Control-Allow-Methods", "PUT,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "DATE,Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")

	if r.Method == "PUT" {
		fmt.Println("Stock Price (+)")
		var resp StockResp
		resp.Status = "S"

		lStock, err := GetStockPrice()
		if err != nil {
			fmt.Println("SP001", err)
			resp.Status = "E"
			resp.ErrMsg = err.Error()
			fmt.Fprintln(w, resp)
			fmt.Println(resp)
		} else {
			data, err := json.Marshal(lStock)
			if err != nil {
				fmt.Println("SP002", err)
				resp.Status = "E"
				resp.ErrMsg = err.Error()
				fmt.Fprintln(w, resp)
				fmt.Println(resp)
			} else {

				fmt.Fprintln(w, string(data))
				fmt.Println("Stock Price(-)")

			}
		}
	}
}

func GetStockPrice() (StockPriceStruct, error) {
	var resp StockResp
	var StockRec StockPriceStruct
	resp.Status = "S"
	db, err := LocalDBConnect()
	if err != nil {
		fmt.Println("GSP001", err)
		resp.Status = "E"
		resp.ErrMsg = err.Error()
		fmt.Println(resp)
	} else {
		defer db.Close()
		lStockPrice := `select  unit_price,quantity  from  loganathan.medapp_stock ms`
		rows, err := db.Query(lStockPrice)
		if err != nil {
			fmt.Println("GSP002", err)
			resp.Status = "E"
			resp.ErrMsg = err.Error()
		} else {
			for rows.Next() {
				err := rows.Scan(&StockRec.Amount, &StockRec.Quantity)
				if err != nil {
					resp.ErrMsg = err.Error()
					fmt.Println("GSP003", err)
					resp.Status = "E"
				} else {
					StockRec.Total += StockRec.Amount * StockRec.Quantity

				}
			}
		}
	}
	return StockRec, err

}
