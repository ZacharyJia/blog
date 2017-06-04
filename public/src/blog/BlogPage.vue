<template>
  <el-row>
    <el-col :span="20" :offset="2">
      <el-col class="article-title">{{ blog.title }}</el-col>
      <el-col class="article-date"><i class="el-icon-date"></i>{{ " " + blog.date }}</el-col>
      <el-col>
        <el-tag class="article-tag" v-for="tag in blog.tags" type="gray" :key="tag">{{ tag.name }}</el-tag>
      </el-col>
      <el-col>
        {{ blog.content }}
      </el-col>
    </el-col>
  </el-row>
</template>

<script>
  export default {
    data: function () {
      return {
        blog: {}
      }
    },
    mounted: function () {
      let id = this.$route.params.id
      this.$http.get('http://localhost:8000/blog/' + id).then(response => {
        this.blog = response.body
      }, response => {
        alert('failed')
      })
    }
  }
</script>

<style scoped="true">

  .article-tag {
    margin-left: 4px;
    margin-right: 4px;
  }

  .article-title {
    color: #565a5f;
    font-size: 2.2em;
    transition: color 0.2s;
  }

  .article-date {
    font-size: 13px;
    color: #9a9ea3;
  }

  .article-brief {
    margin-top: 10px;
    margin-bottom: 10px;
  }

</style>
