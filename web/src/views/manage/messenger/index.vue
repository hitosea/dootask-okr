<template >
    <div class="page-messenger">
        <div class="messenger-wrapper">
            <div class="messenger-select">
                <div class="messenger-search">
                    <div class="search-wrapper">
                        <n-input v-if="tabActive === 'dialog'" v-model:value="dialogSearchKey" type="text"
                            :placeholder="$t(loadDialogs ? '更新中...' : '搜索消息')" clearable>
                            <template #prefix>
                                <Loading v-if="loadDialogs || dialogSearchLoad > 0" />
                                <n-icon v-else :component="SearchOutline" />
                            </template>
                        </n-input>
                        <n-input v-else v-model:value="contactsKey" type="text" :placeholder="$t('搜索联系人')" clearable>
                            <template #prefix>
                                <n-icon :component="SearchOutline" />
                            </template>
                        </n-input>
                    </div>
                </div>
                <div class="messenger-nav" v-if="tabActive === 'dialog' && !dialogSearchKey">
                    <n-dropdown :options="dialogMenus" placement="bottom-start" trigger="click">
                        <div class="nav-icon"><i class="taskfont">&#xe634;</i></div>
                    </n-dropdown>
                    <div v-for="(item, key) in typeItems" :key="key" :class="{ active: dialogActive == item.key }"
                        @click="onActive(item.key)">
                        <div class="nav-title">
                            <em>{{ $t(item.label) }}</em>
                            <n-badge class="nav-num" :max="999" :value="msgUnread(item.key)" />
                        </div>
                    </div>
                </div>
                <!-- <div v-if="isEEUiApp && !appNotificationPermission" class="messenger-notify-permission"
                    @click="onOpenAppSetting">
                    {{ $t('未开启通知权限') }}<i class="taskfont">&#xe733;</i>
                </div> -->
                <n-scrollbar :on-scroll="listScroll">
                    <ul v-if="tabActive === 'dialog'" class="dialog">
                        <li>
                            <n-avatar class="img-avatar" round :size="42"
                                src="https://07akioni.oss-cn-beijing.aliyuncs.com/07akioni.jpeg" />
                            <div class="dialog-box">
                                <div class="dialog-title"><span>tes***@dootask.com</span><i
                                        class="ivu-icon ivu-icon-md-done-all"></i><em>04-24</em></div>
                                <div class="dialog-text no-dark-content">
                                    <div class="last-self">你</div>
                                    <div class="last-text"><span>关注一下消息哈</span></div>
                                </div>
                            </div>
                        </li>
                    </ul>
                    <ul v-else class="contacts">

                        <li>
                            <div class="label">#</div>
                            <ul>
                                <li>
                                    <div class="avatar">
                                        <n-avatar class="img-avatar" round :size="30"
                                            src="http://127.0.0.1:2222/avatar/869%2A%2A%2A%40qq.com.png" />
                                    </div>
                                    <div class="nickname"><em>869***@qq.com</em>
                                        <div class="tags"><span>啊啊啊负责人</span></div>
                                    </div>
                                </li>
                                <li>
                                    <div class="avatar">
                                        <n-avatar class="img-avatar" round :size="30"
                                            src="http://127.0.0.1:2222/avatar/869%2A%2A%2A%40qq.com.png" />
                                    </div>
                                    <div class="nickname">
                                        <em>6***@qq.com</em>
                                        <div class="tags"></div>
                                    </div>
                                </li>
                            </ul>
                        </li>

                        <li class="loaded">共4位联系人</li>
                    </ul>
                </n-scrollbar>
                <div class="messenger-menu">
                    <div class="menu-icon">
                        <Icon @click="onActive(null)" :class="{ active: tabActive === 'dialog' }" type="ios-chatbubbles" />
                        <n-icon @click="onActive(null)" :class="{ active: tabActive === 'dialog' }" size="24"
                            :component="ChatbubblesSharp" />
                    </div>
                    <div class="menu-icon">
                        <n-icon @click="tabActive = 'contacts'" :class="{ active: tabActive === 'contacts' }" size="24"
                            :component="Person" />
                    </div>
                </div>
            </div>
            <div class="messenger-msg" v-if="activeNum > 0 && routeName === 'manage-messenger'">
                <div class="msg-dialog-bg">
                    <div class="msg-dialog-bg-icon">
                        <n-icon  size="46" :component="ChatbubblesSharp" />
                    </div>
                    <div class="msg-dialog-bg-text">{{$t('选择一个会话开始聊天')}}</div>
                </div>
                <dialog-wrapper />
            </div>
        </div>
    </div>
