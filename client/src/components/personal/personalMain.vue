<template>
    <div>
        <PersonalAddPers/>
        <PersonalShowPers/>
    </div>

</template>

<script>

    import PersonalAddPers from './personalAddPers.vue'
    import PersonalShowPers from './personalShowPers.vue'
import DBService from '../DBService';

    export default {
        name: 'personalMain',
        components: {
            PersonalAddPers,
            PersonalShowPers
        },
        data() {
            return {Firmen: [], Abteilungen: [], Processors: [], Items: []}
        },
        methods: {
            async updateAll() {
                //alert("in update all")
                console.log("updateAll() aus personalMain")
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

            }
        }
    }
</script>