<template>
    <div id='servers' class="container">
        <table class='table table-striped'>
            <thead>
                <th scope="col">Server</th>
                <th scope="col">Grade</th>
            </thead>
            <tbody>
                <tr v-for="server in servers" :key="server.addres">
                    <td>
                        <server-element :server="server"/>
                    </td>
                    <td>
                        <h3>{{server.ssl_grade}}</h3>
                    </td>
                </tr>
            </tbody>
        </table>
    </div>
</template>
<script>
import ServerElement from "./ServerElement.vue"
export default {
    name: 'servers',
    props: {
        servers: Array
    },
    components: {
        ServerElement
    }
}
</script>

