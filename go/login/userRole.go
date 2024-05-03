package login

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type UserResp struct {
	UserId string `json:"userid"`
	Role   string `json:"role"`
	Status string `json:"status"`
	ErrMsg string `json:"errMeg"`
}

func FetchRole(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Credentials", "false")
	(w).Header().Set("Access-Control-Allow-Methods", "PUT,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "User,Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")

	if r.Method == "PUT" {
		var resp HistoryResp
		resp.Status = "S"

		Result, err := useridRole()
		if err != nil {
			fmt.Println("FR001", err)
			resp.Status = "E"
			resp.ErrMsg = err.Error()
			fmt.Fprintln(w, resp)
			fmt.Println(resp)
		} else {
			data, err := json.Marshal(Result)
			if err != nil {
				fmt.Println("FR002", err)
				resp.Status = "E"
				resp.ErrMsg = err.Error()
				fmt.Fprintln(w, resp)
				fmt.Println(resp)
			} else {
				fmt.Fprintln(w, string(data))
				fmt.Println(resp)

			}
		}
	}
}
func useridRole() (UserResp, error) {
	var resp UserResp
	resp.Status = "S"
	db, err := LocalDBConnect()
	if err != nil {
		log.Println(err)
		resp.ErrMsg = err.Error()
		fmt.Println("UIDR001", err)
		resp.Status = "E"
	} else {
		defer db.Close()
		luserID := `select user_id  from loganathan.medapp_login_history mlh 
		where login_history_id  = (select count(*) from loganathan.medapp_login_history mlh2)
		`
		rows, err := db.Query(luserID)
		if err != nil {
			resp.ErrMsg = err.Error()
			fmt.Println("UIDR002", err)
			resp.Status = "E"
			fmt.Println(resp)
		} else {
			for rows.Next() {
				err := rows.Scan(&resp.UserId)
				if err != nil {
					resp.ErrMsg = err.Error()
					fmt.Println("UIDR003", err)
					resp.Status = "E"
					fmt.Println(resp)
				} else {
					resp.ErrMsg = ""
					resp.Status = "S"
					fmt.Println(resp)
				}
			}
		}

		//

		luserRole := `select role  from loganathan.medapp_login mlh 
		where user_id = ?`
		rows1, err := db.Query(luserRole, &resp.UserId)
		if err != nil {
			resp.ErrMsg = err.Error()
			fmt.Println("UIDR004", err)
			resp.Status = "E"
			fmt.Println(resp)
		} else {
			for rows1.Next() {
				err := rows1.Scan(&resp.Role)
				if err != nil {
					resp.ErrMsg = err.Error()
					fmt.Println("UIDR005", err)
					resp.Status = "E"
					fmt.Println(resp)
				} else {
					resp.ErrMsg = ""
					resp.Status = "S"
					fmt.Println(resp)
				}
			}
		}

	}

	return resp, err
}
