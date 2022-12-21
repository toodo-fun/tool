<template>
    <div class="downloadSpeedUp">
        <div class="header">
            <div class="title">🅳🅾🆆🅽🅻🅾🅰🅳&nbsp;&nbsp;&nbsp;&nbsp;🆂🅿🅴🅴🅳 &nbsp;&nbsp;🆄🅿</div>
            <div class="desc">该功能主要用于文件下载加速，对下载慢的资源进行下载加速，支持国内外http/https链接</div>
        </div>
        <el-input v-model="input" clearable
            placeholder="请输入要下载的URL(如：https://download.jetbrains.com/idea/ideaIU-2022.2.3.exe)">
            <template #append>
                <el-button :loading="loading" @click="parseLink">
                    <el-icon v-if="!loading">
                        <svg t="1668516973955" class="icon" viewBox="0 0 1024 1024" version="1.1"
                            xmlns="http://www.w3.org/2000/svg" p-id="9466" width="200" height="200">
                            <path
                                d="M400.085333 151.210667c-89.059556 0-172.672 47.616-218.197333 124.273777a14.222222 14.222222 0 1 0 24.462222 14.535112c40.433778-68.067556 114.659556-110.364444 193.735111-110.364445a14.222222 14.222222 0 1 0 0-28.444444zM170.197333 336.298667a14.236444 14.236444 0 0 0-17.080889 10.609777c-4.252444 18.147556-6.656 29.397333-6.656 49.592889a14.222222 14.222222 0 1 0 28.444445 0c0-17.137778 1.92-26.083556 5.902222-43.121777a14.222222 14.222222 0 0 0-10.609778-17.080889z"
                                fill="" p-id="9467"></path>
                            <path
                                d="M947.384889 821.944889L717.809778 592.384a70.584889 70.584889 0 0 0-16.64-12.472889 346.254222 346.254222 0 0 0 47.36-175.089778c0-192.142222-156.302222-348.444444-348.444445-348.444444-192.128 0-348.444444 156.302222-348.444444 348.444444s156.316444 348.430222 348.444444 348.430223c68.408889 0 132.209778-19.882667 186.083556-54.058667 2.915556 4.821333 6.570667 9.472 10.951111 13.852444l229.575111 229.589334a84.849778 84.849778 0 0 0 60.359111 24.974222 84.821333 84.821333 0 0 0 60.344889-24.974222 84.778667 84.778667 0 0 0 24.974222-60.330667 84.920889 84.920889 0 0 0-24.988444-60.359111zM108.529778 404.835556c0-160.768 130.787556-291.555556 291.555555-291.555556 160.782222 0 291.555556 130.787556 291.555556 291.555556 0 160.782222-130.801778 291.541333-291.555556 291.541333-160.768 0-291.555556-130.759111-291.555555-291.541333z m798.620444 497.578666c-10.766222 10.723556-29.496889 10.723556-40.248889 0L637.340444 672.839111c-1.991111-1.976889-2.730667-3.299556-2.830222-3.299555 0.042667-0.611556 0.824889-6.471111 17.080889-22.698667 16.213333-16.256 22.072889-16.995556 22.257778-17.109333 0.369778 0.113778 1.692444 0.853333 3.740444 2.887111l229.546667 229.560889c5.390222 5.390222 8.334222 12.529778 8.334222 20.138666a28.401778 28.401778 0 0 1-8.32 20.096z"
                                fill="" p-id="9468"></path>
                        </svg>
                    </el-icon>
                    <span v-if="!loading">解析</span>
                    <span v-else>解析中</span>
                </el-button>
            </template>
        </el-input>
        <div class="downloadSpeedUp-result" v-if="result.status === 'success' || loading">
            <el-card class="box-card" v-if="!loading">
                <template #header>
                    <div class="card-header">
                        <span>{{ result.filename }}</span>
                    </div>
                </template>
                <div class="downloadSpeedUp-result-item">下载链接：{{ result.dlurl }}</div>
                <div class="downloadSpeedUp-result-item">文件大小：{{ result.size }} {{ result.sizeunit }}</div>
                <div class="downloadSpeedUp-result-item">文件MD5：{{ result.md5 }}</div>
                <div style="float: right; margin: 16px 0;">
                    <el-button-group>
                        <el-button color="#626aef" plain @click="copyLink">复制链接</el-button>
                        <el-button type="primary" plain @click="downloadLink">下载文件</el-button>
                    </el-button-group>
                </div>
            </el-card>
            <el-card class="box-card" v-else>
                <template #header>
                    <div class="card-header">
                        <el-skeleton :rows="0" animated />
                    </div>
                </template>
                <el-skeleton :rows="3" animated />
            </el-card>
        </div>
    </div>
</template>

<script>
import { ref } from 'vue';
import copy from 'copy-to-clipboard'
import { ElMessage } from 'element-plus';

export default {
    data: function () {
        return {
            input: "",
            loading: ref(false),
            result: {},
        }
    },
    methods: {
        parseLink() {
            if (!this.input) {
                ElMessage.error("输入不能为空")
                return
            }
            this.loading = true
            const url = "/universal/downloadSpeedUp?url=" + this.input
            this.$service.get(url).then((res) => {
                this.loading = false
                if (res.status !== "success") {
                    ElMessage.error("解析失败")
                    this.result = {}
                } else {
                    this.result = res
                }
            })
        },
        copyLink() {
            copy(this.result.dlurl)
            ElMessage.success("复制成功")
        },
        downloadLink() {
            window.openDefaultBrowser(this.result.dlurl)
        }
    }
}
</script>

<style lang="less" scoped>
.header {
    padding: 20px;
    text-align: center;
    .title {
        font-size: 28px;
    }

    .desc {
        font-size: small;
        color: gray;
        padding-top: 12px;
    }
}
.downloadSpeedUp {
    width: 100%;
    height: 100%;
    padding: 16px;
    box-sizing: border-box;
    background: transparent;
}

.downloadSpeedUp-result {
    margin-top: 24px;
}

.downloadSpeedUp-result-item {
    padding: 8px;
    font-size: small;

    &:not(:first-child) {
        border: 1px solid rgba(0, 0, 0, 0.1);
        border-style: solid none none none;
    }
}

.card-header {
    font-weight: bold;
}
</style>