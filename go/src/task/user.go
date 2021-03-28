
package task

import (
    "net/http"
        "database/sql"
    _ "github.com/go-sql-driver/mysql"

            "encoding/json"


    )
    
    
type User struct { 
    Id         int  `json:"id"`
    Login_Id       string `json:"login_id"`
    Full_name string `json:"full_name"`
}


func DbConnect() *sql.DB {

db, _ := sql.Open("mysql", "root:root123@tcp(localhost:3306)/testdb")

	return db


}




    
   func AllUsers(w http.ResponseWriter, r *http.Request) { 

	w.Header().Set("Access-Control-Allow-Origin", "*")
    	
    	em_arr_all := []User{}

		db := DbConnect() 
		defer db.Close()
		
     
    res, _ := db.Query("SELECT id,login_id,full_name FROM users")

    defer res.Close()


    for res.Next() {

      em_arr := User{}
      
        res.Scan(&em_arr.Id, &em_arr.Login_Id, &em_arr.Full_name)

        
        
        em_arr_all = append(em_arr_all,em_arr)

    }
    
    
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(em_arr_all)

    
}
 
   func GetUser(w http.ResponseWriter, r *http.Request) { 

	w.Header().Set("Access-Control-Allow-Origin", "*")
    	id := GetUrlParam(r,"id")
      em_arr := User{}

db := DbConnect() 
		defer db.Close()
		
    res, _ := db.Query("SELECT id,login_id,full_name FROM users where id =" +id)

    defer res.Close()

 
    for res.Next() {

      res.Scan(&em_arr.Id, &em_arr.Login_Id, &em_arr.Full_name)


    }
    
    
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(em_arr)

    
}
   func UserDelete(w http.ResponseWriter, r *http.Request) { 

	w.Header().Set("Access-Control-Allow-Origin", "*")
    	id := GetUrlParam(r,"id")
		db := DbConnect() 
		defer db.Close()
    
    sql := "DELETE FROM users WHERE  id = "+id
		
     db.Exec(sql)

       
         	em_arr := make(map[string]string)

    em_arr["success"]  = "user Deleted"
    
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(em_arr)

    
}
 
    
func AddEditUser(w http.ResponseWriter, r *http.Request) { 
	
	w.Header().Set("Access-Control-Allow-Origin", "*")
	r.ParseForm()   
 
    login_id := r.Form.Get("login_id")
    edit_id := r.Form.Get("edit_id")
    full_name := r.Form.Get("full_name")
    	em_arr := make(map[string]string)

     db := DbConnect() 
		defer db.Close()
    


    if edit_id == "" {
		
		
		var id int

   db.QueryRow("SELECT id from users where login_id='"+login_id+"'").Scan(&id)
   
    if id == 0 { 
    
    
		sql := "INSERT INTO users(login_id, full_name) VALUES ('"+login_id+"', '"+full_name+"')"
		
    db.Exec(sql)
  
    
    	em_arr["success"] = "New User Added"
    	
	} else
	{
		em_arr["error"] = "Login ID already exists"
		
	}
	
	
		
	} else {
		
		
		 
		 var id int

   db.QueryRow("SELECT id from users where login_id='"+login_id+"' AND id != "+edit_id).Scan(&id)
   
    if id == 0 { 
    
    
    
    	sql := "UPDATE users SET login_id = '"+login_id+"' ,full_name = '"+full_name+"' WHERE  id = "+edit_id
		
	db.Exec(sql)

  
    
    	em_arr["success"] = "User Updated"
    	
    	
     
		
	}    else
	{
		em_arr["error"] = "Login ID already exists"
		
	}
}
    
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(em_arr)


} 
