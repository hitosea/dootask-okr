<template >
    <n-scrollbar>
        <div class="add-main-box pr-14">

            <n-form ref="formRef" :model="formValue" :rules="rules" size="medium" :label-align="props.labelAlign"
                :label-placement="props.labelPlacement" label-width="auto" require-mark-placement="left">

                <n-form-item class="w-full" :label="$t('目标（O）')" path="title">
                    <n-input v-model:value="formValue.title" :maxlength="255" :placeholder="$t('请输入目标')" />
                </n-form-item>

                <n-form-item :label="$t('类型')" path="type">
                    <n-radio-group v-model:value="formValue.type" name="radiogroup1">
                        <n-space>
                            <n-radio :value="1">{{ $t('承诺型') }}</n-radio>
                            <n-radio :value="2">{{ $t('挑战性') }}</n-radio>
                        </n-space>
                    </n-radio-group>
                </n-form-item>

                <n-form-item :label="$t('优先级')" path="priority">
                    <n-radio-group v-model:value="formValue.priority" name="radiogroup2" :disabled="edit">
                        <n-space>
                            <template v-if="formValue.type == 1">
                                <n-radio value="P0">
                                    <span class="span span-1">P0</span>
                                </n-radio>
                                <n-radio value="P1">
                                    <span class="span span-2">P1</span>
                                </n-radio>
                            </template>
                            <template v-else>
                                <n-radio value="P2">
                                    <span class="span span-3">P2</span>
                                </n-radio>
                            </template>
                        </n-space>
                    </n-radio-group>
                </n-form-item>

                <n-form-item v-if="departmentOwner" :label="$t('归属')" path="ascription">
                    <n-radio-group v-model:value="formValue.ascription" name="radiogroup3" :disabled="edit">
                        <n-space>
                            <n-radio :value="1">{{ $t('部门') }}</n-radio>
                            <n-radio :value="2">{{ $t('个人') }}</n-radio>
                        </n-space>
                    </n-radio-group>
                </n-form-item>


                <n-form-item :label="$t('周期')" path="time">
                    <n-date-picker class="w-full" v-model:value="formValue.time" value-format="yyyy.MM.dd HH:mm:ss"
                        type="daterange" clearable size="medium" />
                </n-form-item>

                <n-form-item :label="$t('可见范围')">
                    <n-select v-model:value="formValue.visible_range" :placeholder="$t('请选择可见范围')"
                        :options="generalOptions" />
                </n-form-item>

                <n-form-item :label="$t('对齐目标')">
                    <div @click="handleGoal"
                        class="w-full h-[30px] bg-[#F4F5F7] border-[1px] border-[#F4F5F7] border-solid rounded cursor-pointer pl-12 pr-8 flex items-center">
                        <p :class="formValue.align_objective?.length > 0 ? 'text-text-li' : 'text-[rgba(194,194,194,1)]'"
                            class=" text-14">{{ formValue.align_objective?.length > 0
                                ? `${$t('已选')}${formValue.align_objective.length}${$t('项')}` : $t('请选择对齐目标') }}</p>
                        <i class="taskfont text-[rgba(194,194,194,1)] ml-auto">&#xe72b;</i>
                    </div>


                </n-form-item>

                <n-form-item :label="$t('关联项目')">
                    <ItemList :edit="props.edit" v-model:value="formValue.project_id" />
                </n-form-item>
            </n-form>
            <div>
                <div class="flex items-center justify-between">
                    <h3 class="text-14 text-text-li font-medium flex items-center">
                        {{ $t('关键KR') }}
                        <n-popover trigger="hover" placement="right-start" :width="220">
                            <template #trigger>
                                <n-icon class=" cursor-pointer text-text-tips ml-8" size="18"
                                    :component="AlertCircleOutline" />
                            </template>
                            <div class="">
                                <p class="text-14 text-[#FFBD23] font-medium">{{ $t('KR1：动词+你要追踪的内容+从到Y/或者具体值') }}</p>
                                <p class="text-14 text-[#FFBD23] mt-8 font-medium">{{ $t('KR2：动词+什么时间节点+达成什么关键成果') }}</p>
                                <p class="text-12 text-title-color mt-12">{{ $t('示例') }}</p>
                                <p class="text-12 text-text-li">{{ $t('将客户续约率从70%提高到90%') }}</p>

                                <p class="text-12 text-text-li mt-8 flex items-center"><img class="mr-4" src="@/assets/images/icon/addIcon1.png" />{{ $t('至少包含1个关键KR') }}</p>
                                <p class="text-12 text-text-li mt-4 flex items-center"><img class="mr-4" src="@/assets/images/icon/addIcon1.png" />{{ $t('不建议超过5个关键KR') }}</p>
                                <p class="text-12 text-text-li mt-4 flex items-center"><img class="mr-4" src="@/assets/images/icon/addIcon2.png" />{{ $t('关键成功定量可衡量') }}</p>
                            </div>
                        </n-popover>
                    </h3>
                    <div class="flex items-center cursor-pointer" @click="handleAddKr">
                        <i class="taskfont text-14 text-primary-color">&#xe731;</i>
                        <p class="text-14 text-primary-color ml-4">{{ $t('添加KR') }}</p>
                    </div>
                </div>

                <div v-for="(item, index) in formKRValue" class="border-[1px] border-solid border-[#F2F2F2] mt-16 rounded">
                    <div
                        class="flex items-center justify-between px-[12px] py-[8px] bg-[#FAFAFA] border-0 border-b-[1px] border-solid border-[#F2F2F2]">
                        <h3 class="text-14 text-text-li font-medium">KR{{ index + 1 }}</h3>
                        <div class="flex items-center cursor-pointer" @click="handleRemoveKr(index)">
                            <i class="taskfont text-14 text-text-tips">&#xe787;</i>
                        </div>
                    </div>

                    <div class="p-24 py-20 pb-8">
                        <n-form ref="formRefs" :model="formKRValue[index]" size="medium" :rules="timeRule"
                            :label-align="props.labelAlign" :label-placement="props.labelPlacement" label-width="auto"
                            require-mark-placement="left">
                            <n-grid :cols="4" :x-gap="16">
                                <n-form-item-gi :span="4" class="w-full" label="KR" path="title">
                                    <n-input v-model:value="item.title" :maxlength="255" :placeholder="$t('请输入')" />
                                </n-form-item-gi>

                                <n-form-item-gi :span="4" :label="$t('时间')" path="time">
                                    <n-date-picker class="w-full" v-model:value="item.time" type="daterange" clearable
                                        size="medium" />
                                </n-form-item-gi>

                                <n-form-item-gi class="hidden md:block" :span="2" :label="$t('参与人')">
                                    <div v-if="showUserSelect" class="w-full min-h-[32px] bg-[#F4F5F7] rounded-[4px]">
                                        <UserSelects :formkey="index" />
                                    </div>
                                    <UserList :edit="props.edit" v-if="!showUserSelect" v-model:value="item.participant">
                                    </UserList>
                                </n-form-item-gi>
                                <!-- <n-form-item-gi class="block md:hidden" :span="4" :label="$t('参与人')">
                                    <div v-if="showUserSelect" class="w-full min-h-[32px] bg-[#F4F5F7] rounded-[4px]">
                                        <UserSelects :formkey="index" />
                                    </div>
                                    <UserList :edit="props.edit" v-if="!showUserSelect" v-model:value="item.participant">
                                    </UserList>
                                </n-form-item-gi> -->


                                <n-form-item-gi class="hidden md:block" :span="2" :label="$t('信心')">
                                    <n-input-number class="w-full" :max="100" :min="0" :precision="0"
                                        v-model:value="item.confidence" placeholder="0-100" :show-button="false" />
                                </n-form-item-gi>
                                <n-form-item-gi class="block md:hidden" :span="4" :label="$t('信心')">
                                    <n-input-number class="w-full" :max="100" :min="0" :precision="0"
                                        v-model:value="item.confidence" placeholder="0-100" :show-button="false" />
                                </n-form-item-gi>
                            </n-grid>
                        </n-form>
                    </div>
                </div>
            </div>

        </div>
    </n-scrollbar>
    <SelectAlignment :value="selectAlignmentShow" :editData="formValue.align_objective"
        @close="() => { selectAlignmentShow = false }" @submit="submitSelectAlignment"></SelectAlignment>
