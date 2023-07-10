<template>
    <div class="setting-item">
        <n-form :rules="rules" :model="formData" label-placement="left" label-width="auto">
            <n-form-item :label="$t('头像')">
                <img-up-load :maxNumber="1" v-model:value="formData.userimg"/>
            </n-form-item>
            <n-form-item :label="$t('邮箱')" path="email">
                <n-input v-model:value="formData.email" disabled />
            </n-form-item>
            <n-form-item :label="$t('电话')" path="phone">
                <n-input v-model:value="formData.phone" clearable />
            </n-form-item>
            <n-form-item :label="$t('昵称')" path="nickname">
                <n-input v-model:value="formData.nickname" clearable />
            </n-form-item>
            <n-form-item :label="$t('职位/职称')" path="profession">
                <n-input v-model:value="formData.profession" clearable />
            </n-form-item>
        </n-form>
        <div class="setting-footer">
            <n-button type="primary">{{ $t('提交') }}</n-button>
            <n-button>{{ $t('重置') }}</n-button>
        </div>
    </div>
</template>
<script setup lang="ts">
import { ref } from "vue"
import ImgUpLoad from "@/components/ImgUpLoad.vue"
const formData = ref({
    userimg: "",
    email: "",
    phone: "",
    nickname: "",
    profession: "",
})

const rules = {
    email: {
        required: true,
        message: "请输入邮箱",
        trigger: ["blur", "input"],
    },
    phone: {
        required: true,
        message: "请输入电话号码",
        trigger: ["blur", "input"],
    },
    nickname: {
        required: true,
        message: "请输入昵称",
        trigger: ["blur", "input"],
    },
    profession: {
        required: true,
        message: "请输入职位/职称",
        trigger: ["blur", "input"],
    },
}



</script>
<style lang="less"></style>
