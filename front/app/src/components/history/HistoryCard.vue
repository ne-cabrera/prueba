<template>
        <div class="card">
        <div class="card-header">
            {{host.hostName}}
        </div>
    <div class="card-body">
        <host-info :server="host"/>
        <servers :servers="host.servers"/>
    </div>
    </div>
    
</template>

<script>
import HostInfo from "../serversInfo/HostInfo.vue";
import Servers from "../serversInfo/Servers.vue";
export default {
    name: "history-card",
    components:{
        HostInfo,
        Servers
    },
    props:{
        host: {}
    }
}
</script>
