 

new Vue({
    el: '#tab',

    data: {
        number :"",
        prime :""
    },
    methods: {
		crawl : function(event) {
			this.prime = ""
         
			if (this.number !="" ){
         
              axios.get('http://localhost:1010/prime', {params: {
    number: this.number
  }} ).then(response => {
			this.prime = "Prime Numbers : " + response.data
            });
		}
                 
               }
    }
});
