<template >
    <div class="relative mr-8 shrink-0" :class="!props.userUrl || props.disable != '' ? ' opacity-60' : ''">
        <n-avatar v-if="props.userUrl" round :size="28" class="" :src="props.userUrl" />
        <div class="w-[28px] h-[28px] rounded-full flex justify-center items-center bg-primary-color text-[#fff] text-[14px]" v-if="!props.userUrl">D</div>
    </div>
</template>
<script setup lang="ts">
const props = defineProps({
    userUrl: {
        type: String,
        default: '',
    },
    disable: {
        type: String,
        default: '',
    }
})
</script>
