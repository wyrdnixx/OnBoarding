<template>
<div id="global page" style ="background-color:rgb(179, 233, 255)">
  <div id="app">
    <!-- <img alt="Vue logo" src="./assets/logo.png"> -->

             
         <!--     <div class="shadow-lg p-3 mb-5 bg-blue rounded" @click='MtoggleViewCompany'><h4>Firmen bearbeiten</h4></div> -->
              <div class="shadow-lg p-3 mb-5 rounded hoverable" @click='MtoggleViewCompany'><h4>Firmen bearbeiten</h4></div> 
              <div v-show="toggleViewCompany" >
                <ManageCompany msg="Firma bearbeiten" />                 
              </div>
           
              <div class="shadow-lg p-3 mb-5 bg-blue rounded hoverable" @click='MtoggleViewDepartments'><h4>Abteilungen bearbeiten</h4></div>
              <div v-show="toggleViewDepartments" >
                <ManageDepartments msg="Abteilungen bearbeiten" /> 
              </div>

              <div class="shadow-lg p-3 mb-5 bg-blue rounded hoverable" @click='MtoggleViewProcessors'><h4>Bearbeiter bearbeiten</h4></div>
              <div v-show="toggleViewProcessors" >
                <ManageProcessors msg="Bearbeiter bearbeiten" /> 
              </div>

              <div class="shadow-lg p-3 mb-5 bg-blue rounded hoverable" @click='MtoggleViewItems'><h4>Items bearbeiten</h4></div>
              <div v-show="toggleViewItems" >
                <ManageItems msg="Items bearbeiten" /> 
              </div>
    
  </div>
</div>
</template>

<script>
import ManageCompany from './components/manageCompany.vue'
import ManageDepartments from './components/manageDepartments.vue'
import ManageItems from './components/manageItems.vue'
import ManageProcessors from './components/manageProcessors.vue'
//import ListUsers from './components/ListUsers.vue'

export default {
  name: 'app',  
  components: {
    ManageCompany,
    ManageDepartments,
    ManageItems,
    ManageProcessors
    
  },
  data(){
    return {
       toggleViewCompany: ``,
       toggleViewDepartments: ``,
       toggleViewProcessors: true,
       toggleViewItems: true,
             hover: false,
    }
  },
  methods: { 
    async MtoggleViewCompany() {
    this.toggleViewCompany = !this.toggleViewCompany;
    },
    async MtoggleViewDepartments() {
    this.toggleViewDepartments = !this.toggleViewDepartments;
    },
    async MtoggleViewItems() {
    this.toggleViewItems = !this.toggleViewItems;
    },
    async MtoggleViewProcessors() {
    this.toggleViewProcessors = !this.toggleViewProcessors;
    },
  }
} 
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #153e66;
  margin-top: 60px;
  background-color: rgb(99, 145, 247);
}

  td.disabledColumn {
    color:#2e2e2e;
    background-color: #696868;
  }
  div.hoverable:hover{
    color: #12406e;
    background-color: rgb(56, 115, 243);
  }

</style>
