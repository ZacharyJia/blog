<template>
  <div id="app">

    <el-menu :default-active="activeIndex" class="el-menu-demo" router="true" mode="horizontal" @select="handleSelect">
      <el-menu-item index="0" route="/" disabled="disabled">Zachary's Blog</el-menu-item>
      <el-menu-item index="1" route="/">首页</el-menu-item>
      <el-menu-item index="2" route="/about">关于</el-menu-item>
    </el-menu>
    <router-view></router-view>
  </div>
</template>

<script>
  import ElMenuItem from '../node_modules/element-ui/packages/menu/src/menu-item'
  export default {
    components: {ElMenuItem},
    name: 'app',
    data: function () {
      return {
        activeIndex: '1'
      }
    }
  }
</script>

<style scoped="">
  #app{
    margin: 0;
  }

  @import "assets/css/app.css";
</style>
