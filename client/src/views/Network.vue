<template>
  <div style="margin-top: 50px;background-color: black;border-color: black">
    <el-table
        :data="$store.state.network.peerList"
        style="width: 90vw;margin-left: 5vw;margin-right: 5vw;background-color: black;border-color: black; color: #DEB887"
        :header-cell-style="{background:'black',color:'#DEB887', borderColor: '#DEB887'}"
    >
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

<br><br><br>
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
import 'animate.css';

export default {
  name: "Network"
}
</script>

<style scoped>

div{
  animation: fadeInUp; /* referring directly to the animation's @keyframe declaration */
  animation-duration: 1s; /* don't forget to set a duration! */
}

el-table-column{
  border-color: black;
  background-color: black;
}
</style>