<template >
    <div class="dialog-wrapper">
        <div class="dialog-nav">
            <div class="nav-wrapper">
                <div class="dialog-block">
                    <div class="dialog-avatar">
                        <n-avatar class="img-avatar" round :size="42"
                            src="https://07akioni.oss-cn-beijing.aliyuncs.com/07akioni.jpeg" />
                    </div>

                    <div class="dialog-title">
                        <div class="main-title">
                            <h2>任务提醒</h2>
                        </div>
                        <ul class="title-tags">
                            <li class="msg active">
                                <i class="no-dark-content"></i>
                                <span>消息</span>
                            </li>
                            <li class="task">
                                <i class="no-dark-content"></i>
                                <span>打开任务</span>
                            </li>
                        </ul>
                    </div>
                </div>

                <n-dropdown trigger="click" placement="bottom-end" :options="options" @select="handleSelect">
                    <i class="taskfont dialog-menu-icon">&#xe6e9;</i>
                </n-dropdown>
            </div>
        </div>
        <div class="dialog-scroller">

        </div>
        <!--底部输入-->
        <div ref="footer" class="dialog-footer" :class="footerClass" :style="footerStyle" @click="onActive">
            <div class="dialog-newmsg" @click="onToBottom">{{ $t(`有${msgNew}条新消息`) }}</div>
            <div class="dialog-goto" @click="onToBottom"><i class="taskfont">&#xe72b;</i></div>
            <DialogUpload ref="chatUpload" class="chat-upload" :dialog-id="props.dialogId"
                @on-success="chatFile('success', $event)" @on-error="chatFile('error', $event)"
                @on-progress="chatFile('error', $event)" />
            <div v-if="todoShow" class="chat-bottom-menu">
                <div class="bottom-menu-label">{{ $t('待办') }}:</div>
                <ul class="scrollbar-hidden">
                    <li v-for="item in todoList" @click.stop="onViewTodo(item)">
                        <div class="bottom-menu-desc no-dark-content">{{ webTs.getMsgSimpleDesc(item.msg_data) }}</div>
                    </li>
                </ul>
            </div>
            <div v-else-if="quickShow" class="chat-bottom-menu">
                <ul class="scrollbar-hidden">
                    <li v-for="item in quickMsgs" @click.stop="sendQuick(item)">
                        <div class="bottom-menu-desc no-dark-content" :style="item.style || null">{{ item.label }}</div>
                    </li>
                </ul>
            </div>
            <div v-if="isMute" class="chat-mute">
                {{ $t('禁言发言') }}
            </div>
        </div>
    </div>
</template>
<script setup lang="ts">
defineOptions({
    name: "DialogWrapper"
})

import { ref, defineEmits } from "vue"
import { DiaLogStore } from "@/store/dialog"
import { GlobalStore } from "@/store/index"
import utils from "@/utils/utils";
import DialogUpload from "./DialogUpload.vue";
import { UserStore } from "@/store/user";
import webTs from "@/utils/web";


const allMsgs = ref([])
const tempMsgs = ref([])
const msgType = ref('')
const msgText = ref('')
const msgNew = ref(0)
const scrollTail = ref(0)
const todoViewMid = ref(0)
const todoViewId = ref(0)
const operateVisible = ref(false)
const todoViewShow = ref(false)
const scroller = ref(null)
const input = ref(null)
const windowScrollY = GlobalStore().windowScrollY
const keyboardType = GlobalStore().keyboardType
const keyboardHeight = GlobalStore().keyboardHeight
const cacheDialogs = DiaLogStore().cacheDialogs
const dialogTodos = DiaLogStore().dialogTodos
const tempId = ref(utils.randNum(1000000000, 9999999999))
const userId = UserStore().info.id
const userIsAdmin = UserStore().info.userIsAdmin

const props = defineProps({
    dialogId: {
        type: Number,
        default: 0
    },
    msgId: {
        type: Number,
        default: 0
    },
    autoFocus: {
        type: Boolean,
        default: false
    },
    beforeBack: Function
})

