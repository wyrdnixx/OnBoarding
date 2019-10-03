<template>

    <div class="entryRow form-group">
        <h3>------ Status - Personal --- ----</h3>
        <table  border="1" class="table  table-striped table-dark">
            <thead class="table table-dark">
                <tr>
                    <th>Firma</th>
                    <th>Pers.Nr.</th>
                    <th>Nachname</th>
                    <th>Vorname</th>
                    <th>Titel</th>
                    <th>Geschlecht</th>
                    <th>Eintritt</th>
                    <th>Austritt</th>
                    <th>Abteilung</th>
                    <th>Geburtsdatum</th>
                    <th>Dienstart</th>
                    <th>Berufsbezeichnung</th>
                    <th>Bearbeitungsstatus</th>
                </tr>
            </thead>
            <tbody>
                 <tr v-for="per in this.Personal" v-bind:key="per.id">
                    <td>{{per.firma}}</td>
                    <td>{{per.persnr}}</td>
                    <td>{{per.nachname}}</td>
                    <td>{{per.vorname}}</td>
                    <td>{{per.titel}}</td>
                    <td>{{per.geschlecht}}</td>
                    <td>{{per.eintritt}}</td>
                    <td>{{per.austritt}}</td>                    
                    <td v-if="per.abteilungStatus == 1">{{per.abteilung}} - <span class="badge badge-success">OK</span></td>
                    <td v-else>{{per.abteilung}} - <span class="badge badge-warning">Offen</span></td>
                    <td>{{per.gebdat}}</td>
                    <td>{{per.dienstart}}</td>
                    <td>{{per.berufsbezeichnung}}</td>                    
                    <td v-if="per.processorStatus == 1"><span class="badge badge-success">OK</span></td>
                    <td v-else><span class="badge badge-warning">Offen</span></td>
                 </tr>
            </tbody>
        </table>

    </div>

</template>
<script>

import DBService from '../DBService';

export default {
    name: 'personalShowPers',
    components: {    
    },
    data() {
        return {
            Personal: [],
        }
    },
    created () {
        this.updateAll()
    },
    methods: {
       async updateAll() {
            await DBService
            // 'Range' wird bei Personal-Get nicht benötigt.
            // wurde aber der Einfachheit halber gelassen, da sonnst in der App ein 
            // extra DBGet gemacht werden müsste
            .DBget("Personal","All")
            .then((_personal) => {
                this.Personal = _personal
            })
            .catch((error) => alert(error));
        }
    }
}
</script>


<style>

</style>

