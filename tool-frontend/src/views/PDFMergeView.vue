<template>
    <div class="PDF">
        <div class="header">
            <div class="title">🅿🅳🅵&nbsp;&nbsp;&nbsp;&nbsp;🅼🅴🆁🅶🅴</div>
            <div class="desc">PDF合并工具，支持上下拖拽排序</div>
        </div>
        <draggable :list="inFiles" item-key="filename">
            <template #item="{ element }">
                <CellCardItem :title="element.filename" :desc="element.desc"
                    avatar="https://i0.wp.com/null48.com/wp-content/uploads/2018/01/PDF-Reader-Pro-Ipa-App-iOS-Free-Download.jpg">
                    <template #right>
                        <el-tooltip content="移除该文档" placement="bottom" effect="light">
                            <el-button type="primary" circle text plain @click="deletePDF(element.md5)">
                                <el-icon>
                                    <Delete />
                                </el-icon>
                            </el-button>
                        </el-tooltip>
                    </template>
                </CellCardItem>
            </template>
        </draggable>
        <div style="text-align: center;">
            <el-tooltip content="添加文件" placement="bottom" effect="light">
                <el-button v-if="inFiles.length > 0" type="primary" circle plain @click="addPDF">
                    <el-icon>
                        <Plus />
                    </el-icon>
                </el-button>
                <el-button v-else type="primary" plain @click="addPDF">点击添加第一个PDF文件</el-button>
            </el-tooltip>
            <el-tooltip content="开始合并" placement="bottom" effect="light" v-if="inFiles.length >= 2">
                <el-button type="success" circle plain @click="mergePDF">
                    <el-icon>
                        <HelpFilled />
                    </el-icon>
                </el-button>
            </el-tooltip>
            <el-tooltip content="清空文件" placement="bottom" effect="light" v-if="inFiles.length >= 2">
                <el-button type="danger" circle plain @click="clearPDF">
                    <el-icon>
                        <Close />
                    </el-icon>
                </el-button>
            </el-tooltip>
        </div>
    </div>
</template>
  
<script>
import { dialog, Notification } from '@electron/remote'
import draggable from 'vuedraggable'
import CellCardItem from '@/components/CellCardItem.vue'
import path from 'path'

export default {
    components: {
        draggable,
        CellCardItem
    },
    data: function () {
        return {
            input: "",
            inFiles: [

            ],
        }
    },
    mounted() { },
    methods: {
        async addPDF() {
            dialog.showOpenDialog({
                title: "选择要合并的PDF文件",
                properties: ["multiSelections"],
                filters: [
                    { name: 'pdf', extensions: ['pdf'] }
                ]
            }).then(result => {
                result.filePaths.forEach((item) => {
                    const url = "/pdf/info?filepath=" + item
                    this.$service.get(url).then((res) => {
                        // 支持同一文件多次添加
                        this.inFiles.push(
                            { filename: res.filename, desc: "文档页数: " + res.page + " md5: " + res.md5, md5: res.md5 }
                        )
                        // if (!this.inFiles.find((item) => item.md5 === res.md5) || true) {

                        // } else {
                        //     ElMessage.warning("该文件已经添加")
                        // }
                    })
                })
            })
        },
        async mergePDF() {
            dialog.showOpenDialog({
                title: "选择要保存的位置",
                properties: ["openDirectory"],
            }).then(result => {
                if (result.canceled) {
                    return
                }
                let inFs = []
                this.inFiles.forEach((item) => {
                    inFs.push(item.filename)
                })
                const payload = {
                    inFiles: inFs,
                    outFile: result.filePaths[0] + path.sep + "result.pdf"
                }
                const url = "/pdf/merge"
                this.$service.post(url, payload).then((res) => {
                    // ElMessage.success("合并成功")
                    new Notification({ title: "文件合并成功", body: "文件路径: " + res }).show();
                    window.openDefaultBrowser(res)
                })
            })
        },
        deletePDF(md5) {
            const tmp = this.inFiles.filter((item) => item.md5 !== md5)
            this.inFiles = tmp
        },
        clearPDF() {
            this.inFiles = []
        }
    }
}
</script>

<style lang="less" scoped>
.PDF {
    background: rgba(255, 255, 255, 0.8);
    height: 100%;
    width: 100%;
    box-sizing: border-box;
    overflow: auto;
    padding: 8px;
}

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
</style>