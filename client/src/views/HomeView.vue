<template>
  <div class="home">
    <Header></Header>
    <router-view></router-view>
  </div>
</template>

<script>

import Header from "@/components/HeaderNav";

export default {
  name: 'HomeView',
  components: {
    Header,
  },
  methods:{
    init(){
      this.$store.dispatch('home/getPeerNameList').catch(error=>{
        console.log(error)
      })
      this.$store.dispatch('home/getChannelNameList').catch(error=>{
        console.log(error)
      })
      this.$store.dispatch('home/getHighestBlock').catch(error=>{
        console.log(error)
      })
      this.$store.dispatch('chaincode/getChainCodeList').catch(error=>{
        console.log(error)
      })
      this.$store.dispatch('transaction/getTransactionList').catch(error=>{
        console.log(error)
      })
      this.$store.dispatch('blocks/getBlockList').catch(error=>{
        console.log(error)
      })
      this.$store.dispatch('channel/getChannelList').catch(error=>{
        console.log(error)
      })
      this.$store.dispatch('network/getPeerList').catch(error=>{
        console.log(error)
      })
    }
  },
  created() {
    this.init()
  },
}
</script>

<style>
</style>
