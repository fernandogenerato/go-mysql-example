package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go-mysql-example/database"
	"go-mysql-example/models/user"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("error to read request"))
	}

	user := user.ToUser(body)

	db := database.Connection()
	defer db.Close()

	st, err := db.Prepare("insert into user(username, email) values (?,?)")
	if err != nil {
		log.Fatal(err)
	}
	defer st.Close()

	result, err := st.Exec(user.Username, user.Email)
	if err != nil {
		log.Fatal(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("cannot find inserted ID"))
		return
	}

	w.WriteHeader(201)
	w.Write([]byte(fmt.Sprintf("id :%d created", id)))

}

func FindUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseInt(params["id"], 10, 32)
	if err != nil {
		w.Write([]byte("err to converter param to integer"))
		return
	}

	db := database.Connection()
	row, err := db.Query("select * from user where id = ?", userID)
	if err != nil {
		panic(err)
	}

	var u user.User
	if row.Next() {
		if err := row.Scan(&u.Username, &u.Email, &u.Id); err != nil {
			w.Write([]byte(fmt.Sprintf("cannot find userID: %d", userID)))
		}
	}

	response, err := json.Marshal(u)
	if err != nil {
		w.Write([]byte("Err to cast users to json"))
		return
	}
	w.Write(response)
}

func FindUsers(w http.ResponseWriter, r *http.Request) {
	db := database.Connection()
	defer db.Close()

	rows, err := db.Query("select * from user")
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("error to find all users"))
	}
	defer rows.Close()

	var users []user.User
	for rows.Next() {
		var user user.User
		if err := rows.Scan(&user.Username, &user.Email, &user.Id); err != nil {
			w.WriteHeader(500)
			w.Write([]byte("error to scan users data"))
			return
		}
		users = append(users, user)

	}
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(users); err != nil {
		w.Write([]byte("Err to cast users to json"))
		return
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseInt(params["id"], 10, 32)
	if err != nil {
		w.Write([]byte("err to converter param to integer"))
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("error to read request"))
	}

	user := user.ToUser(body)

	db := database.Connection()
	st, err := db.Prepare("update user set username = ?, email = ?  where id = ? ")
	if err != nil {
		panic(err)
	}
	defer st.Close()

	_, err = st.Exec(&user.Username, &user.Email, userID)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(204)

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseInt(params["id"], 10, 32)
	if err != nil {
		w.Write([]byte("err to converter param to integer"))
		return
	}

	db := database.Connection()
	_, err = db.Query("delete from user where id = ?", userID)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(204)
}
