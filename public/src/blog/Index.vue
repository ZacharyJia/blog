<template>
  <el-row>
    <el-col :span="20" :offset="2">
      <BlogList></BlogList>
    </el-col>
  </el-row>
</template>

<script>
  import BlogList from '../components/BlogList.vue'
  export default{
    components: {BlogList}
  }
</script>
