import Vue from 'vue';
import VueRouter from 'vue-router';
import ServersInfo from './components/serversInfo/ServersInfo.vue';
import SearchHistory from './components/history/History.vue';

Vue.use(VueRouter);

export default new VueRouter({
    routes: [
      {path: '/', component: ServersInfo},
      {path: '/history', component: SearchHistory}
    ]
   });