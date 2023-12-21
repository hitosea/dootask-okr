<template >
    <div class="relative mr-8 shrink-0" :class="props.disable != '' ? ' opacity-80' : ''">
        <n-avatar v-if="props.userUrl" round :size="28" class="" :src="props.userUrl" />
        <div v-if="props.userUrl" :class="props.disable != '' ? 'disabled-avatar' : ''"></div>
        <div class="w-[28px] h-[28px] rounded-full flex justify-center items-center bg-primary-color text-[#fff] text-[14px]" v-if="!props.userUrl">
            D
        </div>
        <div v-if="!props.userUrl" class="del-avatar"></div>
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
<style lang="less" scoped>
.disabled-avatar {
    position: absolute;
    left: 1px;
    top: 1px;
    right: 1px;
    bottom: 1px;
    background-image: url(data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIGNsYXNzPSJpb25pY29uIiB2aWV3Qm94PSIwIDAgNTEyIDUxMiI+PHRpdGxlPkJhbjwvdGl0bGU+PGNpcmNsZSBmaWxsPSJub25lIiBzdHJva2U9IiNjY2NjY2MiIHN0cm9rZS1taXRlcmxpbWl0PSIxMCIgc3Ryb2tlLXdpZHRoPSI0OCIgY3g9IjI1NiIgY3k9IjI1NiIgcj0iMjAwIi8+PHBhdGggc3Ryb2tlPSIjY2NjY2NjIiBzdHJva2UtbWl0ZXJsaW1pdD0iMTAiIHN0cm9rZS13aWR0aD0iNDgiIGQ9Ik0xMTQuNTggMTE0LjU4bDI4Mi44NCAyODIuODQiLz48L3N2Zz4=);
    background-repeat: no-repeat;
}
.del-avatar {
    position: absolute;
    width: 26px;
    height: 26px;
    left: 1px;
    top: 1px;
    right: 1px;
    bottom: 1px;
    transform: rotate(-45deg);
    background-image: url(data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIGNsYXNzPSJpb25pY29uIiB2aWV3Qm94PSIwIDAgNTEyIDUxMiI+PHRpdGxlPkJhbjwvdGl0bGU+PGNpcmNsZSBmaWxsPSJub25lIiBzdHJva2U9IiNjY2NjY2MiIHN0cm9rZS1taXRlcmxpbWl0PSIxMCIgc3Ryb2tlLXdpZHRoPSI0OCIgY3g9IjI1NiIgY3k9IjI1NiIgcj0iMjAwIi8+PHBhdGggc3Ryb2tlPSIjY2NjY2NjIiBzdHJva2UtbWl0ZXJsaW1pdD0iMTAiIHN0cm9rZS13aWR0aD0iNDgiIGQ9Ik0xMTQuNTggMTE0LjU4bDI4Mi44NCAyODIuODQiLz48L3N2Zz4=);
    background-repeat: no-repeat;
}
</style>