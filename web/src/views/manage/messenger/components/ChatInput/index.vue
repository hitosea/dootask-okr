<template>
    <div class="chat-input-box" :class="boxClass" v-clickoutside="hidePopover">
        <!-- 快速表情 -->
        <div class="chat-input-quick-emoji">
            <n-popover ref="emojiQuickRef" :show="emojiQuickShow" :show-arrow="false" placement="top-end" trigger="manual"
                popperClass="chat-quick-emoji-popover">
                <template #trigger>
                    <div slot="reference"></div>
                </template>
                <ul class="chat-quick-emoji-wrapper">
                    <li v-for="item in emojiQuickItems" @click="onEmojiQuick(item)">
                        <img :title="item.name" :alt="item.name" :src="item.src" />
                    </li>
                </ul>
            </n-popover>
        </div>

        <div class="chat-input-wrapper" @click.stop="focus">
            <!-- 回复、修改 -->
            <div v-if="quoteData" class="chat-quote">
                <div v-if="quoteUpdate" class="quote-label">{{ $t('编辑消息') }}</div>
                <UserAvatar v-else :userid="quoteData.userid" :show-icon="false" :show-name="true"
                    :tooltip-disabled="true" />
                <div class="quote-desc">{{ webTs.getMsgSimpleDesc(quoteData) }}</div>
                <i class="taskfont" @click.stop="cancelQuote">&#xe6e5;</i>
            </div>

            <!-- 输入框 -->
            <div ref="editor" class="no-dark-content" :style="editorStyle" @click.stop="onClickEditor" @paste="handlePaste">
            </div>

            <!-- 工具栏 -->
            <ul class="chat-toolbar" @click.stop>
                <!-- 桌面端表情（漂浮） -->
                <li>
                    <n-popover v-if="!emojiBottom" :show="showEmoji" :show-arrow="false" placement="top" trigger="manual"
                        popperClass="chat-input-emoji-popover">
                        <template #trigger>
                            <n-tooltip trigger="hover" :disabled="windowTouch || showEmoji" placement="top"
                                :content="$t('表情')">
                                <template #trigger>
                                    <i class="taskfont">&#xe7ad;</i>
                                </template>
                            </n-tooltip>
                            <ChatEmoji v-if="showEmoji" @on-select="onSelectEmoji" :searchKey="emojiQuickKey" />
                        </template>
                    </n-popover>
                    <n-tooltip trigger="hover" v-else ref="emojiTip" :disabled="windowTouch || showEmoji" placement="top"
                        :content="$t('表情')">
                        <template #trigger>
                            <i class="taskfont" @click="showEmoji = !showEmoji">&#xe7ad;</i>
                        </template>
                    </n-tooltip>
                </li>

                <!-- @ # -->
                <li>
                    <n-tooltip trigger="hover" placement="top" :disabled="windowTouch" :content="$t('选择成员')">
                        <template #trigger>
                            <i class="taskfont" @click="onToolbar('user')">&#xe78f;</i>
                        </template>
                    </n-tooltip>
                </li>
                <li>
                    <n-tooltip trigger="hover" placement="top" :disabled="windowTouch" :content="$t('选择任务')">
                        <template #trigger>
                            <i class="taskfont" @click="onToolbar('task')">&#xe7d6;</i>
                        </template>
                    </n-tooltip>
                </li>

                <!-- 图片文件 -->
                <li>
                    <n-popover :show="showMore" :show-arrow="false" placement="top" trigger="manual"
                        popperClass="chat-input-more-popover">
                        <template #trigger>
                            <n-tooltip trigger="hover" ref="moreTip" :disabled="windowTouch || showMore" placement="top"
                                :content="$t('展开')">
                                <i class="taskfont">&#xe790;</i>
                            </n-tooltip>
                        </template>
                        <div v-if="recordReady" class="chat-input-popover-item" @click="onToolbar('meeting')">
                            <i class="taskfont">&#xe7c1;</i>
                            {{ $t('新会议') }}
                        </div>
                        <div v-if="canCall" class="chat-input-popover-item" @click="onToolbar('call')">
                            <i class="taskfont">&#xe7ba;</i>
                            {{ $t('拨打电话') }}
                        </div>
                        <div class="chat-input-popover-item" @click="onToolbar('image')">
                            <i class="taskfont">&#xe7bc;</i>
                            {{ $t('发送图片') }}
                        </div>
                        <div class="chat-input-popover-item" @click="onToolbar('file')">
                            <i class="taskfont">&#xe7c0;</i>
                            {{ $t('上传文件') }}
                        </div>
                        <div v-if="canAnon" class="chat-input-popover-item" @click="onToolbar('anon')">
                            <i class="taskfont">&#xe690;</i>
                            {{ $t('匿名消息') }}
                        </div>
                    </n-popover>
                </li>

                <!-- 发送按钮 -->
                <li ref="chatSend" class="chat-send" :class="sendClass" v-touchmouse="clickSend"
                    v-longpress="{ callback: longSend, delay: 300 }">
                    <n-popover :show="showMenu" :show-arrow="false" trigger="manual" placement="top"
                        popperClass="chat-input-more-popover">
                        <template #trigger>
                            <n-tooltip trigger="hover" ref="sendTip" placement="top" :disabled="windowTouch || showMenu"
                                :content="$t(sendContent)">
                                <div v-if="loading">
                                    <div class="chat-load">
                                        <Loading />
                                    </div>
                                </div>
                                <div v-else>
                                    <transition name="mobile-send">
                                        <i v-if="sendClass === 'recorder'" class="taskfont">&#xe609;</i>
                                    </transition>
                                    <transition name="mobile-send">
                                        <i v-if="sendClass !== 'recorder'" class="taskfont">&#xe606;</i>
                                    </transition>
                                </div>
                            </n-tooltip>
                        </template>

                        <div class="chat-input-popover-item" @click="onSend('silence')">
                            <i class="taskfont">&#xe7d7;</i>
                            {{ $t('无声发送') }}
                        </div>
                        <div class="chat-input-popover-item" @click="onSend('md')">
                            <i class="taskfont">&#xe647;</i>
                            {{ $t('Markdown 格式发送') }}
                        </div>
                    </n-popover>
                </li>

                <!-- 录音效果 -->
                <li class="chat-record-recwave">
                    <div ref="recwave"></div>
                </li>
            </ul>

            <!-- 覆盖层 -->
            <div class="chat-cover" @click.stop="onClickCover"></div>
        </div>

        <!-- 移动端表情（底部） -->
        <ChatEmoji v-if="emojiBottom && showEmoji" @on-select="onSelectEmoji" :searchKey="emojiQuickKey" />

        <!-- 录音浮窗 -->
        <transition name="fade">
            <div v-if="['ready', 'ing'].includes(recordState)" v-transfer-dom :data-transfer="true"
                class="chat-input-record-transfer" :class="{ cancel: touchLimitY }" :style="recordTransferStyle"
                @click="stopRecord">
                <div v-if="recordDuration > 0" class="record-duration">{{ recordFormatDuration }}</div>
                <div v-else class="record-loading">
                    <Loading />
                </div>
                <div class="record-cancel" @click.stop="stopRecord(true)">{{ $t(touchLimitY ? '松开取消' : '向上滑动取消') }}</div>
            </div>
        </transition>
    </div>