const options = ref([
    { label: $t('搜索消息'), key: 'searchMsg' },
    { label: $t('修改资料'), key: 'modifyNormal' },
    { label: $t('创建群组'), key: 'openCreate' },
    { label: $t('群组设置'), key: 'groupInfo' },
    { label: $t('修改资料'), key: 'modifyAdmin' },
    { label: $t('退出群组'), key: 'exit' },
    { label: $t('修改资料'), key: 'modifyNormal' },
    { label: $t('转让群主'), key: 'transfer' },
    { label: $t('解散群组'), key: 'disband' },

])

const emit = defineEmits(['on-active'])

const dialogData = computed(() => {
    return cacheDialogs.find(({ id }) => id == props.dialogId) || {};
})

const quoteId = computed(() => {
    if (props.msgId > 0) {
        return props.msgId
    }
    return dialogData.value.extra_quote_id || 0
})

const footerClass = computed(() => {
    if (msgNew.value > 0 && allMsgs.value.length > 0) {
        return 'newmsg'
    }
    if (scrollTail.value > 500) {
        return 'goto'
    }
    return null
})

const footerStyle = computed(() => {
    const style = {
        paddingBottom: ''
    };
    if (windowScrollY === 0
        && keyboardType === "show"
        && keyboardHeight > 0
        && keyboardHeight < 120) {
        style.paddingBottom = (keyboardHeight) + 'px';
    }
    return style;
})

const isMute = computed(() => {
    if (dialogData.value.group_type === 'all') {
        if (dialogData.value.all_group_mute === 'all') {
            return true
        } else if (dialogData.value.all_group_mute === 'user') {
            if (!userIsAdmin) {
                return true
            }
        }
    }
    return false
})

const todoList = computed(() => {
    if (!dialogData.value.todo_num) {
        return []
    }
    return dialogTodos.filter(item => !item.done_at && item.dialog_id == props.dialogId).sort((a, b) => {
        return b.id - a.id;
    });
})

const todoShow = computed(() => {
    return todoList.value.length > 0 && windowScrollY === 0 && quoteId.value === 0
})

const quoteUpdate = computed(() => {
    return dialogData.value.extra_quote_type === 'update'
})

const quoteData = computed(() => {
    return quoteId.value ? allMsgs.value.find(({ id }) => id === quoteId.value) : null
})

const quickShow = computed(() => {
    return quickMsgs.value.length > 0 && windowScrollY === 0 && quoteId.value === 0
})

const quickMsgs = computed(() => {
    return dialogData.value.quick_msgs || []
})

/**
 * 发送消息
 * @param text
 * @param type
 */
const sendMsg = (text: any, type: any) => {
    let textBody,
        textType = "text",
        silence = "no",
        emptied = false;
    if (typeof text === "string" && text) {
        textBody = text;
    } else {
        textBody = msgText.value;
        emptied = true;
    }
    if (type === "md") {
        textBody = input.value.getText()
        textType = "md"
    } else if (type === "silence") {
        silence = "yes"
    }
    if (textBody == '') {
        inputFocus();
        return;
    }
    if (textType === "text") {
        textBody = textBody.replace(/<\/span> <\/p>$/, "</span></p>")
    }
    //
    if (quoteUpdate.value) {
        // 修改
        if (textType === "text") {
            textBody = textBody.replace(new RegExp(`src=(["'])${webTs.apiUrl('../')}`, "g"), "src=$1{{RemoteURL}}")
        }
        const update_id = quoteId.value
        // 重写接口
        // this.$store.dispatch("setLoad", {
        //     key: `msg-${update_id}`,
        //     delay: 600
        // })
        cancelQuote()
        onActive()

        // 重写接口
        // this.$store.dispatch("call", {
        //     url: 'dialog/msg/sendtext',
        //     data: {
        //         dialog_id: this.dialogId,
        //         update_id,
        //         text: textBody,
        //         text_type: textType,
        //         silence,
        //     },
        //     method: 'post',
        //     complete: _ => this.$store.dispatch("cancelLoad", `msg-${update_id}`)
        // }).then(({ data }) => {
        //     this.sendSuccess(data)
        //     this.onPositionId(update_id)
        // }).catch(({ msg }) => {
        //     $A.modalError(msg)
        // });
    } else {
        // 发送
        const typeLoad = utils.stringLength(textBody.replace(/<img[^>]*?>/g, '')) > 5000
        const tempMsg = {
            id: getTempId(),
            dialog_id: dialogData.value.id,
            reply_id: quoteId.value,
            reply_data: quoteData.value,
            type: typeLoad ? 'loading' : 'text',
            userid: userId,
            msg: {
                text: typeLoad ? '' : textBody,
                type: textType,
            },
        }
        tempMsgs.value.push(tempMsg)
        msgType.value = ''
        cancelQuote()
        onActive()
        nextTick(() => { onToBottom })
        //重写接口
        // this.$store.dispatch("call", {
        //     url: 'dialog/msg/sendtext',
        //     data: {
        //         dialog_id: tempMsg.dialog_id,
        //         reply_id: tempMsg.reply_id,
        //         text: textBody,
        //         text_type: textType,
        //         silence,
        //     },
        //     method: 'post',
        // }).then(({ data }) => {
        //     this.tempMsgs = this.tempMsgs.filter(({ id }) => id != tempMsg.id)
        //     this.sendSuccess(data)
        // }).catch(error => {
        //     this.$set(tempMsg, 'error', true)
        //     this.$set(tempMsg, 'errorData', { type: 'text', content: error.msg, msg: textBody })
        // });
    }
    if (emptied) {
        requestAnimationFrame(_ => msgText.value = '')
    }
}


