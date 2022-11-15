<template>
  <div class="cell-card-item">
    <div style="display: flex;">
      <el-avatar v-if="avatar" shape="square" :src="avatar" style="margin-right: 8px" />
      <div>
        <div class="cell-card-item-title">{{ title }}</div>
        <div class="cell-card-item-desc" v-if="desc">{{ desc }}</div>
      </div>
    </div>
    <div>
      <slot name="right"></slot>
    </div>
  </div>
</template>

<script>
export default {
  name: 'CellCardItem',
  props: {
    title: String,
    desc: String,
    avatar: String
  }
}
</script>

<style scoped lang="less">
.cell-card-item {
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

  .cell-card-item-title {
    font-weight: bold;
    font-size: small;
  }

  .cell-card-item-desc {
    font-size: small;
    color: gray;
  }
}
</style>