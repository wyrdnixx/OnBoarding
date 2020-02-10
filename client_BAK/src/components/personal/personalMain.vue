<template>
    <div>
        <PersonalAddPers/>
        <PersonalShowPers/>
    </div>

</template>

<script>

import PersonalAddPers from './personalAddPers.vue'
import PersonalShowPers from './personalShowPers.vue'
export default {
    name: 'personalMain',
    components: {
    PersonalAddPers,
    PersonalShowPers
    }
}
</script>