</template>
<script setup lang="ts">
import { watch } from "vue";

import SelectAlignment from '@/views/components/SelectAlignment.vue'
import ItemList from "./ItemList.vue";
import UserList from "./UserList.vue";
import { addOkr, upDateOkr } from '@/api/modules/created'
import { useMessage } from "naive-ui"
import utils from "@/utils/utils";
import { UserStore } from '@/store/user'
import { AlertCircleOutline } from '@vicons/ionicons5'

const emit = defineEmits(['close', 'loadIng'])

const departmentOwner = UserStore().info.department_owner


const message = useMessage()
const loadIng = ref(false)
const formRef = ref()
const formRefs = ref()
const selectAlignmentShow = ref(false)
const userSelectApps = ref([]);
const showUserSelect = computed(() => window.Vues?.components?.UserSelect ? 1 : 0)


const props = defineProps({
    edit: {
        type: Boolean,
        default: false,
    },
    editData: {
        type: Object,
        default: {},
    },
    labelPlacement: {
        type: undefined,
        default: 'left',
    },
    labelAlign: {
        type: undefined,
        default: 'right',
    },
})



const formValue = ref<any>({
    title: "",
    type: 1,
    priority: "P0",
    ascription: 1,
    visible_range: 1,
    time: null,
    align_objective: null,
    project_id: null,
})


