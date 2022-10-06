package routes

import (
	"encoding/json"
	"fmt"
	"golang_basic_rest_api/config"
	"golang_basic_rest_api/models"
	"log"
	"net/http"
)

func GetDepartment(w http.ResponseWriter, r *http.Request) {
	db, err := config.Connect()
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query("select id, name, code from departments")
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer rows.Close()

	var result models.Departments

	for rows.Next() {
		var each = models.Department{}
		var err = rows.Scan(&each.ID, &each.Name, &each.Code)

		if err != nil {
			log.Println(err.Error())
			return
		}

		result = append(result, each)
	}

	json.NewEncoder(w).Encode(result)
}

func PostDepartment(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		db, err := config.Connect()
		if err != nil {
			log.Println(err.Error())
			return
		}
		defer db.Close()

		var department models.Department

		err = json.NewDecoder(r.Body).Decode(&department)
		if err != nil {
			log.Println(err.Error())
			return
		}

		log.Println(department)

		_, err = db.Exec("insert into departments (name, code) values (?, ?)",
			department.Name, department.Code)
		if err != nil {
			log.Println(err.Error())
			return
		}
		log.Println("insert success!")

		json.NewEncoder(w).Encode(department)

	} else {
		http.Error(w, "invalid method", http.StatusMethodNotAllowed)
	}

}

func PutDepartment(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		db, err := config.Connect()
		if err != nil {
			log.Println(err.Error())
			return
		}
		defer db.Close()

		var department models.Department

		err = json.NewDecoder(r.Body).Decode(&department)
		if err != nil {
			log.Println(err.Error())
			return
		}

		log.Println(department)

		_, err = db.Exec("UPDATE departments SET name = ?, code = ? WHERE id = ?",
			department.Name, department.Code, department.ID)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("update success!")

		json.NewEncoder(w).Encode(department)

	} else {
		http.Error(w, "invalid method", http.StatusMethodNotAllowed)
	}

}

func DeleteDepartment(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		db, err := config.Connect()
		if err != nil {
			log.Println(err.Error())
			return
		}
		defer db.Close()

		var department models.Department

		err = json.NewDecoder(r.Body).Decode(&department)
		if err != nil {
			log.Println(err.Error())
			return
		}

		log.Println(department)

		_, err = db.Exec("DELETE FROM departments WHERE id = ?", department.ID)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println("Delete Success")

		json.NewEncoder(w).Encode(department)

	} else {
		http.Error(w, "invalid method", http.StatusMethodNotAllowed)
	}
}
