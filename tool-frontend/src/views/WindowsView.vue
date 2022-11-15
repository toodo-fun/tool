<template>
  <div class="windows" v-loading="loading">
    <CellCardItem title="windows11折叠/展开右键菜单" desc="展开/折叠windows11系统的右键菜单"
      avatar="https://img2.baidu.com/it/u=3916745668,2257269414&fm=253&fmt=auto&app=138&f=JPEG?w=889&h=500">
      <template #right>
        <el-switch size="small" active-text="展开" inactive-text="折叠" v-model="contextMenuFold"
          :before-change="contextMenuFoldChange" />
      </template>
    </CellCardItem>
  </div>
</template>

<script>
import { ref } from 'vue';
import { ElMessage } from 'element-plus'
import CellCardItem from '@/components/CellCardItem.vue';

export default {
  components: {
    CellCardItem
  },
  data: function () {
    return {
      contextMenuFold: ref(true),
      loading: ref(false)
    };
  },
  mounted() {
    this.initContextMenuFold();
  },
  methods: {
    initContextMenuFold() {
      this.loading = true;
      const url = "/windows/contextMenu/status";
      this.$service.get(url).then((res) => {
        this.contextMenuFold = res;
        this.loading = false;
      });
    },
    async contextMenuFoldChange() {
      this.loading = true;
      let url = "";
      if (this.contextMenuFold) {
        url = "/windows/contextMenu/fold";
      }
      else {
        url = "/windows/contextMenu/unfold";
      }
      const res = await this.$service.get(url);
      this.loading = false;
      if (!res.success) {
        ElMessage.error(res.message);
      }
      return res.success;
    }
  },
}
</script>

<style lang="less" scoped>
.windows {
  width: 100%;
  height: 100%;
  overflow: auto;
}

.windows-item {
  width: 100%;
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 16px;
  background: rgba(255, 255, 255, 0.8);
  box-sizing: border-box;

  &:not(:first-child) {
    border: 1px solid rgba(0, 0, 0, 0.1);
    border-style: solid none none none;
  }

  .windows-item-title {
    font-weight: bold;
    font-size: small;
  }

  .windows-item-desc {
    font-size: small;
    color: gray;
  }
}
</style>