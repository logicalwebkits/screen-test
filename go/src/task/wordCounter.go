package task

import (
    "net/http"
        "regexp"
    "io/ioutil"
    "strings"
    htmlstrip "github.com/grokify/html-strip-tags-go"
        "encoding/json"

)
 
 
 
 
func WordCounter(w http.ResponseWriter, r *http.Request){  
	
	 w.Header().Set("Access-Control-Allow-Origin", "*")
	 
	 var words []WC
	 var words_all [][]WC
	 var cols = 5
	 
	 url := GetUrlParam(r,"url")
	 	
		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		
		body, _ := ioutil.ReadAll(resp.Body)	
		
		bodyfull := string(body) 
		
		
		g := htmlstrip.StripTags(bodyfull)  
		
		reger, _ := regexp.Compile("[^a-zA-Z]+")

		for i,x := range wrdC(g){
			sp := reger.ReplaceAllString(i, "")
			
			if sp !="" {
			words = append(words,WC{N:sp,C:x})
			
				if len(words) == cols {
				
					words_all = append(words_all,words)
					words = nil 
 
				}
			}
			
		} 
		 
		if len(words) > 0  {
			words_all = append(words_all,words)
		}
		
		
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(words_all)

		
}



type WC struct {
	N string  `json:"name"`
	C int	  `json:"count"`
} 



func GetUrlParam(r *http.Request, val string) string {
	
	value := r.URL.Query()[val]
	if len(value) != 0 {	
	return value[0]
	} else
	{
		return ""
	}
}


func wrdC(str string) map[string]int {
	
    cc := make(map[string]int)
    
    for _, w := range strings.Fields(str) {
        _, ok := cc[w]
        
        if ok {
            cc[w] += 1
        } else {
            cc[w] = 1
        }
    }
    return cc
}
