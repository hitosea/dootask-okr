import vue from '@vitejs/plugin-vue'
import {defineConfig, loadEnv} from 'vite'
import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import {NaiveUiResolver} from 'unplugin-vue-components/resolvers'
import * as path from 'path';

export default defineConfig(({command, mode}) => {
    const env = loadEnv(mode, process.cwd(), '')
    const devPort: any = env['DEV_PORT'] || 5567
    const devProxyTarget: string = env['DEV_PROXY_TARGET'] || 'http://127.0.0.1:5566'
    const viteApiUrl: string = env['VITE_API_URL'] || '/'
    const appname: any = '/manage/microapp/okr/'

    return {
        base: `${process.env.NODE_ENV === 'production' ? '' : ''}${appname}`,
        server: {
            host: '0.0.0.0',
            port: devPort,
            proxy: {
                [viteApiUrl]: {
                    target: devProxyTarget,
                    changeOrigin: true,
                }
            }
        },
        resolve: {
            extensions: ['.mjs', '.js', '.ts', '.jsx', '.tsx', '.json', '.vue'],
            alias: [
                {
                    find: '@',
                    replacement: path.resolve(__dirname, './src/'),
                },
                {
                    find: 'vue-i18n',
                    replacement: 'vue-i18n/dist/vue-i18n.cjs.js',
                }
            ],
        },
        plugins: [
            vue({
                template: {
                    compilerOptions: {
                        isCustomElement: (tag) => tag.includes('DialogWrappers')
                    }
                }
            }),
            AutoImport({
                imports: [
                    'vue',
                    {
                        'naive-ui': [
                            'useDialog',
                            'useMessage',
                            'useNotification',
                            'useLoadingBar'
                        ]
                    }
                ]
            }),
            Components({
                resolvers: [NaiveUiResolver()]
            }),
            // 微应用处理
            (function () {
                let basePath = ''
                return {
                    name: "vite:micro-app",
                    apply: 'build',
                    configResolved(config:any) {
                        basePath = `${config.base}${config.build.assetsDir}/`
                    },
                    renderChunk(code:any, chunk:any) {
                        if (chunk.fileName.endsWith('.js')) {
                            code = code.replace(/(from|import\()(\s*['"])(\.\.?\/)/g, (all, $1, $2, $3) => {
                                if( basePath.indexOf('http') == -1){
                                    let str = all.replace($3, basePath)
                                    if( str.indexOf('import(') !== -1){
                                        str = str.replace('import(','import(window.location.origin+')
                                    }
                                    return str
                                }else{
                                    return all.replace($3, new URL($3, basePath))
                                }
                            })
                        }
                        return code
                    },
                }
            })() as any,
        ],
        build: {
            chunkSizeWarningLimit: 3000
        }
    }
})
