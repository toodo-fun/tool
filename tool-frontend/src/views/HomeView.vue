<template>
  <div class="home">
    <el-timeline>
      <el-timeline-item center v-for="(item, index) in changelog" :timestamp="item.version" :key="index"
        size="large" color="#0bbd87" placement="top">
        <el-card>
          <h4>{{ item.version }}</h4>
          <p v-for="(i, idx) in item.changelog" :key="idx">{{ i }}</p>
        </el-card>
      </el-timeline-item>
      <el-timeline-item center timestamp="Initial" placement="top">
        <el-card>
          <h4>Initial</h4>
          <p>初始化项目</p>
        </el-card>
      </el-timeline-item>
    </el-timeline>
  </div>
</template>

<script>
export default {
  name: 'HomeView',
  components: {
  },
  data: function () {
    return {
      changelog: []
    }
  },
  mounted() {
    const url = "/platform/changelog"
    this.$service.get(url).then((res) => {
      this.changelog = res
    })
  },
}
</script>

<style lang="less" scoped>
.home {
  background: #f8f8f8;
  height: 100%;
  width: 100%;
  padding: 16px;
  box-sizing: border-box;
}
</style>
