<template>
<div id="container" @wheel="handleScroll">
  <nodes-card id="nodeCard"></nodes-card>
  <img id="down" src="svg/down.svg" class="image" style="margin-left: 50%; width: 5%;margin-top: 5%">
  <img id="up" src="svg/up.svg" class="image" style="margin-left: 49%; width: 5%;margin-top: 5%">
  <div id="nodeTable">
  <div class="dashboard-tables" style="margin-left: 5vw;margin-right: 5vw;display: flex; margin-top: 3%">
    <el-table
        :data="$store.state.home.peerNameList"
        :header-cell-style="{background:'black',color:'#DEB887', borderColor: '#DEB887'}"
        style="width: 30vw;margin-right: 1vw;background-color: black;border-color: black; color: #DEB887">
      <el-table-column
          prop="peerName"
          label="PEER NAME">
      </el-table-column>
    </el-table>
    <el-table
        :data="$store.state.home.channelNameList"
        :header-cell-style="{background:'black',color:'#DEB887', borderColor: '#DEB887'}"
        style="width: 30vw;margin-left: 1vw;background-color: black;border-color: black; color: #DEB887">
      <el-table-column
          prop="channelName"
          label="CHANNEL NAME">
      </el-table-column>
    </el-table>
  </div>
  </div>

</div>
</template>

<script>
import Card from "@/components/Card";
import NodesCard from "@/NewCards/nodesCard";

export default {
  name: "DashBoard",
  components: { NodesCard, Card},
  methods: {
    handleScroll(e) {
      if (e.deltaY > 0) {
        // 向下滚动...
        var t =document.getElementById("nodeCard")
        t.style.display = "none"
        var x =document.getElementById("nodeTable")
        x.style.display = "inline"
        var d =document.getElementById("down")
        d.style.display = "none"
        var u =document.getElementById("up")
        u.style.display = "inline"
      } else {
        // 向上滚动...
        var t =document.getElementById("nodeCard")
        t.style.display = "inline"
        var x =document.getElementById("nodeTable")
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
#nodeTable{
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