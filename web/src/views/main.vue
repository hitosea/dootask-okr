<template>
    <div>
        <h1 class="text-center mt-[12%]">{{ $t("欢迎使用 ORK 系统") }}</h1>
        <div class="text-center mt-[20px]">
            <router-link to="/apps/list" class="mr-[10px]">{{ $t("OKR 列表") }}</router-link> |
            <router-link to="/apps/analysis" class="ml-[10px]">{{ $t("OKR 汇总") }}</router-link>
        </div>
    </div>

    <OkrDetailsModal
        :id="globalStore.okrDetailId"
        :show="detailsShow"
        @edit="handleEdit"
        @close="handleCloseDetails"
        @openDetail="handleOpenDetail"/>
    <AddOkrsDrawer
        v-model:show="addShow"
        :edit="isEdit"
        :editData="editData"
        @close="handleCloseAdd" />
</template>

<script lang="ts" setup>
import { watch } from "vue"
import OkrDetailsModal from "@/views/components/OkrDetailsModal.vue"
import AddOkrsDrawer from "@/views/components/AddOkrsDrawer.vue"
import { GlobalStore } from "@/store"
import { useRouter } from 'vue-router'
import { handleCloseApp } from "@/utils/app"

const router = useRouter()
const detailsShow = ref(false)
const globalStore = GlobalStore()

const addShow = ref(false)
const isEdit = ref(false)
const editData = ref({})

/**
 * 监听打开
 */
watch(
    () => globalStore.okrDetailId,
    (id) => {
        if (!id) {
            detailsShow.value = false
            handleCloseApp()
            return
        }
        if (window.innerWidth < 768) {
            router.push({
                name: 'okrDetails',
                query: { id },
            })
        } else {
            detailsShow.value = true
        }
    },
)

/**
 * 点击编辑
 * @param data
 */
const handleEdit = (data: any) => {
    addShow.value = true
    isEdit.value = true
    editData.value = data
}

/**
 * 点击对齐目标名字再次打开
 * @param id
 */
const handleOpenDetail = (id: any) => {
    globalStore.okrDetailId = id
}

/**
 * 关闭详情
 */
const handleCloseDetails = () => {
    globalStore.okrDetailId = 0
}

/**
 * 关闭添加
 */
const handleCloseAdd = () => {
    isEdit.value = false
    addShow.value = false
    editData.value = {}
}
</script>
