<template>
    <div class=" w-full h-full flex mt-[100px] md:mt-[10%] items-center flex-col">
        <div v-if="props.types">
            <img class="mr-6 -mt-2 w-[120px]" :src="utils.apiUrl(notSearchSvg)" />
        </div>
        <div v-if="!props.types && !props.loadIng">
            <img class="mr-6 -mt-2 w-[120px]" :src="utils.apiUrl(notDataSvg)" />
        </div>
        <div v-if="props.types" class="text-[#515A6E] text-14 text-opacity-50 text-center">
            <div class="mb-10" v-if="props.msg">{{ $t('没有找到匹配的结果') }}</div>
            <slot name="content"></slot>
        </div>
        <div v-if="!props.types && !props.loadIng" class="text-[#515A6E] text-14 text-opacity-50 text-center">
            <div class="mb-10" v-if="props.msg">{{props.msg}}</div>
            <slot name="content"></slot>
        </div>
    </div>
</template>

<script lang="ts" setup>
import utils from "@/utils/utils";
import notSearchSvg from "@/assets/images/icon/notSearch.svg";
import notDataSvg from "@/assets/images/icon/notData.svg";

const props = defineProps({
    msg: {
        type: String,
        default: ()=> $t('暂无数据'),
    },
    types: {
        type: Boolean,
        default: ()=> false,
    },
    loadIng: {
        type: Boolean,
        default: ()=> false,
    }
});
</script>
