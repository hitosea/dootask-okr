<template>
    <div class="dialog-view" :class="viewClass" :data-id="props.msgData.id">
        <!--昵称-->
        <div v-if="props.dialogType === 'group'" class="dialog-username">
            <UserAvatar :userid="props.msgData.userid" :show-icon="false" :show-name="true" click-open-dialog />
        </div>

        <div class="dialog-head" :class="headClass" v-longpress="{ callback: handleLongpress, delay: 300 }">
            <!--回复-->
            <div v-if="!props.hideReply && props.msgData.reply_data" class="dialog-reply no-dark-content"
                @click="viewReply">
                <UserAvatar :userid="props.msgData.reply_data.userid" :show-icon="false" :show-name="true"
                    :tooltip-disabled="true" />
                <div class="reply-desc">{{ webTs.getMsgSimpleDesc(props.msgData.reply_data) }}</div>
            </div>
            <!--详情-->
            <div ref="content" class="dialog-content" :class="contentClass">
                <!--文本-->
                <div v-if="props.msgData.type === 'text'" class="content-text no-dark-content">
                    <DialogMarkdown v-if="msgData.msg.type === 'md'" @click="viewText" :text="msgData.msg.text" />
                    <pre v-else @click="viewText" v-html="webTs.formatTextMsg(props.msgData.msg.text, userId)"></pre>
                </div>
                <!--文件-->
                <div v-else-if="props.msgData.type === 'file'" :class="`content-file ${msgData.msg.type}`">
                    <div class="dialog-file">
                        <img v-if="props.msgData.msg.type === 'img'" class="file-img" :style="imageStyle(props.msgData.msg)"
                            :src="props.msgData.msg.thumb" @click="viewFile" />
                        <div v-else class="file-box" @click="downFile">
                            <img class="file-thumb" :src="props.msgData.msg.thumb" />
                            <div class="file-info">
                                <div class="file-name">{{ props.msgData.msg.name }}</div>
                                <div class="file-size">{{ utils.bytesToSize(props.msgData.msg.size) }}</div>
                            </div>
                        </div>
                    </div>
                </div>
                <!--录音-->
                <div v-else-if="props.msgData.type === 'record'" class="content-record no-dark-content">
                    <div class="dialog-record" :class="{ playing: audioPlaying === props.msgData.msg.path }"
                        :style="recordStyle(props.msgData.msg)" @click="playRecord">
                        <div class="record-time">{{ recordDuration(props.msgData.msg.duration) }}</div>
                        <div class="record-icon taskfont"></div>
                    </div>
                </div>
                <!--会议-->
                <div v-else-if="props.msgData.type === 'meeting'" class="content-meeting no-dark-content">
                    <ul class="dialog-meeting">
                        <li>
                            <em>{{ $t('会议主题') }}</em>
                            {{ props.msgData.msg.name }}
                        </li>
                        <li>
                            <em>{{ $t('会议创建人') }}</em>
                            <UserAvatar :userid="props.msgData.msg.userid" :show-icon="false" :show-name="true"
                                tooltip-disabled />
                        </li>
                        <li>
                            <em>{{ $t('频道ID') }}</em>
                            {{ props.msgData.msg.meetingid.replace(/^(.{3})(.{3})(.*)$/, '$1 $2 $3') }}
                        </li>
                        <li class="meeting-operation" @click="openMeeting">
                            {{ $t('点击加入会议') }}
                            <i class="taskfont">&#xe68b;</i>
                        </li>
                    </ul>
                </div>
                <!--等待-->
                <div v-else-if="props.msgData.type === 'loading'" class="content-loading">
                    <Icon v-if="props.msgData.error === true" type="ios-alert-outline" />
                    <Loading v-else />
                </div>
                <!--未知-->
                <div v-else class="content-unknown">{{ $t("未知的消息类型") }}</div>
            </div>
            <!--emoji-->
            <ul v-if="utils.arrayLength(props.msgData.emoji) > 0" class="dialog-emoji">
                <li v-for="(item, index) in props.msgData.emoji" :key="index"
                    :class="{ hasme: item.userids.includes(userId) }">
                    <div class="emoji-symbol no-dark-content" @click="onEmoji(item.symbol)">{{ item.symbol }}</div>
                    <div class="emoji-users" @click="onShowEmojiUser(item)">
                        <ul>
                            <template v-for="(uitem, uindex) in item.userids">
                                <li v-if="uindex < emojiUsersNum" :class="{ bold: uitem == userId }">
                                    <UserAvatar :userid="uitem" tooltip-disabled show-name :show-icon="false" />
                                </li>
                                <li v-else-if="uindex == emojiUsersNum">+{{ item.userids.length - emojiUsersNum }}位</li>
                            </template>
                        </ul>
                    </div>
                </li>
            </ul>
        </div>

        <div class="dialog-foot">
            <!--回复数-->
            <div v-if="!props.hideReply && props.msgData.reply_num > 0" class="reply" @click="replyList">
                <i class="taskfont">&#xe6eb;</i>
                {{ props.msgData.reply_num }}条回复
            </div>
            <!--标注-->
            <div v-if="props.msgData.tag" class="tag">
                <i class="taskfont">&#xe61e;</i>
            </div>
            <!--待办-->
            <div v-if="props.msgData.todo" class="todo" @click="openTodo">
                <EPopover v-model="todoShow" ref="todo" popper-class="dialog-wrapper-read-poptip"
                    :placement="props.isRightMsg ? 'bottom-end' : 'bottom-start'">
                    <div class="read-poptip-content">
                        <ul class="read scrollbar-overlay">
                            <li class="read-title"><em>{{ todoDoneList.length }}</em>{{ $t('完成') }}</li>
                            <li v-for="item in todoDoneList">
                                <UserAvatar :userid="item.userid" :size="26" showName tooltipDisabled />
                            </li>
                        </ul>
                        <ul class="unread scrollbar-overlay">
                            <li class="read-title"><em>{{ todoUndoneList.length }}</em>{{ $t('待办') }}</li>
                            <li v-for="item in todoUndoneList">
                                <UserAvatar :userid="item.userid" :size="26" showName tooltipDisabled />
                            </li>
                        </ul>
                    </div>
                    <div slot="reference" class="popover-reference"></div>
                </EPopover>
                <Loading v-if="todoLoad > 0" />
                <i v-else class="taskfont">&#xe7b7;</i>
            </div>
            <!--编辑-->
            <div v-if="props.msgData.modify" class="modify">
                <i class="taskfont">&#xe779;</i>
            </div>
            <!--错误/等待/时间/阅读-->
            <div v-if="props.msgData.error === true" class="error" @click="onError">
                <Icon type="ios-alert" />
            </div>
            <Loading v-else-if="isLoading" />
            <template v-else>
                <!--时间-->
                <div v-if="timeShow" class="time" @click="timeShow = false">{{ props.msgData.created_at }}</div>
                <div v-else class="time" :title="props.msgData.created_at" @click="timeShow = true">
                    {{ webTs.formatTime(props.msgData.created_at) }}</div>
                <!--阅读-->
                <template v-if="!props.hidePercentage">
                    <div v-if="props.msgData.send > 1 || props.dialogType === 'group'" class="percent"
                        @click="openReadPercentage">
                        <EPopover v-model="percentageShow" ref="percent" popper-class="dialog-wrapper-read-poptip"
                            :placement="props.isRightMsg ? 'bottom-end' : 'bottom-start'">
                            <div class="read-poptip-content">
                                <ul class="read scrollbar-overlay">
                                    <li class="read-title"><em>{{ readList.length }}</em>{{ $t('已读') }}</li>
                                    <li v-for="item in readList">
                                        <UserAvatar :userid="item.userid" :size="26" showName tooltipDisabled />
                                    </li>
                                </ul>
                                <ul class="unread scrollbar-overlay">
                                    <li class="read-title"><em>{{ unreadList.length }}</em>{{ $t('未读') }}</li>
                                    <li v-for="item in unreadList">
                                        <UserAvatar :userid="item.userid" :size="26" showName tooltipDisabled />
                                    </li>
                                </ul>
                            </div>
                            <div slot="reference" class="popover-reference"></div>
                        </EPopover>
                        <Loading v-if="percentageLoad > 0" />
                        <WCircle v-else :percent="msgData.percentage" :size="14" />
                    </div>
                    <Icon v-else-if="props.msgData.percentage === 100" class="done" type="md-done-all" />
                    <Icon v-else class="done" type="md-checkmark" />
                </template>
            </template>
        </div>
    </div>
