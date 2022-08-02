<template>
<div>
  <Card></Card>
  <div class="dashboard-tables" style="margin-left: 5vw;margin-right: 5vw;display: flex">
    <el-table
        :data="$store.state.home.peerNameList"
        border
        style="width: 30vw;margin-right: 1vw">
      <el-table-column
          prop="peerName"
          label="PEER NAME">
      </el-table-column>
    </el-table>
    <el-table
        :data="$store.state.home.channelNameList"
        border
        style="width: 30vw;margin-left: 1vw">
      <el-table-column
          prop="channelName"
          label="CHANNEL NAME">
      </el-table-column>
    </el-table>
  </div>
</div>
</template>

<script>
import Card from "@/components/Card";

export default {
  name: "DashBoard",
  components: {Card}
}
</script>

<style scoped>

</style>