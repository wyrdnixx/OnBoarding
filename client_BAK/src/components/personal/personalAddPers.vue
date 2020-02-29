<template>

    <div>

        <div>
            <span class="border border-success">
                <h3>Personal manuell anlegen - Abteilungsauswahl muss noch auf ausgewählte Firma eingeschr. werden.</h3>
                <table class="table">

                    <tbody>
                        <tr>
                            <td>Firma:
                            </td>
                            <td>
                                <select
                                    name="perComp"
                                    v-model="perComp"
                                    class="btn btn-secondary dropdown-toggle">
                                    <!-- <option >Männlich</option>  <option >Weiblich</option> -->
                                    <option v-for="comp in this.Firmen" v-bind:key="comp.id" :value="comp.id">{{comp.name}}
                                    </option>
                                </select>
                            </td>
                            <td>Personalnummer:
                            </td>
                            <td>
                                <input type="text" id="perNr" v-model="perNr"></td>
                                <td>Abteilung:
                                </td>
                                <td>
                                    <select
                                        name="perAbt"
                                        v-model="perAbt"
                                        class="btn btn-secondary dropdown-toggle"
                                         @click="this.filterDepartments()" >
                                        <option v-for="dep in this.Abteilungen" v-bind:key="dep.id" :value="dep.id">{{dep.name}}
                                    </option>
                                    </select>
                                </td>
                                <td>Geschlecht:</td>
                                <td>
                                    <select
                                        name="perSex"
                                        v-model="perSex"
                                        class="btn btn-secondary dropdown-toggle">
                                        <option >Männlich</option>
                                        <option >Weiblich</option>
                                    </select>
                                </td>
                            </tr>
                            <tr>
                                <td>Titel:
                                </td>
                                <td>
                                    <input type="text" id="perTitle" v-model="perTitle"></td>
                                    <td>Nachname:
                                    </td>
                                    <td>
                                        <input type="text" id="perNachname" v-model="perNachname"></td>
                                        <td>Vorname:</td>
                                        <td>
                                            <input type="text" id="perVorname" v-model="perVorname"></td>
                                            <td>Geburtsdatum:</td>
                                            <td>
                                                <input type="text" id="perGebDat" v-model="perGebDat"></td>
                                            </tr>
                                            <tr>
                                                <td>Eintritt:
                                                </td>
                                                <td>
                                                    <input type="text" id="perEintritt" v-model="perEintritt"></td>
                                                    <td>Austritt:
                                                    </td>
                                                    <td>
                                                        <input type="text" id="perAustritt" v-model="perAustritt"></td>
                                                        <td>Dienstart:
                                                        </td>
                                                        <td>
                                                            <input type="text" id="perDienstart" v-model="perDienstart"></td>
                                                            <td>Berufsbezeichnung:
                                                            </td>
                                                            <td>
                                                                <input type="text" id="perBeruf" v-model="perBeruf"></td>
                                                            </tr>
                                                        </tbody>
                                                    </table>
                                                    <button type="button" class="btn btn-success" @click='AddPersonal'>Personal anlegen</button>
                                                </span>
                                            </div>
                                            <br>
                                                <div>
                                                    <h3>-------- Personal Import ---------</h3>
                                                </div>

                                            </div>

                                        </template>

                                        <script>
                                        import DBService from '../DBService';

                                            export default {
                                                name: 'personalMain',
                                                components: {},
                                                data() {
                                                    return {
                                                        Firmen: [],
                                                        Abteilungen: [],                
                                                        perComp: "",
                                                        perNr: "",
                                                        perAbt: "",
                                                        perSex: "",
                                                        perTitle: "",
                                                        perNachname: "",
                                                        perVorname: "",
                                                        perGebDat: "",
                                                        perEintritt: "",
                                                        perAustritt: "",
                                                        perDienstart: "",
                                                        perBeruf: ""
                                                    }
                                                },
                                                 created() {
                                                    this.updateAll();                                                    
                                                },
                                                methods: {
                                                    async updateAll() {                                                        
                                                        await DBService
                                                            .DBget("Firmen", "enabled")
                                                            .then((_firmen) => {
                                                                this.Firmen = _firmen
                                                            });

                                                        await DBService
                                                            .DBget("Abteilungen", "enabled")
                                                            .then((_abteilungen) => {
                                                                this.Abteilungen = _abteilungen
                                                            });
                                                    },

                                                    AddPersonal() {
                                                        DBService
                                                        .AddPersonal(                                                                                                                    
                                                        this.perComp,
                                                        this.perNr,
                                                        this.perAbt,
                                                        this.perSex,
                                                        this.perTitle,
                                                        this.perNachname,
                                                        this.perVorname,
                                                        this.perGebDat,
                                                        this.perEintritt,
                                                        this.perAustritt,
                                                        this.perDienstart,
                                                        this.perBeruf
                                                        )
                                                        .then(() => {
                                                        this.perComp = "",
                                                        this.perNr = "",
                                                        this.perAbt = "",
                                                        this.perSex = "",
                                                        this.perTitle = "",
                                                        this.perNachname = "",
                                                        this.perVorname = "",
                                                        this.perGebDat = "",
                                                        this.perEintritt = "",
                                                        this.perAustritt = "",
                                                        this.perDienstart = "",
                                                        this.perBeruf = ""
                                                        })
                                                        .then(this.updateAll())
                                                        .catch((error) => alert(error));
                                                    },
                                                    filterDepartments() {
                                                        perComp

                                                    }
                                                }
                                            }
                                        </script>