</template>

<script setup lang="ts">
import { Store } from "le5le-store";
import WCircle from "@/components/WCircle.vue";
import DialogMarkdown from "./DialogMarkdown.vue";
import utils from "@/utils/utils";
import webTs from "@/utils/web";
import { ref, defineProps, onBeforeUnmount, computed, watch, defineEmits } from "vue"
import { UserStore } from "@/store/user";
import { DiaLogStore } from "@/store/dialog"

const audioPlaying = DiaLogStore().audioPlaying
const isLoad = DiaLogStore().isLoad
const userId = UserStore().info.id
const timeShow = ref<boolean>(false)
const operateEnter = ref(false)
const percentageLoad = ref(0)
const percentageShow = ref(false)
const percentageList = ref([])
const todoLoad = ref(0)
const todoShow = ref(false)
const todoList = ref([])
const content = ref(null)
const emojiUsersNum = ref(5)
const emit = defineEmits(['on-longpress', 'on-view-reply', 'on-view-text', 'on-view-file', 'on-down-file', 'on-reply-list', 'on-error', 'on-emoji', 'on-show-emoji-user'])


const props = defineProps({
    msgData: {
        type: Object,
        default: () => {
            return {};
        }
    },
    dialogType: {
        type: String,
        default: ''
    },
    hidePercentage: {
        type: Boolean,
        default: false
    },
    hideReply: {
        type: Boolean,
        default: false
    },
    operateVisible: {
        type: Boolean,
        default: false
    },
    operateAction: {
        type: Boolean,
        default: false
    },
    isRightMsg: {
        type: Boolean,
        default: false
    },
})


