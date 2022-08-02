import Vue from 'vue'
import VueRouter from 'vue-router'
import HomeView from '@/views/HomeView.vue'
import DashBoard from "@/views/DashBoard";
import Transaction from "@/views/Transaction";
import Network from "@/views/Network";
import Block from "@/views/Block";
import Chaincode from "@/views/Chaincode";
import Channel from "@/views/Channel";

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView,
    redirect: '/dashboard',
    children:[
      {
        path: '/dashboard',
        name: 'dashboard',
        component: DashBoard,
      },
      {
        path: '/network',
        name: 'network',
        component: Network
      },
      {
        path: '/block',
        name: 'block',
        component: Block,
      },
      {
        path: '/transaction',
        name: 'transaction',
        component: Transaction,
      },
      {
        path: '/chaincode',
        name: 'chaincode',
        component: Chaincode,
      },
      {
        path: '/channel',
        name: 'channel',
        component: Channel,
      },
    ]
  }

]

const router = new VueRouter({
  routes
})

export default router
