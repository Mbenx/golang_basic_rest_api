package routes

import (
	"encoding/json"
	"fmt"
	"golang_basic_rest_api/config"
	"golang_basic_rest_api/models"
	"log"
	"net/http"
)

func GetEmployee(w http.ResponseWriter, r *http.Request) {
	db, err := config.Connect()
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query("select id, name, address, position from employees")
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer rows.Close()

	var result models.Employees

	for rows.Next() {
		var each = models.Employee{}
		var err = rows.Scan(&each.ID, &each.Name, &each.Address, &each.Position)

		if err != nil {
			log.Println(err.Error())
			return
		}

		result = append(result, each)
	}

	json.NewEncoder(w).Encode(result)
}

func PostEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		db, err := config.Connect()
		if err != nil {
			log.Println(err.Error())
			return
		}
		defer db.Close()

		var employee models.Employee

		err = json.NewDecoder(r.Body).Decode(&employee)
		if err != nil {
			log.Println(err.Error())
			return
		}

		log.Println(employee)

		_, err = db.Exec("insert into employees (name, address, position) values (?, ?, ?)",
			employee.Name, employee.Address, employee.Position)
		if err != nil {
			log.Println(err.Error())
			return
		}
		log.Println("insert success!")

		json.NewEncoder(w).Encode(employee)

	} else {
		http.Error(w, "invalid method", http.StatusMethodNotAllowed)
	}

}

func PutEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		db, err := config.Connect()
		if err != nil {
			log.Println(err.Error())
			return
		}
		defer db.Close()

		var employee models.Employee

		err = json.NewDecoder(r.Body).Decode(&employee)
		if err != nil {
			log.Println(err.Error())
			return
		}

		log.Println(employee)

		_, err = db.Exec("UPDATE employees SET name = ?, address = ?, position = ? WHERE id = ?",
			employee.Name, employee.Address, employee.Position, employee.ID)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("update success!")

		json.NewEncoder(w).Encode(employee)

	} else {
		http.Error(w, "invalid method", http.StatusMethodNotAllowed)
	}

}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		db, err := config.Connect()
		if err != nil {
			log.Println(err.Error())
			return
		}
		defer db.Close()

		var employee models.Employee

		err = json.NewDecoder(r.Body).Decode(&employee)
		if err != nil {
			log.Println(err.Error())
			return
		}

		log.Println(employee)

		_, err = db.Exec("DELETE FROM employees WHERE id = ?", employee.ID)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println("Delete Success")

		json.NewEncoder(w).Encode(employee)

	} else {
		http.Error(w, "invalid method", http.StatusMethodNotAllowed)
	}
}
