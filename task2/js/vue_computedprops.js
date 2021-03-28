 

new Vue({
    el: '#tab',

    data: {
        posts: [],
        s :"",
        r :"",
        c :""
    },
    methods: {
		crawl : function(event) {
			this.posts = [] 
         
			if (this.s !="" && this.c !="" && this.r !="" ){
         
              axios.get('http://localhost:1010/excelMaker', {params: {
    s: this.s,
    c: this.c,
    r: this.r
  }} ).then(response => {
	  console.log(response)
			this.posts = response.data
            });
		}
                 
               }
    }
});