// emojiUsersNum.value = Math.min(6, Math.max(2, Math.floor((this.windowWidth - 180) / 52))) this.windowWidth全局的



onBeforeUnmount(() => {
    // this.$store.dispatch("audioStop", props.msgData.msg.path) 终止
})

const isLoading = computed(() => {
    if (!props.msgData.created_at) {
        return true;
    }
    return isLoad(`msg-${props.msgData.id}`)
})

const viewClass = computed(() => {
    const { msgData, operateAction } = props;
    const array = [];
    if (msgData.type) {
        array.push(msgData.type)
    }
    if (msgData.reply_data) {
        array.push('reply-view')
    }
    if (operateAction) {
        array.push('operate-action')
        if (operateEnter.value) {
            array.push('operate-enter')
        }
    }
    return array
})

const readList = computed(() => {
    return percentageList.value.filter(({ read_at }) => read_at)
})

const unreadList = computed(() => {
    return percentageList.value.filter(({ read_at }) => !read_at)
})

const todoDoneList = computed(() => {
    return todoList.value.filter(({ done_at }) => done_at)
})

const todoUndoneList = computed(() => {
    return todoList.value.filter(({ done_at }) => !done_at)
})

const headClass = computed(() => {
    const { reply_id, type, msg, emoji } = props.msgData;
    const array = [];
    if (reply_id === 0 && utils.arrayLength(emoji) === 0) {
        if (type === 'text') {
            if (/^<img\s+class="emoticon"[^>]*?>$/.test(msg.text)
                || /^\s*<p>\s*([\uD800-\uDBFF][\uDC00-\uDFFF]){1,3}\s*<\/p>\s*$/.test(msg.text)) {
                array.push('transparent')
            }
        }
    }
    return array;
})

