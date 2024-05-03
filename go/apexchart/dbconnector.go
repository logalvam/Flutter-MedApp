package apexchart

import (
	"database/sql"
	"fmt"
	"log"
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
