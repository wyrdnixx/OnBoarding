<template>
<div id="global page" style ="background-color:rgb(179, 233, 255)">
    <!-- <img alt="Vue logo" src="./assets/logo.png"> -->

        Hello
        <subTest />
</div>
</template>

<script>
import subTest from './subTest.vue'


export default {
  name: 'index',    
  components: {
    subTest,
        
  },
  data(){
    return {
        someData: "in parents data f"  
    }
  },
  methods: { 
    
    },
  }

</script>