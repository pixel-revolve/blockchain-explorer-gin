<template>
  <div style="margin-top: 2vh">
    <el-table
        :data="$store.state.chaincode.chainCodeList"
        border
        style="width: 90vw;margin-left: 5vw;margin-right: 5vw">
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
</template>

<script>
export default {
  name: "Chaincode"
}
</script>

<style scoped>

</style>