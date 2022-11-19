<template>
  <div class="common-layout">
    <el-container class="main-container">
      <el-header>
        <el-row justify="space-between">
          <el-col :span="3" style="text-align: left">
            <div class="header-title">
              <el-avatar class="header-avatar" size="small" src="/favicon.ico" />
              <div>兔儿工具箱</div>
            </div>
          </el-col>
          <el-col :span="18" class="header-message">
            <!-- 也许命运扼住你的咽喉，只是让你少吃两口。 -->
          </el-col>
          <el-col :span="3" class="header-tool" style="text-align: right; padding-right: 6px;">
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
          <el-menu :default-active="$route.path" class="el-menu-vertical-demo" :collapse="isCollapse" router>
            <el-menu-item index="/">
              <el-icon size="24">
                <svg t="1668430775193" class="icon" viewBox="0 0 1024 1024" version="1.1"
                  xmlns="http://www.w3.org/2000/svg" p-id="1288" width="200" height="200">
                  <path
                    d="M779.994172 65.344241 243.982292 65.344241c-98.714256 0-178.692798 79.979566-178.692798 178.715311l0 536.012903c0 98.602715 79.978543 178.583305 178.692798 178.583305l536.01188 0c98.627275 0 178.715311-79.979566 178.715311-178.583305L958.709483 244.059552C958.709483 145.323806 878.621447 65.344241 779.994172 65.344241zM869.308337 735.405139c0 73.979922-59.957557 133.958968-133.958968 133.958968L288.650631 869.364107c-74.002434 0-133.958968-59.979046-133.958968-133.958968L154.691663 288.705378c0-74.001411 59.956533-134.046972 133.958968-134.046972l446.698738 0c74.001411 0 133.958968 60.045561 133.958968 134.046972L869.308337 735.405139zM243.982292 467.310172l0 44.75532 44.667316 0 0 178.583305-44.667316 0 0 89.423659 536.01188 0 0-89.423659-44.645826 0L735.348346 512.065492l44.645826 0 0-44.75532L511.988232 244.059552 243.982292 467.310172zM467.342406 690.648796 377.94126 690.648796 377.94126 512.065492l89.400123 0L467.341383 690.648796zM556.656571 512.065492l89.291652 0 0 178.583305-89.291652 0L556.656571 512.065492z"
                    p-id="1289"></path>
                </svg>
              </el-icon>
              <template #title>主页</template>
            </el-menu-item>
            <!-- <el-menu-item index="1">
              <el-icon size="24">
                <svg t="1668430791791" class="icon" viewBox="0 0 1024 1024" version="1.1"
                  xmlns="http://www.w3.org/2000/svg" p-id="1430" width="200" height="200">
                  <path
                    d="M467.106022 594.041528l-74.343195 74.299193 74.343195 74.299193-37.128619 37.214577-74.297147-74.299193-74.387197 74.299193-37.126572-37.214577 74.341149-74.299193-74.341149-74.299193 37.126572-37.170574 74.387197 74.255191 74.297147-74.255191L467.106022 594.041528zM593.975525 467.127l74.429153-74.385151 74.255191 74.385151 37.258579-37.170574-74.429153-74.299193 74.429153-74.429153-37.258579-37.08257-74.255191 74.297147-74.429153-74.297147-37.08257 37.08257 74.297147 74.429153-74.297147 74.299193L593.975525 467.127zM958.70846 243.968477l0 536.060999c0 98.642624-80.013335 178.744987-178.615027 178.744987L243.991502 958.774463c-98.774631 0-178.700985-80.101339-178.700985-178.744987L65.290517 243.968477c0-98.772584 79.925331-178.744987 178.700985-178.744987l536.102954 0C878.695125 65.22349 958.70846 145.19487 958.70846 243.968477zM869.401458 288.599977c0-73.994248-59.989279-134.069485-134.069485-134.069485L288.667004 154.530492c-73.994248 0-134.115534 60.075237-134.115534 134.069485l0 446.797999c0 73.994248 60.121286 133.939525 134.115534 133.939525l446.664969 0c74.080206 0 134.069485-59.945277 134.069485-133.939525L869.401458 288.599977zM467.279984 355.656209c0 61.604056-49.953715 111.688755-111.688755 111.688755-61.646012 0-111.600751-50.083675-111.600751-111.688755 0-61.778019 49.953715-111.688755 111.600751-111.688755C417.325245 243.968477 467.279984 293.878191 467.279984 355.656209zM421.818583 355.656209c0-36.647665-29.623691-66.271356-66.227354-66.271356-36.517705 0-66.141396 29.623691-66.141396 66.271356 0 36.473703 29.623691 66.141396 66.141396 66.141396C392.194892 421.797605 421.818583 392.129912 421.818583 355.656209zM780.093433 668.340721c0 61.646012-50.083675 111.688755-111.688755 111.688755-61.775972 0-111.688755-50.04172-111.688755-111.688755s49.91176-111.688755 111.688755-111.688755C730.009757 556.652989 780.093433 606.694709 780.093433 668.340721zM734.546074 668.340721c0-36.517705-29.667693-66.141396-66.141396-66.141396-36.647665 0-66.271356 29.623691-66.271356 66.141396s29.623691 66.183352 66.271356 66.183352C704.87838 734.524073 734.546074 704.858426 734.546074 668.340721z"
                    p-id="1431"></path>
                </svg>
              </el-icon>
              <template #title>通用工具</template>
            </el-menu-item> -->
            <el-menu-item index="/windows">
              <el-icon size="24">
                <svg t="1668430732657" class="icon" viewBox="0 0 1024 1024" version="1.1"
                  xmlns="http://www.w3.org/2000/svg" p-id="1004" width="200" height="200">
                  <path
                    d="M511.955486 288.62249l268.051989-44.654013 0 223.376487L511.955486 467.344964 511.955486 288.62249zM422.647461 556.652989 243.903497 556.652989l0 134.069485 264.306688 44.029796-85.561701-14.756075L422.648484 556.652989zM511.955486 735.397976l0-0.021489-3.745301-0.623193L511.955486 735.397976zM511.955486 735.375463l268.051989 44.654013L780.007475 556.652989 511.955486 556.652989 511.955486 735.375463zM243.903497 467.344964l178.744987 0L422.648484 303.499315l-178.744987 29.776164L243.903497 467.344964zM958.70846 243.968477l0 536.060999c0 98.642624-79.971379 178.744987-178.700985 178.744987L243.903497 958.774463c-98.642624 0-178.61298-80.101339-178.61298-178.744987L65.290517 243.968477c0-98.772584 79.969333-178.744987 178.61298-178.744987l536.105001 0C878.73708 65.22349 958.70846 145.19487 958.70846 243.968477zM869.44546 288.599977c0-73.994248-60.119239-134.069485-134.113487-134.069485L288.667004 154.530492c-74.082252 0-134.071532 60.075237-134.071532 134.069485l0 446.797999c0 73.994248 59.989279 133.983527 134.071532 133.983527l446.664969 0c73.994248 0 134.113487-59.989279 134.113487-133.983527L869.44546 288.599977zM422.647461 302.560944l0 0.938372 89.307002-14.877849 0-0.021489L422.647461 302.560944z"
                    p-id="1005"></path>
                </svg>
              </el-icon>
              <template #title>Windows工具</template>
            </el-menu-item>
            <el-menu-item index="/extraTools">
              <el-icon size="24">
                <svg t="1668430758746" class="icon" viewBox="0 0 1024 1024" version="1.1"
                  xmlns="http://www.w3.org/2000/svg" p-id="1146" width="200" height="200">
                  <path
                    d="M377.558543 489.629503l67.034742 66.925249-134.025483 133.787052-67.014276-66.838268L377.558543 489.629503zM556.630988 690.733731l223.376487 0 0-89.436962L556.630988 601.296769 556.630988 690.733731zM958.70846 243.936755l0 536.105001c0 98.730629-79.927377 178.744987-178.700985 178.744987L243.903497 958.786743c-98.599645 0-178.61298-80.013335-178.61298-178.744987L65.290517 243.936755c0-98.730629 80.013335-178.722474 178.61298-178.722474l536.105001 0C878.781082 65.213257 958.70846 145.205103 958.70846 243.936755zM869.44546 288.612257c0-74.015737-60.119239-133.981481-134.113487-133.981481L288.578999 154.630776c-73.994248 0-134.027529 59.965743-134.027529 133.981481l0 446.775486c0 74.015737 60.033281 133.961014 134.027529 133.961014l446.752973 0c73.994248 0 134.113487-59.945277 134.113487-133.961014L869.44546 288.612257zM243.553527 355.842451l134.005017 133.787052 0.021489-0.021489 67.01223 66.946738 67.100234-66.946738L310.567803 289.00623 243.553527 355.842451z"
                    p-id="1147"></path>
                </svg>
              </el-icon>
              <template #title>开发者工具</template>
            </el-menu-item>
          </el-menu>
          <div class="aside-tool">
            <el-button type="danger" v-if="!isCollapse" circle @click="handleCollapse">
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
  -webkit-user-select: none;

  --header-height: 32px;
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
    font-weight: bold;
    align-items: center;

    .header-avatar {
      margin: 0 4px;
    }
  }

  .header-message {
    font-size: small;
    // border: 1px solid #fff;
    // border-style: none solid;
    padding: 0 8px;
  }

  .el-header,
  .header-message,
  .header-avatar,
  .header-title {
    -webkit-app-region: drag;
  }

  .header-tool {
    -webkit-app-region: none;
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

  .el-menu {
    background: transparent;
  }

  .el-main {
    background-color: var(--el-color-primary-light-9);
    color: var(--el-text-color-primary);
    height: calc(100vh - var(--header-height));
    overflow: hidden;
    padding: 0;
  }


  // 滚动条宽度
  div::-webkit-scrollbar {
    width: 6px;
    height: 6px;
  }

  // 滚动条轨道
  div::-webkit-scrollbar-track {
    background: rgb(239, 239, 239);
    border-radius: 2px;
  }

  // 小滑块
  div::-webkit-scrollbar-thumb {
    background: #40a0ff49;
    border-radius: 10px;
  }

  div::-webkit-scrollbar-thumb:hover {
    background: #40a0ff;
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
  setup() {

  },
  mounted() {
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
