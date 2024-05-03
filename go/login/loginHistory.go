package login

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type HistoryStruct struct {
	UserId      string `json:"userid"`
	LoginDate   string `json:"logindate"`
	LoginTime   string `json:"logintime"`
	LogoutTime  string `json:"logouttime"`
	CreatedBy   string `json:"createdBy"`
	CreatedDate string `json:"createdDate"`
	UpdatedBy   string `json:"updatedBy"`
	UpdatedDate string `json:"updatedDate"`
}
type HistoryResp struct {
	Role   string `json:"role"`
	Status string `json:"status"`
	ErrMsg string `json:"errMeg"`
}

func LoginHistory(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Login History Insert")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Credentials", "false")
	(w).Header().Set("Access-Control-Allow-Methods", "PUT,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "USER,Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")

	if r.Method == "PUT" {
		var HistoryRec HistoryStruct
		var resp LoginResp
		resp.Status = "S"
		User := r.Header.Get("USER")
		fmt.Println("****")
		fmt.Println(User)
		HistoryRec.UserId = User

		Result, err := insertLoginHistory(User)
		if err != nil {
			fmt.Println("LH001", err)
			resp.Status = "E"
			resp.ErrMsg = err.Error()
			fmt.Fprintln(w, resp)
			fmt.Println(resp)
		} else {
			data, err := json.Marshal(Result)
			if err != nil {
				fmt.Println("LH002", err)
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

func insertLoginHistory(User string) (HistoryResp, error) {
	var resp HistoryResp
	// var UserRec UserStruct

	resp.Status = "S"
	db, err := LocalDBConnect()
	if err != nil {
		log.Println(err)
		resp.ErrMsg = err.Error()
		fmt.Println("ILH001", err)
		resp.Status = "E"
	} else {
		defer db.Close()
		lLoginHis := `insert into loganathan.medapp_login_history(login_date,login_time,logout_time,created_by,created_date,updated_by,updated_date,user_id)values(curdate(),Now(),null,?,curdate(),?,curdate(),?)`
		_, err := db.Exec(lLoginHis, User, User, User)
		if err != nil {
			resp.ErrMsg = err.Error()
			fmt.Println("ILH002", err)
			resp.Status = "E"
			fmt.Println(resp)
		} else {
			resp.Status = "S"
			fmt.Println(resp)
		}
	}
	return resp, err
}

func LogoutHistory(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Logout History Insert")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Credentials", "false")
	(w).Header().Set("Access-Control-Allow-Methods", "PUT,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "User,Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")

	if r.Method == "PUT" {
		var resp HistoryResp
		resp.Status = "S"

		Result, err := insertLogoutHistory()
		if err != nil {
			fmt.Println("LOH001", err)
			resp.Status = "E"
			resp.ErrMsg = err.Error()
			fmt.Fprintln(w, resp)
			fmt.Println(resp)
		} else {
			data, err := json.Marshal(Result)
			if err != nil {
				fmt.Println("LOH002", err)
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

func insertLogoutHistory() (HistoryResp, error) {
	var resp HistoryResp
	// var UserRec UserStruct

	resp.Status = "S"
	db, err := LocalDBConnect()
	if err != nil {
		log.Println(err)
		resp.ErrMsg = err.Error()
		fmt.Println("ILOH001", err)
		resp.Status = "E"
	} else {
		defer db.Close()
		lLoginHis := `update  loganathan.medapp_login_history mlh 
		set  logout_time  = curtime()
		where  login_history_id = (select count(*) from loganathan.medapp_login_history mlh2)`
		_, err := db.Exec(lLoginHis)
		if err != nil {
			resp.ErrMsg = err.Error()
			fmt.Println("ILOH002", err)
			resp.Status = "E"
			fmt.Println(resp)
		} else {
			resp.Status = "S"
			fmt.Println(resp)
		}
		lgetRole := `select user_id  from loganathan.medapp_login_history mlh 
		where login_history_id  = (select count(*) from loganathan.medapp_login_history mlh2)
		`
		rows, err := db.Query(lgetRole)
		if err != nil {
			resp.ErrMsg = err.Error()
			fmt.Println("ILOH003", err)
			resp.Status = "E"
			fmt.Println(resp)
		} else {
			for rows.Next() {
				err := rows.Scan(&resp.Role)
				if err != nil {
					resp.ErrMsg = err.Error()
					resp.Status = "E"
					fmt.Println("ILOH004", err)
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

type LogHistoryView struct {
	UserId     string `json:"userid"`
	Date       string `json:"date"`
	LoginTime  string `json:"logintime"`
	LogoutTime string `json:"logouttime"`
}

type LogHisResp struct {
	LogHisList []LogHistoryView `json:"hislist"`
	Status     string           `json:"status"`
	ErrMsg     string           `json:"errmsg"`
}

func ViewLogHistory(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Logout History Insert")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Credentials", "false")
	(w).Header().Set("Access-Control-Allow-Methods", "GET,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "User,Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")

	if r.Method == "GET" {
		var LoghisRec LogHistoryView
		var resp LogHisResp
		resp.Status = "S"

		Result, err := HistoryList(LoghisRec)
		if err != nil {
			resp.ErrMsg = err.Error()
			fmt.Println("VLH001", err)

			resp.Status = "E"
		} else {
			data, err := json.Marshal(Result)
			if err != nil {
				resp.Status = "E"
				fmt.Println("VLH002", err)

				resp.ErrMsg = err.Error()
			} else {
				fmt.Println("Histor of login list Geted")
				fmt.Fprintln(w, string(data))
			}
		}

	}
}

func HistoryList(LoghisRec LogHistoryView) (LogHisResp, error) {
	var resp LogHisResp
	resp.Status = "S"
	db, err := LocalDBConnect()
	if err != nil {
		log.Println(err)
		resp.ErrMsg = err.Error()
		fmt.Println("HL001", err)

		resp.Status = "E"
	} else {
		defer db.Close()

		getLogHis := `SELECT ifnull(user_id,""),login_date,login_time ,ifnull(logout_time ,"")   from loganathan.medapp_login_history mlh `
		rows, err := db.Query(getLogHis)
		if err != nil {
			resp.ErrMsg = err.Error()
			fmt.Println("HL002", err)

			resp.Status = "E"
		} else {
			for rows.Next() {
				err := rows.Scan(&LoghisRec.UserId, &LoghisRec.Date, &LoghisRec.LoginTime, &LoghisRec.LogoutTime)
				if err != nil {
					resp.ErrMsg = err.Error()
					resp.Status = "E"
					fmt.Println("HL003", err)

				} else {
					resp.LogHisList = append(resp.LogHisList, LoghisRec)
				}
			}

		}
	}

	return resp, err
}
