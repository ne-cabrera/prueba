<template>
    <div class="container">
        <div class="row itemHistory" v-for="(host, index) in hosts.items" :key="index">
            <history-card :host="host"/>
        </div>
    </div>
</template>

<script>
import HistoryCard from "./HistoryCard.vue"
export default {
    name: "history",
    components:{
        HistoryCard
    },
    data(){
        return{
            hosts:{}
        }
    },
    mounted(){
        this.getHosts()
    },
    methods:{
        async getHosts(){
            try {
                const response = await fetch('http://localhost:8081/history',{
                    method: 'GET',
                    headers: { "Content-Type": "application/json"}
                })
                const data = await response.json()
                this.hosts = data
            } catch (error) {
                // eslint-disable-next-line
                console.log(error)
            }
        }
    }
}
</script>
<style scoped>
.itemHistory{
    margin-top: 10px;
}
</style>