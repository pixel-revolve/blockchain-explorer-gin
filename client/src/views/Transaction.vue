<template>
  <div style="margin-top: 2vh">
    <el-table
        :data="$store.state.transaction.transactionList"
        border
        style="width: 90vw;margin-left: 5vw;margin-right: 5vw">
      <el-table-column
          prop="CreatorMsp"
          label="Creator">
      </el-table-column>
      <el-table-column
          prop="Channel Name"
          label="Channel Name">
      </el-table-column>
      <el-table-column
          prop="TxID"
          label="Tx Id">
      </el-table-column>
      <el-table-column
          prop="Type"
          label="Type">
      </el-table-column>
      <el-table-column
          prop="ChaincodeID"
          label="Chaincode">
      </el-table-column>
      <el-table-column
          prop="CreateTime"
          label="Timestamp">
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
export default {
  name: "Transaction"
}
</script>

<style scoped>

</style>