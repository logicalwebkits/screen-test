 

new Vue({
    el: '#tab',

    data: {
		users : [],
        login_id: "",
        full_name :"",
        submission :"Add",
        error :"",
        success :"",
        edit_id :""
    },
     created: function(){
    this.allusers();
  },
    methods: {
		allusers : function() {
			
			 axios.get('http://localhost:1010/allUsers').then(response => {
				 
					this.users = response.data
            });
            
			
		},
		deleteRecord : function(id) {
			
			 axios.get('http://localhost:1010/userDelete', {params: {
    id: id
  }} ).then(response => { 
				 
				   if (response.data.error) {
			  this.error = response.data.error
		  } else {
			   this.success = response.data.success 
		  }
							this.allusers();

            });
            
			
		},
		updateRecord : function(id) {
			
		
		  axios.get('http://localhost:1010/getUser', {params: {
    id: id
  }} ).then(response => { 
	 
			this.full_name = response.data.full_name
			this.login_id = response.data.login_id
			this.edit_id = id
			this.submission = "Edit"
			
            });
            
            
			
		},
		crawl : function() {
			
			this.error = ""
			this.success = ""
			if (this.login_id == "" ){
				
				this.error = "Login Id Required"
				return false
				
			}
			if (this.full_name == "") {
				
				this.error = "Full Name Required"
				return false;
				
			}
			
			const params = new URLSearchParams()
		params.append('login_id', this.login_id)
		params.append('full_name', this.full_name)
		params.append('edit_id', this.edit_id)
		


              axios.post('http://localhost:1010/addEditUser', params ,{ headers: {
      'Content-type': 'application/x-www-form-urlencoded',
      } } ).then(response => {
		  
		  if (response.data.error) {
			  this.error = response.data.error
		  } else {
			   this.success = response.data.success 
		  }
		  this.edit_id = ""
		  this.full_name = ""
		  this.login_id = ""
		  this.submission = "Add"
			this.allusers();
		
            });
            
            
            
            
		
 
               }
    }
});
