<template>
    <div class="manage-menu">
        <n-dropdown class="manage-menu-dropdown" trigger="click" :options="options" @select="handleSelect"
            :onUpdateShow="updateShow">
            <div :class="['manage-box-title', showDropdown ? 'menu-visible' : '']">
                <div class="avatar-box">
                    <n-avatar round :size="36" src="https://07akioni.oss-cn-beijing.aliyuncs.com/07akioni.jpeg" />
                    <span class="menu-user-name">{{ userStore.info.email }}</span>
                    <p class="manage-box-arrow">
                        <n-icon size="12" :component="ChevronUpOutline" />
                        <n-icon size="12" :component="ChevronDownOutline" />
                    </p>
                </div>
            </div>
        </n-dropdown>
        <ul class="scrollbar-overlay">
            <li :class="classNameRoute('dashboard')" @click="toggleRoute('/manage/dashboard')"
                >
                <i class="taskfont">&#xe6fb;</i>{{ $t("仪表盘") }}
            </li>
            <li :class="classNameRoute('calendar')" @click="toggleRoute('/manage/calendar')"
               ><i class="taskfont">&#xe6f5;</i>{{ $t("日历") }}</li>
            <li :class="classNameRoute('messenger')" @click="toggleRoute('/manage/messenger')"
               ><i class="taskfont">&#xe6eb;</i>{{ $t("消息") }}</li>
            <li :class="classNameRoute('file')" @click="toggleRoute('/manage/file')"
                ><i class="taskfont">&#xe6f3;</i>{{ $t("文件") }}</li>
            <li :class="classNameRoute('microapp', menu)" v-for="menu in getPluginMenu" @click="goToMicroapp(menu)"><i
                    class="taskfont">&#xe777;</i>{{ menu.title }}</li>
            <li class="menu-project">
                <n-scrollbar>
                    <ul class="scrollbar-overlay">
                        <li @contextmenu="handleContextMenu('1', $event)">1231232</li>
                        <n-dropdown placement="bottom-start" trigger="manual" :x="xRef" :y="yRef" :options="optionsUp"
                :show="showDropdownRef" :on-clickoutside="onClickoutside" @select="handleSelect" />
                    </ul>
                </n-scrollbar>
            </li>
        </ul>
    </div>
</template>

<script setup lang="ts">
import { ref, nextTick } from "vue"
import { ChevronUpOutline, ChevronDownOutline } from "@vicons/ionicons5"
import { useRoute, useRouter } from "vue-router"
import { UserStore } from "@/store/user"
import { GlobalStore } from "@/store"
import utils from "@/utils/utils"

const userStore = UserStore()
const globalStore = GlobalStore()
const router = useRouter()
const route = useRoute()
const showDropdown = ref<Boolean>(false)
const xRef = ref(0)
const yRef = ref(0)
const showDropdownRef = ref(false)
const taskBrowseChildren = ref([{ label: $t("暂无打开记录"), key: '1', disabled: true }])
const options = [
    { label: $t("最近打开的任务"), key: "taskBrowse", children: taskBrowseChildren.value },
    { type: 'divider', key: '' },
    { label: $t("个人设置"), key: "PersonalSettings" },
    { label: $t("系统设置"), key: "systemSettings" },
    { type: 'divider', key: '' },
    { label: $t("清除缓存"), key: "clearCache" },
    { label: $t("退出登录"), key: "Logout" },
]

const optionsUp = [{ label: $t("置顶该项目"), key: "up" }]



// 下拉选中事件
const handleSelect = (event: any, option: any) => {
    switch (event) {
        case "Logout":
            router.push("/login")
            userStore.setUserInfo()
            break
        case "clearCache":
            utils.IDBSet("clearCache", "handle").then(_ => {
                utils.reloadUrl()
            });
            break
        case "systemSettings":
            router.push("/manage/setting/system")
            break
        case "PersonalSettings":
            router.push("/manage/setting/personal")
            break
    }
}

//更新下拉显示
const updateShow = (e: any) => {
    showDropdown.value = e
}

//菜单事件
const handleContextMenu = (id: any, e: any) => {
    e.preventDefault()
    showDropdownRef.value = false
    nextTick().then(() => {
        showDropdownRef.value = true
        xRef.value = e.clientX
        yRef.value = e.clientY
    })
}

//点击下拉
const onClickoutside = () => {
    showDropdownRef.value = false
}

// 获取插件菜单
const getPluginMenu = computed(() => {
    return globalStore.pluginConfig.menu.filter((item: any) => item.platform == 'desktop')
})

// 去往微应用
const goToMicroapp = (menu: any) => {
    router.push(`/manage/microapp/${menu.code}/${menu.id}#/`)
}

// 当前激活页面
const classNameRoute = (path: any, menu: any = {}) => {
    if (path == "microapp") {
        return {
            "active": route.path === `/manage/microapp/${menu.code}/${menu.id}`,
        };
    } else {
        return {
            "active": route.name === `manage-${path}`,
        };
    }
}

// 路由跳转
const toggleRoute = (path: any) => {
    router.push(path)
}
</script>

<style lang="less" scoped>
.manage-menu {
    @apply flex flex-col h-full w-255 bg-bg-manage-menu duration-300 ease-in-out;

    .manage-box-title {
        @apply w-86p m-auto mt-27 rounded-lg overflow-hidden;

        .avatar-box {
            @apply flex items-center cursor-pointer bg-bg-login-box py-6 px-10 duration-300 ease-in-out;

            .menu-user-name {
                @apply flex-1 text-16 font-semibold truncate pl-12;
            }

            .manage-box-arrow {
                @apply flex flex-col pl-16;

                i {
                    margin: -1px;
                }
            }
        }
    }

    .menu-visible {
        box-shadow: 0 1px 6px rgba(0, 0, 0, 0.2);
    }

    .scrollbar-overlay {
        @apply flex-1 flex flex-col overflow-hidden mt-16 w-full;

        >li {
            @apply items-center rounded text-manage-menu-color flex cursor-pointer flex-shrink-0 h-36 m-auto my-5 px-4p max-w-full w-72p;

            >i {
                @apply opacity-30 text-20 mr-10;
            }
        }

        >li:first-child {
            @apply mt-12;
        }

        >.active {
            @apply bg-white;
        }

        >li.menu-project {
            @apply items-center cursor-default flex flex-1 flex-col pt-12 pb-0 px-0 w-full;

            .scrollbar-overlay {
                @apply m-0;

                >li {
                    @apply flex items-start flex-col list-none cursor-pointer w-72p m-auto my-2;
                }
            }
        }
    }
}
</style>
<style>
.manage-menu-dropdown {
    @apply w-219;
}
</style>