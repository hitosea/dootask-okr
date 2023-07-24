<template>
    <div class="teditor-wrapper" @click="onClickWrap" @touchstart="onTouchstart">
        <div class="teditor-box" :class="[!props.inline && spinShow ? 'teditor-loadstyle' : 'teditor-loadedstyle']">
            <template v-if="props.inline">
                <div ref="myTextarea" :id="props.id" v-html="spinShow ? '' : content"></div>
                <Icon v-if="spinShow" type="ios-loading" :size="18" class="icon-loading icon-inline"></Icon>
            </template>
            <template v-else>
                <textarea ref="myTextarea" :id="props.id">{{ content }}</textarea>
                <Spin fix v-if="spinShow">
                    <Icon type="ios-loading" :size="18" class="icon-loading"></Icon>
                    <div>{{ $t('加载组件中...') }}</div>
                </Spin>
            </template>
            <ImgUpload ref="myUpload" class="upload-control" type="callback" :uploadIng.sync="uploadIng"
                @on-callback="editorImage" num="50" />
            <Upload name="files" ref="fileUpload" class="upload-control" :action="actionUrl" :headers="headers" multiple
                :format="uploadFormat" :show-upload-list="false" :max-size="maxSize" :on-progress="handleProgress"
                :on-success="handleSuccess" :on-error="handleError" :on-format-error="handleFormatError"
                :on-exceeded-size="handleMaxSize" :before-upload="handleBeforeUpload" />
            <div class="teditor-operate" :style="operateStyles" v-show="operateVisible">
                <Dropdown trigger="custom" :visible="operateVisible" @on-clickoutside="operateVisible = false" transfer>
                    <div :style="{ userSelect: operateVisible ? 'none' : 'auto', height: operateStyles.height }"></div>
                    <DropdownMenu slot="list">
                        <DropdownItem @click.native="onFull">{{ $t('编辑') }}</DropdownItem>
                        <DropdownItem v-if="operateLink" @click.native="onLinkPreview">{{ $t('打开链接') }}</DropdownItem>
                        <DropdownItem v-if="operateImg" @click.native="onImagePreview">{{ $t('查看图片') }}</DropdownItem>
                    </DropdownMenu>
                </Dropdown>
            </div>
        </div>
        <Spin fix v-if="uploadIng > 0">
            <Icon type="ios-loading" class="icon-loading"></Icon>
            <div>{{ $t('正在上传文件...') }}</div>
        </Spin>
        <Modal v-model="transfer" class="teditor-transfer" @on-visible-change="transferChange" footer-hide fullscreen
            transfer>
            <div slot="close">
                <Button type="primary" size="small">{{ $t('完成') }}</Button>
            </div>
            <div class="teditor-transfer-body">
                <textarea :id="'T_' + props.id">{{ content }}</textarea>
            </div>
            <Spin fix v-if="uploadIng > 0">
                <Icon type="ios-loading" class="icon-loading"></Icon>
                <div>{{ $t('正在上传文件...') }}</div>
            </Spin>
        </Modal>
    </div>
</template>

<script setup lang="ts">
import tinymce from as tinymce;
import ImgUpload from "./ImgUpLoad.vue";
import { languageType } from "../language";
import utils from "@/utils/utils";
import { UserStore } from '@/store/user';

import { GlobalStore } from '@/store';

const props = defineProps({
    id: {
        type: String,
        default: () => {
            return "tinymce_" + Math.round(Math.random() * 10000);
        }
    },
    value: {
        default: ''
    },
    height: {
        default: 360,
    },
    minHeight: {
        type: Number,
        default: 0,
    },
    htmlClass: {
        default: '',
        type: String
    },
    plugins: {
        type: Array,
        default: () => {
            return [
                'advlist autolink lists link image charmap print preview hr anchor pagebreak',
                'searchreplace visualblocks visualchars code',
                'insertdatetime media nonbreaking save table directionality',
                'emoticons paste codesample'
            ];
        }
    },
    toolbar: {
        type: String,
        default: ' undo redo | styleselect | uploadImages | uploadFiles | bold italic underline forecolor backcolor | alignleft aligncenter alignright | bullist numlist outdent indent | link image emoticons media codesample | preview screenload',
    },
    options: {
        type: Object,
        default: () => {
            return {};
        }
    },
    optionFull: {
        type: Object,
        default: () => {
            return {};
        }
    },
    inline: {
        type: Boolean,
        default: false
    },
    readOnly: {
        type: Boolean,
        default: false  // windowTouch 为 true 时，非全屏模式下，readOnly 为 true
    },
    autoSize: {
        type: Boolean,
        default: false
    },
    placeholder: {
        type: String,
        default: ''
    },
    placeholderFull: {
        type: String,
        default: ''
    },
    scrollHideOperateClassName: {   // 触发隐藏操作菜单的滚动区域class
        type: String,
    },
})

