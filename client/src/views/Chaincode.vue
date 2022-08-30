<template>
  <div style="margin-top: 2vh" @wheel="handleScroll">
    <chaincode-card id="chaincodeCard"></chaincode-card>
    <img id="down" src="svg/down.svg" class="image" style="margin-left: 50%; width: 5%;margin-top: 5%">
    <img id="up" src="svg/up.svg" class="image" style="margin-left: 49%; width: 5%;margin-top: 5%">
    <div id="chaincodeTable">
    <el-table
        :data="$store.state.chaincode.chainCodeList"
        :header-cell-style="{background:'black',color:'#DEB887', borderColor: '#DEB887'}"
        style="width: 90vw;margin-left: 5vw;margin-right: 5vw;background-color: black;border-color: black; color: #DEB887;margin-top: 3%">
      <el-table-column
          prop="chaincodeName"
          label="Chaincode Name">
      </el-table-column>
      <el-table-column
          prop="channelName"
          label="Channel Name">
      </el-table-column>
      <el-table-column
          prop="path"
          label="Path">
      </el-table-column>
      <el-table-column
          prop="transactionCount"
          label="Transaction Count">
      </el-table-column>
      <el-table-column
          prop="version"
          label="Version">
      </el-table-column>
    </el-table>
    </div>
  </div>
</template>

<script>
import ChaincodeCard from "@/NewCards/chaincodeCard";
export default {
  name: "Chaincode",
  components: {ChaincodeCard},
  methods: {
    handleScroll(e) {
      if (e.deltaY > 0) {
        // 向下滚动...
        var t =document.getElementById("chaincodeCard")
        t.style.display = "none"
        var x =document.getElementById("chaincodeTable")
        x.style.display = "inline"
        var d =document.getElementById("down")
        d.style.display = "none"
        var u =document.getElementById("up")
        u.style.display = "inline"
      } else {
        // 向上滚动...
        var t =document.getElementById("chaincodeCard")
        t.style.display = "inline"
        var x =document.getElementById("chaincodeTable")
        x.style.display = "none"
        var d =document.getElementById("down")
        d.style.display = "inline"
        var u =document.getElementById("up")
        u.style.display = "none"
      }
    }
  }
}
</script>

<style scoped>
#chaincodeTable{
  display: none;
  margin-top: 40%;
  animation: fadeInUp; /* referring directly to the animation's @keyframe declaration */
  animation-duration: 1s; /* don't forget to set a duration! */
}
#down{
  animation: fadeInUp; /* referring directly to the animation's @keyframe declaration */
  animation-duration: 1s; /* don't forget to set a duration! */
}
#up{
  display: none;
  animation: fadeInUp; /* referring directly to the animation's @keyframe declaration */
  animation-duration: 1s; /* don't forget to set a duration! */
}
</style>