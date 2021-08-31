module.exports = {
    lintOnSave: false,
    publicPath: process.env.VUE_BASE_URL || '/',
    productionSourceMap: false,
    pluginOptions: {
        electronBuilder: {
          preload: 'src/preload.js',
        }
    },
    transpileDependencies: [
        'resize-detector' // vue-echarts
    ]
}
