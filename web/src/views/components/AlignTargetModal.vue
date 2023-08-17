<template >
    <n-modal v-model:show="props.value" transform-origin="center" :mask-closable="false">
        <n-card class="w-[640px]" :bordered="false" size="huge" role="dialog" aria-modal="true">
            <template #header>
                {{ $t('对齐目标') }}
                <i v-if="props.eidtItem.canceled == '0' && props.eidtItem.completed == '0' && userInfo.userid == props.eidtItem.userid"
                    class="taskfont cursor-pointer text-[#A7ABB5] text-[18px]" @click="handleOpenSelect">&#xe779;</i>
            </template>
            <template #header-extra>
                <i class="taskfont text-16 cursor-pointer text-[#A7ABB5]" @click="handleClose">&#xe6e5;</i>
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

<style lang="less" ></style>