const windowTouch = GlobalStore().windowTouch
const userToken = UserStore().info.token
const content = ref('')
const editor = ref(null)
const editorT = ref(null)
const checkerTimeout = ref(null)
const isTyping = ref(false)

const spinShow = ref(true)
const transfer = ref(false)

const uploadIng = ref(0)
const uploadFormat = ref(['jpg', 'jpeg', 'webp', 'png', 'gif', 'doc', 'docx', 'xls', 'xlsx', 'ppt', 'pptx', 'txt', 'esp', 'pdf', 'rar', 'zip', 'gz', 'ai', 'avi', 'bmp', 'cdr', 'eps', 'mov', 'mp3', 'mp4', 'pr', 'psd', 'svg', 'tif'])
const actionUrl = ref($A.apiUrl('system/fileupload'))
const maxSize = ref(10240)

const operateStyles = ref({})
const operateVisible = ref(false)
const operateLink = ref(null)
const operateImg = ref(null)

const timer = ref(null)
const listener = ref(null)

onMounted(() => {
    content.value = props.value;
    init();
    //
    if (props.scrollHideOperateClassName) {
        let parent = this.$parent.$el.parentNode;
        while (parent) {
            if (parent.classList.contains(props.scrollHideOperateClassName)) {
                listener.value = parent;
                parent.addEventListener("scroll", this.onTouchstart);
                break;
            }
            parent = parent.parentNode;
        }
    }
})
onActivated(() => {
    content.value = props.value;
    init();
})
onDeactivated(() => {
    destroy();
})
onBeforeUnmount(() => {
    listener.value?.removeEventListener("scroll", this.onTouchstart);
})


onUnmounted(() => {
    destroy();
})




const headers = computed(() => {
    return {
        fd: utils.getSessionStorageString("userWsFd"),
        token: userToken,
    }
})

