package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Employee struct {
	Name string
}

func employeeHandler(w http.ResponseWriter, r *http.Request) {
	var emp = new(Employee)
	emp.Name = "Gaurav Handa"
	out, err := json.Marshal(*emp)
	responseString := "{}"
	if err == nil {
		responseString = string(out)
	} else {
		responseString = `{"error" : "` + err.Error() + `"}`
	}

	w.WriteHeader(200)
	w.Header().Set("content-type", "application/json")
	fmt.Fprintf(w, responseString)
}

func main() {
	http.HandleFunc("/Employees", employeeHandler)
	http.ListenAndServe(":8080", nil)
}
