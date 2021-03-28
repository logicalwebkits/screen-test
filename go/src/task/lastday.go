package task

import (
    "net/http"
    "time" 
    "fmt" 
                  "strconv"
            "encoding/json"


    )
func Lastday(w http.ResponseWriter, r *http.Request) { 
			  w.Header().Set("Access-Control-Allow-Origin", "*")
 
	r.ParseForm()   
    myDateString := r.Form.Get("Date") 
    fmt.Println(myDateString)

myDate, _ := time.Parse("2006-01-02 15:04", myDateString)
getlastdate := time.Date(myDate.Year() , myDate.Month()+1, 0, 10, 0, 0, 0, time.Local)

	em_arr := make(map[string]string)
	em_arr["Date"] = myDateString
	em_arr["LastDayOfMonth"] = strconv.Itoa((getlastdate.Day()))
	
	w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(em_arr)


} 
