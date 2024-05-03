package login

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type NewUserStruct struct {
	LoginId     int    `json:"loginid"`
	UserId      string `json:"userid"`
	PassWord    string `json:"password"`
	Role        string `json:"role"`
	CreatedBy   string `json:"createdBy"`
	CreatedDate string `json:"createdDate"`
	UpdatedBy   string `json:"updatedBy"`
	UpdatedDate string `json:"updatedDate"`
}

type NewUserResp struct {
	Status string `json:"status"`
	ErrMsg string `json:"errMeg"`
}

func AddNewUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Insert New User")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Credentials", "false")
	(w).Header().Set("Access-Control-Allow-Methods", "PUT,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "User,Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")
	if r.Method == "PUT" {
		var UserRec NewUserStruct
		var resp NewUserResp
		resp.Status = "S"
		body, err := io.ReadAll(r.Body)
		if err != nil {
			resp.Status = "E"
			fmt.Println("ANU001", err)

			resp.ErrMsg = err.Error()
			fmt.Fprintln(w, resp)
			fmt.Println(resp)
		} else {
			err := json.Unmarshal(body, &UserRec)
			if err != nil {
				resp.Status = "E"
				fmt.Println("ANU002", err)
				resp.ErrMsg = err.Error()
				fmt.Println(resp)
			} else {
				Result, err := AddUser(UserRec)
				if err != nil {
					fmt.Println("ANU003", err)
					resp.Status = "E"
					resp.ErrMsg = err.Error()
					fmt.Fprintln(w, resp)
					fmt.Println(resp)
				} else {
					data, err := json.Marshal(Result)
					if err != nil {
						fmt.Println("ANU004", err)
						resp.Status = "E"
						resp.ErrMsg = err.Error()
						fmt.Fprintln(w, resp)
						fmt.Println(resp)
					} else {
						fmt.Fprintln(w, string(data))
						fmt.Println("User Added")
					}
				}
			}
		}

	}
}

func AddUser(UserRec NewUserStruct) (NewUserResp, error) {
	var resp NewUserResp
	resp.Status = "S"
	db, err := LocalDBConnect()
	if err != nil {
		resp.ErrMsg = err.Error()
		fmt.Println("AU001", err)
		resp.Status = "E"
	} else {
		defer db.Close()

		loginId := `select count(*) from loganathan.medapp_login`
		rows, err := db.Query(loginId)
		if err != nil {
			resp.ErrMsg = err.Error()
			fmt.Println("AU002", err)
			resp.Status = "E"
			fmt.Println(resp)
		} else {
			for rows.Next() {
				err := rows.Scan(&UserRec.LoginId)
				if err != nil {
					resp.ErrMsg = err.Error()
					fmt.Println("AU003", err)
					resp.Status = "E"
					fmt.Println(resp)
				} else {

					InserUser := `insert  into  loganathan.medapp_login (login_id,user_id,password,role,created_by,created_date,updated_by,updated_date)
		select ?,?,?,?,"admin" ,curdate(),"admin",curdate()
		from dual
		where not exists (select 1 from loganathan.medapp_login
		where user_id = ?)`
					loginid := UserRec.LoginId + 1
					_, err := db.Exec(InserUser, loginid, &UserRec.UserId, &UserRec.PassWord, &UserRec.Role, &UserRec.UserId)
					// pENDING
					if err != nil {
						resp.ErrMsg = err.Error()
						fmt.Println("AU004", err)
						resp.Status = "E"
						fmt.Println(resp)
					} else {
						resp.Status = "S"
						fmt.Println("AU004", err)
						fmt.Println(resp)
					}
				}
			}
		}

	}
	return resp, err
}
