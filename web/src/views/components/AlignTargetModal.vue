<template>
    <n-modal v-model:show="props.value" transform-origin="center" :mask-closable="false" :on-after-leave="closeDrawer" @after-enter="showDrawer">
        <n-card class="w-[640px]" :bordered="false" size="huge" role="dialog" aria-modal="true">
            <template #header>
                <div class="headr-nav gap-[32px]">
                    <h3 :class="active == 0 ? 'active' : ''">
                        <span @click="handleActive(0)">{{ $t('对齐目标') }}</span>
                        <i v-if="props.editItem.canceled == '0' && props.editItem.completed == '0' && userInfo.userid == props.editItem.userid" class="okrfont cursor-pointer text-[#A7ABB5] text-[18px]" @click="handleOpenSelect">&#xe779;</i>
                    </h3>
                    <h3 :class="active == 1 ? 'active' : ''">
                        <span @click="handleActive(1)">{{ $t('被对齐目标') }}</span>
                    </h3>
                </div>
            </template>
            <template #header-extra>
                <i class="okrfont text-16 cursor-pointer text-[#999]" @click="handleClose">&#xe6e5;</i>
            </template>
            <AlignTarget :value="props.value" :active="active" :id="props.editItem.id" :userid="props.editItem.userid" :cancelShow="props.editItem.canceled == '0' && props.editItem.completed == '0' && userInfo.userid == props.editItem.userid" @unalign="handleUnalign" @openDetail="openDetail"></AlignTarget>
        </n-card>
    </n-modal>
</template>

<script setup lang="ts">
import AlignTarget from "@/views/components/AlignTarget.vue";
import { UserStore } from '@/store/user'

const userInfo = UserStore().info
const active = ref(0)

const props = defineProps({
    value: {
        type: Boolean,
        default: false,
    },
    editItem: {
        type: undefined,
        default: {},
    },
})


const emit = defineEmits(['close', 'upData', 'openSelectAlignment', 'openDetail'])


const handleOpenSelect = () => {
    emit('openSelectAlignment', props.editItem)
    emit('close')
}

const handleActive = (e) => {
    active.value = e
}
const handleClose = () => {
    emit('close')
    active.value = 0
}

const closeDrawer = () => {
}

const showDrawer = () => {
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
:deep(.n-card__content) {
    @apply pb-32;
}

.headr-nav {
    @apply flex items-center;

    h3 {
        @apply text-16 text-text-li cursor-pointer;
    }

    h3.active {
        @apply text-18 text-title-color;
    }
}
</style>
