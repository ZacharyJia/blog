<template>
  <div>
    <el-row class="article-block" v-for="item in blogs" :key="item.id">
      <el-col class="article-title">{{ item.title }}</el-col>
      <el-col class="article-date"><i class="el-icon-date"></i>{{ " " + item.date }}</el-col>
      <el-col class="article-brief">{{ item.content }}</el-col>
      <el-button class="btn-readmore" size="small">阅读全文</el-button>
      <el-col>
        <el-tag class="article-tag" v-for="tag in item.tags" type="gray" :key="tag">{{ tag.name }}</el-tag>
      </el-col>
    </el-row>
  </div>
</template>

<script>
  import ElButton from '../../node_modules/element-ui/packages/button/src/button'
  export default {
    components: {ElButton},
    data: function () {
      return {
        blogs: []
      }
    },
    mounted: function () {
      this.$http.get('http://localhost:8000/blog/list').then(response => {
        this.blogs = response.body
      }, response => {
        alert('failed')
      })
    }
  }
</script>

<style scoped="true">
  .article-block {
    margin-bottom: 40px;
    margin-top: 40px;
  }

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

  .btn-readmore {
    margin-top: 5px;
    margin-bottom: 10px;
  }
</style>
