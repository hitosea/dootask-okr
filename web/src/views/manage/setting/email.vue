<template>
    <div class="setting-item">
        <n-form :rules="rules" :model="formData" label-placement="left" label-width="auto">
            <n-form-item :label="$t('新邮箱地址')" path="email">
                <n-input-group>
                    <n-input v-model:value="formData.email" clearable />
                    <n-button type="primary" :loading="loadIng" @click="sendEmailCode">{{ $t("发送验证码") }}</n-button>
                </n-input-group>
            </n-form-item>
            <n-form-item :label="$t('验证码')" path="code">
                <n-input v-model:value="formData.code" clearable />
            </n-form-item>
        </n-form>
        <div class="setting-footer">
            <n-button type="primary" :loading="loadIng" @click="handleSubmit">{{ $t("提交") }}</n-button>
            <n-button :loading="loadIng" @click="handleReset">{{ $t("重置") }}</n-button>
        </div>
    </div>
</template>
<script setup lang="ts">
import { ref } from "vue"
import { useMessage } from "naive-ui"
import utils from "@/utils/utils"
import { sendEmail, changeEmail } from "@/api/modules/user"
import { ResultDialog } from "@/api"
import { UserStore } from "@/store/user"

const userState = UserStore()
const message = useMessage()
const loadIng = ref<boolean>(false)
const count = ref(120)
const sendBtnText = ref('')

const formData = ref({
    email: "",
    code: "",
})
const rules = {
    email: { required: true, message: $t("请输入新邮箱地址"), trigger: ["blur", "input"], },
    code: { required: true, message: $t("请输入验证码"), trigger: ["blur", "input"], }
}

// 发送邮箱验证码
const sendEmailCode = () => {
    if (formData.value.email) message.success($t("请输入新邮箱地址"));
    if (!utils.isEmail(formData.value.email)) return message.info($t("请填写正确邮箱"))
    loadIng.value = true
    const upData = {
        email: formData.value.email,
    }
    sendEmail(upData)
        .then(({ msg }) => {
            message.success(msg)
                sendBtnText.value = count.value + $t('秒');
                let times = setInterval(() => {
                    count.value--; 
                    sendBtnText.value = count.value + $t('秒');
                    if (count.value <= 0) {
                        sendBtnText.value =  $t('发送验证码')
                        clearInterval(times);
                    }
                }, 1000);
        })
        .catch(ResultDialog)
        .finally(() => {
            loadIng.value = false
        })
}

// 修改邮箱
const handleSubmit = () => {
    loadIng.value = true
    const upData = {
        email: formData.value.email,
        id: userState.info.id,
    }
    changeEmail(upData)
        .then(({ msg }) => {
            message.success(msg)
        })
        .catch(ResultDialog)
        .finally(() => {
            loadIng.value = false
        })
}

//重置
const handleReset = () => {
    formData.value = {
        email: "",
        code: "",
    }
}
</script>
<style lang="less"></style>
