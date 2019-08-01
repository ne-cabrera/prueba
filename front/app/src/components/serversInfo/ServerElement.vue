<template>
    <div id='server-element'>
        <h3>IP: {{server.addres}}</h3>
        <h4>Country: {{server.country}}</h4>
        <h5>Owner: {{server.owner}}</h5>
    </div>
</template>
<script>
    export default {
    name: 'server-element',
    props: {
        server: Object
    }
}
</script>
