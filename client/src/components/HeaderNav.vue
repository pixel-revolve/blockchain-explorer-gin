<template>
  <div style="top: 0">
    <el-menu
        class="el-menu-demo"
        mode="horizontal"
        :ellipsis="false"
        @select="handleSelect"
        background-color="rgba(0,0,0,0.84)"
        text-color="#DEB887"
        active-text-color="rgba(255,255,255,1)"
        style="border-color: black">
      <el-menu-item style="font-family: Copperplate,Fantasy,SansSerif;font-size: 2vw" @click.native="goToHomePage" index="0">BLOCKCHAIN-EXPLORER</el-menu-item>
      <div class="flex-grow" />
      <el-menu-item index="1">NODES</el-menu-item>
      <el-menu-item index="3">BLOCKS</el-menu-item>
      <el-menu-item index="4">TRANSACTIONS</el-menu-item>
      <el-menu-item index="5">CHAINCODES</el-menu-item>
      <el-menu-item index="2">NETWORK</el-menu-item>
    </el-menu>
  </div>
</template>

<script>
export default {
  name: "Header",
  methods: {
    goToHomePage(){
      this.$router.push("/")
    },
    handleSelect(key, keyPath) {
      switch (key){
        case "1":{
          this.$router.push("/dashboard")
          break
        }
        case "2":{
          this.$router.push("/network")
          break
        }
        case "3":{
          this.$router.push("/block")
          break
        }
        case "4":{
          this.$router.push("/transaction")
          break
        }
        case "5":{
          this.$router.push("/chaincode")
          break
        }
        case "6":{
          this.$router.push("/channel")
          break
        }
      }
    }
  }
}
</script>

<style scoped>
.flex-grow {
  flex-grow: 1;
}
</style>