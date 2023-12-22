<template >
    <div class="setting-main-okr">

        <p class="font-medium text-[#303133]">{{ $t('评分部门OKR') }}</p>
        <div class="mt-16" v-if="showUserSelect">
            <UserSelects/>
        </div>
        <p class="mt-8 text-xs text-[#515A6E] opacity-50">{{ $t('单选：指定可评分部门及部门负责人OKR的人员。') }}</p>

        <p class="mt-24 font-medium text-[#303133]">{{ $t('评分规则') }}</p>
        <div class="mt-16 font-medium">
            <n-form ref="settingFormRef"
                :model="formModel"
                label-placement="left"
                label-width="auto"
                require-mark-placement="right-hanging"
                :size="'medium'"
                :style="{ maxWidth: '640px'}"
            >
                <n-form-item :label="$t('自评占比')" path="self_score_weight">
                    <n-input-number v-model:value="formModel.self_score_weight" placeholder="0" :show-button="false" :min="0" :max="100" :on-update:value="onSelfScoreWeight">
                        <template #suffix>%</template>
                    </n-input-number>
                </n-form-item>
                <n-form-item :label="$t('上级占比')" path="superior_score_weight">
                    <n-input-number v-model:value="formModel.superior_score_weight" placeholder="0" :show-button="false" :min="0" :max="100" :on-update:value="onSuperiorScoreWeight">
                        <template #suffix>%</template>
                    </n-input-number>
                </n-form-item>
            </n-form>
        </div>
        <p class=" text-xs text-[#515A6E] opacity-50">{{ $t('自评与上级的占比总和需等于100%。') }}</p>
    </div>
</template>
<script setup lang="ts">
import { useMessage } from "@/utils/messageAll"
import { okrSetting } from '@/api/modules/system'

const message = useMessage()
const emit = defineEmits(['close', 'loadIng','submit'])

const showUserSelect = ref(window.Vues?.components?.UserSelect ? 1 : 0)
const userSelectApps = ref([]);
const loadIng = ref(false);
const formModel = ref({
    score_department_user: 0,
    self_score_weight: 0,
    superior_score_weight: 0,
    type: 'save'
})

// 加载选择用户组件
const loadUserSelects = () => {
    unmountUserSelectsApps()
    nextTick(() => {
        if (!window.Vues) return false;
        document.querySelectorAll('userselects').forEach(e => {
            let app = new window.Vues.Vue({
                el: e,
                store: window.Vues.store,
                render: (h: any) => {
                    return h(window.Vues?.components?.UserSelect, {
                        class: "okr-user-selects",
                        props: {
                            value: formModel.value.score_department_user ? [formModel.value.score_department_user] : [],
                            title: $t('选择用户'),
                            border: false,
                            avatarSize: 36,
                            addIcon: false,
                            avatarName: true,
                            forcedRadio: true,
                            showDisable: true,
                            multipleMax: 100,
                        },
                        on: {
                            "on-show-change": (show: any, values: any) => {
                                if (!show) {
                                    formModel.value.score_department_user = values[0]
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

// 卸载用户组件
const unmountUserSelectsApps = () => {
    if(userSelectApps.value){
        userSelectApps.value.forEach(app => {
            let dom = document.createElement("UserSelects")
            app.$el.replaceWith(dom);
            app.$destroy()
        })
        userSelectApps.value = [];
    }
}

// 监听输入
const onSelfScoreWeight = (val: number) => {
    if(val){
        formModel.value.self_score_weight = Number((val).toFixed(0));
        formModel.value.superior_score_weight = Number((100 - val).toFixed(0));
    }else{
        formModel.value.self_score_weight = val
    }
}

// 监听输入
const onSuperiorScoreWeight = (val: number) => {
    if(val){
        formModel.value.superior_score_weight = Number((val).toFixed(0));
        formModel.value.self_score_weight = Number((100 - val).toFixed(0));
    }else{
        formModel.value.superior_score_weight = val
    }
}

// 提交
const handleSubmit = (type) => {
    if(type != 'get'){
        if(!formModel.value.score_department_user){
            message.error($t('请选择评分人员'))
            return;
        }
        if(formModel.value.self_score_weight + formModel.value.superior_score_weight != 100){
            message.error($t('自评与上级的占比总和需等于100%'))
            return;
        }
        formModel.value.type = 'save';
    }else{
        formModel.value.type = 'get';
    }
    loadIng.value = true;
    emit('loadIng', loadIng.value)
    okrSetting(formModel.value).then(({ data }) => {
        if(formModel.value.type != 'get'){
            message.success($t('修改成功'))
            emit('submit')
            emit('close', 2)
        }else{
            formModel.value.score_department_user = data.score_department_user
            formModel.value.self_score_weight = data.self_score_weight
            formModel.value.superior_score_weight = data.superior_score_weight
            loadUserSelects()
        }
    })
    .catch(({ msg }) => {
        message.error(msg)
    })
    .finally(() => {
        loadIng.value = false
        emit('loadIng', loadIng.value)
    })
}

// 关闭Drawer
const closeDrawer = () => {
    unmountUserSelectsApps()
    emit('close')
}

// 卸载
window.addEventListener('apps-unmount', function () {
    closeDrawer()
})
onBeforeUnmount(() => {
    closeDrawer()
})

//
onMounted(() => {
    nextTick(()=>{
        showUserSelect.value = window.Vues?.components?.UserSelect ? 1 : 0
        handleSubmit('get')
        loadUserSelects()
    })
})

defineExpose({
    handleSubmit,
    closeDrawer,
    showDrawer,
})
</script>

<style lang="less" scoped>
:deep(.n-input-number){
    width: 100%;
}
</style>
