<template>
  <el-row :gutter="20">
    <el-col :span="11"><div class="grid-content ep-bg-purple"/>
      <el-card class="box-card">
        <img src="svg/node.svg" class="image" style="margin-left: 50%; margin-top: 20%">
      </el-card>
    </el-col>
    <el-col :span="11"><div class="grid-content ep-bg-purple" />
      <el-card class="box-card" style="margin-top: 20%">
        <h1>{{this.$store.state.home.cardList[2].title}}</h1>
        <br><br>
       <h3>{{this.$store.state.home.cardList[2].number}}</h3>
      </el-card>
    </el-col>
  </el-row>
</template>

<script>
import 'animate.css';
export default {
  name: "nodesCard"
}
</script>

<style scoped>
.el-row {
  margin-bottom: 20px;
  color: #DEB887;

  animation: fadeInUp; /* referring directly to the animation's @keyframe declaration */
  animation-duration: 1s; /* don't forget to set a duration! */
}
.el-row:last-child {
  margin-bottom: 0;
}
.el-col {
  border-radius: 4px;
}

.grid-content {
  border-radius: 4px;
  min-height: 36px;
}
.el-card{
  background-color: black;
  border-color: black;
  color: #DEB887;
}


img{
  width: 40%;
}

h1{
  font-size: 80px;
  text-align: center;
}
h3{
  font-size: 40px;
  text-align: center;
}
</style>