const formKRValue = ref<any>([
    {
        title: null,
        time: null,
        confidence: null,
        participant: null,
    },

])


//编辑
watch(() => props.edit, (newValue) => {
    if (newValue) {
        formValue.value = utils.cloneJSON(props.editData)
        formKRValue.value = utils.cloneJSON(props.editData.key_results)
        if (formValue.value.project_id == 0) formValue.value.project_id = null
        props.editData.key_results.map((item, index) => {
            formKRValue.value[index].time = [Date.parse(item.start_at), Date.parse(item.end_at)]
            formKRValue.value[index].participant = item.participant.split(",").map(Number)
        })
        formValue.value.time = [Date.parse(props.editData.start_at), Date.parse(props.editData.end_at)]
    }
}, { immediate: true })

watch(() => formValue.value.type, (newValue) => {
    if (newValue) {
        formValue.value.type == 1 ? formValue.value.priority = 'P0' : formValue.value.priority = 'P2'
    }
})


const generalOptions = ref([
    { label: $t('全公司'), value: 1 },
    { label: $t('仅相关成员'), value: 2 },
    { label: $t('仅部门成员'), value: 3 },
])

const rules = <any>{
    title: {
        required: true,
        message: $t('请输入目标'),
        trigger: ['input']
    },
    type: {
        type: 'number',
        required: true,
        trigger: 'change',
        message: $t('请选择类型')
    },
    priority: {
        required: true,
        trigger: 'change',
        message: $t('请选择优先级')
    },
    ascription: {
        type: 'number',
        required: true,
        trigger: 'change',
        message: $t('请选择归属')
    },
    time: {
        type: 'array',
        required: true,
        trigger: 'change',
        message: $t('请选择周期')
    }
}

const timeRule = <any>{
    title: {
        required: true,
        message: $t('请输入KR内容'),
        trigger: ['input']
    },
    time: {
        type: 'array',
        required: true,
        trigger: 'change',
        message: $t('请选择时间')
    }
}

const submitSelectAlignment = (e) => {
    formValue.value.align_objective = e
}

//提交
const handleSubmit = () => {
    formRef.value?.validate((errors) => {
        formRefs.value?.forEach(element => {
            element.validate((errors) => {
                if (errors) return false;
            })
        });
        if (errors) return false;
        const keyResults = []
        for (let index = 0; index < formKRValue.value.length; index++) {
            if (!formKRValue.value[index].title || !formKRValue.value[index].time) return
            keyResults.push({
                id: formKRValue.value[index].id || 0,
                title: formKRValue.value[index].title,
                confidence: formKRValue.value[index].confidence == null ? 0 : formKRValue.value[index].confidence,
                participant: formKRValue.value[index].participant == null ? "" : formKRValue.value[index].participant.join(','),
                start_at: utils.formatDate('Y-m-d 00:00:00', formKRValue.value[index].time[0] / 1000),
                end_at: utils.formatDate('Y-m-d 23:59:59', formKRValue.value[index].time[1] / 1000),
            })
        }
        const upData = {
            id: 0,
            title: formValue.value.title,
            type: formValue.value.type,
            priority: formValue.value.priority,
            ascription: formValue.value.ascription,
            visible_range: formValue.value.visible_range,
            start_at: utils.formatDate('Y-m-d 00:00:00', formValue.value.time[0] / 1000),
            end_at: utils.formatDate('Y-m-d 23:59:59', formValue.value.time[1] / 1000),
            align_objective: formValue.value.align_objective == null ? "" : formValue.value.align_objective.join(','),
            project_id: formValue.value.project_id == null ? 0 : formValue.value.project_id,
            key_results: keyResults,
        }
        loadIng.value = true
        emit('loadIng', loadIng.value)
        if (props.edit) {
            upData.id = formValue.value.id
            upDateOkr(upData)
                .then(({ msg }) => {
                    message.success($t('修改成功'))
                    emit('close', 2, formValue.value.id)
                })
                .catch(({ msg }) => {
                    message.error(msg)
                })
                .finally(() => {
                    loadIng.value = false
                    emit('loadIng', loadIng.value)
                })
        } else {
            addOkr(upData)
                .then(({ msg }) => {
                    message.success($t('添加成功'))
                    emit('close', 1)
                })
                .catch(({ msg }) => {
                    message.error(msg)
                })
                .finally(() => {
                    loadIng.value = false
                    emit('loadIng', loadIng.value)
                })
        }

    })
}

