<template>
    <div id="app">
        <p>
            <!-- <router-link to="/foo">Foo</router-link>
            <router-link to="/bar">Bar</router-link> -->
            <br>
            <router-link type="button" class="btn btn-success" to="/Personal">Personal </router-link>
            <router-link type="button" class="btn btn-success" to="/Manage">Manage </router-link>
            <router-link  type="button" class="btn btn-warning" to="/Test"> ## Test ## </router-link>
        </p>
        <router-view></router-view>
    </div>
</template>

<script>
export default {
    // empty
}
</script>

<style>
 

  td.disabledColumn {
    color:#2e2e2e;
    background-color: #696868;
  }
  div.hoverable:hover{
    color: #12406e;
    background-color: rgb(56, 115, 243);
  }

</style>
