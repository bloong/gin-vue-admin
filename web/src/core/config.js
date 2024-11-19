/**
 * 网站配置文件
 */
const greenText = (text) => `\x1b[32m${text}\x1b[0m`;

const config = {
  appName: 'EMS3000',
  appLogo: 'logo.png',
  showViteLogo: true,
  logs: [],
}

export const viteLogo = (env) => {
  if (config.showViteLogo) {
    console.log(greenText(`> 当前版本:v2.7.6`));
    console.log(greenText(`> 默认自动化文档地址:http://127.0.0.1:${env.VITE_SERVER_PORT}/swagger/index.html`));
    console.log(greenText(`> 默认前端文件运行地址:http://127.0.0.1:${env.VITE_CLI_PORT}`));
    console.log('\n');
  }
}

export default config
