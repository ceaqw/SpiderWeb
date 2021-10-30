'use strict'
const path = require('path')

function resolve(dir) {
    return path.join(__dirname, dir)
}

const name = process.env.VUE_APP_TITLE || 'Nint Spider Monitor' // 网页标题

const port = process.env.port || process.env.npm_config_port || 8001 // 端口

module.exports = {
    publicPath: process.env.NODE_ENV === "production" ? "/" : "/",
    lintOnSave: true,
    productionSourceMap: false,
    devServer: {
        port: port, // 端口号
        proxy: {
            "^/api": {
                target: `http://localhost:8000`,
                changeOrigin: true,
                ws: false,
                pathRewrite: {
                    '^/api': '/api'
                }
            }
        }
    },
    configureWebpack: {
        name: name,
        resolve: {
            alias: {
                '@': resolve('src')
            }
        }
    },
}