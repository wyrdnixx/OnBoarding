<template>

    <div >

        <!-- <h1> {{msg}} </h1> -->
        <section class="form">
            <table class="table  table-striped table-dark">
                <thead class="thead-dark">
                                                </thead>
                    <tr>
                        <td>
                                                <input
                        class="form-control"
                        v-model="newProcessorName"
                        placeholder="Bearbeitende Abteilung"/>
                        </td>
                    <td>
                    <input
                        class="form-control"
                        v-model="newProcessorMail"
                        placeholder="Notify Email"/>
                    </td>
                    </tr>
                    <button type="button" class="btn btn-success" @click='AddProcessor'>Anlegen</button>
                    
            </table>
                    <hr/>
                    <div >
                        <table class="table  table-striped table-dark">
                            <thead class="thead-dark">
                                <tr>
                                    <th scope="col">Bearbeiter</th>

                                    <th scope="col">Mail</th>
                                    <th scope="col">(de-)aktivieren</th>
                                </tr>
                            </thead>
                            <tbody >
                                <tr v-for="processor in this.$parent.Processors" v-bind:key="processor.id">
                                    <template v-if="processor.enabled === 1">
                                        <td>{{processor.name}}</td>
                                        <td>{{processor.notifyMail}}</td>
                                        <td align="right">
                                            <button
                                                type="button"
                                                class="btn-disable"
                                                @click='toggleProcessors(processor.id, processor.enabled)'>disable</button>
                                        </td>
                                    </template>
                                    <template v-if="processor.enabled === 0">
                                        <td class="disabledColumn">{{processor.name}}</td>
                                        <td class="disabledColumn">{{processor.notifyMail}}</td>
                                        <td class="disabledColumn" align="right">
                                            <button
                                                type="button"
                                                class="btn-enable"
                                                @click='toggleProcessors(processor.id, processor.enabled)'>Enable</button>
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
            //import TodoItem from './TodoItem.vue';
            export default {
                name: 'ManageProcessors',
                components: {},
                data() {
                    return {newProcessorName: "", newProcessorMail: ""}
                },
                props: {
                    msg: String
                },
                created() {
                    //this.updateAllData()
                    this
                        .$parent
                        .updateAll();
                },
                methods: {

                    async AddProcessor() {
                        const _NewProc = {
                            NewProcessorName: this.newProcessorName,
                            NewProcessorMail: this.newProcessorMail
                        };

                        //alert ("Name: " + _NewProc.NewProcessorName)
                        await DBService
                            .DBAddProcessor(_NewProc)
                            .then(() => this.$parent.updateAll())
                            .then(() => {
                                this.newProcessorMail = ""
                                this.newProcessorName = ""
                            })
                            .catch((error) => alert(("Server returned an Error:\n" + error.response.data)));
                    },

                    async toggleProcessors(id, enabled) {
                        await DBService
                            .toggleEntry("Processors", id, enabled)
                            .then(() => this.$parent.updateAll())
                            .catch((error) => alert(("Server returned an Error:\n" + error.response.data)));
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