<template>
  <div style="margin-top: 2vh">
    <el-table
        :data="$store.state.blocks.blockList"
        border
        style="width: 90vw;margin-left: 5vw;margin-right: 5vw">
      <el-table-column
          prop="BlockHeight"
          label="Block Number">
      </el-table-column>
      <el-table-column
          prop="Channel Name"
          label="Channel Name">
      </el-table-column>
      <el-table-column
          prop="TxNumber"
          label="Number of Tx">
      </el-table-column>
      <el-table-column
          prop="DataHash"
          label="Data Hash">
      </el-table-column>
      <el-table-column
          prop="CurrentBlockHash"
          label="Block Hash">
      </el-table-column>
      <el-table-column
          prop="PreviousBlockHash"
          label="Previous Hash">
      </el-table-column>
      <el-table-column
          prop="Transactions"
          label="Transactions">
      </el-table-column>
      <el-table-column
          prop="Size"
          label="Size(KB)">
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
export default {
  name: "Block"
}
</script>

<style scoped>

</style>