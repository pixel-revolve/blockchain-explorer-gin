<template>
  <div id="app">
    <router-view></router-view>
  </div>
</template>

<script>

export default {
  name: 'app',
}
</script>

<style>
*{
  margin: 0;
  padding: 0;
}

body{
  background-color: rgba(0,0,0,1);
}

router-view{
  border-color: #DEB887;
}
</style>
