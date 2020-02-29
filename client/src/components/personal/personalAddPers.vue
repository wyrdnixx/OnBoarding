<template>

    <div>

        <div>
            <span class="border border-success">
                <h3>Personal manuell anlegen - Abteilungsauswahl muss noch auf ausgewählte Firma
                    eingeschr. werden.</h3>

                <!-- <li v-for="firma in this.$parent.Firmen" v-bind:key="firma.id">Test:
                {{firma.name}}</li> -->
                <li v-for="abt in this.AvailableAbteilungen" v-bind:key="abt.id">Test:
                    {{abt.name}}</li>

                <table class="table">

                    <tbody>
                        <tr>
                            <td>Firma:
                            </td>
                            <td>
                                <select
                                    name="perComp"
                                    v-model="perComp"
                                    class="btn btn-secondary dropdown-toggle"
                                    @click="filterDepartments()">
                                    <option v-for="firma in this.$parent.Firmen" v-bind:key="firma.id">{{firma.name}}
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
                                    <!-- <select name="perAbt" v-model="perAbt" class="btn btn-secondary
                                    dropdown-toggle"> <option v-for="dep in this.AvailableAbteilungen"
                                    v-bind:key="dep.id" :value="dep.id">{{dep.name}} </option> </select> -->
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
                                                    <!-- <button type="button" class="btn btn-success" @click='AddPersonal'>Personal
                                                    anlegen</button> -->
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
                                                        perBeruf: "",
                                                        AvailableAbteilungen: []
                                                    }
                                                },
                                                async created() {
                                                    console.log(JSON.stringify(this.$parent.Abteilungen), null, '    ');
                                                    this
                                                        .$parent
                                                        .updateAll()
                                                        .catch((error) => alert(("Server returned an Error:\n" + error.response.data)))

                                                    },
                                                methods: {

                                                    ClearFields() {
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
                                                    },
                                                    filterDepartments() {
                                                        //his.AvailableAbteilungen = this.$parent.Abteilungen

                                                        this
                                                            .$parent
                                                            .updateAll()
                                                            .catch((error) => alert(("Server returned an Error:\n" + error.response.data)))

                                                            this.AvailableAbteilungen =[];
                                                            
                                                            // Info - Besser mal mit Lodash probieren
                                                            this
                                                            .$parent
                                                            .Abteilungen
                                                            .forEach((abt, index) => {
                                                                console.log(JSON.stringify(abt))
                                                                if (abt.enabled == 1) {
                                                                    this
                                                                        .AvailableAbteilungen
                                                                        .push(abt)
                                                                }

                                                            });

                                                        //console.log(this.AvailableAbteilungen)
                                                        /*
                                                        this.$parent.Abteilungen.forEach((abteilung, index) => {
                                                            if (abteilung.enabled === 0) {
                                                                this.$parent.Abteilungen.splice(index, 1);
                                                            }
                                                        });
*/

                                                        console.log("AvailableAbteilungen: " + JSON.stringify(this.AvailableAbteilungen))

                                                }
                                            }
                                        }
                                        </script>