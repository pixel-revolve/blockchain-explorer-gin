<template>
  <div class="root" id="root">

    <el-row>
      <el-col :span="8" :offset="8">
        <h1 style="font-family: Copperplate,Fantasy,SansSerif;font-size: 2vw;color: #DEB887;text-align: center" >BLOCKCHAIN-EXPLORER</h1>


        <el-card class="box-card">
          <h1 class="title"> Register </h1>

          <el-col :span="18" :offset="3">
            <el-form
                class="form"
                label-width="100px"
            >
              <el-form-item label="Name" class="form-item" >
                <el-input v-model="name" placeholder="Please input name..." class="input"/>
              </el-form-item>
              <el-form-item label="Password" class="form-item">
                <el-input v-model="password" type="password" placeholder="Please input password..." show-password class="input"/>
              </el-form-item>
              <el-form-item>
                <el-button @click="onSubmit">Sign Up</el-button>
                &nbsp &nbsp
                <router-link to="login" class="switch">Go sign in...</router-link>
              </el-form-item>
            </el-form>

          </el-col>

        </el-card>
      </el-col>
    </el-row>

  </div>
</template>

<script>
import 'animate.css';

export default {
  name: "RegisterView",
  data() {
    return {
      "name": "",
      "password": "",
    }
  },
  methods: {
    onSubmit() {
      console.log('submit!')
    }
  }
}
</script>

<style scoped>
#root{
  animation: fadeIn;
  animation-duration: 1s; /* don't forget to set a duration! */
}

.root{
  padding: 100px 0;

}
.box-card{
  background-color: black;
  border-color: #DEB887;
  text-align: center;
}
.form{
  max-width: 460px;
  text-align: center;
  margin-top: 5%;
}
.form-item{
  color: #DEB887;

}
.el-button{
  background-color: black;
  color: #DEB887;
  border-color: #DEB887;
}
.el-button:hover{
  background-color: #DEB887;
  color: aliceblue;
  border-color: #DEB887;
}
.title{
  text-align: center;
  color: #DEB887;
}
.input{
  text-align: center;
}
</style>

<style>
.form-item .el-form-item__label {
  color: #DEB887;
}
.input .el-input__inner{
  border-color: #DEB887;
  color: #DEB887;
  background-color: black;
  text-align: center;
}
.el-input__inner::placeholder {
  color: #DEB887;
  text-align:center;
}
/* 谷歌 */
.el-input__inner::-webkit-input-placeholder {
  color: #DEB887;
  text-align:center;
}
/* 火狐 */
.el-input__inner:-moz-placeholder {
  color: #DEB887;
  text-align:center;
}
.switch{
  color: aliceblue;
  text-decoration: aliceblue;
}
</style>
