<template>
    <div id="app" style="background-color:rgb(40, 40, 40); padding-left: 3%; padding-right: 3%;" >
        <p>
            <!-- <router-link to="/foo">Foo</router-link>
            <router-link to="/bar">Bar</router-link> -->
            <br>
            <router-link type="button" class="btn btn-success" to="/Personal">Personal </router-link>
            <router-link type="button" class="btn btn-success" to="/Manage">Manage </router-link>
            <router-link  type="button" class="btn btn-warning" to="/Test"> ## Test ## </router-link>
        </p>
        <router-view></router-view>
    </div>
</template>

<script>
export default {
    // empty
}
</script>

<style>
 

  td.disabledColumn {
    color:#757575;
    background-color: #353535;
  }
  td {
    color:#7881fa;
    background-color: #3f3e3e;
  }
  
  div.hoverable:hover{
    color: #12406e;
    background-color: rgb(56, 115, 243);
  }


    #app {
        font-family: 'Avenir', Helvetica, Arial, sans-serif;
        -webkit-font-smoothing: antialiased;
        -moz-osx-font-smoothing: grayscale;
        text-align: left;
        color: #3793ee;        
        margin-top: 1%;        
        background-color: rgb(99, 145, 247);
    }

    div {
        margin-top: 1%;
    }
  
  
    input {
        background-color: rgb(116, 156, 243);
    }

</style>
