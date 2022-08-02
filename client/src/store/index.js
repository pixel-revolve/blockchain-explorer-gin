import Vue from 'vue'
import Vuex from 'vuex'
import home from "./modules/home"
import chaincode from "@/store/modules/chaincode";
import transaction from "@/store/modules/transaction";
import blocks from "@/store/modules/blocks";
import channel from "@/store/modules/channel";
import network from "@/store/modules/network";


Vue.use(Vuex)

const store = new Vuex.Store({
  modules:{
    home,
    chaincode,
    transaction,
    blocks,
    channel,
    network,
  }
})

export default store

