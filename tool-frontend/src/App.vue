<template>
  <div class="common-layout">
    <el-container class="main-container">
      <el-header>
        <el-row justify="space-between">
          <el-col :span="4" style="text-align: left">
            <div class="header-title">
              <el-avatar class="header-avatar" size="small" src="/favicon.ico" />
              <div>兔儿工具箱</div>
            </div>
          </el-col>
          <el-col :span="16" style="text-align: center">
            啦啦啦啦啦
          </el-col>
          <el-col :span="4" style="text-align: right">
            <el-button-group class="ml-4" size="small">
              <el-button color="transparent" text @click="bthMinimize">
                <el-icon color="gray">
                  <SemiSelect />
                </el-icon>
              </el-button>
              <el-button color="transparent" text @click="btnMaximize">
                <el-icon color="gray">
                  <FullScreen />
                </el-icon>
              </el-button>
              <el-button color="transparent" text @click="btnWinClose">
                <el-icon color="gray">
                  <CloseBold />
                </el-icon>
              </el-button>
            </el-button-group>
          </el-col>
        </el-row>
      </el-header>
      <el-container>
        <el-aside>
          <el-menu default-active="2" class="el-menu-vertical-demo" :collapse="isCollapse">
            <el-sub-menu index="1">
              <template #title>
                <el-icon>
                  <Plus />
                </el-icon>
                <span>Navigator One</span>
              </template>
              <el-menu-item-group>
                <template #title><span>Group One</span></template>
                <el-menu-item index="1-1">item one</el-menu-item>
                <el-menu-item index="1-2">item two</el-menu-item>
              </el-menu-item-group>
              <el-menu-item-group title="Group Two">
                <el-menu-item index="1-3">item three</el-menu-item>
              </el-menu-item-group>
              <el-sub-menu index="1-4">
                <template #title><span>item four</span></template>
                <el-menu-item index="1-4-1">item one</el-menu-item>
              </el-sub-menu>
            </el-sub-menu>
            <el-menu-item index="2">
              <el-icon>
                <TrendCharts />
              </el-icon>
              <template #title>Navigator Two</template>
            </el-menu-item>
            <el-menu-item index="3" disabled>
              <el-icon>
                <document />
              </el-icon>
              <template #title>Navigator Three</template>
            </el-menu-item>
            <el-menu-item index="4">
              <el-icon>
                <setting />
              </el-icon>
              <template #title>Navigator Four</template>
            </el-menu-item>
          </el-menu>
          <div class="aside-tool">
            <el-button type="danger" circle @click="handleCollapse">
              <el-icon v-if="isCollapse">
                <DArrowRight />
              </el-icon>
              <el-icon v-else>
                <DArrowLeft />
              </el-icon>
            </el-button>
          </div>
        </el-aside>
        <el-main>
          <router-view />
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<style lang="less">
html body {
  margin: 0;
  padding: 0;
}

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
  height: 100%;
  margin: 0;
  padding: 0;

  --header-height: 24px;
}

.common-layout {
  .main-container {
    position: absolute;
    top: 0;
    bottom: 0;
    left: 0;
    right: 0;
  }

  .aside-tool {
    padding: 8px;
    text-align: center;
  }

  .header-title {
    line-height: var(--header-height);
    display: flex;
    font-size: small;
    font-weight: 600;
    align-items: center;

    .header-avatar {
      margin: 0 4px;
    }
  }

  .el-header,
  .el-footer {
    background-color: var(--el-color-primary-light-7);
    color: var(--el-text-color-primary);
    height: var(--header-height);
    line-height: var(--header-height);
    margin: 0;
    padding: 0;
  }

  .el-aside {
    background-color: var(--el-color-primary-light-8);
    color: var(--el-text-color-primary);
    width: fit-content;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
  }

  .el-main {
    background-color: var(--el-color-primary-light-9);
    color: var(--el-text-color-primary);
    height: 100%;
  }
}
</style>

<script>
import { ref } from 'vue'
const ipcRenderer = window.electron.ipcRenderer

export default {
  data: function () {
    return {
      isCollapse: ref(true)
    }
  },
  methods: {
    handleCollapse() {
      console.log(123);
      this.isCollapse = !this.isCollapse
    },
    forceBlur(evt) {
      let target = evt.target
      while (target.nodeName !== "BUTTON") {
        target = target.parentNode
      }
      target.blur()
    },
    bthMinimize(evt) {
      this.forceBlur(evt)
      ipcRenderer.send('windowMin')
    },
    btnMaximize(evt) {
      this.forceBlur(evt)
      ipcRenderer.send('windowMax')
    },
    btnWinClose(evt) {
      this.forceBlur(evt)
      ipcRenderer.send('windowClose')
    }
  }
}
</script>
