<template>
  <div style="margin-top: 2vh">
    <el-table
        :data="$store.state.network.peerList"
        border
        style="width: 90vw;margin-left: 5vw;margin-right: 5vw">
      <el-table-column
          prop="PeerName"
          label="Peer Name">
      </el-table-column>
      <el-table-column
          prop="PeerName"
          label="Request Url">
      </el-table-column>
      <el-table-column
          prop="Type"
          label="PeerType">
      </el-table-column>
      <el-table-column
          prop="Org"
          label="MSPID">
      </el-table-column>
      <el-table-column
          prop="High"
          label="High">
      </el-table-column>
      <el-table-column
          prop="Low"
          label="Low">
      </el-table-column>
      <el-table-column
          prop="Unsigned"
          label="Unsigned">
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
export default {
  name: "Network"
}
</script>

<style scoped>

</style>