</template>

<script setup lang="ts">
import Quill from 'quill';
import "quill-mention-hi";
import ChatEmoji from "./emoji";
import { Store } from "le5le-store";
import { defineProps, ref, computed, getCurrentInstance, onBeforeUnmount, watch, toRefs, defineEmits } from "vue"
import { DiaLogStore } from "@/store/dialog"
import { TaskStore } from '@/store/task';
import { GlobalStore } from "@/store/index"
import webTs from '@/utils/web';
import utils from '@/utils/utils';
import { useMessage } from 'naive-ui'
import { UserStore } from '@/store/user';

const props = defineProps({
    value: {
        type: [String, Number],
        default: ''
    },
    dialogId: {
        type: Number,
        default: 0
    },
    taskId: {
        type: Number,
        default: 0
    },
    placeholder: {
        type: String,
        default: ''
    },
    disabled: {
        type: Boolean,
        default: false
    },
    disabledRecord: {
        type: Boolean,
        default: false
    },
    loading: {
        type: Boolean,
        default: false
    },
    enterSend: {
        type: [String, Boolean],
        default: null
    },
    emojiBottom: {
        type: Boolean,
        default: false
    },
    options: {
        type: Object,
        default: () => ({})
    },
    maxlength: {
        type: Number
    },
    defaultMenuOrientation: {
        type: String,
        default: "top"
    },
})

const message = useMessage()
const quill = ref(null)
const isFocus = ref(false)
const rangeIndex = ref(0)
const _content = ref('')
const _options = ref({})
const mentionMode = ref('')
const userList = ref(null)
const userCache = ref(null)
const taskList = ref(null)
const fileList = ref({})
const showMenu = ref(false)
const showMore = ref(false)
const showEmoji = ref(false)
const emojiQuickShow = ref(false)
const emojiQuickTimer = ref(null)
const emojiQuickKey = ref('')
const emojiQuickItems = ref([])
const observer = ref(null)
const wrapperWidth = ref(0)
const wrapperHeight = ref(0)
const editorHeight = ref(0)
const recordReady = ref(false)
const recordRec = ref(null,)
const recordBlob = ref(null)
const recordWave = ref(null)
const recordInter = ref(null)
const recordState = ref("stop")
const recordDuration = ref(0)
const touchStart = ref<any>({})
const touchLimitX = ref(false)
const touchLimitY = ref(false)
const pasteClean = ref(true)
const isSpecVersion = ref()
const instance = ref<any>(getCurrentInstance())
const editor = ref(null)
const sendTip = ref(null)
const recwave = ref(null)
const emojiQuickRef = ref(null)
const dialogMsgs = DiaLogStore().dialogMsgs
const taskStore = TaskStore()
const cacheProjects = TaskStore().cacheProjects
const cacheTasks = TaskStore().cacheTasks
const cacheUserBasic = TaskStore().cacheUserBasic
const cacheDialogs = TaskStore().cacheDialogs
const windowTouch = GlobalStore().windowTouch
const windowScrollY = GlobalStore().windowScrollY
const windowLandscape = GlobalStore().windowLandscape
const windowPortrait = GlobalStore().windowPortrait
const userId = UserStore().info.id
const { value, dialogId, taskId, disabled, loading, emojiBottom } = toRefs(props)

const emit = defineEmits(['on-emoji-visible-change', 'on-focus', 'on-blur', 'on-record-state', 'on-height-change', 'input', 'on-ready', 'on-send', 'on-record', 'on-more','on-file'])

