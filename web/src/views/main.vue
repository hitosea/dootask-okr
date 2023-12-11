<template>

    <div v-if="show">
        <h1 class="text-center mt-[12%]">{{ $t('欢迎使用okr子系统') }}</h1>
        <div class="text-center mt-[20px]">
            <router-link to="/apps/list" class="mr-[10px]">{{ $t('前往Okr列表') }}</router-link> |
            <router-link to="/apps/analysis" class="ml-[10px]">{{ $t('前往Okr汇总') }}</router-link>
        </div>
    </div>

    <OkrDetailsModal ref="RefOkrDetails" @edit="handleEdit" :id="globalStore.okrDetail.id" :show="okrDetailsShow" @close="close" @openDetail="handleOpenDetail" />
    <AddOkrsDrawer v-model:show="addShow" :edit="edit" :editData="editData" @close="handleClose" />
</template>

<script lang="ts" setup>
import { watch } from 'vue';
import OkrDetailsModal from '@/views/components/OkrDetailsModal.vue';
import AddOkrsDrawer from '@/views/components/AddOkrsDrawer.vue'
import { GlobalStore } from "@/store"
import { useRouter } from 'vue-router'

const router = useRouter()
const okrDetailsShow = ref(false)
const globalStore = GlobalStore()

const show = ref(false)
const addShow = ref(false)
const edit = ref(false)
let editData = {}

setTimeout(()=>{
    show.value = true
},1000);

//编辑
const handleEdit = (data) => {
    okrDetailsShow.value = false
    addShow.value = true
    edit.value = true
    editData = data
}

// 监听打开
watch(() => globalStore.okrDetail, (newValue) => {
    if (newValue.show) {
        if (window.innerWidth < 910) {
            router.push({
                path: globalStore.baseRoute + '/okrDetails',
                query: { id: globalStore.okrDetail.id },
            })
        }
        else {
            okrDetailsShow.value = false
            nextTick(() => {
                okrDetailsShow.value = true;
            })
        }
    }
})

//点击对齐目标名字再次打开
const handleOpenDetail = (id)=>{
    if (window.innerWidth < 910) {
            router.push({
                path: globalStore.baseRoute + '/okrDetails',
                query: { data: id },
            })
        }
        else {
            globalStore.okrDetail.id = id
            okrDetailsShow.value = false
            nextTick(() => {
                okrDetailsShow.value = true;
            })
        }
}

// 关闭
const close = () => {
    okrDetailsShow.value = false
    globalStore.okrDetail.id = 0
    globalStore.okrDetail.show = false
}

const handleClose = (e,id) => {
    edit.value = false
    editData = {}
    addShow.value = false
}
</script>
