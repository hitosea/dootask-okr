<template>
    <div class="setting-item">
        <n-form ref="formRef" :rules="rules" :model="formData" label-placement="left" label-width="auto" >
            <n-form-item label="旧密码" path="password">
                <n-input type="password" v-model:value="formData.password" clearable />
            </n-form-item>
            <n-form-item label="新密码" path="newPassword">
                <n-input type="password" v-model:value="formData.newPassword" clearable />
            </n-form-item>
            <n-form-item label="确认新密码" path="newPassword2">
                <n-input type="password" v-model:value="formData.newPassword2" clearable />
            </n-form-item>
        </n-form>
        <div class="setting-footer">
            <n-button type="primary" :loading="loadIng" @click="handleSubmit">提交</n-button>
            <n-button :loading="loadIng" @click="handleReset">重置</n-button>
        </div>
    </div>
</template>
<script setup lang="ts">
import { ref } from "vue"
import { FormItemRule, FormInst, useMessage } from "naive-ui"
import { editPassword } from "@/api/modules/user"
import { ResultDialog } from "@/api"
import { UserStore } from "@/store/user"

const formRef = ref<FormInst | null>()
const userInfo = UserStore()
const message = useMessage()
const loadIng = ref<boolean>(false)
const formData = ref({
    password: "",
    newPassword: "",
    newPassword2: "",
})
// 表单验证
const rules = {
    password: {
        required: true,
        trigger: ["blur", "input"],
        validator(rule: FormItemRule, value: string) {
            if (value.length == 0) {
                return new Error($t('请输入旧密码！'))
            }
            else if (value.length < 6) {
                return new Error($t('密码长度至少6位！'))
            }
            return true
        },
    },
    newPassword: {
        required: true,
        trigger: ["blur", "input"],
        validator(rule: FormItemRule, value: string) {
            if (value.length == 0) {
                return new Error($t('请输入新密码！'))
            }
            else if (value.length < 6) {
                return new Error($t('密码长度至少6位！'))
            }
            return true
        },
    },
    newPassword2: {
        required: true,
        trigger: ["blur", "input"],
        validator(rule: FormItemRule, value: string) {
            if (value.length == 0) {
                return new Error($t('请重新输入新密码！'))
            }
            else if (value != formData.value.newPassword) {
                return new Error($t('两次密码输入不一致！'))
            }
            return true
        },
    },
}

// 提交
const handleSubmit = () => {
    formRef.value?.validate((errors) => {
        if (!errors) {
            const upData = {
                id: userInfo.info.id,
                old_password: formData.value.password,
                new_password: formData.value.newPassword,
            }
            loadIng.value = true
            editPassword(upData).then(({ msg }) => {
                message.success(msg)
            })
            .catch(ResultDialog)
            .finally(() => {
                loadIng.value = false
            })
        } 
    })

}

// 重置
const handleReset = () => {
    formData.value = {
        password: "",
        newPassword: "",
        newPassword2: "",
    }
}
</script>
<style lang="less"></style>