/**
         * 发送快捷消息
         * @param item
         */
const sendQuick = (item: any) => {
    sendMsg(`<p><span data-quick-key="${item.key}">${item.label}</span></p>`, '')
}

const inputFocus = () => {
    nextTick(() => {
        input && input.value.focus()
    })
}

const handleSelect = () => {

}

const onToBottom = () => {
    msgNew.value = 0;
    if (scroller) {
        scroller.value.scrollToBottom();
        requestAnimationFrame(_ => scroller.value.scrollToBottom())    // 确保滚动到
    }
}

const getTempId = () => {
    return tempId.value++
}

const chatFile = (type: any, file: any) => {
    switch (type) {
        case 'progress':
            const tempMsg = {
                id: file.tempId,
                dialog_id: dialogData.value.id,
                reply_id: quoteId,
                type: 'loading',
                userid: userId,
                msg: {},
            }
            tempMsgs.value.push(tempMsg)
            msgType.value = ''
            cancelQuote()
            onActive()
            nextTick(() => { onToBottom })
            break;

        case 'error':
            tempMsgs.value = tempMsgs.value.filter(({ id }) => id != file.tempId)
            break;

        case 'success':
            tempMsgs.value = tempMsgs.value.filter(({ id }) => id != file.tempId)
            sendSuccess(file.data)
            break;
    }
}

const onViewTodo = (item: any) => {
    if (operateVisible.value) {
        return
    }
    todoViewId.value = item.id
    todoViewMid.value = item.msg_id
    todoViewShow.value = true
    //
    const index = allMsgs.value.findIndex(item => item.id === todoViewMid.value)
    if (index === -1) {
        // 重写接口
        // this.$store.dispatch("call", {
        //     url: 'dialog/msg/one',
        //     data: {
        //         msg_id: this.todoViewMid
        //     },
        // }).then(({ data }) => {
        //     this.todoViewData = data
        // })
    }
}

const sendSuccess = (data: any) => {
    if (utils.isArray(data)) {
        data.some(sendSuccess)
        return;
    }
    // this.$store.dispatch("saveDialogMsg", data); 重写接口
    if (!quoteUpdate.value) {
        // this.$store.dispatch("increaseTaskMsgNum", data); 重写接口
        // this.$store.dispatch("increaseMsgReplyNum", data);
        // this.$store.dispatch("updateDialogLastMsg", data);
    }
    cancelQuote();
    onActive();
}

const cancelQuote = () => {
    input.value.cancelQuote()
}

const onActive = () => {
    emit("on-active");
}

