<template >
    <n-modal v-model:show="show" transform-origin="center"  @after-leave="closeModal" >
        <n-card class="w-[480px]" :title="$t('评分')" :bordered="false" size="huge" role="dialog" aria-modal="true">
            <template #header-extra>
                <n-icon class="cursor-pointer text-[#A7ACB6]" size="24" :component="Close" @click="handleClose" />
            </template>
            <div>
                <div class="flex items-center">
                    <p v-if="Mark > -1" class="flex-1 mb-12 text-title-color text-14">{{ $t('负责人自评分数：') }}<span class=" text-primary-color">{{ Mark }}</span></p>
                <p v-if="superiorScore > -1" class="flex-1 mb-12 text-title-color text-14">{{ $t('上级评分分数：') }}<span class=" text-primary-color">{{ superiorScore }}</span></p>
                </div>

                <n-form v-if="Mark <= 0 || superiorScore <= 0" ref="formRef" :model="formValue" size="medium" label-placement="left" label-width="auto">
                    <n-form-item>
                        <n-select v-model:value="formValue.score" :placeholder="$t('请选择评分')" :options="markOptions" />
                    </n-form-item>
                </n-form>
            </div>

            <template #footer>
                <div class="button-box flex justify-end mt-0">
                    <div class="flex items-center gap-4">
                        <n-button :loading="loadIng" type="default" @click="handleClose">
                            {{ $t('取消') }}
                        </n-button>
                        <n-button :loading="loadIng" type="primary" @click="handleSubmit">
                            {{ $t('确定') }}
                        </n-button>
                    </div>

                </div>
            </template>
        </n-card>
    </n-modal>
</template>
<script setup lang="ts">
import { Close } from "@vicons/ionicons5"
import { okrScore } from '@/api/modules/okrList'
import { useMessage } from "naive-ui"
import { ResultDialog } from "@/api"

const message = useMessage()
const show = ref(false)
const loadIng = ref(false)
const formValue = ref({
    score: null,
})

const Mark = ref(-1);
const superiorScore = ref(-1);

const markOptions = ref([
    {
        label: $t('0 未达成目标，态度问题'),
        value: 0,
    },
    {
        label: $t('1 未达成目标'),
        value: 1,
    },
    {
        label: $t('2 目标达成效果较差，未及时调整'),
        value: 2,
    },
    {
        label: $t('3 目标达成效果不佳 '),
        value: 3,
    },
    {
        label: $t('4 目标勉强达成，还需努力'),
        value: 4,
    },
    {
        label: $t('5 目标基本达成'),
        value: 5,
    },
    {
        label: $t('6 目标达成，效果一般'),
        value: 6,
    },
    {
        label: $t('7 目标达成，效果良好'),
        value: 7,
    },
    {
        label: $t('8 100%达成目标'),
        value: 8,
    },
    {
        label: $t('9 110%达成目标'),
        value: 9,
    },
    {
        label: $t('10 超出预期'),
        value: 10,
    },

])



const props = defineProps({
    id: {
        type: Number,
        default: 0,
    },
    score: {
        type: Number,
        default: 0,
    },
    superior_score: {
        type: Number,
        default: 0,
    },
})

watch(() => props.id, (newValue) => {
    if (newValue) {
        Mark.value = props.score
        superiorScore.value = props.superior_score
    }
}, { immediate: true })


const emit = defineEmits(['close'])

const handleSubmit = () => {
    const upData = {
        id: props.id,
        score: formValue.value.score,
    }
    loadIng.value = true
    okrScore(upData)
        .then(({ msg }) => {
            message.success(msg)
        })
        .catch(ResultDialog)
        .finally(() => {
            emit('close', 1)
            loadIng.value = false
        })
}

const handleClose = () => {
    emit('close', 2)
}
const closeModal = () => {
    formValue.value.score = null
}
</script>
<style lang="less" scoped>
:deep(.n-card__content) {
    @apply pb-0 !important;
}
</style>
