/*
 * gin-vue-admin web框架组
 *
 * */
// 加载网站配置文件夹
import { register } from './global'

export default {
  install: (app) => {
    register(app)
    console.log(`
<<<<<<< HEAD
       欢迎使用 EMS3000
       当前版本:v2.7.6
=======
       欢迎使用 Gin-Vue-Admin
       当前版本:v2.7.7
       加群方式:微信：shouzi_1994 QQ群：622360840
       项目地址：https://github.com/flipped-aurora/gin-vue-admin
       插件市场:https://plugin.gin-vue-admin.com
       GVA讨论社区:https://support.qq.com/products/371961
>>>>>>> origin/main
       默认自动化文档地址:http://127.0.0.1:${import.meta.env.VITE_SERVER_PORT}/swagger/index.html
       默认前端文件运行地址:http://127.0.0.1:${import.meta.env.VITE_CLI_PORT}
  
    `)
  }
}
