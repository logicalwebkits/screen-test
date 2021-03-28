package task

import (
    "net/http"
            "encoding/json"
            "strconv"

    )
    
    
    
func Prime(w http.ResponseWriter, r *http.Request) {  
	
				 w.Header().Set("Access-Control-Allow-Origin", "*")

 number ,_:= strconv.Atoi(GetUrlParam(r,"number"))
var output []int

for i:=1; i <=number ; i++ {

	ch := make(chan int)

	go CheckPrime(i,ch)
	v := <-ch
	if v > 0 {
	output = append(output,v)
	}
}


w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(output)

}


func CheckPrime(number int,chanl chan int) {  
 isPrime := true  
  
  for i := 2; i <= number/2; i++ {  
   if number%i == 0 {  
    isPrime = false  
    break  
   }  
  }  
  
if isPrime == true && number !=1 {
chanl <- number
} else
{
chanl <- 0
}

}  

  
