package main

import (
    "task"
    "github.com/gorilla/mux"
    "net/http"
)
 
func main() {
   
	r := mux.NewRouter().StrictSlash(true)
    r.HandleFunc("/wordCounter", task.WordCounter)
    r.HandleFunc("/excelMaker", task.ExcelMaker)  
    r.HandleFunc("/lastday", task.Lastday).Methods("POST")
    r.HandleFunc("/addEditUser", task.AddEditUser).Methods("POST")
    r.HandleFunc("/prime", task.Prime)  
    r.HandleFunc("/allUsers", task.AllUsers)  
    r.HandleFunc("/getUser", task.GetUser)  
    r.HandleFunc("/userDelete", task.UserDelete)  
    
  	if err := http.ListenAndServe(":1010", r); err != nil {
		panic(err)
	}   
}
