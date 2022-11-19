<template>
    <div class="extraTools" v-loading="loading">
        <el-row :gutter="8">
            <el-col :span="8" v-for="item, index in extraTools" :key="index">
                <CardItem :title="item.title" :describe="item.describe" :icon="item.icon">
                    <template #footer>
                        <el-progress v-if="item.status == 'installing'" :percentage="getPercent(item.downloadInfo)" />
                        <div style="display: flex; justify-content: flex-end;">
                            <el-button v-if="item.status == 'uninstall'" type="info" text plain @click="install(item)">
                                安装</el-button>
                            <el-button v-if="item.status == 'installed'" type="primary" text plain
                                @click="start(item.startCmd)">启动</el-button>
                        </div>
                    </template>
                </CardItem>
            </el-col>
        </el-row>
    </div>
</template>

<script>
import CardItem from '@/components/CardItem.vue';
import { ref } from 'vue'

export default {
    data: function () {
        return {
            loading: ref(false),
            extraTools: []
        }
    },
    components: {
        CardItem
    },
    mounted() {
        this.init()
    },
    unmounted() {
        clearTimeout(this.initST)
    },
    methods: {
        init() {
            this.loading = true
            const url = "/extraTool/list"
            this.$service.get(url).then((res) => {
                this.extraTools = res
                this.loading = false
                this.initST = setTimeout(() => {
                    clearTimeout(this.initST)
                    this.init()
                }, 1000)
            })
        },
        getPercent(info) {
            try {
                const i = JSON.parse(info)
                return Math.trunc(i.downloaded / i.total * 100) ? Math.trunc(i.downloaded / i.total * 100) : 0
            } catch (error) {
                return 0
            }
        },
        install(item) {
            this.loading = true
            const url = "/extraTool/install"
            this.$service.post(url, item).then(() => {
                this.loading = false
            })
        },
        start(cmd) {
            window.openDefaultBrowser("./resources/server" + cmd)
        }
    }
}
</script>

<style lang="less" scoped>
.extraTools {
    height: 100%;
    width: 100%;
    box-sizing: border-box;
    overflow: auto;
    padding: 0 8px;
}
</style>