// 加载选择用户组件
const loadUserSelects = () => {
    nextTick(() => {
        if (!window.Vues) return false;
        document.querySelectorAll('userselects').forEach(e => {
            let item = formKRValue.value[e.getAttribute('formkey')];
            let app = new window.Vues.Vue({
                el: e,
                store: window.Vues.store,
                render: (h: any) => {
                    return h(window.Vues?.components?.UserSelect, {
                        class: "okr-user-selects",
                        formkey: e.getAttribute('formkey'),
                        props: {
                            value: item.participant || [],
                            title: $t('选择参与人'),
                            border: true,
                            avatarSize: 23,
                        },
                        on: {
                            "on-show-change": (show: any, values: any) => {
                                if (!show) {
                                    item.participant = values;
                                }
                            }
                        }
                    })
                },
            });
            userSelectApps.value.push(app)
        })
    })
}

// 清除数据
const handleClear = () => {
    formValue.value = {
        title: "",
        type: 1,
        priority: "P0",
        ascription: 1,
        visible_range: 1,
        time: null,
        align_objective: null,
        project_id: null,
    }

    formKRValue.value = [
        {
            title: null,
            time: null,
            confidence: null,
            participant: null,
        },
    ]
}

//  添加kr
const handleAddKr = () => {
    formKRValue.value.push(
        {
            title: null,
            time: null,
            confidence: null,
            participant: null,
        },
    )
    loadUserSelects()
}

// 删除kr
const handleRemoveKr = (index) => {
    if (formKRValue.value.length == 1) return message.warning($t('至少需要一个KR！'))
    formKRValue.value.splice(index, 1)
    userSelectApps.value.forEach(app => {
        let dom = document.createElement("UserSelects")
        dom.setAttribute('formkey',app._vnode.data.formkey)
        app.$el.replaceWith(dom);
        app.$destroy()
    })
    userSelectApps.value = [];
    loadUserSelects()
}

// 对齐目标
const handleGoal = () => {
    selectAlignmentShow.value = true
}

// 关闭Drawer
const closeDrawer = () => {
    handleClear()
    userSelectApps.value.forEach(app => app.$destroy())
    emit('close')
}

// 显示
const showDrawer = () => {
    loadUserSelects()
}

// 卸载
window.addEventListener('apps-unmount', function () {
    userSelectApps.value.forEach(app => app.$destroy())
})

//判断初始化归属
nextTick(()=>{
    if(!departmentOwner){
        formValue.value.ascription = 2
    }
})


defineExpose({
    handleSubmit,
    closeDrawer,
    showDrawer,
})

</script>
<style lang="less" scoped>
:deep(.n-drawer-body-content-wrapper) {
    @apply relative;
}

.add-main-box {
    @apply flex-auto shrink-0;
}

.button-box {
    @apply flex gap-2 mt-32 flex-initial;
}

:deep(.n-drawer-header__close) {
    @apply absolute -left-36;

    &:focus {
        @apply bg-none;
    }

    i {
        @apply text-[#fff];
    }
}

:deep(.n-base-close:not(.n-base-close--disabled):focus::before) {
    @apply bg-transparent;
}

.span {
    @apply text-14 text-white px-6 py-4 rounded-full flex items-center leading-3 shrink-0;
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

.okr-user-selects {
    @apply w-full bg-none border-none !important;
}

</style>

<style >
:deep(.n-popover){
    padding: 24px !important;
}
</style>