<template>
    <div class="card-item">
        <div>
            <div class="card-item-head">
                <div>{{ title }}</div>
                <el-avatar v-if="icon" shape="square" :src="icon" :size="28" style="background: transparent;" />
            </div>
            <div class="card-item-body">
                {{ describe }}
            </div>
        </div>
        <div>
            <slot name="footer"></slot>
        </div>
    </div>
</template>

<script>
export default {
    name: 'CardItem',
    props: {
        title: String,
        describe: String,
        icon: String
    }
}
</script>

<style lang="less" scoped>
.card-item {
    background: rgba(255, 255, 255, 0.8);
    padding: 8px;
    box-sizing: border-box;
    border-radius: 8px;
    margin: 12px 0 0 0;
    height: 100%;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
}

.card-item-head {
    display: flex;
    justify-content: space-between;
    align-items: center;
    font-weight: bold;
}

.card-item-body {
    font-size: small;
    color: gray;
    margin: 8px 0;
}
</style>