onMounted(() => {
    init();
    //
    observer.value = new ResizeObserver(entries => {
        entries.some(({ target, contentRect }) => {
            if (target === instance.value.vnode.el) {
                wrapperWidth.value = contentRect.width;
                wrapperHeight.value = contentRect.height;
            } else if (target === editor.value) {
                editorHeight.value = contentRect.height;
            }
        })
    });
    observer.value.observe(instance.value.vnode.el);
    observer.value.observe(editor.value);
    //
    recordInter.value = setInterval(_ => {
        if (recordState.value === 'ing') {
            // 录音中，但录音时长不增加则取消录音
            if (instance.value.vnode.__recordDuration && instance.value.vnode.__recordDuration === recordDuration.value) {
                instance.value.vnode.__recordDuration = null;
                stopRecord(true);
                message.error($t('录音失败，请重试'))
            } else {
                instance.value.vnode.__recordDuration = recordDuration.value;
            }
        }
    }, 1000)
    //isEEUiApp 专用
    // if (this.$isEEUiApp) {
    //     window.__onPermissionRequest = (type, result) => {
    //         if (type === 'recordAudio' && result === false) {
    //             // Android 录音权限被拒绝了
    //             this.stopRecord(true);
    //         }
    //     }
    // }
})


onBeforeUnmount(() => {
    if (quill.value) {
        quill.value = null
    }
    if (recordRec.value) {
        recordRec.value = null
    }
    if (observer.value) {
        observer.value.disconnect()
        observer.value = null
    }
    if (recordInter.value) {
        clearInterval(recordInter.value)
    }
})



const isEnterSend = computed(() => {
    if (typeof props.enterSend === "boolean") {
        return props.enterSend;
    } else {
        return !windowTouch
    }
})

const canCall = computed(() => {
    return dialogData.value.type === 'user' && !dialogData.value.bot
    // && this.$isEEUiApp 客户端专用
})

const canAnon = computed(() => {
    return dialogData.value.type === 'user' && !dialogData.value.bot
})

const editorStyle = computed(() => {
    if (wrapperWidth.value > 0
        && editorHeight.value > 0
        && (wrapperWidth.value < 280 || editorHeight.value > 40)) {
        return {
            width: '100%'
        };
    } else {
        return {};
    }
})

const recordTransferStyle = computed(() => {
    return windowScrollY > 0 ? {
        marginTop: (windowScrollY / 2) + 'px'
    } : null
})

const boxClass = computed(() => {
    const array = [];
    if (['ready', 'ing'].includes(recordState.value)) {
        if (recordState.value === 'ing' && recordDuration.value > 0) {
            array.push('record-progress');
        } else {
            array.push('record-ready');
        }
    }
    if (showMenu.value) {
        array.push('show-menu');
    }
    if (showMore.value) {
        array.push('show-more');
    }
    if (showEmoji.value) {
        array.push('show-emoji');
    }
    if (mentionMode.value) {
        array.push(mentionMode.value);
    }
    return array
})

const sendClass = computed(() => {
    if (props.value) {
        return 'sender';
    }
    if (recordReady.value) {
        return 'recorder'
    }
    return ''
})

const sendContent = computed(() => {

    if (sendTip && sendTip.$refs.popper) {
        sendTip.$refs.popper.style.visibility = 'hidden'
        sendTip.showPopper = false
        setTimeout(_ => {
            if (sendTip.$refs.popper) {
                sendTip.$refs.popper.style.visibility = 'visible'
            }
        }, 300)
    }
    return sendClass.value === 'recorder' ? $t('长按录音') : $t('发送')
})

const recordFormatDuration = computed(() => {
    let minute = Math.floor(recordDuration.value / 60000),
        seconds = Math.floor(recordDuration.value / 1000) % 60,
        millisecond = ("00" + recordDuration.value % 1000).substr(-2)
    if (minute < 10) minute = Number(`0${minute}`)
    if (seconds < 10) seconds = Number(`0${seconds}`)
    return `${minute}:${seconds}″${millisecond}`
})

const dialogData = computed(() => {
    return props.dialogId > 0 ? (cacheDialogs.find(({ id }) => id == props.dialogId) || {}) : {};
})

const quoteUpdate = computed(() => {
    return dialogData.value.extra_quote_type === 'update'
})

const quoteData = computed(() => {
    const { extra_quote_id } = dialogData.value;
    if (extra_quote_id) {
        return dialogMsgs.find(item => item.id === extra_quote_id)
    }
    return null;
})


// Watch content change
watch(value, (val: any) => {
    if (quill.value) {
        if (val && val !== _content.value) {
            _content.value = val
            setContent(val)
        } else if (!val) {
            quill.value.setText('')
        }
    }
    // this.$store.dispatch("saveDialogDraft", { id: props.dialogId, extra_draft_content: val })
})

// Watch disabled change
watch(disabled, (val) => {
    quill.value?.enable(!val)
})

// Reset lists
watch(dialogId, () => {
    userList.value = null;
    userCache.value = null;
    taskList.value = null;
    fileList.value = {};
    loadInputDraft()
})
watch(taskId, () => {
    userList.value = null;
    userCache.value = null;
    taskList.value = null;
    fileList.value = {};
    loadInputDraft()
})

watch(dialogData.value.extra_draft_content, () => {
    if (isFocus.value) {
        return
    }
    loadInputDraft()
})

watch(showMenu, (val) => {
    if (val) {
        // this.showMenu = false;
        showMore.value = false;
        showEmoji.value = false;
        emojiQuickShow.value = false;
    }
})

watch(showMore, (val) => {
    if (val) {
        showMenu.value = false;
        // this.showMore = false;
        showEmoji.value = false;
        emojiQuickShow.value = false;
    }
})

