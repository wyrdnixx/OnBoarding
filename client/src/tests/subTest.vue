
<template>
<div id="global page" style ="background-color:rgb(179, 233, 255)">
    <!-- <img alt="Vue logo" src="./assets/logo.png"> -->

  test: {{this.$parent.someData}}
</div>
</template> 

<script>


export default  {
  name: 'subTest',      
  components: {        
  },
  
   // kebab case here available as camelcase in template
data(){
    return {
    //  parentData: this.$parent.someData
    }
  },
  methods: { 
    
    },
  }

</script>