</script>
<style lang="less">
.dialog-wrapper {
    @apply absolute top-0 bottom-0 left-0 right-0 flex flex-col bg-white z-10;

    .dialog-nav {
        @apply w-full;

        .nav-wrapper {
            @apply flex items-center px-22 h-68 relative;

            &:before {
                @apply absolute left-0 right-0 bottom-0 w-full h-1 border-solid border-0 border-b border-border-color;
                content: '';
            }

            &.completed {
                &:after {
                    @apply pointer-events-none absolute top-1/2 right-52 text-40 text-completed-color opacity-20;
                    content: "\f373";
                    font-family: Ionicons, serif;
                    transform: translateY(-50%);
                    z-index: 1;
                }

                .dialog-title {
                    @apply pr-52;
                }
            }

            .dialog-block {
                @apply flex w-0 flex-1 items-center;

                .dialog-avatar {
                    @apply flex-shrink-0 mr-12;

                    .img-avatar,
                    .user-avatar,
                    .icon-avatar {
                        @apply w-42 h-42 mr-2 flex-grow-0 flex-shrink-0;
                    }

                    .img-avatar {
                        @apply flex items-center justify-center;

                        >img {
                            @apply w-full h-full;
                        }
                    }

                    .icon-avatar {
                        @apply flex items-center justify-center rounded-full text-26 text-white bg-icon-avatar;

                        &.department {
                            @apply bg-icon-avatar-department;
                        }

                        &.project {
                            @apply bg-icon-avatar-project;
                        }

                        &.task {
                            @apply bg-icon-avatar-task text-24;
                        }
                    }
                }

                .dialog-title {
                    @apply flex-1 w-0 flex flex-col justify-center;

                    .main-title {
                        @apply flex items-center leading-22 max-w-full;

                        h2 {
                            @apply text-17 font-semibold truncate;
                        }
                    }

                    .title-tags {
                        @apply flex items-center;

                        li {
                            @apply text-12 list-none leading-24 mt-4 mr-6 -mb-6 truncate cursor-pointer flex items-center px-6 rounded-md;

                            i {
                                @apply flex-shrink-0 w-14 h-14 leading-14 text-14 mr-4 bg-no-repeat bg-contain bg-center;
                            }
                        }

                        li.active {
                            @apply font-medium text-primary-color bg-messenger-active;
                        }

                        li.msg {
                            i {
                                @apply bg-msg;
                            }
                        }

                        li.task {
                            i {
                                @apply bg-task;
                            }
                        }
                    }
                }
            }

            .dialog-menu-icon {
                @apply cursor-pointer ml-22 text-22 text-text-li;
            }
        }
    }

    .dialog-scroller {
        @apply flex-1 relative px-32 py-16;
    }

    .dialog-footer {
        @apply relative px-24 mb-16;

        .dialog-newmsg,
        .dialog-goto {
            @apply absolute right-30 text-white cursor-pointer select-none pointer-events-none opacity-0 bg-dialog-goto;
            z-index: 2;
            transition: all 0.2s;
            transform: scale(0);
        }

        .dialog-newmsg {
            @apply block -top-44 h-30 leading-30 text-12 px-12 rounded-2xl;
        }

        .dialog-goto {
            @apply flex items-center justify-center text-text-li bg-white -top-48 w-40 h-40 leading-38 rounded-full;
            border: 1px solid #eeeeee;
            box-shadow: 0 4px 8px 0 rgba(--primary-text-color, 0.2);

            .taskfont {
                @apply text-24;
            }
        }

        .chat-upload {
            @apply hidden w-0 h-0 overflow-hidden;
        }

        .chat-bottom-menu {
            @apply flex items-center py-8;

            .bottom-menu-label {
                @apply flex-shrink-0 pr-8;
            }

            >ul {
                @apply flex-1 flex items-center overflow-x-auto;

                >li {
                    @apply flex-shrink-0 list-none mr-8 py-12 rounded-xl leading-26 text-13 flex items-center;
                    background-color: #f0f1f3;

                    .bottom-menu-desc {
                        @apply max-w-150 truncate;
                    }
                }
            }
        }

        .chat-mute {
            @apply text-text-tips bg-bg-manage-menu py-8 px-12 rounded-lg text-center;

        }

        .chat-input-box {
            .chat-input-wrapper {
                @apply bg-bg-manage-menu py-8 px-12 rounded-lg;

                .ql-container {
                    .ql-editor {
                        @apply my-4 mx-12;

                        &.ql-blank {
                            &::before {
                                @apply left-12 right-12;

                            }
                        }
                    }
                }
            }
        }

        &.newmsg {
            .dialog-newmsg {
                @apply pointer-events-auto opacity-100;
                transform: scale(1);
            }
        }

        &.goto {
            .dialog-goto {
                @apply pointer-events-auto opacity-100;
                transform: scale(1);
            }
        }
    }


}
</style>