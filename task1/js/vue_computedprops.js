 

new Vue({
    el: '#tab',

    data: {
        posts: [],
        url :""
    },
    methods: {
		crawl : function(event) {
			this.posts = []
         
			if (this.url !="" ){
         
              axios.get('http://localhost:1010/wordCounter', {params: {
    url: this.url
  }} ).then(response => {
	  console.log(response)
			this.posts = response.data
            });
		}
                 
               }
    }
});
