<template >
    <n-modal v-model:show="show" transform-origin="center" @after-leave="closeModal" :mask-closable="false">
        <n-card class="w-[480px] max-w-[90%]" :title="$t('评分')" :bordered="false" size="huge" role="dialog"
            aria-modal="true">
            <template #header>
                <h3 class="text-16 md:text-18 text-title-color ">{{ $t('评分') }}</h3>
            </template>
            <template #header-extra>
                <n-icon class="cursor-pointer text-[#A7ACB6]" size="24" :component="Close" @click="handleClose" />
            </template>
            <div>
                <div class="flex items-center">
                    <p v-if="Mark > -1 && !changeShow" class="flex-1 mb-12 text-title-color text-14">{{ $t('负责人自评分数：')
                    }}<span class=" text-primary-color">{{ Mark }}</span></p>
                    <p v-if="superiorScore > -1 && !changeShow" class="flex-1 mb-12 text-title-color text-14">{{
                        $t('上级评分分数：') }}<span class=" text-primary-color">{{ superiorScore }}</span></p>
                </div>

                <n-form v-if="props.inputShow" ref="formRef" :rules="rules" :model="formValue" size="medium"
                    label-placement="left" label-width="auto">
                    <n-form-item path="score">
                        <n-select v-model:value="formValue.score" :placeholder="$t('请选择评分')" :options="markOptions" />
                    </n-form-item>
                </n-form>

                <n-form v-if="changeShow" ref="formRef" :rules="rules" :model="formValue" size="medium"
                    label-placement="left" label-width="auto">
                    <n-form-item path="score">
                        <n-select v-model:value="formValue.score" :placeholder="$t('请选择评分')" :options="markOptions" />
                    </n-form-item>
                </n-form>

                <n-form v-if="selfShow" ref="formRef" :rules="rules" :model="formValue" size="medium"
                    label-placement="left" label-width="auto">
                    <n-form-item path="score">
                        <n-select v-model:value="formValue.score" :placeholder="$t('请选择评分')" :options="markOptions" />
                    </n-form-item>
                </n-form>
            </div>

            <template #footer>
                <div class="button-box flex justify-end mt-0">
                    <div class="flex flex-1 md:flex-initial items-center gap-4">
                        <n-button class="flex-1 md:flex-initial md:block hidden" :loading="loadIng" strong secondary @click="handleClose">
                            {{ props.inputShow || changeShow || selfShow? $t('取消') : $t('关闭') }}
                        </n-button>
                        <n-button class="flex-1 md:flex-initial" :loading="loadIng"
                            v-if="props.canOwnerUpdateScore && !selfShow && !props.inputShow && !changeShow && props.superiorUser == userInfo.userid && props.userid == userInfo.userid"
                            strong secondary @click="handleChange">
                            {{ $t('修改') }}
                        </n-button>
                        <n-button class="flex-1 md:flex-initial" v-if="props.inputShow || changeShow || selfShow" :loading="loadIng" type="primary"
                            @click="handleSubmit">
                            {{ $t('确定') }}
                        </n-button>

                        <n-button class="flex-1 md:flex-initial" v-if="!selfShow && superior_score < 0 && !props.canSuperiorUpdateScore && !props.inputShow && !changeShow && props.superiorUser == userInfo.userid && props.userid == userInfo.userid" :loading="loadIng" type="primary"
                            @click="handleSelf">
                            {{ $t('评分') }}
                        </n-button>
                        <n-button class="flex-1 md:flex-initial"
                            v-if="(props.canOwnerUpdateScore || props.canSuperiorUpdateScore) && !props.inputShow && !changeShow && props.superiorUser != userInfo.userid && props.userid == userInfo.userid"
                            :loading="loadIng" type="primary" @click="handleChange">
                            {{ $t('修改') }}
                        </n-button>
                        <n-button class="flex-1 md:flex-initial"
                            v-if=" props.canSuperiorUpdateScore && !props.inputShow && !changeShow && props.superiorUser == userInfo.userid && props.userid != userInfo.userid"
                            :loading="loadIng" type="primary" @click="handleChange">
                            {{ $t('修改') }}
                        </n-button>
                        <n-button class="flex-1 md:flex-initial"
                            v-if=" props.canSuperiorUpdateScore && !props.inputShow && !changeShow && props.superiorUser == userInfo.userid && props.userid == userInfo.userid"
                            :loading="loadIng" type="primary" @click="handleChange">
                            {{ $t('修改') }}
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
import { useMessage } from "@/utils/messageAll"
import { ResultDialog } from "@/api"
import { UserStore } from '@/store/user'

const userInfo = UserStore().info
const message = useMessage()
const show = ref(false)
const loadIng = ref(false)

const formRef = ref(null)
const changeShow = ref(false)
const selfShow = ref(false)

const formValue = ref({
    score: null,
})

const Mark = ref(-1);
const superiorScore = ref(-1);

const rules = <any>{
    score: {
        type: 'number',
        required: true,
        message: $t('请选择评分'),
        trigger: ['blur', 'change'],
    },
}

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
        label: $t('3 目标达成效果不佳'),
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
    userid: {
        type: Number,
        default: 0,
    },
    superiorUser: {
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
    inputShow: {
        type: Boolean,
        default: false,
    },
    canOwnerUpdateScore: {
        type: Boolean,
        default: false,
    },
    canSuperiorUpdateScore: {
        type: Boolean,
        default: false,
    }
})

watch(() => props.id, (newValue) => {
    if (newValue) {
        Mark.value = props.score
        superiorScore.value = props.superior_score
    }
}, { immediate: true })


const emit = defineEmits(['close'])

//评分
const handleSubmit = () => {
    formRef.value?.validate((errors) => {
        if (errors) return false;
        let type = null;
        if (props.userid == userInfo.userid) { type = 1 }
        if (props.superiorUser == userInfo.userid) { type = 2 }
        if (props.userid == userInfo.userid && props.superiorUser == userInfo.userid && props.canOwnerUpdateScore && !props.canSuperiorUpdateScore) { type = 1 }
        if (selfShow.value ){  type = 2 }
        const upData = {
            id: props.id,
            score: formValue.value.score,
            type: type,
        }
        loadIng.value = true
        okrScore(upData)
            .then(({ msg }) => {
                message.success($t('操作成功'))
            })
            .catch(ResultDialog)
            .finally(() => {
                emit('close', 1)
                changeShow.value = false
                loadIng.value = false
            })
    })
}


const handleChange = () => {
    changeShow.value = true
    if (Mark.value > -1 && superiorScore.value == -1) {
        formValue.value.score = Mark.value
    }
    if (Mark.value > -1 && superiorScore.value > -1) {
        formValue.value.score = superiorScore.value
    }
}
const handleSelf = () => {
    selfShow.value = true
    formValue.value.score = null
}

const handleClose = () => {
    changeShow.value = false
    selfShow.value = false
    emit('close', 2)
}
const closeModal = () => {
    formValue.value.score = null
    changeShow.value = false
    selfShow.value = false
}
</script>
<style lang="less" scoped>
:deep(.n-card__content) {
    @apply pb-0 !important;
}
</style>