const contentClass = computed(() => {
    const { type, msg } = props.msgData;
    const classArray = [];
    if (type === 'text') {
        if (/^<img\s+class="emoticon"[^>]*?>$/.test(msg.text)) {
            classArray.push('an-emoticon')
        } else if (/^\s*<p>\s*([\uD800-\uDBFF][\uDC00-\uDFFF]){3}\s*<\/p>\s*$/.test(msg.text)) {
            classArray.push('three-emoji')
        } else if (/^\s*<p>\s*([\uD800-\uDBFF][\uDC00-\uDFFF]){2}\s*<\/p>\s*$/.test(msg.text)) {
            classArray.push('two-emoji')
        } else if (/^\s*<p>\s*[\uD800-\uDBFF][\uDC00-\uDFFF]\s*<\/p>\s*$/.test(msg.text)) {
            classArray.push('an-emoji')
        }
    }
    return classArray;
})


watch(props, (newValue, oldValue) => {
    operateEnter.value = false;
    if (newValue) {
        setTimeout(_ => operateEnter.value = true, 500)
    }
})



const handleLongpress = (event, el) => {
    emit("on-longpress", { event, el, msgData: props.msgData })
}

const openTodo = () => {
    if (todoLoad.value > 0) {
        return;
    }
    if (todoShow.value) {
        todoShow.value = false;
        return;
    }
    // 重新写接口
    // todoLoad.value++;
    // this.$store.dispatch("call", {
    //     url: 'dialog/msg/todolist',
    //     data: {
    //         msg_id: props.msgData.id,
    //     },
    // }).then(({ data }) => {
    //     todoList.value = data;
    // }).catch(() => {
    //     todoList.value = [];
    // }).finally(_ => {
    //     setTimeout(() => {
    //         todoLoad.value--;
    //         todoShow.value = true
    //     }, 100)
    // });
}

const openReadPercentage = () => {
    if (percentageLoad.value > 0) {
        return;
    }
    if (percentageShow.value) {
        percentageShow.value = false;
        return;
    }
    // 重新写接口
    // percentageLoad.value++;
    // this.$store.dispatch("call", {
    //     url: 'dialog/msg/readlist',
    //     data: {
    //         msg_id: props.msgData.id,
    //     },
    // }).then(({ data }) => {
    //     percentageList.value = data;
    // }).catch(() => {
    //     percentageList.value = [];
    // }).finally(_ => {
    //     setTimeout(() => {
    //         percentageLoad.value--;
    //         percentageShow.value = true
    //     }, 100)
    // });
}

const recordStyle = (info) => {
    const { duration } = info;
    const width = 50 + Math.min(180, Math.floor(duration / 150));
    return {
        width: width + 'px',
    };
}

const recordDuration = (duration) => {
    const minute = Math.floor(duration / 60000),
        seconds = Math.floor(duration / 1000) % 60;
    if (minute > 0) {
        return `${minute}:${seconds}″`
    }
    return `${Math.max(1, seconds)}″`
}

const imageStyle = (info) => {
    const { width, height } = info;
    if (width && height) {
        let maxW = 220,
            maxH = 220,
            tempW = width,
            tempH = height;
        if (width > maxW || height > maxH) {
            if (width > height) {
                tempW = maxW;
                tempH = height * (maxW / width);
            } else {
                tempW = width * (maxH / height);
                tempH = maxH;
            }
        }
        return {
            width: tempW + 'px',
            height: tempH + 'px',
        };
    }
    return {};
}