</template>
<script setup lang="ts">
import { ref, computed  } from "vue"
import { DiaLogStore } from "@/store/dialog"
import { SearchOutline, ChatbubblesSharp, Person } from "@vicons/ionicons5"
import Loading from "@/components/Loading.vue"
import utils from "@/utils/utils";
import webTs from "@/utils/web";
import { useRoute } from "vue-router"
import DialogWrapper from "./components/DialogWrapper.vue"


const routeName = useRoute().name
const activeNum = ref(1)
const tabActive = ref('dialog')
const dialogSearchKey = ref('')
const loadDialogs = DiaLogStore().loadDialogs
const dialogId = DiaLogStore().dialogId
const cacheDialogs = DiaLogStore().cacheDialogs
// const appNotificationPermission = diaLogStore().appNotificationPermission
const dialogSearchLoad = ref(0)
const contactsKey = ref('')
const MessengerObject = ref({ menuHistory: [] })
let dialogHistory = MessengerObject.value.menuHistory
const dialogActive = ref('')
const dialogSearchList = ref([])
const operateVisible = ref(false)
const operateItem = ref({ id: 0 })
const operateStyles = ref({})

const dialogMenus = ref([
    { key: '', label: $t('全部') },
    { key: 'project', label: $t('项目') },
    { key: 'task', label: $t('任务') },
    { key: 'user', label: $t('单聊') },
    { key: 'group', label: $t('群聊') },
    { key: 'bot', label: $t('机器人') },
])

// 选项的列表
const typeItems = computed(() => {
    const types = []
    if (dialogHistory.includes(dialogActive.value)) {
        types.push(...dialogHistory)
    } else {
        types.push('')
        if (dialogActive.value) {
            types.push(dialogActive.value)
        }
        dialogHistory.some(item => {
            if (!types.includes(item)) {
                types.push(item)
            }
        });
        ['project', 'task', 'user'].some(item => {
            if (!types.includes(item)) {
                types.push(item)
            }
        })
        dialogHistory = types.slice(0, 4)
        utils.IDBSave("dialogMenuHistory", dialogHistory)

        //
        return dialogHistory.map(key => {
            return dialogMenus.value.find(item => item.key == key)
        })
    }
})

const listScroll = (res) => {
    // if (res.scrollE < 10) {
    //     this.getContactsNextPage()
    // }
    // 判断翻页
    operateVisible.value = false;
}


//未读消息
const msgUnread = computed(() => {
    return function (type) {
        let num = 0
        this.cacheDialogs.some((dialog) => {
            switch (type) {
                case 'project':
                case 'task':
                    if (type != dialog.group_type) {
                        return false
                    }
                    break;
                case 'user':
                    if (type != dialog.type || dialog.bot) {
                        return false
                    }
                    break;
                case 'group':
                    if (type != dialog.type || ['project', 'task'].includes(dialog.group_type)) {
                        return false
                    }
                    break;
                case 'bot':
                    if (!dialog.bot) {
                        return false;
                    }
                    break;
            }
            num += webTs.getDialogNum(dialog);
        });
        return num;
    }
})


//聊天列表 
const dialogList = computed(() => {
    if (dialogSearchList.value.length > 0) {
        return dialogSearchList.value.sort((a, b) => {
            // 搜索结果排在后面
            return (a.is_search === true ? 1 : 0) - (b.is_search === true ? 1 : 0)
        })
    }
    if (dialogActive.value == '' && dialogSearchKey.value == '') {
        return cacheDialogs.filter(dialog => filterDialog(dialog)).sort(dialogSort);
    }
    const list = cacheDialogs.filter(dialog => {
        if (filterDialog(dialog)) {
            return false;
        }
        if (dialogSearchKey.value) {
            const { name, pinyin, last_msg } = dialog;
            let searchString = `${name} ${pinyin}`
            if (last_msg) {
                switch (last_msg.type) {
                    case 'text':
                        searchString += ` ${last_msg.msg.text.replace(/<[^>]+>/g, "")}`
                        break
                    case 'meeting':
                    case 'file':
                        searchString += ` ${last_msg.msg.name}`
                        break
                }
            }
            if (!utils.strExists(searchString, dialogSearchKey.value)) {
                return false;
            }
        } else if (dialogActive.value) {
            switch (dialogActive.value) {
                case 'project':
                case 'task':
                    if (dialogActive.value != dialog.group_type) {
                        return false;
                    }
                    break;
                case 'user':
                    if (dialogActive.value != dialog.type || dialog.bot) {
                        return false;
                    }
                    break;
                case 'group':
                    if (dialogActive.value != dialog.type || ['project', 'task'].includes(dialog.group_type)) {
                        return false;
                    }
                    break;
                case 'bot':
                    if (!dialog.bot) {
                        return false;
                    }
                    break;
                default:
                    return false;
            }
        }
        return true;
    })
    return list.sort(dialogSort)
})

