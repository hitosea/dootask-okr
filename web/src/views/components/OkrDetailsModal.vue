<template >
    <n-modal
        v-model:show="props.show"
        transform-origin="center"
        :mask-closable="false"
        @after-leave="closeDrawer"
        @after-enter="showDrawer"
        :z-index="nextZIndex()"
        :trap-focus="false">
        <n-card
            class="w-[90%] max-w-[1200px] max-md:fixed max-md:inset-0 max-md:w-full max-md:h-full max-md:rounded-none max-md:overflow-hidden"
            :bordered="false"
            size="huge"
            role="dialog"
            aria-modal="true">
            <OkrDetailsMain
                ref="OkrDetailsMainRef"
                :show="props.show"
                :id="props.id"
                @close="()=>{ emit('close') }"
                @edit="(e)=>{ emit('edit',e) }"
                @getList="(e)=>{ emit('getList', e) }"
                @upData="(e)=>{ emit('upData', e) }"
                @openDetail="(id, userid)=>{ emit('openDetail', id, userid)}"/>
        </n-card>
    </n-modal>
</template>
<script setup lang="ts">
import OkrDetailsMain from '@/views/components/OkrDetailsMain.vue';
import { nextZIndex } from "@dootask/tools"

const emit = defineEmits(['close', 'edit', 'upData','getList','openDetail'])

const OkrDetailsMainRef = ref(null)

const props = defineProps({
    show: {
        type: Boolean,
        default: false,
    },
    id: {
        type: Number,
        default: 0,
    },
})

// 关闭
const closeDrawer = () => {
    OkrDetailsMainRef.value.closeDrawer()
}

// 打开
const showDrawer = () => {
}


const getDetail = (type) => {
    OkrDetailsMainRef.value.getDetail(type)
}

defineExpose({
    getDetail
})
</script>
<style lang="less" scoped>
.icon-title {
    @apply text-16 cursor-pointer;
}

.icon-item {
    @apply text-18 mr-8;
}

.span {
    @apply text-14 text-white px-6 py-4 rounded-full flex items-center leading-3 shrink-0;
}

.span-1 {
    @apply bg-[#FF7070];
}

.span-2 {
    @apply bg-[#FC984B];
}

.span-3 {
    @apply bg-[#72A1F7];
}

.li-nav {
    @apply list-none cursor-pointer text-text-li opacity-50 text-14;
}

.li-nav.active {
    @apply text-title-color text-16 opacity-100;
}
</style>
