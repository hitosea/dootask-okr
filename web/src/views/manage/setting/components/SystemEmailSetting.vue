<template>
    <div class="setting-component-item">
        <n-scrollbar>
            <n-form :model="formData" label-placement="left" label-width="auto">
                <div class="block-setting-box">
                    <h3>{{ $t("帐号相关") }}</h3>
                    <div class="form-box">
                        <n-form-item :label="$t('SMTP服务器')">
                            <div class="select-box">
                                <n-input v-model:value="formData.smtp_server" type="text" />
                            </div>
                        </n-form-item>
                        <n-form-item :label="$t('端口')">
                            <div class="select-box">
                                <n-input v-model:value="formData.port" type="text" />
                            </div>
                        </n-form-item>
                        <n-form-item :label="$t('帐号')">
                            <div class="select-box">
                                <n-input v-model:value="formData.account" type="text" />
                            </div>
                        </n-form-item>
                        <n-form-item :label="$t('密码')">
                            <div class="select-box">
                                <n-input v-model:value="formData.password" type="text" />
                            </div>
                        </n-form-item>
                        <n-form-item label=" ">
                            <div class="select-box">
                                <n-button>{{ $t("邮件发送测试") }}</n-button>
                            </div>
                        </n-form-item>
                    </div>
                </div>

                <div class="block-setting-box">
                    <h3>{{ $t("帐号相关") }}</h3>
                    <div class="form-box">
                        <n-form-item class="ridio-from-box" :label="$t('邮件通知设置')">
                            <div class="select-box">
                                <n-radio-group v-model:value="formData.reg_verify" name="reg_identity">
                                    <n-space>
                                        <n-radio v-for="item in regVerifyList" :key="item.value" :value="item.value">
                                            {{ item.label }}
                                        </n-radio>
                                    </n-space>
                                </n-radio-group>
                                <div class="form-tip" v-if="formData.reg_verify == 'open'">
                                    {{ $t('开启后：') }}<br />
                                    ① {{ $t('帐号需验证通过才可登录') }}<br />
                                    ② {{ $t('修改邮箱和删除帐号需要邮箱验证码') }}
                                </div>
                            </div>
                        </n-form-item>

                        <n-form-item class="ridio-from-box" :label="$t('消息提醒')">
                            <div class="select-box">
                                <n-radio-group v-model:value="formData.notice_msg" name="reg">
                                    <n-space>
                                        <n-radio v-for="item in noticeMsgList" :key="item.value" :value="item.value">
                                            {{ item.label }}
                                        </n-radio>
                                    </n-space>
                                </n-radio-group>
                                <n-form class="mt-12" v-if="formData.notice_msg == 'open'" @submit.native.prevent
                                    label-placement="left" label-width="auto">
                                    <n-form-item class="input-from-box" :label="$t('未读个人消息')" prop="msg_unread_user_minute">
                                        <div class="input-number-box">
                                            <n-input-number v-model:value="formData.msg_unread_user_minute"
                                                :show-button="false" :min="0" :step="1">
                                                <template #suffix>
                                                    {{ $t('分钟') }}(m)
                                                </template>
                                            </n-input-number>
                                        </div>
                                    </n-form-item>
                                    <n-form-item class="input-from-box" :label="$t('未读群聊消息')"
                                        prop="msg_unread_group_minute">
                                        <div class="input-number-box">
                                            <n-input-number v-model:value="formData.msg_unread_group_minute"
                                                :show-button="false" :min="0" :step="1">
                                                <template #suffix>
                                                    {{ $t('分钟') }}(m)
                                                </template>
                                            </n-input-number>
                                        </div>
                                    </n-form-item>
                                    <div class="form-tip">{{ $t('填写-1则不通知，误差±10分钟') }}</div>
                                </n-form>
                            </div>
                        </n-form-item>
                    </div>
                </div>

                <div class="block-setting-box">
                    <h3>{{ $t("忽略邮箱地址") }}</h3>
                    <div class="form-box">
                        <n-form-item :label="$t('忽略邮箱')">
                            <div class="select-box">
                                <n-input v-model:value="formData.ignore_addr" :rows="2"
                                    :autosize="{ minRows: 3, maxRows: 8 }" type="textarea" />
                                <div class="form-tip">{{ $t('不会向忽略的邮箱地址发送邮件，可使用换行分割多个地址。') }}</div>
                            </div>

                        </n-form-item>

                    </div>
                </div>
            </n-form>
        </n-scrollbar>

        <div class="setting-footer">
            <n-button type="primary" :loading="loadIng" @click="handleSubmit">{{ $t("提交") }}</n-button>
            <n-button :loading="loadIng" @click="getSystemSettingMail">{{ $t("重置") }}</n-button>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref } from "vue"
import { systemSettingMail } from "@/api/modules/system"
import { useMessage } from "naive-ui"
import { ResultDialog } from "@/api"

const loadIng = ref(false)
const message = useMessage()

const formData = ref<any>({
    smtp_server: "",
    port: null,
    account: "",
    password: "",
    reg_verify: "",
    notice_msg: "",
    msg_unread_user_minute: 0,
    msg_unread_group_minute: 0,
    ignore_addr: "",
})

const regVerifyList = [
    { value: "open", label: $t("开启") },
    { value: "close", label: $t("关闭") },
]
const noticeMsgList = [
    { value: "open", label: $t("开启") },
    { value: "close", label: $t("关闭") },
]

// 获取邮箱设置
const getSystemSettingMail = () => {
    loadIng.value = true
    const upData = {
        type: "get"
    }
    systemSettingMail(upData).then(({ data }) => {
        formData.value = data
    })
        .catch(ResultDialog)
        .finally(() => {
            loadIng.value = false
        })
}
getSystemSettingMail()

const handleSubmit = () => {
    loadIng.value = true
    const upData = { ...formData.value}
    upData.port = Number(upData.port)
    upData.msg_unread_user_minute = Number(upData.msg_unread_user_minute)
    upData.msg_unread_group_minute = Number(upData.msg_unread_group_minute)
    upData.type = "save"
    systemSettingMail(upData).then(({ data, msg }) => {
        formData.value = data
        message.success(msg)
    })
        .catch(ResultDialog)
        .finally(() => {
            loadIng.value = false
        })
}
</script>
<style lang="less" scoped>
.ridio-from-box {
    :deep(.n-form-item-blank) {
        @apply pt-6 items-start !important;
    }
}

.input-from-box {
    :deep(.n-form-item-blank) {
        @apply pt-0 items-center !important;
    }
}
</style>