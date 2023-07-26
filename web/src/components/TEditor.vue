<template>
    <div class="tinymce-box">
        <div ref="myTextarea" :id="tinymceId">{{ props.value }}</div>
    </div>
</template>
  
<script setup lang="ts">
import tinymce from 'tinymce/tinymce';
import { GlobalStore } from "../store"

const globalStore = GlobalStore()
const { themeName } = globalStore.appSetup()

const tinymceId = ref("apps_tinymce_" + Math.round(Math.random() * 10000))
const editorVue = ref()
const spinShow = ref(false)

const emit = defineEmits(['update:value','editorInit','keyup','change','focus','blur'])
const props = defineProps({
    value: {
        type: String,
        default: '',
    },
    readOnly: {
        type: Boolean,
        default: false,
    },
});

// 加载
nextTick(() => {
    let lang = globalStore.language;
    switch (lang) {
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

    tinymce.init({
        selector: '#' + tinymceId.value,
        base_url: location.origin + '/js/tinymce/',
        language: lang,
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
        // 设置工具栏
        toolbar: [
            " undo redo | styleselect | uploadImages | uploadFiles | bold italic underline forecolor backcolor | alignleft aligncenter alignright | bullist numlist outdent indent | link image emoticons media codesample | preview screenload"
        ],
        // 设置插件
        plugins: 'codesample lists advlist link autolink charmap emoticons fullscreen preview code searchreplace table visualblocks wordcount insertdatetime image',
        toolbar_mode: "sliding",
        resize: true,
        paste_data_images: true,
        inline: false,
        content_css: themeName.value  == 'dark' ? 'dark' : 'default',
        convert_urls: false,
        height:'100%',
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
        setup: (editor) => {
            editor.on('Init', (e) => {
                spinShow.value = false;
                editorVue.value = editor;
                editorVue.value.setContent(props.value);
                if (props.readOnly) {
                    editorVue.value.setMode('readonly');
                } else {
                    editorVue.value.setMode('design');
                }
                emit('editorInit')
            });
            editor.on('KeyUp', (e) => {
                emit('update:value', editorVue.value.getContent())
                emit('keyup', e)
            });
            editor.on('Change', (e) => {
                emit('change', e)
            });
            editor.on('focus', (e) => {
                emit('focus', e)
            });
            editor.on('blur', (e) => {
                emit('blur', e)
            });
        }
    });
})


</script>
  
<style lang="less" scoped>
.tinymce-box {
    height: 100%;
    :deep(.tox-tinymce) {
        box-shadow: none;
        box-sizing: border-box;
        border-color: #dddee1;
        border-radius: 4px;
        overflow: hidden;
        .tox-statusbar {
            span.tox-statusbar__branding {
                a {
                    display: none;
                }
            }
        }
        .tox-tbtn--bespoke {
            .tox-tbtn__select-label {
                width: auto;
            }
        }
    }
}
</style>
  

<style lang="less">
.okr-theme-light .tox, .okr-theme-dark .tox{
    &.tox-silver-sink {
        position: initial !important;
    }
    .tox-dialog-wrap__backdrop--opaque{
        background-color: rgba(255,255,255,.75) !important;
    }
}
</style>