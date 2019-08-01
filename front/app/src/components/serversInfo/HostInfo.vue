<template>
    <div id='host-info' class="container">
        <div class="row">
            <div class="col-sm-12">
                <img v-bind:src="server.logo">
                <h4>{{server.title}}</h4>
            </div>
        </div>
        <div class="row">
            <div class="col-sm-6">
                <h4>Grade: {{server.ssl_grade}}</h4>
            </div>
            <div class="col-sm-6">
                <h4>Previous Grade: {{server.previous_sslGrade}}</h4>
            </div>
        </div>
        <div class="row">
            <div class="col-sm-6">
                <h4>Grade changed: {{server.servers_changed}}</h4>
            </div>
            <div class="col-sm-6">
                <h4>Host is down: {{server.is_down}}</h4>
            </div>
        </div>
    </div>
</template>
<script>
export default {
    name: "host-info",
    props: {
        server: Object
    }
}
</script>

