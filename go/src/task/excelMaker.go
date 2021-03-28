package task

import (
    "net/http"
            "encoding/json"
            "strconv"
            "strings"

    )
func ExcelMaker(w http.ResponseWriter, r *http.Request) {  
		 w.Header().Set("Access-Control-Allow-Origin", "*")

 col ,_:= strconv.Atoi(GetUrlParam(r,"c"))
 row,_ := strconv.Atoi(GetUrlParam(r,"r"))
 start := strings.ToUpper(GetUrlParam(r,"s"))
 
 
 total_values := col * row
 var values []string
 var tmp_va []string
 var values_all [][]string
 

 for i := 0; i < total_values; i++ {
	 
	 if(len(values) > 0) {
		last := values[len(values)-1]
		strin := updateChar(last)
		values = append(values,strin)
		
		}	else
		{
			values = append(values,start)
		}
		
		
	}
	
	
	for _,i := range values {
		
		tmp_va = append(tmp_va,i)
		
	if len(tmp_va) == col {
			values_all = append(values_all,tmp_va) 
			tmp_va = nil
		}
	
	}
	
    w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(values_all)
    
} 




func updateChar(s string) string {
	
	
		var arr = [...]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
	"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	
	em_arr := make(map[string]string)
	
	for i, s := range arr {
		if s !="Z" {
			em_arr[s] = arr[i+1]
		} else
		{
			em_arr[s] = "A"
		}
		
	}
	
	var sp string
	
	var brk = 0;
	
	for i := (len(s)-1); i >= 0; i-- {
		 
		active_var := string(s[i])
	

		if brk == 0 {
		sp = em_arr[active_var] +sp
		} else
		{
			
			sp= active_var +sp
			
		} 
		 
		if i==0 && brk ==0 && active_var == "Z" {
			sp= "A"+sp
		} 
		
		if active_var != "Z" {
			brk = 1
		} 
		
		
		
	}
	
	
	
	
	return sp
	
	
	
	
}
