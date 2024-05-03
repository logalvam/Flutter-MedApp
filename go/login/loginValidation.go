package login

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type LoginStruct struct {
	UserId   string `json:"userid"`
	PassWord string `json:"password"`
}
type LoginResp struct {
	Role   string `json:"role"`
	Status string `json:"status"`
	ErrMsg string `json:"errMeg"`
}

func LoginValidation(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Login Validation(+)")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Credentials", "false")
	(w).Header().Set("Access-Control-Allow-Methods", "PUT,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")

	if r.Method == "PUT" {
		var LoginRec LoginStruct
		var resp LoginResp
		resp.Status = "S"

		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println("LV001", err)
			resp.Status = "E"
			resp.ErrMsg = err.Error()
			fmt.Fprintln(w, resp)
			fmt.Println(resp)
		} else {
			err := json.Unmarshal(body, &LoginRec)
			if err != nil {
				fmt.Println("LV002", err)
				resp.Status = "E"
				resp.ErrMsg = err.Error()
				fmt.Fprintln(w, resp)
				fmt.Println(resp)
			}
			role, err := userRole(LoginRec)
			if err != nil {
				fmt.Fprintln(w, "error in get data marchal"+err.Error())
			} else {
				data, err := json.Marshal(role)
				if err != nil {
					fmt.Println("LV003", err)
					resp.Status = "E"
					resp.ErrMsg = err.Error()
					fmt.Fprintln(w, resp)
					fmt.Println(resp)
				} else {

					fmt.Fprintln(w, string(data))

				}
			}
		}
	}
	fmt.Println("Login Validation(-)")

}

func userRole(LoginRec LoginStruct) (LoginResp, error) {
	var resp LoginResp
	resp.Status = "S"
	db, err := LocalDBConnect()
	if err != nil {
		log.Println(err)
		resp.ErrMsg = err.Error()
		fmt.Println("UR001", err)
		resp.Status = "E"
	} else {
		defer db.Close()
		role := `select role from loganathan.medapp_login where user_id = ? and password = ?`
		rows, err := db.Query(role, &LoginRec.UserId, &LoginRec.PassWord)
		if err != nil {
			resp.ErrMsg = err.Error()
			fmt.Println("UR002", err)
			resp.Status = "E"
			fmt.Println(resp)
		} else {
			for rows.Next() {
				err := rows.Scan(&resp.Role)
				if err != nil {
					resp.ErrMsg = err.Error()
					fmt.Println("UR003", err)
					resp.Status = "E"
					fmt.Println(resp)
				} else {
					resp.Status = "S"
					// fmt.Println(resp)

				}
			}
		}
	}
	return resp, err
}
