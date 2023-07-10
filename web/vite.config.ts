import vue from '@vitejs/plugin-vue'
import {defineConfig, loadEnv} from 'vite'
import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import {NaiveUiResolver} from 'unplugin-vue-components/resolvers'
import * as path from 'path';

export default defineConfig(({command, mode}) => {
    const env = loadEnv(mode, process.cwd(), '')
    const devPort: any = env['DEV_PORT'] || 4444
    const devProxyTarget: string = env['DEV_PROXY_TARGET'] || 'http://127.0.0.1:5566'
    const viteApiUrl: string = env['VITE_API_URL'] || '/'

    return {
        base: '/',
        server: {
            port: devPort,
            proxy: {
                [`${viteApiUrl}/ws`]: {
                    target: devProxyTarget,
                    changeOrigin: true,
                    ws: true
                },
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
                        isCustomElement: (tag) => tag.includes('micro-app')
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
            })
        ],
        build: {
            chunkSizeWarningLimit: 3000
        }
    }
})