watch(() => props.value, (newValue) => {
    if (newValue == null) {
        newValue = "";
    }
    if (!isTyping.value) {
        setContent(newValue);
    }
})
watch(() => props.readOnly, (value) => {
    if (editor.value !== null) {
        if (windowTouch) {
            return;
        }
        if (value) {
            editor.value.setMode('readonly');
        } else {
            editor.value.setMode('design');
        }
    }
})


    init() {
        timer.value && clearTimeout(timer.value);
        this.$nextTick(() => {
            tinymce.init(this.concatAssciativeArrays(this.option(false), props.options));
        });
    },

    initTransfer() {
        this.$nextTick(() => {
            tinymce.init(this.concatAssciativeArrays(this.option(true), props.optionFull));
        });
    },

    destroy() {
        timer.value && clearTimeout(timer.value);
        timer.value = setTimeout(_ => {
            if (editor.value !== null) {
                editor.value.destroy()
                editor.value = null
            }
            if (editorT.value !== null) {
                editorT.value.destroy();
                editorT.value = null;
            }
            spinShow.value = true;
            operateVisible.value = false;
            $A(this.$refs.myTextarea).show();
        }, 500);
    },

    plugin(isFull) {
        if (isFull) {
            return props.plugins.filter((val) => val != 'autoresize');
        } else {
            return props.plugins;
        }
    },

    option(isFull) {
        let lang = languageType;
        switch (languageType) {
            case 'zh':
                lang = "zh_CN";
                break;
            case 'zh-CHT':
                lang = "zh-TW";
                break;
            case 'fr':
                lang = "fr_FR";
                break;
            case 'ko':
                lang = "ko_KR";
                break;
        }
        const optionInfo = {
            inline: isFull ? false : this.inline,
            selector: (isFull ? '#T_' : '#') + props.id,
            base_url: $A.originUrl('js/tinymce'),
            language: lang,
            toolbar: props.toolbar,
            plugins: this.plugin(isFull),
            placeholder: isFull && props.placeholderFull ? props.placeholderFull : props.placeholder,
            save_onsavecallback: (e) => {
                this.$emit('editorSave', e);
            },
            paste_data_images: true,
            menu: {
                view: {
                    title: 'View',
                    items: 'code | visualaid visualchars visualblocks | spellchecker | preview fullscreen screenload | showcomments'
                },
                insert: {
                    title: "Insert",
                    items: "image link media addcomment pageembed template codesample inserttable | charmap emoticons hr | pagebreak nonbreaking anchor toc | insertdatetime | uploadImages | uploadFiles"
                }
            },
            codesample_languages: [
                { text: "HTML/VUE/XML", value: "markup" },
                { text: "JavaScript", value: "javascript" },
                { text: "CSS", value: "css" },
                { text: "PHP", value: "php" },
                { text: "Ruby", value: "ruby" },
                { text: "Python", value: "python" },
                { text: "Java", value: "java" },
                { text: "C", value: "c" },
                { text: "C#", value: "csharp" },
                { text: "C++", value: "cpp" }
            ],
            height: isFull ? '100%' : ($A.rightExists(props.height, '%') ? props.height : ($A.runNum(props.height) || 360)),
            resize: !isFull,
            convert_urls: false,
            toolbar_mode: 'sliding',
            content_css: this.themeIsDark ? 'dark' : 'default',
            setup: (editor) => {
                editor.ui.registry.addMenuButton('uploadImages', {
                    text: $t('图片'),
                    tooltip: $t('上传/浏览 图片'),
                    fetch: (callback) => {
                        let items = [{
                            type: 'menuitem',
                            text: $t('上传本地图片'),
                            onAction: () => {
                                this.$refs.myUpload.handleClick();
                            }
                        }, {
                            type: 'menuitem',
                            text: $t('浏览已上传图片'),
                            onAction: () => {
                                this.$refs.myUpload.browsePicture();
                            }
                        }];
                        callback(items);
                    }
                });
                editor.ui.registry.addNestedMenuItem('uploadImages', {
                    icon: 'image',
                    text: $t('上传图片'),
                    getSubmenuItems: () => {
                        return [{
                            type: 'menuitem',
                            text: $t('上传本地图片'),
                            onAction: () => {
                                this.$refs.myUpload.handleClick();
                            }
                        }, {
                            type: 'menuitem',
                            text: $t('浏览已上传图片'),
                            onAction: () => {
                                this.$refs.myUpload.browsePicture();
                            }
                        }];
                    }
                });
                editor.ui.registry.addMenuItem('imagePreview', {
                    text: $t('预览图片'),
                    onAction: () => {
                        operateImg.value = null
                        const imgElm = editor.selection.getNode();
                        if (imgElm && imgElm.nodeName === "IMG") {
                            operateImg.value = imgElm.getAttribute("src")
                        }
                        this.onImagePreview()
                    },
                    onSetup: (api) => {
                        const imgElm = editor.selection.getNode();
                        api.setDisabled(!(imgElm && imgElm.nodeName === "IMG"));
                    }
                });
                editor.ui.registry.addButton('uploadFiles', {
                    text: $t('文件'),
                    tooltip: $t('上传文件'),
                    onAction: () => {
                        if (this.handleBeforeUpload()) {
                            this.$refs.fileUpload.handleClick();
                        }
                    }
                });
                editor.ui.registry.addMenuItem('uploadFiles', {
                    text: $t('上传文件'),
                    onAction: () => {
                        if (this.handleBeforeUpload()) {
                            this.$refs.fileUpload.handleClick();
                        }
                    }
                });
                if (isFull) {
                    editor.ui.registry.addButton('screenload', {
                        icon: 'fullscreen',
                        tooltip: $t('退出全屏'),
                        onAction: () => {
                            this.closeFull();
                        }
                    });
                    editor.ui.registry.addMenuItem('screenload', {
                        text: $t('退出全屏'),
                        onAction: () => {
                            this.closeFull();
                        }
                    });
                    editor.on('Init', (e) => {
                        editorT.value = editor;
                        editorT.value.setContent(content.value);
                        if (props.readOnly) {
                            editorT.value.setMode('readonly');
                        } else {
                            editorT.value.setMode('design');
                        }
                    });
                } else {
                    editor.ui.registry.addButton('screenload', {
                        icon: 'fullscreen',
                        tooltip: $t('全屏'),
                        onAction: () => {
                            this.onFull();
                        }
                    });
                    editor.ui.registry.addMenuItem('screenload', {
                        text: $t('全屏'),
                        onAction: () => {
                            this.onFull();
                        }
                    });
                    editor.on('Init', (e) => {
                        spinShow.value = false;
                        editor.value = editor;
                        editor.value.setContent(content.value);
                        if (props.readOnly || windowTouch) {
                            editor.value.setMode('readonly');
                            this.updateTouchContent();
                        } else {
                            editor.value.setMode('design');
                        }
                        this.$emit('editorInit', editor.value);
                    });
                    editor.on('KeyUp', (e) => {
                        if (editor.value !== null) {
                            this.submitNewContent();
                        }
                    });
                    editor.on('KeyDown', (e) => {
                        if (e.metaKey || e.ctrlKey) {
                            if (e.keyCode === 83) {
                                e.preventDefault();
                                this.$emit('editorSave', e);
                            }
                        }
                    });
                    editor.on('Change', (e) => {
                        if (editor.value !== null) {
                            if (this.getContent() !== props.value) {
                                this.submitNewContent();
                            }
                            this.$emit('editorChange', e);
                        }
                    });
                    editor.on('focus', () => {
                        this.$emit('on-focus');
                    });
                    editor.on('blur', () => {
                        this.$emit('on-blur');
                    });
                }
            },
        };
        if (props.autoSize) {
            optionInfo.plugins.push('autoresize')
        }
        if (props.minHeight > 0) {
            optionInfo.min_height = props.minHeight
        }
        return optionInfo;
    },

    onFull() {
        content.value = this.getContent();
        transfer.value = true;
        this.initTransfer();
    },

    closeFull() {
        content.value = this.getContent();
        this.$emit('input', content.value);
        this.$emit('on-blur');
        transfer.value = false;
        if (editorT.value != null) {
            editorT.value.destroy();
            editorT.value = null;
        }
    },

    transferChange(visible) {
        if (!visible && editorT.value != null) {
            content.value = editorT.value.getContent();
            this.$emit('input', content.value);
            editorT.value.destroy();
            editorT.value = null;
            //
            if (windowTouch) {
                this.$nextTick(() => {
                    this.updateTouchContent();
                    this.$emit('on-blur');
                });
            }
        }
    },

    getEditor() {
        return transfer.value ? editorT.value : editor.value;
    },

    concatAssciativeArrays(array1, array2) {
        if (array2.length === 0) return array1;
        if (array1.length === 0) return array2;
        let dest = [];
        for (let key in array1) {
            if (array1.hasOwnProperty(key)) {
                dest[key] = array1[key];
            }
        }
        for (let key in array2) {
            if (array2.hasOwnProperty(key)) {
                dest[key] = array2[key];
            }
        }
        return dest;
    },

    submitNewContent() {
        isTyping.value = true;
        if (checkerTimeout.value !== null) {
            clearTimeout(checkerTimeout.value);
        }
        checkerTimeout.value = setTimeout(() => {
            isTyping.value = false;
        }, 300);
        this.$emit('input', this.getContent());
    },

    insertContent(content) {
        if (this.getEditor() !== null) {
            this.getEditor().insertContent(content);
        } else {
            content.value += content;
        }
    },

    getContent() {
        if (this.getEditor() === null) {
            return "";
        }
        return this.getEditor().getContent();
    },

    setContent(content) {
        if (this.getEditor() === null) {
            content.value = content;
        } else if (content != this.getEditor().getContent()) {
            this.getEditor().setContent(content);
        }
    },

    focus() {
        if (this.getEditor() === null) {
            return "";
        }
        return this.getEditor().focus();
    },

    insertImage(src) {
        this.insertContent('<img src="' + src + '">');
    },

    editorImage(lists) {
        for (let i = 0; i < lists.length; i++) {
            let item = lists[i];
            if (typeof item === 'object' && typeof item.url === "string") {
                this.insertImage(item.url);
            }
        }
    },

    getValueImages() {
        const imgs = [];
        const imgReg = /<img.*?(?:>|\/>)/gi,
            srcReg = new RegExp("src=([\"'])([^'\"]*)\\1"),
            widthReg = new RegExp("original-width=\"(\\d+)\""),
            heightReg = new RegExp("original-height=\"(\\d+)\"")
        const array = (this.getContent() + "").match(imgReg);
        if (array) {
            for (let i = 0; i < array.length; i++) {
                const src = array[i].match(srcReg);
                const width = array[i].match(widthReg);
                const height = array[i].match(heightReg);
                if (src) {
                    imgs.push({
                        src: src[2],
                        width: width ? width[1] : -1,
                        height: height ? height[1] : -1,
                    });
                }
            }
        }
        return imgs;
    },

    onLinkPreview() {
        if (operateLink.value) {
            window.open(operateLink.value);
        }
    },

    onImagePreview() {
        const array = this.getValueImages();
        if (array.length === 0) {
            $A.messageWarning("没有可预览的图片")
            return;
        }
        let index = Math.max(0, array.findIndex(item => item.src === operateImg.value));
        this.$store.dispatch("previewImage", { index, list: array })
    },

    onClickWrap(event) {
        if (!windowTouch) {
            return
        }
        event.stopPropagation()
        operateVisible.value = false;
        operateLink.value = event.target.tagName === "A" ? event.target.href : null;
        operateImg.value = event.target.tagName === "IMG" ? event.target.src : null;
        this.$nextTick(() => {
            const rect = this.$el.getBoundingClientRect();
            operateStyles.value = {
                left: `${event.clientX - rect.left}px`,
                top: `${event.clientY - rect.top}px`,
            }
            operateVisible.value = true;
        })
    },

    onTouchstart() {
        if (!windowTouch) {
            return
        }
        operateVisible.value = false;
    },

    updateTouchContent() {
        if (!windowTouch) {
            return
        }
        this.$nextTick(_ => {
            if (!editor.value) {
                return;
            }
            if (!props.placeholder || content.value) {
                editor.value.bodyElement.removeAttribute("data-mce-placeholder");
                editor.value.bodyElement.removeAttribute("aria-placeholder");
            } else {
                editor.value.bodyElement.setAttribute("data-mce-placeholder", props.placeholder);
                editor.value.bodyElement.setAttribute("aria-placeholder", props.placeholder);
            }
            this.updateTouchLink(0);
        })
    },

    updateTouchLink(timeout) {
        if (!windowTouch) {
            return
        }
        setTimeout(_ => {
            if (!editor.value) {
                return;
            }
            editor.value.bodyElement.querySelectorAll("a").forEach(item => {
                if (item.__dataMceClick !== true) {
                    item.__dataMceClick = true;
                    item.addEventListener("click", event => {
                        event.preventDefault();
                        event.stopPropagation();
                        this.onClickWrap(event);
                    })
                }
            })
            if (timeout < 300) {
                this.updateTouchLink(timeout + 100);
            }
        }, timeout)
    },

    /********************文件上传部分************************/

    handleProgress(event, file) {
        //开始上传
        if (file._uploadIng === undefined) {
            file._uploadIng = true;
            uploadIng.value++;
        }
    },

    handleSuccess(res, file) {
        //上传完成
        uploadIng.value--;
        if (res.ret === 1) {
            this.insertContent(`<a href="${res.data.url}" target="_blank">${res.data.name} (${$A.bytesToSize(res.data.size * 1024)})</a>`);
        } else {
            $A.noticeWarning({
                title: $t('上传失败'),
                desc: $t('文件 ' + file.name + ' 上传失败，' + res.msg)
            });
        }
    },

    handleError() {
        //上传错误
        uploadIng.value--;
    },

    handleFormatError(file) {
        //上传类型错误
        $A.noticeWarning({
            title: $t('文件格式不正确'),
            desc: $t('文件 ' + file.name + ' 格式不正确，仅支持上传：' + uploadFormat.value.join(','))
        });
    },

    handleMaxSize(file) {
        //上传大小错误
        $A.noticeWarning({
            title: $t('超出文件大小限制'),
            desc: $t('文件 ' + file.name + ' 太大，不能超过：' + $A.bytesToSize(maxSize.value * 1024))
        });
    },

    handleBeforeUpload() {
        //上传前判断
        return true;
    },


</script>
