<template>
    <OkrNotDatas v-if="items.length == 0">
        <template v-slot:content>
            <p>{{$t('暂无复盘')}}</p>
        </template>
    </OkrNotDatas>
    <n-scrollbar v-else trigger="hover">
        <div class="replay mt-[1%]">
            <div
                v-for="(item, index) in items"
                :key="index"
                :class="{ 'replay-item mt-20': true, 'replay-item-active': item.isActive }"
                @click="openMultiple"
            >
                <div class="replay-item-head">
                    <div>
                        <span class="replay-item-okr-level scale-[0.8333]" :class="pStatus(item.priority)">{{ item.priority }}</span>
                        <span class="text-[14px] m-[5px] text-[#333333]" ><b>{{ item.replayName }}</b></span>
                        <span class="text-[#515a6e] text-12">{{ $t('的目标复盘') }}</span>
                    </div>
                    <div class="replay-item-head-right cursor-pointer" @click="() => (item.isActive = !item.isActive)">
                        <span class="mr-[10px]">{{ item.isActive == true ? $t('收起') : $t('展开') }}</span>
                        <i class="replay-item-head-icon replay-item-head-right taskfont">&#xe705;</i>
                    </div>
                </div>
                <div class="flex">
                    <div class="replay-item-okr">
                        <div class="replay-item-okr-icon w-[25px] h-[15px]">O</div>
                        <div class="text-[#515A6E] text-14">{{ item.okrName }}</div>
                    </div>
                </div>
                <div class="replay-item-body" v-if="item.isActive">
                    <OkrReplayDetail :okrReplayList="item"></OkrReplayDetail>
                </div>
            </div>
        </div>
    </n-scrollbar>
    <AddMultiple v-model:show="addMultipleShow" @close="() => { addMultipleShow = false }"></AddMultiple>
</template>

<script lang="ts" setup>
import { ref } from "vue"
import OkrReplayDetail from "@/views/components/OkrReplayDetails.vue"
import OkrNotDatas from "@/views/components/OkrNotDatas.vue"
import * as http from "../../api/modules/replay"

const addMultipleShow = ref(false)
const items = ref([])

//封装复盘列表
const loadResplayList = (data)=>{
    items.value = data.map((itemData) => ({
        id:itemData.okr_id,
        isActive: false,
        replayName: itemData.okr_alias.join(","),
        okrName: itemData.okr_title,
        priority: itemData.okr_priority,
        krList: (itemData.kr_history?itemData.kr_history:[]),
        okrProgress:itemData.okr_progress,
        review: itemData.review
  }));    
}

// 获取数据
const getData = () => {
    // 获取复盘列表
    http.getReplayList().then(({ data }) => {
        loadResplayList(data.data)
    })
}

// 状态
const pStatus = (p) => {
    return p == 'P0' ? 'span-1' : p == 'P1' ? 'span-2' : 'span-3'
}


// 打开复盘
const openMultiple = () => {
    addMultipleShow.value = true;
}


nextTick(()=>{
    getData()
})
</script>
<style lang="less" scope>
.replay {
    margin-top: 1%;
    width: 100%;
    height: 100%;
    &-item {
        @apply bg-white rounded-lg h-auto p-24;
        &-head {
            @apply flex items-center justify-between text-[#87d068] pl-[10px];
            &-icon {
                display: inline-block;
                transform: rotate(0deg);
                transition: transform 0.5s;
            }
        }
        &-okr {
            background: #f4f5f7;
            padding: 12px 16px;
            @apply mt-10 h-auto flex rounded-lg items-center;
            &-icon {
                @apply text-center text-10 leading-4 rounded-lg bg-[#87D068] text-[#87d068] mr-10;
                background: rgba(135, 208, 104, 0.2);
            }
            &-level {
                @apply bg-[#FF7070] text-10 text-white px-6 py-2 rounded-full scale-[0.8333] origin-center;
            }
        }
        &-body {
            @apply mt-14;
        }
    }
    &-item-active &-item-head {
        &-icon {
            transform: rotate(-180deg);
        }
    }
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
</style>
