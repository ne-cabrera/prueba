<template>
    <nav class="navbar navbar-expand-lg navbar navbar-dark bg-dark">
        <div class="navbar-brand">
            <router-link to="/">Search</router-link>
        </div>
        <div class="collapse navbar-collapse" id="navbarText">
            <ul class="navbar-nav mr-auto">
      <li class="nav-item active">
        <div class="nav-link">
            <router-link to="/history">History</router-link>
        <span class="sr-only">(current)</span></div>
      </li>
            </ul>
        </div>
    </nav>
</template>

<script>
export default {
    name: "nav-bar"
}
</script>

<style scoped>
a{
    color: white;
}
a:hover{
    color: white;
}
</style>
