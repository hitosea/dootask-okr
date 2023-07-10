<template>
    <div @click="onCLick" ref="markdownBody" class="markdown-body" v-html="html"></div>
</template>

<script setup lang="ts">
import { defineProps, ref, onUpdated, computed, defineEmits } from 'vue'
import '@/styles/dialog-markdown/markdown.less'
import MarkdownIt from 'markdown-it'
import mdKatex from 'markdown-it-katex'
import mila from 'markdown-it-link-attributes'
import hljs from 'highlight.js'

const markdownIt = MarkdownIt()
const MdKatex = mdKatex
const Mila = mila
const Hljs = hljs
const mdi = ref(null)
const emit = defineEmits(['click'])
const markdownBody = ref()

const props = defineProps({
    text: {
        type: String,
        default: ''
    },
})

onMounted(() => {
    copyCodeBlock()
})

onUpdated(() => {
    copyCodeBlock()
})

const html = computed(() => {
    if (mdi.value === null) {
        mdi.value = new markdownIt({
            linkify: true,
            highlight(code, language) {
                const validLang = !!(language && Hljs.getLanguage(language))
                if (validLang) {
                    const lang = language ?? ''
                    return highlightBlock(Hljs.highlight(code, { language: lang }).value, lang)
                }
                return highlightBlock(Hljs.highlightAuto(code).value, '')
            },
        })
        mdi.value.use(Mila, { attrs: { target: '_blank', rel: 'noopener' } })
        mdi.value.use(MdKatex, { blockClass: 'katexmath-block rounded-md p-[10px]', errorColor: ' #cc0000' })
    }
    return formatMsg(mdi.value.render(props.text))
})



const highlightBlock = (str, lang = '') => {
    return `<pre class="code-block-wrapper"><div class="code-block-header"><span class="code-block-header__lang">${lang}</span><span class="code-block-header__copy">${$t('复制代码')}</span></div><code class="Hljs code-block-body ${lang}">${str}</code></pre>`
}

const formatMsg = (text) => {
    const array = text.match(/<img\s+[^>]*?>/g);
    if (array) {
        array.some(res => {
            text = text.replace(res, `<div class="no-size-image-box">${res}</div>`);
        })
    }
    return text
}

const copyCodeBlock = () => {
    const codeBlockWrapper = markdownBody.value.querySelectorAll('.code-block-wrapper')
    codeBlockWrapper.forEach((wrapper) => {
        const copyBtn = wrapper.querySelector('.code-block-header__copy')
        const codeBlock = wrapper.querySelector('.code-block-body')
        if (copyBtn && codeBlock && copyBtn.getAttribute("data-copy") !== "click") {
            copyBtn.setAttribute("data-copy", "click")
            copyBtn.addEventListener('click', () => {
                if (navigator.clipboard?.writeText)
                    navigator.clipboard.writeText(codeBlock.textContent ?? '')
                else
                    copyText({ text: codeBlock.textContent ?? '', origin: true })
            })
        }
    })
}

const copyText = (options) => {
    const props = { origin: true, ...options }

    let input

    if (props.origin)
        input = document.createElement('textarea')
    else
        input = document.createElement('input')

    input.setAttribute('readonly', 'readonly')
    input.value = props.text
    document.body.appendChild(input)
    input.select()
    if (document.execCommand('copy'))
        document.execCommand('copy')
    document.body.removeChild(input)
}

const onCLick = (e) => {
    emit('click', e)
}


</script>
<style lang="less"></style>