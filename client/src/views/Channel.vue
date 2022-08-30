<template>
  <div style="margin-top: 2vh">
    <el-table
        :data="$store.state.channel.channelList"
        :header-cell-style="{background:'black',color:'#DEB887', borderColor: '#DEB887'}"
        style="width: 90vw;margin-left: 5vw;margin-right: 5vw;background-color: black;border-color: black; color: #DEB887">
      <el-table-column
          prop="ID"
          label="ID">
      </el-table-column>
      <el-table-column
          prop="id"
          label="Channel Name">
      </el-table-column>
      <el-table-column
          prop="blockNumber"
          label="Blocks">
      </el-table-column>
      <el-table-column
          prop="Transactions"
          label="Transactions">
      </el-table-column>
      <el-table-column
          prop="Timestamp"
          label="Timestamp">
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
export default {
  name: "Channel"
}
</script>

<style scoped>

</style>