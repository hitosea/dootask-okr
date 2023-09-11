<template >
    <n-modal v-model:show="props.value" transform-origin="center" :mask-closable="false"  :on-after-leave="closeDrawer" @after-enter="showDrawer">
        <n-card class="w-[640px]" :bordered="false" size="huge" role="dialog" aria-modal="true">
            <template #header>
                {{ $t('对齐目标') }}
                <i v-if="props.eidtItem.canceled == '0' && props.eidtItem.completed == '0' && userInfo.userid == props.eidtItem.userid"
                    class="okrfont cursor-pointer text-[#A7ABB5] text-[18px]" @click="handleOpenSelect">&#xe779;</i>
            </template>
            <template #header-extra>
                <i class="okrfont text-16 cursor-pointer text-[#A7ABB5]" @click="handleClose">&#xe6e5;</i>
            </template>
            <AlignTarget :value="props.value" :id="props.eidtItem.id" :userid="props.eidtItem.userid"
                :cancelShow="props.eidtItem.canceled == '0' && props.eidtItem.completed == '0' && userInfo.userid == props.eidtItem.userid"
                @unalign="handleUnalign" @openDetail="openDetail"></AlignTarget>
        </n-card>
    </n-modal>
</template>

<script setup lang="ts">
import AlignTarget from "@/views/components/AlignTarget.vue";
import { UserStore } from '@/store/user'

const userInfo = UserStore().info

const props = defineProps({
    value: {
        type: Boolean,
        default: false,
    },
    eidtItem: {
        type: undefined,
        default: {},
    },
})


const emit = defineEmits(['close', 'upData', 'openSelectAlignment', 'openDetail'])


const handleOpenSelect = () => {
    emit('openSelectAlignment', props.eidtItem)
    emit('close')
}

const handleClose = () => {
    emit('close')
}

const closeDrawer = () => {
    document.removeEventListener('keydown', handleKeydown);
}

const showDrawer = () => {
    document.addEventListener('keydown', handleKeydown);
}

// ESC
const handleKeydown = (event) => {
    if (event.key === 'Escape') {
        // 执行ESC键按下时的逻辑
        emit('close')
    }
}

// 取消
const handleUnalign = (id) => {
    emit('upData', id)
}
// 打开详情
const openDetail = (id, userid) => {
    emit('close')
    emit('openDetail', id, userid)
}

</script>

<style lang="less" scoped>
:deep(.n-card__content){
    @apply pb-32;
}
</style>