watch(showEmoji, (val) => {
    if (props.emojiBottom) {
        if (val) {
            quill.value.enable(false)
        } else if (!props.disabled) {
            quill.value.enable(true)
        }
    }
    if (val) {
        let text = props.value.toString().replace(/&nbsp;/g, " ")
        text = text.replace(/<[^>]+>/g, "")
        if (text
            && text.indexOf(" ") === -1
            && text.length >= 1
            && text.length <= 4) {
            emojiQuickKey.value = text;
        } else {
            emojiQuickKey.value = "";
        }
        //
        showMenu.value = false;
        showMore.value = false;
        // this.showEmoji = false;
        emojiQuickShow.value = false;
        if (quill.value) {
            const range = quill.value.selection.savedRange;
            rangeIndex.value = range ? range.index : 0
        }
    } else if (rangeIndex.value > 0) {
        quill.value.setSelection(rangeIndex.value)
    }
    emit('on-emoji-visible-change', val)
})

watch(emojiQuickShow, (val) => {
    if (val) {
        showMenu.value = false;
        showMore.value = false;
        showEmoji.value = false;
        // this.emojiQuickShow = false;
    }
})

watch(isFocus, (val) => {
    if (timerScroll) {
        clearInterval(timerScroll);
    }
    if (val) {
        emit('on-focus')
        hidePopover
        if (isSpecVersion.value) {
            // ios11.0-11.3 对scrollTop及scrolIntoView解释有bug
            // 直接执行会导致输入框滚到底部被遮挡
        } else if (windowPortrait) {
            var timerScroll = setInterval(() => {
                if (quill.value?.hasFocus()) {
                    // windowScrollY > 0 && $A.scrollIntoViewIfNeeded(this.$refs.editor); 滚动重写
                } else {
                    clearInterval(timerScroll);
                }
            }, 200);
        }
    } else {
        emit('on-blur')
    }
})

watch(recordState, (state) => {
    if (state === 'ing') {
        recordWave.value = window.Recorder.FrequencyHistogramView({
            elem: recwave.value,
            lineCount: 90,
            position: 0,
            minHeight: 1,
            stripeEnable: false
        })
    } else {
        recordWave.value = null
        recwave.value.innerHTML = ""
    }
    emit('on-record-state', state)
})

watch(wrapperHeight, (newVal, oldVal) => {
    emit('on-height-change', { newVal, oldVal })
})


