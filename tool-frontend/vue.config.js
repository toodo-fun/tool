module.exports = {
  pluginOptions: {
    electronBuilder: {
      builderOptions: {
        appId: 'fun.toodo.tool',
        productName: '兔儿工具箱',
        copyright: 'Copyright © 2022',
        directories: {
          output: './dist_electron'
        },
        extraResources: {
          from: './server',
          to: 'server'
        },
        win: {
          icon: './public/favicon.ico',
          target: [
            {
              target: 'nsis',
              arch: [ 'x64' ]
            }
          ]
        },
        nsis: {
          oneClick: false,
          allowElevation: true,
          allowToChangeInstallationDirectory: true
        }
      },
      preload: './src/preload.js'
    }
  },
  chainWebpack: (config) => {
    config.plugin("html").tap((args) => {
      args[0].title = "兔儿工具箱";
      return args;
    });
  },
}