// 点击菜单
const onActive = (key) => {
    if (key === null) {
        if (tabActive.value !== 'dialog') {
            tabActive.value = 'dialog'
            return;
        }
        key = dialogActive.value
    }
    if (dialogActive.value == key) {
        shakeUnread()  // 再次点击滚动到未读条目
    } else {
        dialogActive.value = key
    }
}

// 点击滚动到未读条目
const shakeUnread = () => {
    let index = dialogList.value.findIndex(dialog => webTs.getDialogNum(dialog) > 0)
    if (index === -1) {
        index = dialogList.value.findIndex(dialog => webTs.getDialogUnread(dialog, true) > 0)
    }
    if (index > -1) {
        // 需要重新写
        // const el = this.$refs[`dialog_${dialogList[index]?.id}`]
        // if (el && el[0]) {
        //     if (el[0].classList.contains("common-shake")) {
        //         return
        //     }
        //     utils.scrollIntoViewIfNeeded(el[0])
        //     requestAnimationFrame(_ => {
        //         el[0].classList.add("common-shake")
        //         setTimeout(_ => {
        //             el[0].classList.remove("common-shake")
        //         }, 600)
        //     })
        // }
    }
}

// 打开聊天
const openDialog = (dialogId) => {
    if (operateVisible.value) {
        return
    }
    //
    if (utils.isJson(dialogId) && utils.leftExists(dialogId.dialog_id, "u:")) {
        // this.$store.dispatch("openDialogUserid", $A.leftDelete(dialogId.dialog_id, "u:")).catch(({ msg }) => {
        //     $A.modalError(msg)
        // })
    } else {
        // this.$store.dispatch("openDialog", dialogId)
    }
}

//返回聊天的样式
const dialogClass = (dialog) => {
    if (dialogSearchKey.value) {
        return null
    }
    return {
        top: dialog.top_at,
        active: dialog.id == dialogId,
        operate: operateVisible.value && dialog.id == operateItem.value.id,
        completed: webTs.dialogCompleted(dialog)
    }
}

//聊天过滤
const filterDialog = (dialog: any) => {
    if ((dialog.id > 0 && dialog.id == dialogId) || dialog.top_at || dialog.todo_num > 0 || webTs.getDialogNum(dialog) > 0) {
        return true
    }
    if (dialog.name === undefined || dialog.dialog_delete === 1) {
        return false;
    }
    if (!dialog.last_at) {
        return false;
    }
    if (dialog.type == 'group') {
        if (['project', 'task'].includes(dialog.group_type) && utils.isJson(dialog.group_info)) {
            if (dialog.group_type == 'task' && dialog.group_info.complete_at) {
                // 已完成5天后隐藏对话
                let time = Math.max(Number(utils.Date(dialog.last_at, true)), Number(utils.Date(dialog.group_info.complete_at, true)))
                if (5 * 86400 + time < utils.Time()) {
                    return false
                }
            }
            if (dialog.group_info.deleted_at) {
                // 已删除2天后隐藏对话
                let time = Math.max(Number(utils.Date(dialog.last_at, true)), Number(utils.Date(dialog.group_info.deleted_at, true)))
                if (2 * 86400 + time < utils.Time()) {
                    return false
                }
            }
            if (dialog.group_info.archived_at) {
                // 已归档3天后隐藏对话
                let time = Math.max(Number(utils.Date(dialog.last_at, true)), Number(utils.Date(dialog.group_info.archived_at, true)))
                if (3 * 86400 + time < utils.Time()) {
                    return false
                }
            }
        }
    }
    return true;
}

// 聊天排序
const dialogSort = (a, b) => {
    // 根据置顶时间排序
    if (a.top_at || b.top_at) {
        return Number(utils.Date(b.top_at)) - Number(utils.Date(a.top_at));
    }
    // 根据未读数排序
    if (a.todo_num > 0 || b.todo_num > 0) {
        return b.todo_num - a.todo_num;
    }
    // 根据草稿排序
    if (a.extra_draft_has || b.extra_draft_has) {
        return b.extra_draft_has - a.extra_draft_has;
    }
    // 根据最后会话时间排序
    return Number(utils.Date(b.last_at)) - Number(utils.Date(a.last_at));
}

const onOpenAppSetting = () => {
    // EEUI App 专用
    // $A.eeuiAppSendMessage({
    //     action: 'gotoSetting',
    // });
}

