<template>
    <div class="PDF">
        <div @click="onChange">123</div>
    </div>
</template>
  
<script>
import { dialog } from '@electron/remote'

export default {
    data: function () {
        return {
            input: ""
        }
    },
    mounted() { },
    methods: {
        async onChange() {
            dialog.showOpenDialog({
                properties: ["multiSelections"],
            }).then(result => {
                console.log(result.filePaths);
            })
        }
    }
}
</script>