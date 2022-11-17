<template>
    <div class="PDF">
        <div class="header">
            <div class="title">🅿🅳🅵&nbsp;&nbsp;&nbsp;&nbsp;🆂🅿🅻🅸🆃</div>
            <div class="desc">PDF分割工具，支持将一个PDF分割为多个文件</div>
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
                <el-button v-if="inFiles.length == 0" type="primary" plain @click="addPDF">点击添加一个PDF文件</el-button>
            </el-tooltip>
            <el-tooltip content="分割文件" placement="bottom" effect="light" v-if="inFiles.length >= 1">
                <el-button type="success" circle plain @click="dialogFormVisible = true">
                    <el-icon>
                        <Grid />
                    </el-icon>
                </el-button>
            </el-tooltip>
            <el-tooltip content="清空文件" placement="bottom" effect="light" v-if="inFiles.length >= 1">
                <el-button type="danger" circle plain @click="clearPDF">
                    <el-icon>
                        <Close />
                    </el-icon>
                </el-button>
            </el-tooltip>
        </div>
        <el-dialog v-model="dialogFormVisible" :show-close="false">
            每
            <el-input-number v-model="span" :step="1" step-strictly :min="1" :max="inFiles[0].page"
                style="width: fit-content" /> 页进行分割
            <template #footer>
                <span class="dialog-footer">
                    <el-button @click="dialogFormVisible = false">取消</el-button>
                    <el-button type="primary" @click="splitPDF">
                        分割
                    </el-button>
                </span>
            </template>
        </el-dialog>
    </div>
</template>
  
<script>
import { ref } from 'vue'
import { dialog, Notification } from '@electron/remote'
import draggable from 'vuedraggable'
import CellCardItem from '@/components/CellCardItem.vue'

export default {
    components: {
        draggable,
        CellCardItem
    },
    data: function () {
        return {
            dialogFormVisible: ref(false),
            span: 1,
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
                            { filename: res.filename, desc: "文档页数: " + res.page + " md5: " + res.md5, md5: res.md5, page: res.page }
                        )
                        // if (!this.inFiles.find((item) => item.md5 === res.md5) || true) {

                        // } else {
                        //     ElMessage.warning("该文件已经添加")
                        // }
                    })
                })
            })
        },
        async splitPDF() {
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
                    infile: inFs[0],
                    outDir: result.filePaths[0],
                    span: 1
                }
                console.log(payload);
                const url = "/pdf/split"
                this.$service.post(url, payload).then(() => {
                    new Notification({ title: "文件分割成功", body: "文件路径: " + result.filePaths[0] }).show();
                    window.openDefaultBrowser(result.filePaths[0])
                    this.dialogFormVisible = false
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