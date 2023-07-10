<template>
    <div class="page-setting manage-box-view">
        <div class="setting-head">
            <div class="setting-titbox">
                <div class="setting-title">
                    <h2>{{ $t('设置') }}</h2>
                </div>
            </div>
        </div>
        <div class="setting-box">
            <div class="setting-menu">
                <ul>
                    <li v-for="(item, key) in menu" :key="key" :class="classNameRoute(item.path, item.divided)"
                        @click="toggleRoute(item.path)">
                        {{ item.name }}
                    </li>
                </ul>
            </div>
            <div class="setting-content">
                <div class="setting-content-title">{{ titleNameRoute }}</div>
                <div class="setting-content-view">
                    <router-view></router-view>
                </div>
            </div>
        </div>
    </div>
</template>
<script lang="ts" setup>
import { computed } from "vue"
import { useRoute, useRouter } from "vue-router"
import { GlobalStore } from "@/store/index"
import utils from "@/utils/utils";
import { UserStore } from "@/store/user"

const userStore = UserStore()
const globalStore = GlobalStore()
const route = useRoute()
const router = useRouter()

const menu = [
    { path: "/manage/setting/personal", name: $t('个人设置'), divided: false },
    { path: "/manage/setting/language", name: $t('语言设置'), divided: false },
    { path: "/manage/setting/theme", name: $t('主题设置'), divided: false },
    { path: "/manage/setting/password", name: $t('密码设置'), divided: false },
    { path: "/manage/setting/email", name: $t('修改邮箱'), divided: false },
    { path: "/manage/setting/system", name: $t('系统设置'), divided: true },
    { path: "clearCache", name: $t('清除缓存'), divided: true },
    { path: "logout", name: $t('退出登录'), divided: false },
]

// 当前页面名称
const titleNameRoute = computed(() => {
    let name = ""
    menu.some((item) => {
        if (route.path === item.path) {
            name = `${item.name}`
            return true
        }
    })
    return name || $t('设置')
})

// 路由跳转
const toggleRoute = (path: any) => {
    if (path == 'clearCache') {
        utils.IDBSet("clearCache", "handle").then(_ => {
            utils.reloadUrl()
        });
        return
    }
    if (path == 'logout') {
        router.push("/manage/setting/personal")
            userStore.setUserInfo()
        return
    }
    router.push(path)
}

// 当前激活页面
const classNameRoute = (path: any, divided: any) => {
    return {
        "active": globalStore.windowLandscape && route.path === `${path}`,
        "divided": !!divided
    }
}
</script>
<style lang="less" scoped>
.page-setting {
    @apply flex flex-col;

    .setting-head {
        @apply flex items-center border-solid border-0 border-b border-border-color mt-32 mb-16 mr-32 ml-32 duration-300 ease-in-out;

        .setting-titbox {
            @apply flex-1 mb-16;

            h2 {
                @apply text-title-color flex-1 text-28 font-semibold;
            }
        }
    }

    .setting-box {
        @apply flex flex-1 pb-16;

        .setting-menu {
            @apply border-solid border-0 border-r border-border-color w-200 overflow-auto duration-300 ease-in-out;

            ul {
                @apply pt-12 pl-32;

                li {
                    @apply text-text-li cursor-pointer list-none my-5 px-20 relative leading-42 truncate duration-300 ease-in-out;

                    &:hover {
                        @apply bg-bg-manage-menu;
                    }
                }

                li.active {
                    @apply bg-bg-manage-menu;
                }

                li.divided {
                    @apply relative pt-10 mt-10;
                }

                li.divided::before {
                    @apply absolute top-0 left-0 right-0 z-10 h-1 bg-bg-manage-menu duration-300 ease-in-out;
                    content: " ";
                }

                li.divided::after {
                    @apply absolute top-1 left-0 right-0 z-10 h-9 bg-white duration-300 ease-in-out;
                    content: " ";
                }
            }
        }

        .setting-content {
            @apply flex flex-1 flex-col overflow-auto relative;

            .setting-content-title {
                @apply text-20 font-medium py-12 px-32;
            }

            .setting-content-view {
                @apply flex-1 relative;

                .setting-item {
                    @apply flex flex-col p-0 absolute bottom-0 top-0 right-0 left-0;

                    :deep(.n-form) {
                        @apply flex-1 overflow-auto py-24 px-40;

                        .n-input-group {
                            @apply max-w-460;
                        }

                        .n-input {
                            @apply max-w-460;
                        }

                        .n-select {
                            @apply max-w-460;
                        }
                    }

                    :deep(.setting-footer) {
                        @apply flex gap-x-2 flex-shrink-0 pt-16 px-24 static border-solid border-0 border-t border-border-color duration-300 ease-in-out;
                    }
                }
            }
        }
    }
}</style>