const playRecord = () => {
    if (props.operateVisible) {
        return
    }
    // this.$store.dispatch("audioPlay", props.msgData.msg.path) 重写全局方法
}

const openMeeting = () => {
    if (props.operateVisible) {
        return
    }
    Store.set('addMeeting', {
        type: 'join',
        name: props.msgData.msg.name,
        meetingid: props.msgData.msg.meetingid,
        meetingdisabled: true,
    });
}

const viewReply = () => {
    emit("on-view-reply", {
        msg_id: props.msgData.id,
        reply_id: props.msgData.reply_id
    })
}

const viewText = (e) => {
    emit("on-view-text", e, content.value)
}

const viewFile = () => {
    emit("on-view-file", props.msgData)
}

const downFile = () => {
    emit("on-down-file", props.msgData)
}

const replyList = () => {
    emit("on-reply-list", {
        msg_id: props.msgData.id,
    })
}

const onError = () => {
    emit("on-error", props.msgData)
}

const onEmoji = (symbol) => {
    emit("on-emoji", {
        msg_id: props.msgData.id,
        symbol
    })
}

const onShowEmojiUser = (item) => {
    emit("on-show-emoji-user", item)
}


</script>
<style lang="less">
.dialog-view {
    @apply flex flex-col items-start ml-8 relative;

    &.text {
        @apply max-w-70p;
    }

    &.operate-action {
        .dialog-head {
            @apply shadow-dialog-head;
        }
    }

    .dialog-username {
        @apply max-w-full h-22 mb-6 opacity-80;
    }

    .dialog-head {
        @apply flex flex-col bg-bg-manage-menu p-8 min-w-32 rounded-lg rounded-tl-sm max-w-full;
        transition: box-shadow 0.3s ease;

        &.transparent {
            @apply bg-transparent !important;
        }

        .dialog-reply {
            @apply cursor-pointer pl-9 mb-4 relative;

            &:after {
                @apply absolute top-0 left-0 bottom-0 w-3 rounded-sm bg-primary-color-80;
                content: "";
                transform: scaleX(0.8);
                transform-origin: left center;
            }

            .common-avatar {
                @apply text-primary-color text-13 font-medium;
            }

            .reply-desc {
                @apply text-13 truncate;
                display: -webkit-box;
                -webkit-line-clamp: 2;
                -webkit-box-orient: vertical;
            }
        }

        .dialog-content {
            @apply flex items-start relative;

            &.an-emoji {
                .content-text {
                    >pre {
                        @apply text-72 leading-none;
                    }
                }
            }

            &.two-emoji {
                .content-text {
                    >pre {
                        @apply text-52 leading-none tracking-4;
                    }
                }
            }

            &.three-emoji {
                .content-text {
                    >pre {
                        @apply text-32 leading-none tracking-4;
                    }
                }
            }

            .content-text {
                @apply text-primary-color p-2 max-w-full;

                .no-size-image-box {
                    @apply inline-block max-w-220;
                }

                >pre {
                    @apply block m-0 p-0 leading-20 whitespace-pre-wrap text-14;
                    word-break: break-word;
                    font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Helvetica, Arial, sans-serif, "Apple Color Emoji", "Segoe UI Emoji";

                    a {
                        color: #436FF6;
                    }

                    ol,
                    ul {
                        li {
                            @apply list-none;

                            &::before {
                                @apply inline-block whitespace-nowrap w-1.2em text-left mr-1.2em;
                            }
                        }
                    }

                    ul {
                        li {
                            &::before {
                                content: '\2022';
                                font-weight: 900;
                            }
                        }
                    }

                    ol {
                        li {
                            counter-reset: list-1 list-2 list-3 list-4 list-5 list-6 list-7 list-8 list-9;
                            counter-increment: list-0;

                            &:before {
                                @apply w-auto min-w-1.2em;
                                content: counter(list-0, decimal) '. ';
                            }
                        }
                    }

                    blockquote {
                        @apply my-5 pl-16 border-solid border-0 border-l border-blockquote;
                    }

                    pre {
                        @apply whitespace-pre-wrap my-5 py-5 px-10 rounded bg-pre-bg text-pre-color overflow-visible;
                    }

                    img {
                        @apply cursor-pointer max-w-full max-h-220 align-bottom;

                        &.emoticon {
                            @apply max-w-100% max-h-150;
                        }
                    }

                    .mention {
                        @apply text-flow-status-end-color bg-transparent p-0 m-0;
                        user-select: inherit;

                        >span {
                            @apply m-0;
                        }

                        &.task {
                            @apply cursor-pointer;
                        }

                        &.file,
                        &[data-denotation-char="~"] {
                            color: #436FF6 !important;
                        }

                        &.me {
                            @apply text-13 font-semibold py-3 px-4 text-white bg-primary-color;

                        }
                    }
                }
            }

            .content-file {
                &.file {
                    @apply inline-block;

                    .file-box {
                        @apply bg-white flex items-center py-10 px-14 rounded-sm w-220 cursor-pointer;

                        .file-thumb {
                            @apply w-36;
                        }

                        .file-info {
                            @apply ml-12 flex flex-col justify-center;

                            .file-name {
                                @apply text-title-color text-14 leading-18 truncate;
                                display: -webkit-box;
                                -webkit-line-clamp: 2;
                                -webkit-box-orient: vertical;
                            }

                            .file-size {
                                @apply pt-4 text-text-li text-14;

                            }
                        }
                    }
                }

                &.img {
                    @apply p-0 flex max-w-220 max-h-220 rounded-md overflow-hidden;

                    .file-img {
                        @apply flex cursor-pointer;
                    }
                }
            }

            .content-record {
                @apply flex text-text-li;

                .dialog-record {
                    @apply flex flex-row-reverse justify-end content-center leading-24 cursor-pointer;

                    .record-time {
                        @apply pl-4;
                    }

                    .record-icon {
                        transform: rotate(180deg) scale(0.9);

                        &:before {
                            content: "\E793";
                        }
                    }

                    &.playing {
                        .record-icon {
                            &:before {
                                animation: record-playing 1s infinite;
                            }
                        }
                    }

                    @keyframes record-playing {
                        0% {
                            content: "\E793";
                        }

                        33% {
                            content: "\E791";
                        }

                        66% {
                            content: "\E792";
                        }

                        100% {
                            content: "\E793";
                        }
                    }
                }
            }

            .content-meeting {
                @apply py-4 px-6 text-title-color;

                .dialog-meeting {
                    @apply min-w-220;

                    >li {
                        @apply list-none flex flex-col items-start mb-16;

                        &.meeting-operation {
                            @apply mb-0 py-12 flxe flex-row items-center text-12 relative cursor-pointer;

                            &:hover {
                                .taskfont {
                                    @apply pl-4;
                                }
                            }

                            &:before {
                                @apply absolute top-0 left-0 right-0 h-1;
                                content: "";
                                background-color: rgba(204, 204, 204, 0.8);
                                transform: scaleY(0.5);
                            }

                            .taskfont {
                                @apply text-12 pl-2;
                                transform: scale(0.8);
                                transition: all 0.2s;
                            }
                        }

                        >em {
                            @apply not-italic font-black pb-2;
                        }
                    }
                }
            }

            .content-loading {
                @apply flex;

                >i {
                    @apply text-20 m-2 text-title-color;
                }

                .common-loading {
                    @apply w-20 h-20 m-4;
                }
            }

            .content-unknown {
                @apply underline text-text-li;
            }
        }

        .dialog-emoji {
            @apply flex flex-wrap items-center;

            >li {
                @apply list-none flex items-center py-2 px-7 mt-6 mr-8 rounded-lg leading-22 cursor-pointer;
                background-color: rgba(#e1e1e1, 0.5);

                &.hasme {
                    background-color: #e1e1e1;
                }

                .emoji-symbol {
                    @apply text-16;
                    transition: transform 0.3s;

                    &:hover {
                        transform: scale(1.5);
                    }
                }

                .emoji-users {
                    @apply relative pl-14;

                    &:before {
                        @apply absolute left-7 top-6 bottom-6 w-1 bg-emoji-users-color;
                        content: "";
                        transform: scaleX(0.5);
                        background-color: rgba(#818181, 0.5);
                    }

                    >ul {
                        @apply flex flex-wrap items-center;

                        >li {
                            @apply flex items-center list-none text-bg-emoji-users-color text-12;

                            &.bold {
                                @apply font-semibold;

                            }

                            &:after {
                                content: "、";
                            }

                            &:last-child {
                                &:after {
                                    @apply hidden;
                                }
                            }
                        }
                    }
                }
            }
        }
    }

    .dialog-foot {
        @apply flex items-center pt-4 h-21 leading-none;

        .error {
            @apply cursor-pointer;
            color: #ED4014;

            >i {
                @apply text-14;
            }
        }

        .common-loading {
            @apply my-2 w-10 g-10;
        }

        .popover-reference {
            @apply absolute left-65p bottom-0 w-0 h-full pointer-events-none;

        }

        .tag,
        .todo,
        .reply,
        .modify {
            @apply flex items-center mr-6;

            >i {
                @apply text-13;

            }
        }

        .todo {
            @apply relative cursor-pointer;

            .common-loading {
                @apply mr-3;
            }
        }

        .reply {
            @apply text-12 text-text-li cursor-pointer;

            >i {
                @apply pr-2;
            }
        }

        .time {
            @apply text-12;
            color: #bbbbbb;
        }

        .done {
            @apply hidden ml-4 text-12 text-text-li;
            transform: scale(0.9);
        }

        .percent {
            @apply hidden ml-4 items-center relative cursor-pointer;

        }
    }
}

.self {
    @apply flex-row-reverse;

    .dialog-view {
        @apply items-end mr-8;
        .dialog-head {
            @apply bg-primary-color rounded-lg rounded-tr-sm;

            .dialog-reply {
                @apply text-white;
                .common-avatar,
                .bot {
                    @apply text-white;
              
                }

                &:after {
                    @apply bg-white;
                }
            }

            .dialog-content {
                .content-text {
                    @apply text-white;

                    >pre {
                        .mention {
                            @apply text-title-color;

                            &.me {
                                font-size: inherit;
                                font-weight: inherit;
                                padding: inherit;
                                background-color: inherit;
                            }
                        }
                    }
                }

                .content-record {
                    @apply text-white;

                    .dialog-record {
                        @apply flex-row;

                        .record-time {
                            @apply pr-4;
                        }

                        .record-icon {
                            transform: rotate(0) scale(0.9);
                        }
                    }
                }

                .content-meeting {
                    @apply text-white;

                    .dialog-meeting {
                        >li {
                            &.meeting-operation {
                                &:before {
                                    @apply bg-white-80;

                                }
                            }
                        }
                    }
                }

                .content-loading {
                    >i {
                        @apply text-white;
                   
                    }
                }

                .content-unknown {
                    @apply text-white;
                 
                }
            }

            .dialog-emoji {
                >li {
                    @apply bg-dialog-emoji-50;

                    &.hasme {
                        @apply bg-dialog-emoji;

                    }

                    .emoji-users {
                        &:before {
                            @apply bg-white-50;
        
                        }

                        >ul {
                            >li {
                                @apply text-white;
                
                            }
                        }
                    }
                }
            }
        }

        .dialog-foot {
            .done {
                @apply inline-block;
            }

            .percent {
                @apply flex;
            }
        }
    }
}</style>