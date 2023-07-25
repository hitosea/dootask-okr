<template>
    <div class="tinymce-box">
        <textarea ref="myTextarea" id="tinymceTextarea">{{ content }}</textarea>
    </div>
</template>
  
<script setup lang="ts">
import tinymce from 'tinymce/tinymce';
import { GlobalStore } from "../store"

const globalStore = GlobalStore()

// const id = ref("apps_tinymce_" + Math.round(Math.random() * 10000)
const content = ref("Hello, World!")

nextTick(()=>{

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
        selector: '#tinymceTextarea',
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
        resize:true,
        paste_data_images: true,
        inline: false,
        content_css: "default",
        convert_urls: false,
        // height:'100%',
        codesample_languages: [
            {text:"HTML/VUE/XML",value:"markup"},
            {text:"JavaScript",value:"javascript"},
            {text:"CSS",value:"css"},
            {text:"PHP",value:"php"},
            {text:"Ruby",value:"ruby"},
            {text:"Python",value:"python"},
            {text:"Java",value:"java"},
            {text:"C",value:"c"},
            {text:"C#",value:"csharp"},
            {text:"C++",value:"cpp"}
        ],
        // setup: (editor) => {
            // editor.on('Init', (e) => {
            //     this.spinShow = false;
            //     this.editor = editor;
            //     this.editor.setContent(this.content);
            //     if (this.readOnly || this.windowTouch) {
            //         this.editor.setMode('readonly');
            //         this.updateTouchContent();
            //     } else {
            //         this.editor.setMode('design');
            //     }
            //     this.$emit('editorInit', this.editor);
            // });
            // editor.on('KeyUp', (e) => {
            //     if (this.editor !== null) {
            //         this.submitNewContent();
            //     }
            // });
            // editor.on('KeyDown', (e) => {
            //     if (e.metaKey || e.ctrlKey) {
            //         if (e.keyCode === 83) {
            //             e.preventDefault();
            //             this.$emit('editorSave', e);
            //         }
            //     }
            // });
            // editor.on('Change', (e) => {
            //     if (this.editor !== null) {
            //         if (this.getContent() !== this.value) {
            //             this.submitNewContent();
            //         }
            //         this.$emit('editorChange', e);
            //     }
            // });
            // editor.on('focus', () => {
            //     this.$emit('on-focus');
            // });
            // editor.on('blur', () => {
            //     this.$emit('on-blur');
            // });
        // }
    });
})


</script>
  
<style lang="less" scoped>
.tinymce-box {
    height: 100%;
    
    :deep(.tox-tinymce){

        &.tox-silver-sink {
            z-index: 13000;
        }

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
  