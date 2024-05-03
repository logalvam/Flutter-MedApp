package main

import (
	"fmt"
	"go/apexchart"
	"go/bills"
	"go/download"
	"go/login"
	"go/sales"
	"go/stock"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Api Start (+)")
	http.HandleFunc("/insertMedicineBrand", stock.InsertBrand)
	http.HandleFunc("/insertMedicineName", stock.AddMedicine)
	http.HandleFunc("/loginValidation", login.LoginValidation)
	http.HandleFunc("/loginhistoryInsert", login.LoginHistory)
	http.HandleFunc("/logouthistoryInsert", login.LogoutHistory)
	http.HandleFunc("/viewLogHistory", login.ViewLogHistory)
	http.HandleFunc("/viewMedicine", stock.ViewMedicine)
	http.HandleFunc("/viewBrand", stock.ViewBrand)
	http.HandleFunc("/getMedicineByBrand", stock.MedicineName)
	http.HandleFunc("/MedicineInput", bills.GetMedicineInput)
	http.HandleFunc("/addNewUser", login.AddNewUser)

	http.HandleFunc("/MedicineOutput", bills.BillMasterDetails)
	http.HandleFunc("/BillDetails", bills.BillDetails)

	http.HandleFunc("/userRoleFetching", login.FetchRole)
	http.HandleFunc("/salesreport", sales.SalesReport)
	http.HandleFunc("/todaysales", sales.TodaSales)
	http.HandleFunc("/comparesales", sales.CompareSales)
	http.HandleFunc("/stockprices", sales.StockPrices)
	http.HandleFunc("/todaysalesmanager", sales.ManagerSales)
	http.HandleFunc("/fetchQuantity", stock.FetchMedQuantity)
	http.HandleFunc("/downloadBill", download.CsvFileWrite)
	http.HandleFunc("/downloadPdf", download.BillPdf)
	http.HandleFunc("/monthsales", apexchart.MonthSales)
	http.HandleFunc("/billersales", apexchart.BillerSales)
	http.ListenAndServe(":8260", nil)
}
