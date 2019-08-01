<template>
    <div id="serversInfo">
        <search-server @add:host="searchHost"/>
        <host-info v-if="seen" :server="hostInfo"/>
        <servers v-if="seen" :servers="hostInfo.servers"/>
    </div>
</template>
<script>
import SearchServer from "./SearchServer.vue"
import Servers from "./Servers.vue"
import HostInfo from "./HostInfo.vue"
export default {
    name: 'servers-info',
    components:{
        SearchServer,
        Servers,
        HostInfo
    },
    data(){
        return{
            hostInfo: {},
            seen: false
        }
    },
    methods:{
        async searchHost(host){
            // eslint-disable-next-line
            console.log(JSON.stringify(host))
            try {
                const response = await fetch('http://localhost:8081/servers',{
                    method: 'POST',
                    body: JSON.stringify(host),
                    headers: { "Content-Type": "application/json"}
                })
                const data = await response.json()
                this.hostInfo = data
                this.seen = true
            } catch (error) {
                // eslint-disable-next-line
                console.log(error)
            }
        }
    }
}
</script>
<style scoped>
#serversInfo{
    margin-top: 5%
}
</style>