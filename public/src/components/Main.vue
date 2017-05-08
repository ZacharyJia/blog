<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
    <el-button type="primary" icon="share" @click="open">分享</el-button>
    <el-alert
      title="成功提示的文案"
      type="success">
    </el-alert>

  </div>
</template>

<script>
export default {
  name: 'main',
  data () {
    return {
      msg: 'This is my main page'
    }
  },
  methods: {
    open () {
      this.$message({
        message: '恭喜恭喜',
        type: 'success'
      })
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
}

ul {
  list-style-type: none;
  padding: 0;
}

li {
  display: inline-block;
  margin: 0 10px;
}

a {
  color: #42b983;
}
</style>
