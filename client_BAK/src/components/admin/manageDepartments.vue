<template>

    <div class="manageDepartments">

        <!-- <h1> {{msg}} </h1> -->

        <section class="form">
            <input class="form-control" v-model="NewDepName" placeholder="Abteilung"/>
            <input class="form-control" v-model="NewDepMail" placeholder="Mail"/>
            Firma, der die Abteilung zugehörig ist:
            <select
                name="NewDepFirma"
                v-model="selected"                
                class="btn btn-secondary dropdown-toggle">
                <!-- <option :value="comp.id">{{comp.id}} </option> -->
                <option v-for="comp in this.$parent.Firmen" v-bind:key="comp.id" :value="comp.id">{{comp.name}}
                </option>
            </select>

            <!-- <select v-model="selected" id="deptList" class="btn btn-secondary
            dropdown-toggle"> <option v-for="comp in Firmen" v-bind:value="comp.name">
            {{comp.name}} </option> </select> -->

            <button id="adfds" type="button" class="btn btn-success" @click='AddDepartment'>Abteilung anlegen</button>

            <table class="table  table-striped table-dark">
                <thead class="thead-dark">
                    <tr>
                        <th scope="col">Abteilung</th>
                        <th scope="col">Firma</th>
                        <th scope="col">Mail</th>
                        <th scope="col">(de-)aktivieren</th>
                    </tr>
                </thead>
                <tbody >
                    <tr v-for="dep in this.$parent.Abteilungen" v-bind:key="dep.id">
                        <template v-if="dep.enabled === 1">
                            <td >{{dep.name}}</td>
                            <td>{{dep.firma}}</td>
                            <td>{{dep.notifyMail}}</td>
                            <td align="right">
                                <button
                                    type="button"
                                    class="btn-disable"
                                    @click='toggleDepartment(dep.id, dep.enabled)'>disable</button>
                            </td>
                        </template>
                        <template v-if="dep.enabled === 0">
                            <td class="disabledColumn">{{dep.name}}</td>
                            <td class="disabledColumn">{{dep.firma}}</td>
                            <td class="disabledColumn">{{dep.notifyMail}}</td>
                            <td class="disabledColumn" align="right">
                                <button
                                    type="button"
                                    class="btn-enable"
                                    @click='toggleDepartment(dep.id, dep.enabled)'>Enable</button>
                            </td>
                        </template>

                    </tr>
                </tbody>
            </table>
        </section>
    </div>

</template>

<script>
    import DBService from '../DBService';
    //import TodoItem from './TodoItem.vue';
    //import manageIndex from './manageIndex.vue'

    export default {
        name: 'ManageDepartments',
        //extends: manageIndex,
        components: {},
        data() {
            return {
                // Firmen: [],
                // Abteilungen: [],
                NewDepName: '',
                NewDepCompany: '',
                NewDepMail: '',
                selected: ""
            }
        },
        props: {
            msg: String
        },
         created() {
        // this.$parent.updateAll()        
        },
        methods: {
                      
           async AddDepartment() {

                // wenn keine Abteilung ausgewählt wurde - nimm default ALL
                if (!this.selected) 
                    this.selected = 0;
                const _NewDep = {
                    newDepName: this.NewDepName,
                    newDepCompany: this.selected,
                    newDepMail: this.NewDepMail
                };

                await DBService
                    .DBAddDepartment(_NewDep)
                    .then(() =>  this.$parent.updateAll())
                    .then(() => {
                        this.newDepName = ""
                        this.newDepCompany = ""
                        this.newDepMail = ""
                        this.selected = ""

                    })
                    .catch((error) => alert(("Server returned an Error:\n" + error.response.data)));

            },
            async toggleDepartment(id, enabled) {
                await DBService.toggleEntry("Abteilungen", id, enabled)
                .then(() =>  this.$parent.updateAll())            
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