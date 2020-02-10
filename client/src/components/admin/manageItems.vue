<template>

    <div class="ManageItems">

        <!-- <h1> {{msg}} </h1> -->
        <section class="form">
            <input class="form-control" v-model="newItemText" placeholder="Item-Text"/>
            Abteilung, für die das Item verfügbar ist ( ALL = Alle Abteilungen):
            <select
                name="NewItemDepartment"
                v-model="selectedDepartment"                 
                 @click="this.$parent.updateAll()"              
                class="btn btn-secondary dropdown-toggle">
                <!-- <option :value="comp.id">{{comp.id}} </option> -->                
                <option v-for="dep in this.$parent.Abteilungen" v-bind:key="dep.id" :value="dep.id" v-if="dep.enabled===1">{{dep.name}}
                </option>
            </select>
            <br>
                Bearbeitende Abteilung:
                <select
                    name="NewItemProcessor"
                    v-model="selectedProcessor"                       
                     @click="this.$parent.updateAll()"                 
                    class="btn btn-secondary dropdown-toggle">
                    <!-- <option :value="comp.id">{{comp.id}} </option> -->
                    <option v-for="proc in this.$parent.Processors" v-bind:key="proc.id" :value="proc.id" v-if="proc.enabled===1">{{proc.name}}
                    </option>
                </select>
                <br>
                    Typ des Items:
                    <select
                        name="NewItemType"
                        v-model="selectedType"
                        class="btn btn-secondary dropdown-toggle">
                        <!-- <option :value="comp.id">{{comp.id}} </option> -->
                        <option >CheckBox</option>
                        <option >Choose</option>
                        <option >Text</option>
                    </select>
                    <input
                        v-if="selectedType === 'Choose'"
                        class="form-control"
                        v-model="newItemChooseOptions"
                        placeholder="Optionen, die zur Auswahl stehen. ( Eins;Zwei;Drei; )"/>

                    <br>
                        <button type="button" class="btn btn-success" @click='AddItem'>Item anlegen</button>
                        <hr/>
                        <div >
                            <table class="table  table-striped table-dark">
                                <thead class="thead-dark">
                                    <tr>
                                        <th scope="col">Item</th>
                                        <th scope="col">Abteilung</th>
                                        <th scope="col">Bearbeiter</th>
                                        <th scope="col">Type/Value</th>
                                        <th scope="col">(de-)aktivieren</th>
                                    </tr>
                                </thead>
                                <tbody >
                                    <tr v-for="item in this.$parent.Items" v-bind:key="item.id">
                                        <template v-if="item.enabled === 1">
                                            <td>{{item.Text}}</td>
                                            <td>{{item.DepId}}</td>
                                            <td>{{item.ProcessorId}}</td>
                                            <td>{{item.ItemType}}</td>
                                            <td align="right">
                                                <button
                                                    type="button"
                                                    class="btn-disable"
                                                    @click='toggleItem(item.id, item.enabled)'>disable</button>
                                            </td>
                                        </template>
                                        <template v-if="item.enabled === 0">
                                            <td class="disabledColumn">{{item.Text}}</td>
                                            <td class="disabledColumn">{{item.DepId}}</td>
                                            <td class="disabledColumn">{{item.ProcessorId}}</td>
                                            <td class="disabledColumn">{{item.ItemType}}</td>
                                            <td class="disabledColumn" align="right">
                                                <button
                                                    type="button"
                                                    class="btn-enable"
                                                    @click='toggleItem(item.id, item.enabled)'>Enable</button>
                                            </td>
                                        </template>
                                    </tr>
                                </tbody>
                            </table>
                        </div>
                    </section>

                </div>
            </template>

            <script>
                import DBService from '../DBService';
                // import {__vlsIterationHelper} from '../../../vue-temp/vue-editor-bridge';
                // import TodoItem from './TodoItem.vue';
                export default {
                    name: 'ManageItems',
                    components: {},
                    data() {
                        return {
//                            Items: [],  
                            newItemText: "",
                         //   Abteilungen: [],
                          //  Processors: [],
                            selectedDepartment: "",
                            selectedProcessor: "",
                            selectedType: "",
                            newItemChooseOptions: ""
                        }
                    },
                    props: {
                        msg: String
                    },
                    created() {
                        
                    },
                    methods: {
                       async AddItem() {
                            if (!this.selectedDepartment) 
                                this.selectedDepartment = 1;
                            
                            if (this.selectedType == "Choose") {
                                this.selectedType = "Choose:" + this.newItemChooseOptions;
                            }

                            if (this.selectedType == "" || this.selectedProcessor == "") {
                                alert("Error: no type or Processor choosen")
                            } else {
                                

                            const _NewItem = {
                                newItemText: this.newItemText, 
                                selectedDepartment: this.selectedDepartment,
                                selectedProcessor: this.selectedProcessor,
                                selectedType: this.selectedType,
                                newItemChooseOptions: this.newItemChooseOptions
                            };

                            await DBService
                                .DBAddItem(_NewItem)
                                .then (() => this.$parent.updateAll())
                                .then (() => {
                                    this.newItemText = "";
                                    this.selectedDepartment = "";
                                    this.selectedProcessor = "";
                                    this.selectedType = "";
                                    this.newItemChooseOptions = "";
                                })
                                .catch((error) => alert(("Server returned an Error:\n" + error.response.data)));                                
                            }
                        },

                        async toggleItem(id, enabled) {
                            await DBService.toggleEntry("Items", id, enabled);
                            //this.updateAllData()
                            this.$parent.updateAll();
                        }
                    }
                }
                
            </script>

            <!-- Add "scoped" attribute to limit CSS to this component only -->
            <style scoped="scoped">


                div.container {
                    max-width: 90%;

                }

                p.error {
                    border: 1px solid #ff5b5f;
                    background-color: #ffc5c1;
                    padding: 10px;
                    margin-bottom: 15px;
                }

                div.entryRow {
                    position: relative;
                    border: 1px solid #5bd658;
                    background-color: #bcffb8;
                    padding: 10px 10px 1px;
                    margin-bottom: 0;
                }

                div.created-at {
                    position: absolute;
                    top: 0;
                    left: 0;
                    padding: 5px 15px;
                    background-color: darkgreen;
                    color: white;
                    font-size: 13px;
                }

                p.text {
                    font-size: 22px;
                    font-weight: 700;
                    margin-bottom: 0;
                }

                p.smalltext {
                    font-size: 14px;
                    font-weight: 500;
                    margin-bottom: 0;
                }
                /* table {
                                      width: 100%;
                                      text-align: left;
                                      border: 1px solid black;
                                      border-collapse: collapse;
                                           
                                    }
                                    th {
                                      padding: 5px;
                                      font-weight: 700;
                                      border: 1px solid black;
                                      vertical-align: middle;
                                    } */
                td {
                    padding: 4px;
                    font-size: 14px;
                    font-weight: 500;
                    margin-bottom: 0;
                    border: 1px solid black;
                    vertical-align: middle;
                    align: left;
                }
                .btn-disable {
                    background-color: rgb(53, 53, 53);
                    color: rgb(236, 48, 48);
                    border: 2px solid rgb(179, 7, 7);
                    /* Green */
                }
                .btn-enable {
                    background-color: rgb(53, 53, 53);
                    color: rgb(124, 223, 148);
                    border: 2px solid rgb(0, 145, 43);
                    /* Green */
                }
            </style>