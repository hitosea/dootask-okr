<template >
    <n-drawer v-model:show="show" style="max-width: 948px;width: 90%;" :on-after-leave="closeDrawer" @after-enter="showDrawer" :z-index="13" :mask-closable="false"
        :trap-focus="false" :on-update-show="()=>{emit('close') }" class="okr">
        <n-drawer-content :title="$t('复盘')" closable>
            <div class="flex flex-col h-full">
                <AddMultipleMain ref="AddMultipleMainRef" :data="props.data" :multipleId="props.multipleId" :superiorUser="props.superiorUser" @canComment="()=>{canComment = true}" @close="()=>{ emit('close') }" @loadIng="(e)=>{ loadIng = e }"></AddMultipleMain>
                <div class="button-box">
                    <n-button  v-if="props.multipleId == 0 " :loading="loadIng" type="primary" @click="handleSubmit">
                        {{ $t('提交') }}
                    </n-button>
                    <n-button  v-if="canComment" :loading="loadIng" type="primary" @click="handleComment" >
                        {{ $t('提交') }}
                    </n-button>
                </div>
            </div>
        </n-drawer-content>
    </n-drawer>
</template>
<script setup lang="ts">
import AddMultipleMain from './AddMultipleMain.vue';

const show = ref(false)
const loadIng = ref(false)
const canComment = ref(false)
const AddMultipleMainRef = ref(null)

const props = defineProps({
    data: {
        type: Object,
        default: () => { },
    },
    multipleId: {
        type: Number,
        default: 0,
    },
    superiorUser: {
        type: undefined,
        default: [],
    },
})


const emit = defineEmits(['close'])

const closeDrawer = () => {
    AddMultipleMainRef.value.closeDrawer()
}

const showDrawer = () => {
}

const handleSubmit = () => {
    AddMultipleMainRef.value.handleSubmit()
}

const handleComment = () => {
    AddMultipleMainRef.value.handleComment()
}

</script>
<style lang="less" scoped>
.button-box {
    @apply flex gap-2 mt-24 flex-initial;
}

:deep(.n-scrollbar-content) {
    @apply h-full;
}
:deep(.n-drawer-header__close) {
    @apply absolute -left-36 ;
    &:focus{
        @apply bg-none;
    }
    i {
        @apply text-[#fff];
    }
}

:deep(.n-base-close:not(.n-base-close--disabled):focus::before) {
        @apply bg-transparent;
}
</style>