const init = () => {
    // Options
    _options.value = Object.assign({
        theme: 'bubble',
        readOnly: false,
        placeholder: props.placeholder,
        modules: {
            toolbar: [
                ['bold', 'strike', 'italic', 'underline', { 'list': 'ordered' }, { 'list': 'bullet' }, 'blockquote', 'code-block']
            ],
            keyboard: {
                bindings: {
                    'short enter': {
                        key: 13,
                        shortKey: true,
                        handler: _ => {
                            if (!isEnterSend.value) {
                                onSend();
                                return false;
                            }
                            return true;
                        }
                    },
                    'enter': {
                        key: 13,
                        shiftKey: false,
                        handler: _ => {
                            if (isEnterSend.value) {
                                onSend('');
                                return false;
                            }
                            return true;
                        }
                    },
                    'esc': {
                        key: 27,
                        shiftKey: false,
                        handler: _ => {
                            if (emojiQuickShow.value) {
                                emojiQuickShow.value = false;
                                return false;
                            }
                            return true;
                        }
                    }
                }
            },
            mention: {
                allowedChars: /^\S*$/,
                mentionDenotationChars: ["@", "#", "~"],
                defaultMenuOrientation: props.defaultMenuOrientation,
                isolateCharacter: true,
                positioningStrategy: 'fixed',
                renderItem: (data) => {
                    if (data.disabled === true) {
                        return `<div class="mention-item-disabled">${data.value}</div>`;
                    }
                    if (data.id === 0) {
                        return `<div class="mention-item-at">@</div><div class="mention-item-name">${data.value}</div><div class="mention-item-tip">${data.tip}</div>`;
                    }
                    if (data.avatar) {
                        const botHtml = data.bot ? `<div class="taskfont mention-item-bot">&#xe68c;</div>` : ''
                        return `<div class="mention-item-img${data.online ? ' online' : ''}"><img src="${data.avatar}"/><em></em></div>${botHtml}<div class="mention-item-name">${data.value}</div>`;
                    }
                    if (data.tip) {
                        return `<div class="mention-item-name" title="${data.value}">${data.value}</div><div class="mention-item-tip">${data.tip}</div>`;
                    }
                    return `<div class="mention-item-name" title="${data.value}">${data.value}</div>`;
                },
                renderLoading: () => {
                    return "Loading...";
                },
                source: (searchTerm, renderList, mentionChar) => {
                    const mentionName = mentionChar == "@" ? 'user-mention' : (mentionChar == "#" ? 'task-mention' : 'file-mention');
                    const containers = document.getElementsByClassName("ql-mention-list-container");
                    for (let i = 0; i < containers.length; i++) {
                        containers[i].classList.remove("user-mention");
                        containers[i].classList.remove("task-mention");
                        containers[i].classList.remove("file-mention");
                        containers[i].classList.add(mentionName);
                        utils.scrollPreventThrough(containers[i]);
                    }
                    let mentionSourceCache = null;
                    getMentionSource(mentionChar, searchTerm, array => {
                        const values = [];
                        array.some(item => {
                            let list = item.list;
                            if (searchTerm) {
                                list = list.filter(({ value }) => utils.strExists(value, searchTerm));
                            }
                            if (list.length > 0) {
                                item.label && values.push(...item.label)
                                values.push(...list)
                            }
                        })
                        if (utils.jsonStringify(values.map(({ id }) => id)) !== mentionSourceCache) {
                            mentionSourceCache = utils.jsonStringify(values.map(({ id }) => id))
                            renderList(values, searchTerm);
                        }
                    })
                }
            }
        }
    }, props.options)

    // Instance
    quill.value = new Quill(editor.value, _options.value)
    quill.value.enable(!props.disabled)

    // Set editor content
    if (props.value) {
        setContent(props.value)
    } else {
        loadInputDraft()
    }

    // Mark model as touched if editor lost focus
    quill.value.on('selection-change', range => {
        if (!range) {
            // 修复光标会超出的问题
            if (quill.value.hasFocus()) {
                quill.value.setSelection(0)
                return
            }
            if (document.activeElement && document.activeElement.className === 'ql-clipboard') {
                quill.value.setSelection(quill.value.getLength())
                return
            }
        }
        isFocus.value = !!range;
    })

    // Update model if text changes
    quill.value.on('text-change', _ => {
        if (props.maxlength > 0 && quill.value.getLength() > props.maxlength) {
            quill.value.deleteText(props.maxlength, quill.value.getLength());
        }
        let html = editor.value.children[0].innerHTML
        html = html.replace(/^(<p>\s*<\/p>)+|(<p>\s*<\/p>)+$/gi, '')
        html = html.replace(/^(<p><br\/*><\/p>)+|(<p><br\/*><\/p>)+$/gi, '')
        updateEmojiQuick(html)
        _content.value = html
        emit('input', _content.value)
        nextTick(() => {
            const range = quill.value.getSelection();
            if (range) {
                const endText = quill.value.getText(range.index);
                /^\n\n$/.test(endText) && quill.value.deleteText(range.index, 1);
            }
        })
    })

    // Clipboard Matcher (保留图片跟空格，清除其余所以样式)
    quill.value.clipboard.addMatcher(Node.ELEMENT_NODE, (node, delta) => {
        if (!pasteClean.value) {
            return delta
        }
        delta.ops = delta.ops.map(op => {
            const obj: any = {
                insert: op.insert
            };
            try {
                // 替换 mention 对象为纯文本
                if (typeof obj.insert.mention === "object" && node.innerHTML) {
                    obj.insert = node.innerHTML.replace(/<[^>]+>/g, "")
                }
            } catch (e) { }
            if (op.attributes) {
                ['bold', 'strike', 'italic', 'underline', 'list', 'blockquote', 'link'].some(item => {
                    if (op.attributes[item]) {
                        if (typeof obj.attributes === "undefined") {
                            obj.attributes = {}
                        }
                        obj.attributes[item] = op.attributes[item]
                    }
                })
            }
            return obj
        })
        return delta
    })

    // Ready event
    emit('on-ready', quill.value)

    // Load recorder
    if (!props.disabledRecord) {
        utils.loadScriptS([
            'js/recorder/recorder.mp3.min.js',
            'js/recorder/lib.fft.js',
            'js/recorder/frequency.histogram.view.js',
        ]).then(_ => {
            if (typeof window.Recorder !== 'function') {
                return;
            }
            recordRec.value = window.Recorder({
                type: "mp3",
                bitRate: 64,
                sampleRate: 32000,
                onProcess: (buffers, powerLevel, duration, sampleRate, newBufferIdx, asyncEnd) => {
                    recordWave.value.input(buffers[buffers.length - 1], powerLevel, sampleRate);
                    recordDuration.value = duration;
                    if (duration >= 3 * 60 * 1000) {
                        // 最长录3分钟
                        stopRecord(false);
                    }
                }
            })
            if (window.Recorder.Support()) {
                recordReady.value = true;
            }
        });
    }
}

const updateEmojiQuick = (text: any) => {
    if (!isFocus.value || !text) {
        emojiQuickShow.value = false
        return
    }
    emojiQuickTimer.value && clearTimeout(emojiQuickTimer.value)
    emojiQuickTimer.value = setTimeout(_ => {
        text = text.replace(/&nbsp;/g, " ")
        text = text.replace(/<[^>]+>/g, "")
        if (text
            && text.indexOf(" ") === -1
            && text.length >= 1
            && text.length <= 4
            && utils.isArray(window.emoticonData)) {
            // 显示快捷选择表情窗口
            emojiQuickItems.value = [];
            let baseUrl = webTs.apiUrl("../images/emoticon") //接口重写
            window.emoticonData.some(data => {
                let item = data.list.find(d => utils.strExists(d.name + (d.key ? ` ${d.key}` : ''), text))
                if (item) {
                    emojiQuickItems.value.push(Object.assign(item, {
                        type: `emoticon`,
                        asset: `images/emoticon/${data.path}/${item.path}`,
                        src: `${baseUrl}/${data.path}/${item.path}`
                    }))
                    if (emojiQuickItems.value.length >= 3) {
                        return true
                    }
                }
            });
            if (emojiQuickItems.value.length > 0) {
                nextTick(() => {
                    emojiQuickShow.value = true
                    emojiQuickRef.value.updatePopper()
                })
                return
            }
        }
        emojiQuickShow.value = false
    }, 100)
}

const getText = () => {
    if (quill.value) {
        return quill.value.getText()
    }
    return "";
}

const setText = (value: any) => {
    if (quill.value) {
        quill.value.setText(value)
    }
}

const setContent = (value) => {
    if (quill.value) {
        quill.value.setContents(quill.value.clipboard.convert(value))
    }
}

const setPasteMode = (bool) => {
    pasteClean.value = bool
}

const loadInputDraft = () => {
    const { extra_draft_content } = dialogData.value;
    if (extra_draft_content) {
        pasteClean.value = false
        emit('input', extra_draft_content)
        nextTick(() => pasteClean.value = true)
    } else {
        emit('input', '')
    }
}

const onClickEditor = () => {
    taskStore.$patch((state) => {
        state.messengerSearchKey = { dialog: '', contacts: '' }
    })
    updateEmojiQuick(props.value)
}

const focus = () => {
    nextTick(() => {
        if (quill.value) {
            quill.value.setSelection(quill.value.getLength())
            quill.value.focus()
        }
    })
}

const blur = () => {
    nextTick(() => {
        quill.value && quill.value.blur()
    })
}

const clickSend = (action: any, event: any) => {
    if (props.loading) {
        return;
    }
    switch (action) {
        case 'down':
            touchLimitX.value = false;
            touchLimitY.value = false;
            touchStart.value = event.type === "touchstart" ? event.touches[0] : event;
            if ((event.button === undefined || event.button === 0) && startRecord()) {
                return;
            }
            break;

        case 'move':
            const touchMove = event.type === "touchmove" ? event.touches[0] : event;
            touchLimitX.value = (touchStart.value.clientX - touchMove.clientX) / window.innerWidth > 0.1
            touchLimitY.value = (touchStart.value.clientY - touchMove.clientY) / window.innerHeight > 0.1
            break;

        case 'up':
            if (showMenu.value) {
                return;
            }
            if (stopRecord(touchLimitY.value)) {
                return;
            }
            if (touchLimitY.value || touchLimitX.value) {
                return; // 移动了 X、Y 轴
            }
            onSend('')
            break;
    }
}

const longSend = () => {
    if (sendClass.value === 'recorder') {
        return;
    }
    showMenu.value = true;
}

const onSend = (type: any) => {
    hidePopover('send')
    rangeIndex.value = 0
    taskStore.$patch((state) => {
        state.messengerSearchKey = { dialog: '', contacts: '' }
    })
    if (type) {
        emit('on-send', null, type)
    } else {
        emit('on-send')
    }
}

const startRecord = () => {
    if (sendClass.value === 'recorder') {
        // this.$store.dispatch("audioStop", true) 重写接口
        recordDuration.value = 0;
        recordState.value = "ready";
        nextTick(() => {
            recordRec.value.open(_ => {
                if (recordState.value === "ready") {
                    recordState.value = "ing"
                    recordBlob.value = null
                    setTimeout(_ => {
                        recordRec.value.start()
                    }, 300)
                } else {
                    recordRec.value.close();
                }
            }, (msg) => {
                recordState.value = "stop";
                message.error(msg || $t('打开录音失败'))
            });
        })
        return true;
    } else {
        return false;
    }
}

const stopRecord = (isCancel: any) => {
    switch (recordState.value) {
        case "ing":
            recordState.value = "stop";
            recordRec.value.stop((blob, duration) => {
                recordRec.value.close();
                if (isCancel === true) {
                    return;
                }
                if (duration < 600) {
                    message.warning($t("说话时间太短")) // 小于 600ms 不发送
                } else {
                    recordBlob.value = blob;
                    uploadRecord(duration);
                }
            }, (msg) => {
                recordRec.close();
                message.error($t('录音失败'))
            });
            return true;

        case "ready":
            recordState.value = "stop";
            return true;

        default:
            recordState.value = "stop";
            return false;
    }
}

const hidePopover = (action: any) => {
    showMenu.value = false;
    showMore.value = false;
    if (action === 'send') {
        return
    }
    showEmoji.value = false;
    emojiQuickShow.value = false;
}

const onClickCover = () => {
    hidePopover('');
    nextTick(() => {
        quill.value?.focus()
    })
}

const uploadRecord = (duration: any) => {
    if (recordBlob.value === null) {
        return;
    }
    const reader = new FileReader();
    reader.onloadend = () => {
        emit('on-record', {
            type: recordBlob.value.type,
            base64: reader.result,
            duration,
        })
    };
    reader.readAsDataURL(recordBlob.value);
}

const onEmojiQuick = (item: any) => {
    if (item.type === 'online') {
        emit('input', "")
        emit('on-send', `<img src="${item.src}"/>`)
    } else {
        emit('input', "")
        emit('on-send', `<img class="emoticon" data-asset="${item.asset}" data-name="${item.name}" src="${item.src}"/>`)
    }
    emojiQuickShow.value = false
    focus()
}

const onSelectEmoji = (item: any) => {
    if (!quill.value) {
        return;
    }
    if (item.type === 'emoji') {
        quill.value.insertText(rangeIndex.value, item.text);
        rangeIndex.value += item.text.length
        if (windowLandscape) {
            showEmoji.value = false;
        }
    } else if (item.type === 'emoticon') {
        emit('on-send', `<img class="emoticon" data-asset="${item.asset}" data-name="${item.name}" src="${item.src}"/>`)
        if (item.asset === "emosearch") {
            emit('input', "")
        }
        if (windowLandscape) {
            showEmoji.value = false;
        }
    }
}

const onToolbar = (action: any) => {
    hidePopover('');
    switch (action) {
        case 'user':
            openMenu("@");
            break;

        case 'task':
            openMenu("#");
            break;

        case 'meeting':
            Store.set('addMeeting', {
                type: 'create',
                dialog_id: props.dialogId,
                userids: [userId],
            });
            break;

        case 'image':
        case 'file':
        case 'call':
        case 'anon':
            emit('on-more', action)
            break;
    }
}

const onMoreVisibleChange = (v: any) => {
    showMore.value = v;
}

const setQuote = (id: any, type = 'reply') => {
    // 重写接口
    // props.dialogId > 0 && this.$store.dispatch("saveDialog", {
    //     id: props.dialogId,
    //     extra_quote_id: id,
    //     extra_quote_type: type === 'update' ? 'update' : 'reply'
    // });
}

const cancelQuote = () => {
    if (quoteUpdate.value) {
        emit('input', '')
    }
    setQuote(0)
}

const openMenu = (char: any) => {
    if (!quill.value) {
        return;
    }
    if (props.value.toString().length === 0 || props.value.toString().endsWith("<p><br></p>")) {
        quill.value.getModule("mention").openMenu(char);
    } else {
        let str = props.value.toString().replace(/<[^>]+>/g, "");
        if (str.length === 0 || str.endsWith(" ")) {
            quill.value.getModule("mention").openMenu(char);
        } else {
            quill.value.getModule("mention").openMenu(` ${char}`);
        }
    }
}

const addMention = (data: any) => {
    if (!quill.value) {
        return;
    }
    quill.value.getModule("mention").insertItem(data, true);
}

const getProjectId = () => {
    let object = null;
    if (props.dialogId > 0) {
        object = cacheProjects.find(({ dialog_id }) => dialog_id == props.dialogId);
        if (object) {
            return object.id;
        }
        object = cacheTasks.find(({ dialog_id }) => dialog_id == props.dialogId);
        if (object) {
            return object.project_id;
        }
    } else if (props.taskId > 0) {
        object = cacheTasks.find(({ id }) => id == props.taskId);
        if (object) {
            return object.project_id;
        }
    }
    return 0;
}

const getMentionSource = (mentionChar: any, searchTerm: any, resultCallback: any) => {
    switch (mentionChar) {
        case "@": // @成员
            mentionMode.value = "user-mention";
            const atCallback = (list) => {
                getMoreUser(searchTerm, list.map(item => item.id)).then(moreUser => {
                    userList.value = list
                    userCache.value = [];
                    if (moreUser.length > 0) {
                        if (list.length > 2) {
                            userCache.value.push({
                                label: null,
                                list: [{ id: 0, value: $t('所有人'), tip: $t('仅提示会话内成员') }]
                            })
                        }
                        userCache.value.push(...[{
                            label: [{ id: 0, value: $t('会话内成员'), disabled: true }],
                            list,
                        }, {
                            label: [{ id: 0, value: $t('会话以外成员'), disabled: true }],
                            list: moreUser,
                        }])
                    } else {
                        if (list.length > 2) {
                            userCache.value.push(...[{
                                label: null,
                                list: [{ id: 0, value: $t('所有人'), tip: $t('提示所有成员') }]
                            }, {
                                label: [{ id: 0, value: $t('会话内成员'), disabled: true }],
                                list,
                            }])
                        } else {
                            userCache.value.push({
                                label: null,
                                list
                            })
                        }
                    }
                    resultCallback(userCache.value)
                })
            }
            //
            if (dialogData.value.people && utils.arrayLength(userList.value) !== dialogData.value.people) {
                userList.value = null;
                userCache.value = null;
            }
            if (userCache.value !== null) {
                resultCallback(userCache.value)
            }
            if (userList.value !== null) {
                atCallback(userList.value)
                return;
            }
            //
            const array = [];
            if (props.dialogId > 0) {
                // 根据会话ID获取成员 重写接口
                // this.$store.dispatch("call", {
                //     url: 'dialog/user',
                //     data: {
                //         dialog_id: props.dialogId,
                //         getuser: 1
                //     }
                // }).then(({ data }) => {
                //     if (cacheDialogs.find(({ id }) => id == props.dialogId)) {
                //         this.$store.dispatch("saveDialog", {
                //             id: props.dialogId,
                //             people: data.length
                //         })
                //     }
                //     if (data.length > 0) {
                //         array.push(...data.map(item => {
                //             return {
                //                 id: item.userid,
                //                 value: item.nickname,
                //                 avatar: item.userimg,
                //                 online: item.online,
                //                 bot: item.bot,
                //             }
                //         }))
                //     }
                //     atCallback(array)
                // }).catch(_ => {
                //     atCallback(array)
                // });
            } else if (props.taskId > 0) {
                // 根据任务ID获取成员
                const task = cacheTasks.find(({ id }) => id == props.taskId)
                if (task && utils.isArray(task.task_user)) {
                    task.task_user.some(tmp => {
                        const item = cacheUserBasic.find(({ userid }) => userid == tmp.userid);
                        if (item) {
                            array.push({
                                id: item.userid,
                                value: item.nickname,
                                avatar: item.userimg,
                                online: item.online,
                                bot: item.bot,
                            })
                        }
                    })
                }
                atCallback(array)
            }
            break;

        case "#": // #任务
            mentionMode.value = "task-mention";
            if (taskList.value !== null) {
                resultCallback(taskList.value)
                return;
            }
            const taskCallback = (list) => {
                taskList.value = [];
                // 项目任务
                if (list.length > 0) {
                    list = list.map(item => {
                        return {
                            id: item.id,
                            value: item.name,
                            tip: item.complete_at ? $t('已完成') : null,
                        }
                    }).splice(0, 100)
                    taskList.value.push({
                        label: [{ id: 0, value: $t('项目任务'), disabled: true }],
                        list,
                    })
                }
                // 待完成任务
                let dataA = this.$store.getters.transforTasks(this.$store.getters.dashboardTask['all']);
                if (dataA.length > 0) {
                    dataA = dataA.sort((a, b) => {
                        return Number(utils.Date(a.end_at || "2099-12-31 23:59:59")) - Number(utils.Date(b.end_at || "2099-12-31 23:59:59"));
                    }).splice(0, 100)
                    taskList.value.push({
                        label: [{ id: 0, value: $t('我的待完成任务'), disabled: true }],
                        list: dataA.map(item => {
                            return {
                                id: item.id,
                                value: item.name
                            }
                        }),
                    })
                }
                // 我协助的任务
                let dataB = this.$store.getters.assistTask;
                if (dataB.length > 0) {
                    dataB = dataB.sort((a, b) => {
                        return Number(utils.Date(a.end_at || "2099-12-31 23:59:59"))  - Number(utils.Date(b.end_at || "2099-12-31 23:59:59"));
                    }).splice(0, 100)
                    taskList.value.push({
                        label: [{ id: 0, value: $t('我协助的任务'), disabled: true }],
                        list: dataB.map(item => {
                            return {
                                id: item.id,
                                value: item.name
                            }
                        }),
                    })
                }
                resultCallback(taskList.value)
            }
            //
            const projectId = getProjectId();
            if (projectId > 0) {
                // 重写接口
                // this.$store.dispatch("getTaskForProject", projectId).then(_ => {
                //     const tasks = cacheTasks.filter(task => {
                //         if (task.archived_at) {
                //             return false;
                //         }
                //         return task.project_id == projectId
                //             && task.parent_id === 0
                //             && !task.archived_at
                //     }).sort((a, b) => {
                //         return $A.Date(b.complete_at || "2099-12-31 23:59:59") - $A.Date(a.complete_at || "2099-12-31 23:59:59")
                //     })
                //     if (tasks.length > 0) {
                //         taskCallback(tasks)
                //     } else {
                //         taskCallback([])
                //     }
                // }).catch(_ => {
                //     taskCallback([])
                // })
                return;
            }
            taskCallback([])
            break;

        case "~": // ~文件
            mentionMode.value = "file-mention";
            if (utils.isArray(fileList.value[searchTerm])) {
                resultCallback(fileList.value[searchTerm])
                return;
            }
            fileTimer && clearTimeout(fileTimer)
            var fileTimer = setTimeout(_ => {
                // 重写接口
                // this.$store.dispatch("searchFiles", searchTerm).then(({ data }) => {
                //     fileList.value[searchTerm] = [{
                //         label: [{ id: 0, value: this.$t('文件分享查看'), disabled: true }],
                //         list: data.filter(item => item.type !== "folder").map(item => {
                //             return {
                //                 id: item.id,
                //                 value: item.ext ? `${item.name}.${item.ext}` : item.name
                //             }
                //         })
                //     }];
                //     resultCallback(this.fileList[searchTerm])
                // }).catch(() => {
                //     resultCallback([])
                // })
            }, 300)
            break;

        default:
            resultCallback([])
            break;
    }
}

const getMoreUser = (key, existIds) => {
    return new Promise(resolve => {
        const { owner_id, type } = dialogData.value
        const permission = type === 'group' && [0, userId].includes(owner_id)
        if (props.taskId > 0 || permission) {
            __getMoreTimer && clearTimeout(__getMoreTimer)
            var __getMoreTimer = setTimeout(_ => {
                // 重写接口
                // this.$store.dispatch("call", {
                //     url: 'users/search',
                //     data: {
                //         keys: {
                //             key,
                //         },
                //         state: 1,
                //         take: 30
                //     },
                // }).then(({ data }) => {
                //     const moreUser = data.filter(item => !existIds.includes(item.userid))
                //     resolve(moreUser.map(item => {
                //         return {
                //             id: item.userid,
                //             value: item.nickname,
                //             avatar: item.userimg,
                //             online: !!item.online,
                //         }
                //     }))
                // }).catch(_ => {
                //     resolve([])
                // });
            }, userCache.value === null ? 0 : 600)
        } else {
            resolve([])
        }
    })
}

const checkIOSVersion = () => {
    let ua = window && window.navigator && window.navigator.userAgent;
    let match = ua.match(/OS ((\d+_?){2,3})\s/i);
    let IOSVersion = match ? match[1].replace(/_/g, ".") : "unknown";
    const iosVsn = IOSVersion.split(".");
    return +iosVsn[0] == 11 && +iosVsn[1] >= 0 && +iosVsn[1] < 3;
}
isSpecVersion.value = checkIOSVersion()

const handlePaste = (e:any) => {
    const files = Array.prototype.slice.call(e.clipboardData.files)
    const postFiles = files.filter(file => !utils.leftExists(file.type, 'image/'));
    if (postFiles.length > 0) {
        e.preventDefault()
        emit('on-file', files)
    } else if (pasteRtf(e)) {
        e.preventDefault()
    }
}

const pasteRtf = (e) => {
    if (e && e.clipboardData && e.clipboardData.items) {
        const imgHtml = (new DOMParser).parseFromString(e.clipboardData.getData("text/html") || "", "text/html").querySelector("img");
        if (!imgHtml) {
            const array = [];
            let image = null;
            if (e.clipboardData.types && -1 != [].indexOf.call(e.clipboardData.types, "text/rtf") || e.clipboardData.getData("text/rtf")) {
                image = e.clipboardData.items[0].getAsFile();
                if (image) {
                    array.push(image)
                }
            } else {
                for (let s = 0; s < e.clipboardData.items.length; s++) {
                    image = e.clipboardData.items[s].getAsFile()
                    if (image) {
                        array.push(image)
                    }
                }
            }
            if (array.length > 0) {
                array.forEach(image => {
                    const t = new FileReader;
                    t.onload = ({ target }) => {
                        const length = quill.value.getSelection(true).index;
                        quill.value.insertEmbed(length, "image", target.result);
                        quill.value.setSelection(length + 1)
                    };
                    t.readAsDataURL(image)
                })
                return true
            }
        }
    }
    return false
}


</script>
