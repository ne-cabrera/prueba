<template>
  <div id="app" class="container-fluid">
    <nav-bar/>
    <router-view></router-view>
  </div>
</template>

<script>
import NavBar from "./components/Nav/NavBar.vue"
export default {
  name: 'app',
  components:{
    NavBar
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 0px;
  width: 100%;
  margin: 0px;
  padding: 0px;
}
</style>
