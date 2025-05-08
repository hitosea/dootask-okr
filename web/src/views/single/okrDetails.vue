<template>
    <div class="py-16 px-24 h-[100vh] overflow-hidden">
        <OkrDetailsMain ref="OkrDetailsMainRef" :show="true" :single="true" :id="id" @openDetail="openDetail" />
    </div>
</template>
<script lang="ts" setup>
import OkrDetailsMain from "@/views/components/OkrDetailsMain.vue"
import { useRoute, useRouter } from "vue-router"

const route = useRoute()
const router = useRouter()

const id = ref(null)

if (route.query.id != undefined) {
    id.value = Number(route.query.id)
}

// 监听打开
const openDetail = (okrId: any, userId: any) => {
    router.replace({
        path: route.path,
        query: {
            id: okrId,
            userid: userId,
        },
    })
    id.value = okrId
}
</script>
