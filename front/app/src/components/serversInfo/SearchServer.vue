<template>
    <div id="search-server" class="container">
        <form @submit.prevent="handleSubmit">
        <div class="form-group row">
            <div class="col-sm-8 textBar">
            <input type="text" v-model="host.hostName" class="form-control" aria-describedby="hostNameHelp" placeholder="Enter host name">
            </div>
            <div class="col-sm-4 boton">
                <button type="submit" class="btn btn-primary">Search</button>
            </div>
        </div>
        </form>
    </div>
</template>
<script>
export default {
    name: "search-server",
    data(){
        return{
            host:{
                hostName: "",
            },
        }
    },
    methods: {
        handleSubmit(){
            this.$emit('add:host', this.host)
        },
    },
}
</script>

<style>
.boton{
    text-align: left;
    padding-left: 0px;
    margin-left: 0px;
}

.textBar{
    padding-right: 5px;
}
</style>