const handleLongpress = (event, el) => {
    if (dialogSearchKey.value) {
        return;
    }
    const dialogId = utils.getAttr(el, 'data-id')
    const dialogItem = dialogList.value.find(item => item.id == dialogId)
    if (!dialogItem) {
        return
    }
    operateVisible.value = false;
    operateItem.value = utils.isJson(dialogItem) ? dialogItem : {};
    nextTick(() => {
        // 需要改
        // const dialogRect = el.getBoundingClientRect();
        // const wrapRect = $refs.list.$el.getBoundingClientRect();
        // operateStyles.value = {
        //     left: `${event.clientX - wrapRect.left}px`,
        //     top: `${dialogRect.top + this.windowScrollY}px`,
        //     height: dialogRect.height + 'px',
        // }
        operateVisible.value = true;
    })
}
</script>
<style lang="less">
.page-messenger {
    @apply absolute left-0 right-0 top-0 bottom-0 overflow-auto flex;

    .messenger-wrapper {
        @apply flex-1 flex items-start overflow-hidden;

        .messenger-select {
            @apply relative h-full w-30p min-w-240 max-w-320 flex-shrink-0 flex flex-col;

            &:after {
                @apply absolute right-0 top-0 h-full w-1 border-solid border-0 border-r border-border-color;
                content: " ";
            }

            .messenger-search {
                @apply flex items-center justify-center h-54 px-12 flex-shrink-0;

                .search-wrapper {
                    @apply flex-1 bg-bg-search px-8 mx-4 rounded-xl overflow-hidden;

                    .n-input,
                    .n-input:focus {
                        @apply bg-bg-none border-0;

                        .n-input__border,
                        .n-input__state-border {
                            @apply border-0 shadow-none;
                        }
                    }
                }
            }

            .messenger-nav {
                @apply flex items-center pt-2 pb-12 px-10;

                >div {
                    @apply flex flex-shrink-0 flex-1 items-center justify-center px-6 cursor-pointer text-messenger-nav not-italic leading-none font-light;

                    em {
                        @apply not-italic;
                    }
                }

                .nav-icon {
                    @apply max-w-34;
                }

                >div.active {
                    @apply font-medium;
                }
            }

            ul.dialog {
                li {
                    @apply flex flex-row items-start py-16 px-12 relative cursor-pointer list-none;

                    .dialog-box {
                        @apply flex-1 flex flex-col pl-12;

                        .dialog-title {
                            @apply flex flex-row items-center justify-between leading-6;

                            span {
                                @apply flex-1 text-title-color text-14 truncate;
                            }

                            em {
                                @apply flex-shrink-0 ml-8 not-italic text-12 text-text-tips;
                            }
                        }
                        .dialog-text{
                            @apply text-12 text-text-tips min-h-24 leading-24 flex items-center;
                            .last-self{
                                @apply flex-shrink-0 pr-4 mr-4 overflow-hidden relative;
                                &:after{
                                    @apply absolute top-1/2 right-0;
                                    content:':';
                                    transform: translateY(-50%);
                                }
                            }
                            .last-text{
                                @apply flex-1 flex items-center;
                                span{
                                    @apply flex-1 w-0 truncate;
                                }
                            }
                        }
                    }
                }
            }

            ul.contacts {
                >li {
                    @apply list-none ml-24 relative;
                    .label{
                        @apply pl-4 my-6 h-34 leading-34 relative;
                        &:after{
                            @apply absolute bottom-0 right-0 left-0 h-1 border-solid border-0 border-b border-border-color;
                            content: " "
                        }
                    }
                    ul {
                        li{
                            @apply list-none flex flex-row items-center h-52 cursor-pointer;
                            .avatar{
                                @apply flex-grow-0 flex-shrink-0 w-30 h-30;
                            }
                            .nickname{
                                @apply flex-1 w-0 pl-12 text-14 flex items-center justify-between;
                                em{
                                    @apply not-italic;
                                }
                                .tags{
                                    @apply pr-12 text-12 truncate text-text-tips;
                                }
                            }
                        }
                    }
                }
                >li.loaded{
                    @apply m-0 h-52 flex items-center justify-center;
                }
            }

            .messenger-menu {
                @apply flex items-center justify-center h-52 flex-shrink-0 border-solid border-0 border-t border-border-color;

                .menu-icon {
                    @apply h-full flex items-center relative;

                    >i {
                        @apply text-menu-color cursor-pointer mx-24 opacity-90;
                    }

                    >i.active {
                        @apply text-primary-color;
                    }
                }
            }
        }

        .messenger-msg{
            @apply flex w-0 h-full flex-1 relative;
            .msg-dialog-bg{
                @apply absolute top-0 bottom-0 right-0 left-0 z-0 flex-1 flex flex-col items-center justify-center;
                .msg-dialog-bg-icon{
                    @apply bg-bg-manage-menu p-20 rounded-full flex;
                    i{
                        @apply text-messenger-msg-icon;
                    }
                }
                .msg-dialog-bg-text{
                    @apply mt-16 bg-bg-manage-menu px-15 py-4 rounded-2xl text-14 text-messenger-msg-icon;
                }
            }
        }
    }
}
</style>