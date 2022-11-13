<template>
  <div class="home">
    <el-button type="primary" style="margin-left: 16px" @click="drawer = true">
    open
  </el-button>

  <el-drawer v-model="drawer" title="I am the title" :with-header="false">
    <span>Hi there!</span>
  </el-drawer>
  </div>
</template>

<script>
// @ is an alias to /src
// import HelloWorld from '@/components/HelloWorld.vue'

import { ref } from 'vue';

export default {
  name: 'HomeView',
  components: {
    // HelloWorld
  },
  data: function() {
    return {
      drawer: ref(false)
    }
  }
}
</script>
