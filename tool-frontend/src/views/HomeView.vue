<template>
  <div class="home" v-loading="loading"
    element-loading-text="服务启动中...">
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
import { ref } from 'vue'

export default {
  name: 'HomeView',
  components: {
  },
  data: function () {
    return {
      changelog: [],
      loading: ref(false)
    }
  },
  mounted() {
    this.loading = true
    const url = "/platform/changelog"
    this.$service.get(url).then((res) => {
      this.changelog = res
      this.loading = false
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
