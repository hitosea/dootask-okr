<template >
    <n-modal v-model:show="props.value" transform-origin="center">
        <n-card class="w-[640px]" :title="$t('选择对齐目标')" :bordered="false" size="huge" role="dialog" aria-modal="true">
            <template #header-extra>
                <n-icon class="cursor-pointer" size="24" :component="Close" @click="handleClose" />
            </template>
            <div>
                <n-input placeholder="搜索">
                    <template #prefix>
                        <i class="taskfont"> &#xe6f8;</i>
                    </template>
                </n-input>

                <n-checkbox-group v-model:value="cities" class="mt-[24px]">
                    <div class="flex flex-col gap-[24px]">
                        <div class="object-field">
                            <n-checkbox value="Beijing" />
                            <i class="taskfont">&#xe745;</i>
                            <span class="span span-1 scale-[0.8333]">P0</span>
                            <h3 class="text-14 font-normal ml-4 truncate">
                                完善团队制度，为了更好的服务客户完善团队制度，为了更好的服务客户完善团队制度，为了更好的服务客户完善团队制度，为了更好的服务客户完善团队制度，为了更好的服务客户</h3>
                        </div>

                    </div>
                </n-checkbox-group>
            </div>
            <template #footer>
                <div class="button-box">
                    <n-button :loading="loadIng" type="default" @click="handleClose">
                        {{ $t('取消') }}
                    </n-button>
                    <n-button :loading="loadIng" type="primary" @click="">
                        {{ $t('确定') }}
                    </n-button>
                </div>
            </template>
        </n-card>
    </n-modal>
</template>
<script setup lang="ts">
import { Close } from "@vicons/ionicons5"

const props = defineProps({
    value: {
        type: Boolean,
        default: true,
    },
})

const emit = defineEmits(['close'])

const loadIng = ref(false)
const cities = ref(null)

const handleClose = () => {
    emit('close')
}
</script>
<style lang="less" scoped>
.object-field {
    @apply flex items-center;

    i {
        @apply text-12 text-text-tips ml-12 cursor-pointer;
    }

    .span {
        @apply text-10 text-white px-6 py-2 rounded-full origin-center flex items-center leading-3 shrink-0;
    }

    .span-1 {
        @apply bg-[#FF7070];
    }
}
</style>