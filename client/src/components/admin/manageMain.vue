<template>
    <div id="global page" style="background-color:rgb(40, 40, 40)">
        <!-- <img alt="Vue logo" src="./assets/logo.png"> -->


        <!-- Test neuer Navbar -->
        
    <!-- <div class="btn-group btn-group-toggle" data-toggle="buttons"> -->

        <div class="nav nav-tabs" >
        <label class="btn btn-secondary" id="btnComp" v-bind:class="{active: NavLblIsActive}" @click='MtoggleViewCompany'>
            Firmen bearbeiten
        </label>
        <label class="btn btn-secondary"  id="btnDep" @click='MtoggleViewDepartments'>
            Abteilungen bearbeiten
        </label>
        <label class="btn btn-secondary"  id="btnProc" @click='MtoggleViewProcessors'>
            Bearbeiter bearbeiten
        </label>
        <label class="btn btn-secondary" id="btnItem" @click='MtoggleViewItems'>
            Items bearbeiten
        </label>
    </div>
        
        <div v-show="toggleViewCompany">
            <ManageCompany msg="Firma bearbeiten"/>
        </div>

        <div v-show="toggleViewDepartments">
            <ManageDepartments msg="Abteilungen bearbeiten"/>
        </div>
        <div v-show="toggleViewProcessors">
            <ManageProcessors msg="Bearbeiter bearbeiten"/>
        </div>
        <div v-show="toggleViewItems">
            <ManageItems msg="Items bearbeiten"/>
        </div>
    </div>
</template>

<script>
    import ManageCompany from './manageCompany.vue'
    import ManageDepartments from './manageDepartments.vue'
    import ManageItems from './manageItems.vue'
    import ManageProcessors from './manageProcessors.vue'

    import DBService from '../DBService';

import axios from 'axios';

    export default {
        name: 'index',
        components: {
            ManageCompany,
            ManageDepartments,
            ManageItems,
            ManageProcessors
        },
        data() {
            return {
                toggleViewCompany: false,
                toggleViewDepartments: false,
                toggleViewProcessors: false,
                toggleViewItems: true,
                hover: false,

                Firmen: [],
                Abteilungen: [],
                Processors: [],
                Items: [],

                NavLblIsActive: false,
            }
        },
        methods: {
            async MtoggleViewCompany() {                
                       this.updateAll()
                this.toggleViewCompany = true;
                this.toggleViewDepartments = false;
                this.toggleViewItems = false;
                this.toggleViewProcessors = false;
                this.toggleButtons("btnComp");
            },
            async MtoggleViewDepartments() {
                this.updateAll()
                this.toggleViewCompany = false;
                this.toggleViewDepartments = true;
                this.toggleViewItems = false;
                this.toggleViewProcessors = false;
                this.toggleButtons("btnDep");
            },
            async MtoggleViewProcessors() {
                this.toggleViewCompany = false;
                this.toggleViewDepartments = false;
                this.toggleViewItems = false;
                this.toggleViewProcessors = true;
                this.toggleButtons("btnProc");
            },
            async MtoggleViewItems() {
                this.toggleViewCompany = false;
                this.toggleViewDepartments = false;
                this.toggleViewItems = true;
                this.toggleViewProcessors = false;                
                this.toggleButtons("btnItem");
            },

            async toggleButtons(_button){
                var btnComp = document.getElementById("btnComp");
                var btnDep = document.getElementById("btnDep");
                var btnProc = document.getElementById("btnProc");                
                var btnItem = document.getElementById("btnItem");
                
                btnComp.classList.remove("active");
                btnDep.classList.remove("active");
                btnProc.classList.remove("active");
                btnItem.classList.remove("active");

                var current = document.getElementById(_button);
                current.classList.add("active");
                
            },

            async updateAll() {
                console.log("updateAll() aus Main")
                await DBService
                    .DBGetCompanys()
                    .then((_firmen) => {
                        this.Firmen = _firmen
                    });
                
                await DBService
                    .DBGetDepartments()
                    .then((_departmens) => {
                        this.Abteilungen = _departmens
                    });

                await DBService
                    .DBGetProcessors()
                    .then((_processors) => {
                        this.Processors = _processors
                    });
                await DBService
                    .DBGetItems()
                    .then((_items) => {
                        this.Items = _items
                    });


                // const apiurlnew = window.location.protocol + "//"+ window.location.hostname +":5000/api/DBGetCompanys"
                // axios .get(apiurlnew) 
                //     .then((response) =>  {
                //         console.log(response.data)
                //         this.Firmen = response.data.Firmen
                //     })  

                // await DBService
                //     .DBgetCompanys()
                //     .then((_firmen) => {
                //         this.Firmen = _firmen
                //     });

                ///////////////////////////////////

                // await DBService
                //     .DBget("Abteilungen", "all")
                //     .then((_abteilungen) => {
                //         this.Abteilungen = _abteilungen
                //     });

                // await DBService
                //     .DBget("Processors", "all")
                //     .then((_processors) => {
                //         this.Processors = _processors
                //     });
                // await DBService
                //     .DBget("Items", "all")
                //     .then((_items) => {
                //         this.Items = _items

                //     });

                
            }
        }
    }
</script>

<style >
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
    td.disabledColumn {
        color: #2e2e2e;
        background-color: #696868;
    }
    div.hoverable:hover {
        color: #12406e;
        
    }
    
    input. {
        background-color: rgb(116, 156, 243);
    }
</style>