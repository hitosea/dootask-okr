<template>
    <div :class="classArray">
        <div v-if="props.source.type === 'tag'" class="dialog-tag" @click="onViewTag">
            <div class="tag-user">
                <UserAvatar :userid="props.source.userid" :tooltipDisabled="props.source.userid == userId" :show-name="true"
                    :show-icon="false" />
            </div>
            {{ $t(props.source.msg.action === 'remove' ? '取消标注' : '标注了') }}
            "{{ webTs.getMsgSimpleDesc(source.msg.data) }}"
        </div>
        <div v-else-if="source.type === 'todo'" class="dialog-todo" @click="onViewTodo">
            <div class="todo-user">
                <UserAvatar :userid="props.source.userid" :tooltipDisabled="props.source.userid == userId" :show-name="true"
                    :show-icon="false" />
            </div>
            {{ $t(props.source.msg.action === 'remove' ? '取消待办' : (props.source.msg.action === 'done' ? '完成' : '设待办')) }}
            "{{ webTs.getMsgSimpleDesc(props.source.msg.data) }}"
            <div v-if="formatTodoUser(props.source.msg.data).length > 0" class="todo-users">
                <span>{{ $t('给') }}</span>
                <template v-for="(item, index) in formatTodoUser(props.source.msg.data)">
                    <div v-if="index < 3" class="todo-user">
                        <UserAvatar :userid="item" :tooltipDisabled="item == userId" :show-name="true" :show-icon="false" />
                    </div>
                    <div v-else-if="index == 3" class="todo-user">+{{ formatTodoUser(props.source.msg.data).length - 3 }}
                    </div>
                </template>
            </div>
        </div>
        <div v-else-if="source.type === 'notice'" class="dialog-notice">
            {{ props.source.msg.notice }}
        </div>
        <template v-else>
            <div class="dialog-avatar">
                <UserAvatar v-longpress="{ callback: onMention, delay: 300 }" @open-dialog="onOpenDialog"
                    :userid="props.source.userid" :size="30" tooltip-disabled />
            </div>
            <DialogView :msg-data="source" :dialog-type="dialogData.type" :hide-percentage="hidePercentage"
                :hide-reply="hideReply" :operate-visible="operateVisible"
                :operate-action="operateVisible && source.id === operateItem.id" :is-right-msg="isRightMsg"
                @on-longpress="onLongpress" @on-view-reply="onViewReply" @on-view-text="onViewText"
                @on-view-file="onViewFile" @on-down-file="onDownFile" @on-reply-list="onReplyList" @on-error="onError"
                @on-emoji="onEmoji" @on-show-emoji-user="onShowEmojiUser" />
        </template>
    </div>
</template>

<script lang="ts" setup>
import { defineProps, computed, watch, defineEmits } from "vue"
import { UserStore } from "@/store/user";
import utils from "@/utils/utils";
import webTs from "@/utils/web";
import DialogView from "./DialogView.vue";
import { GlobalStore } from "@/store";

const windowActive = GlobalStore().windowActive
const userId = UserStore().info.id

const emit = defineEmits(["on-mention", "on-longpress", "on-view-reply", "on-view-text", "on-view-file", "on-down-file", "on-reply-list", "on-error", "on-emoji", "on-show-emoji-user"])

const props = defineProps({
    source: {
        type: Object,
        default() {
            return {}
        }
    },
    dialogData: {
        type: Object,
        default() {
            return {}
        }
    },
    operateVisible: {
        type: Boolean,
        default: false
    },
    operateItem: {
        type: Object,
        default() {
            return {}
        }
    },
    simpleView: {
        type: Boolean,
        default: false
    },
    isMyDialog: {
        type: Boolean,
        default: false
    },
    msgId: {
        type: Number,
        default: 0
    },
})

const isRightMsg = computed(() => {
    return props.source.userid == userId
})

const isReply = computed(() => {
    return props.simpleView || props.msgId === props.source.id
})

const hidePercentage = computed(() => {
    return props.simpleView || props.isMyDialog || isReply.value
})

const hideReply = computed(() => {
    return props.simpleView || props.msgId > 0
})

const classArray = computed(() => {
    return {
        'dialog-item': true,
        'reply-item': isReply.value,
        'self': isRightMsg.value,
    }
})

watch(props.source, (newValue) => {
    msgRead()
}, { immediate: true })

watch(windowActive, (newValue) => {
    if (newValue) {
        msgRead()
    }
})  


const msgRead = () => {
    if (windowActive) { 
        return;
    }
    // 重写接口
    // this.$store.dispatch("dialogMsgRead", this.source);
}

const formatTodoUser = (data) => {
    if (utils.isJson(data)) {
        const { userids } = data
        if (userids) {
            return userids.split(",")
        }
    }
    return []
}

const onViewTag = () => {
    onViewReply({
        msg_id: props.source.id,
        reply_id: props.source.msg.data.id
    })
}

const onViewTodo = () => {
    onViewReply({
        msg_id: props.source.id,
        reply_id: props.source.msg.data.id
    })
}

const onOpenDialog = (userid) => {
    if (props.dialogData.type != 'group') {
        return
    }
    // 重写接口
    // this.$store.dispatch("openDialogUserid", userid).then(_ => {
    //     this.goForward({ name: 'manage-messenger' })
    // }).catch(({ msg }) => {
    //     $A.modalError(msg)
    // });
}

const onMention = () => {
    dispatch("on-mention", props.source)
}

const onLongpress = (e) => {
    dispatch("on-longpress", e)
}

const onViewReply = (data) => {
    dispatch("on-view-reply", data)
}

const onViewText = (e, el) => {
    dispatch("on-view-text", e, el)
}

const onViewFile = (data) => {
    dispatch("on-view-file", data)
}

const onDownFile = (data) => {
    dispatch("on-down-file", data)
}

const onReplyList = (data) => {
    dispatch("on-reply-list", data)
}

const onError = (data) => {
    dispatch("on-error", data)
}

const onEmoji = (data) => {
    dispatch("on-emoji", data)
}

const onShowEmojiUser = (data) => {
    dispatch("on-show-emoji-user", data)
}

const dispatch = (event, ...arg) => {
    if (isReply.value) {
        emit(event, ...arg)
        return
    }
    // 这部分需要修改
    // let parent = getCurrentInstance().parent
    // let name = parent.getCurrentInstance().type.name

    // while (parent && (!name || name !== 'virtual-list')) {
    //     parent = parent.getCurrentInstance().parent
    //     if (parent) {
    //         name = parent.getCurrentInstance().type.name
    //     }
    // }

    // if (parent) {
    //     parent.emit(event, ...arg)
    // }
}


</script>

<style lang="less">
.dialog-item {
    @apply flex flex-row items-start list-none pb-16;

    .dialog-tag,
    .dialog-todo,
    .dialog-notice {
        @apply cursor-pointer text-12 max-w-80p m-auto px-8 py-4 rounded-lg text-text-tips bg-dialog-item-notice;

        .tag-user {
            @apply inline-block;
        }
    }

    .dialog-avatar {
        @apply relative mb-20 flex-shrink-0 w-30 h-30;
    }
}

.dialog-item.self {
    @apply flex-row-reverse;
}
</style>