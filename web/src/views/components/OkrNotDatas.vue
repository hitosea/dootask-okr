<template>
    <div class=" w-full h-full flex mt-[10%] items-center flex-col">
        <div v-if="props.types!=''">
            <img class="mr-6 -mt-2 w-80" src="@/assets/images/icon/notSearch.svg" />
        </div>
        <div v-else>
            <img class="mr-6 -mt-2 w-80" src="@/assets/images/icon/notData.svg" />
        </div>
        <div v-if="props.types!=''" class="text-[#515A6E] text-14 text-opacity-50 text-center"> 
            <div class="mb-10" v-if="props.msg">{{ $t('没有找到匹配的结果') }}</div>           
            <slot name="content"></slot>
        </div>
        <div v-else class="text-[#515A6E] text-14 text-opacity-50 text-center"> 
            <div class="mb-10" v-if="props.msg">{{props.msg}}</div>           
            <slot name="content"></slot>
        </div>
    </div>
</template>

<script lang="ts" setup>
    const props = defineProps({
        msg: {
            type: String,
            default: ()=> $t('暂无OKR'),
        },
        types: {
            type: String,
            default: ()=> $t(''),
        }
    });
</script>