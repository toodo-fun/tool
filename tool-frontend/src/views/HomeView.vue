<template>
  <div class="home" v-loading="loading" element-loading-text="服务启动中...">
    <el-row :gutter="8">
      <el-col :span="16">
        <el-row :gutter="0">
          <el-col :span="12">
            <CellCardItem title="软件中心" desc="实用工具合集"
              avatar="https://ts1.cn.mm.bing.net/th/id/R-C.9d09099ba1df66d3a9ea863aebb19896?rik=u4PHIm89aMvczw&riu=http%3a%2f%2fwww.ambition-soft.com%2fupload%2fimage%2f20200927%2f20200927141007_7963.jpg&ehk=s9qTSE04X0bx%2fKVB9odZtOc76rFsMS5cbKLHx4JU51s%3d&risl=&pid=ImgRaw&r=0">
              <template #right>
                <el-button type="primary" text @click="$router.push('/extraTools')">打开</el-button>
              </template>
            </CellCardItem>
          </el-col>
          <el-col :span="12">
            <CellCardItem title="Windows工具" desc="Windows优化工具合集"
              avatar="https://cdn.windowsreport.com/wp-content/uploads/2019/12/Windows-7-1.jpg">
              <template #right>
                <el-button type="primary" text @click="$router.push('/windows')">打开</el-button>
              </template>
            </CellCardItem>
          </el-col>
          <el-col :span="24">
            <CellCardItem title="文件下载加速" desc="对下载慢的资源进行下载加速"
              avatar="https://tse2-mm.cn.bing.net/th/id/OIP-C.8nh6EUsftutQ7OS27lZ2EAHaE8?pid=ImgDet&rs=1">
              <template #right>
                <el-button type="primary" text @click="$router.push('/downloadSpeedUp')">打开</el-button>
              </template>
            </CellCardItem>
          </el-col>
          <el-col :span="24">
            <CellCardItem title="PDF合并工具" desc="PDF合并工具，支持将多个PDF文件合并为同一个文件"
              avatar="https://i0.wp.com/null48.com/wp-content/uploads/2018/01/PDF-Reader-Pro-Ipa-App-iOS-Free-Download.jpg">
              <template #right>
                <el-button type="primary" text @click="$router.push('/pdfMerge')">打开</el-button>
              </template>
            </CellCardItem>
          </el-col>
          <el-col :span="24">
            <CellCardItem title="PDF分割工具" desc="PDF分割工具，支持将一个PDF分割为多个文件"
              avatar="https://i0.wp.com/null48.com/wp-content/uploads/2018/01/PDF-Reader-Pro-Ipa-App-iOS-Free-Download.jpg">
              <template #right>
                <el-button type="primary" text @click="$router.push('/pdfSplit')">打开</el-button>
              </template>
            </CellCardItem>
          </el-col>
        </el-row>
      </el-col>
      <el-col :span="8" style="background: rgba(255, 255, 255, 0.8); ">
        <div style="display: flex; justify-content: space-between; flex-direction: column;">
          <div>
            <div style="padding: 8px; font-size:small;">获取更多精彩内容</div>
            <el-row style="margin-bottom: 8px;">
              <el-col :span="12" style="text-align: center">
                <el-image style=""
                  :src="require('../assets/qrcode_for_gh_bd8bca8aadc4_258.jpg')" fit="cover" />
                  <div style="font-size: 12px;">草莓橙汁儿</div>
              </el-col>
              <el-col :span="12" style="text-align: center">
                <el-image style=""
                  :src="require('../assets/qrcode_for_gh_d7ec3a9e36b4_258.jpg')" fit="cover" />
                  <div style="font-size: 12px">云原生技术研究院</div>
              </el-col>
            </el-row>
          </div>
          <div>
            <div style="padding: 8px; font-size:small">应用更新日志</div>
            <el-timeline style="padding: 8px;">
              <el-timeline-item center v-for="(item, index) in changelog" :key="index" color="#0bbd87" placement="top">
                <!-- <el-card> -->
                  <div style="font-weight: bold; padding-bottom: 8px;">{{ item.version }}</div>
                  <div style="font-size: small; line-height: 22px;" v-for="(i, idx) in item.changelog" :key="idx">{{ i
                  }}
                  </div>
                <!-- </el-card> -->
              </el-timeline-item>
              <!-- <el-timeline-item center placement="top">
                <el-card>
                  <div style="font-weight: bold; padding-bottom: 8px;">我们开始出发啦</div>
                </el-card>
              </el-timeline-item> -->
            </el-timeline>
          </div>
        </div>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import { ref } from 'vue'
import CellCardItem from '@/components/CellCardItem.vue'

export default {
  name: 'HomeView',
  components: {
    CellCardItem
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
  background: transparent;
  height: 100%;
  width: 100%;
  box-sizing: border-box;
  overflow: auto;
  padding: 8px;
}
</style>
