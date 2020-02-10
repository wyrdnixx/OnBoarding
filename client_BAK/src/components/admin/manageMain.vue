<template>
    <div id="global page" style="background-color:rgb(179, 233, 255)">
        <!-- <img alt="Vue logo" src="./assets/logo.png"> -->

        <!-- <div class="shadow-lg p-3 mb-5 bg-info rounded"
        @click='MtoggleViewCompany'><h4>Firmen bearbeiten</h4></div> -->
        <div class="shadow-lg p-3 mb5 hoverable" @click='MtoggleViewCompany'>
            <h4>Firmen bearbeiten</h4>
        </div>
        <div v-show="toggleViewCompany">
            <ManageCompany msg="Firma bearbeiten"/>
        </div>

        <div
            class="shadow-lg p-3 mb5  rounded hoverable"
            @click='MtoggleViewDepartments'>
            <h4>Abteilungen bearbeiten</h4>
        </div>
        <div v-show="toggleViewDepartments">
            <ManageDepartments msg="Abteilungen bearbeiten"/>
        </div>

        <div
            class="shadow-lg p-3 mb5  rounded hoverable"
            @click='MtoggleViewProcessors'>
            <h4>Bearbeiter bearbeiten</h4>
        </div>
        <div v-show="toggleViewProcessors">
            <ManageProcessors msg="Bearbeiter bearbeiten"/>
        </div>

        <div
            class="shadow-lg p-3 mb5  rounded hoverable"
            @click='MtoggleViewItems'>
            <h4>Items bearbeiten</h4>
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
                Items: []
            }
        },
        methods: {
            async MtoggleViewCompany() {
                       this.updateAll()
                this.toggleViewCompany = !this.toggleViewCompany;
            },
            async MtoggleViewDepartments() {
                this.updateAll()
                this.toggleViewDepartments = !this.toggleViewDepartments;
            },
            async MtoggleViewItems() {
                this.toggleViewItems = !this.toggleViewItems;
            },
            async MtoggleViewProcessors() {
                this.toggleViewProcessors = !this.toggleViewProcessors;
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
        color: #2e2e2e;
        background-color: #696868;
    }
    div.hoverable:hover {
        color: #12406e;
        background-color: rgb(116, 156, 243);
    }
</style>