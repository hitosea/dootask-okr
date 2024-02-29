<template >
    <div class="delete-box nav-top  h-[52px] bg-[#FAFAFA] z-[5]" :style="{ 'z-index': modalTransferIndex + 1 }">
        <i @click="handleReturn" class="okrfont icon-return z-[2]">&#xe676;</i>
        <h2 class="absolute left-0 right-0 text-center text-title-color text-17 font-medium">{{ $t('复盘详情')}}</h2>
    </div>
    <div class="relative pt-[76px] pb-16 pl-16 pr-16 h-full" :style="{ 'z-index': modalTransferIndex }">
        <div class="flex flex-col">
            <div class="flex justify-between ">
                <h3 class=" text-text-li text-14 font-normal">{{ $t('目标（O）') }}</h3>
                <p class=" text-primary-color text-14 font-normal">{{ dataDetail.okr_progress }}%</p>
            </div>
            <div class="mt-8 border-solid border-[1px] border-[#F2F2F2] rounded p-16 text-15 text-text-li">
                {{ dataDetail.okr_title }}
            </div>

            <div class="flex justify-between mt-16">
                <h3 class=" text-text-li text-14 font-normal">{{ $t('关键KR') }}</h3>
            </div>
            <div v-for="(item,index) in dataDetail.key_results" class="mt-8 border-solid border-[1px] border-[#F2F2F2] rounded p-16 ">
                <div class="text-15 text-text-li"><span class=" text-text-tips text-12 mr-4">KR{{ index+1 }}</span>{{ item.title }}</div>
                <div class="flex mt-12 text-text-li text-12">
                    <p class="pr-16 flex-1 text-center border-solid border-0 border-r-[1px] border-[#F2F2F2]">{{ $t('KR完成度：') }}<span class=" text-primary-color">{{ item.progress }}%</span></p>
                    <p class="pl-16 flex-1 text-center">{{ $t('KR评分：') }}<span class=" text-primary-color">{{ item.kr_score }}</span></p>
                </div>
            </div>

            <div v-if="dataDetail">
                <OkrReplayDetail :okrReplayList="dataDetail"></OkrReplayDetail>
            </div>

        </div>
    </div>
</template>
<script lang="ts" setup>
import { getReplayList } from '@/api/modules/replay'
import {useRoute, useRouter } from 'vue-router';
import OkrReplayDetail from "@/views/components/OkrReplayDetails.vue"

const route = useRoute()
const router = useRouter()
const id = ref(0)
const dataDetail = ref<any>("")

const modalTransferIndex = window.modalTransferIndex = window.modalTransferIndex + 1

if (route.query.id != undefined) {
    id.value = Number(route.query.id)
}


const handleReturn = () => {
    router.go(-1)
}

// 获取数据
const getData = () => {
    const data = {
            okr_id: id.value,
            page: 1,
            page_size: 20,
        }
        getReplayList(data)
        .then(({ data }) => {
            dataDetail.value = data.data[0]
        })
}
onMounted(()=>{
    getData()
})

</script>
<style lang="less" scoped>
.nav-top {
    @apply flex items-center justify-between fixed top-0 left-0 right-0 px-10;

    .icon-return {
        @apply text-26 text-text-tips;
    }
}
</style>
