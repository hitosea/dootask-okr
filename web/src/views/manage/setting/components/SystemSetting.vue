<template>
    <div class="setting-component-item">
        <n-scrollbar>
            <n-form :model="formDatum" label-placement="left" label-width="auto">
                <div class="block-setting-box">
                    <h3>{{ $t("帐号相关") }}</h3>
                    <div class="form-box">
                        <n-form-item :label="$t('允许注册')">
                            <div class="select-box">
                                <n-radio-group v-model:value="formDatum.reg" name="reg">
                                    <n-space>
                                        <n-radio v-for="item in regList" :key="item.value" :value="item.value">
                                            {{ item.label }}
                                        </n-radio>
                                    </n-space>
                                </n-radio-group>
                                <div v-if="formDatum.reg == 'open'" class="form-tip">{{ $t("允许：开放注册功能。") }}</div>
                                <template v-else-if="formDatum.reg == 'invite'">
                                    <div class="form-tip">{{ $t("邀请码：注册时需填写下方邀请码。") }}</div>
                                    <n-input-group>
                                        <n-input-group-label>{{ $t("邀请码") }}</n-input-group-label>
                                        <n-input v-model:value="formDatum.reg_invite" />
                                    </n-input-group>
                                </template>
                            </div>
                        </n-form-item>
                        <n-form-item v-show="['open', 'invite'].includes(formDatum.reg)" label="$t('注册身份')">
                            <div class="select-box">
                                <n-radio-group v-model:value="formDatum.reg_identity" name="reg_identity">
                                    <n-space>
                                        <n-radio v-for="item in regIdentityList" :key="item.value" :value="item.value">
                                            {{ item.label }}
                                        </n-radio>
                                    </n-space>
                                </n-radio-group>
                                <div class="form-tip form-list">
                                    <p>{{ $t("临时帐号：") }}</p>
                                    <ol>
                                        <li>{{ $t("禁止查看共享所有人的文件。") }}</li>
                                        <li>{{ $t("禁止发起会话。") }}</li>
                                        <li>{{ $t("禁止创建群聊。") }}</li>
                                        <li>{{ $t("禁止拨打电话。") }}</li>
                                    </ol>
                                </div>
                            </div>
                        </n-form-item>

                        <n-form-item :label="$t('登录验证码')">
                            <div class="select-box">
                                <n-radio-group v-model:value="formDatum.login_code" name="login_code">
                                    <n-space>
                                        <n-radio v-for="item in loginCodeList" :key="item.value" :value="item.value">
                                            {{ item.label }}
                                        </n-radio>
                                    </n-space>
                                </n-radio-group>
                                <div v-if="formDatum.login_code == 'auto'" class="form-tip">
                                    {{ $t("自动：密码输入错误后必须添加验证码。") }}
                                </div>
                                <div v-else-if="formDatum.login_code == 'open'" class="form-tip">
                                    {{ $t("开启：每次登录都需要图形验证码。") }}
                                </div>
                                <div v-else-if="formDatum.login_code == 'close'" class="form-tip">
                                    {{ $t("关闭：不需要输入图形验证。") }}
                                </div>
                            </div>
                        </n-form-item>
                        <n-form-item :label="$t('密码策略')">
                            <div class="select-box">
                                <n-radio-group v-model:value="formDatum.password_policy" name="password_policy">
                                    <n-space>
                                        <n-radio v-for="item in passwordPolicyList" :key="item.value" :value="item.value">
                                            {{ item.label }}
                                        </n-radio>
                                    </n-space>
                                </n-radio-group>
                                <div v-if="formDatum.password_policy == 'simple'" class="form-tip">
                                    {{ $t("简单：大于或等于6个字符。") }}
                                </div>
                                <div v-else-if="formDatum.password_policy == 'complex'" class="form-tip">
                                    {{ $t("复杂：大于或等于6个字符，包含数字、字母大小写或者特殊字符。") }}
                                </div>
                            </div>
                        </n-form-item>
                    </div>
                </div>

                <div class="block-setting-box">
                    <h3>项目相关</h3>
                    <div class="form-box">
                        <n-form-item label="邀请项目">
                            <div class="select-box">
                                <n-radio-group v-model:value="formDatum.project_invite" name="reg">
                                    <n-space>
                                        <n-radio v-for="item in projectInviteList" :key="item.value" :value="item.value">
                                            {{ item.label }}
                                        </n-radio>
                                    </n-space>
                                </n-radio-group>
                                <div v-if="formDatum.project_invite == 'open'" class="form-tip">
                                    {{ $t("开启：项目管理员可生成链接邀请成员加入项目。") }}
                                </div>
                            </div>
                        </n-form-item>

                        <n-form-item label="自动归档任务">
                            <div class="select-box">
                                <n-radio-group v-model:value="formDatum.auto_archived" name="reg">
                                    <n-space>
                                        <n-radio v-for="item in autoArchivedList" :key="item.value" :value="item.value">
                                            {{ item.label }}
                                        </n-radio>
                                    </n-space>
                                </n-radio-group>
                                <div class="form-tip">{{ $t("任务完成后自动归档。") }}</div>
                                <n-tooltip trigger="hover" placement="right" v-if="formDatum.auto_archived == 'open'">
                                    <template #trigger>
                                        <n-input-group>
                                            <n-input v-model:value="formDatum.archived_day" type="text"
                                                :allow-input="onlyAllowNumber" />
                                            <n-input-group-label>天</n-input-group-label>
                                        </n-input-group>
                                    </template>
                                    {{ $t("任务完成 7 天后自动归档。") }}
                                </n-tooltip>
                            </div>
                        </n-form-item>
                    </div>
                </div>
                <div class="block-setting-box">
                    <h3>{{ $t("消息相关") }}</h3>
                    <div class="form-box">
                        <n-form-item :label="$t('全员群组禁言')">
                            <div class="select-box">
                                <n-radio-group v-model:value="formDatum.all_group_mute" name="all_group_mute">
                                    <n-space>
                                        <n-radio v-for="item in allGroupMuteList" :key="item.value" :value="item.value">
                                            {{ item.label }}
                                        </n-radio>
                                    </n-space>
                                </n-radio-group>
                                <div v-if="formDatum.all_group_mute == 'open'" class="form-tip">
                                    {{ $t("开放：所有人都可以发言。") }}
                                </div>
                                <div v-else-if="formDatum.all_group_mute == 'user'" class="form-tip">
                                    {{ $t("成员禁言：仅管理员可以发言。") }}
                                </div>
                                <div v-else-if="formDatum.all_group_mute == 'all'" class="form-tip">
                                    {{ $t("全部禁言：所有人都禁止发言。") }}
                                </div>
                            </div>
                        </n-form-item>
                        <n-form-item :label="$t('自动进入全员群')">
                            <div class="select-box">
                                <n-radio-group v-model:value="formDatum.all_group_autoin" name="all_group_autoin">
                                    <n-space>
                                        <n-radio v-for="item in allGroupAutoinList" :key="item.value" :value="item.value">
                                            {{ item.label }}
                                        </n-radio>
                                    </n-space>
                                </n-radio-group>
                                <div v-if="formDatum.all_group_autoin == 'yes'" class="form-tip">
                                    {{ $t("自动：注册成功后自动进入全员群。") }}
                                </div>
                                <div v-else-if="formDatum.all_group_autoin == 'no'" class="form-tip">
                                    {{ $t("关闭：其他成员通过@邀请进入。") }}
                                </div>
                            </div>
                        </n-form-item>
                        <n-form-item :label="$t('聊天资料')">
                            <div class="select-box">
                                <n-radio-group v-model:value="formDatum.chat_information" name="chat_information">
                                    <n-space>
                                        <n-radio v-for="item in chatInformationList" :key="item.value" :value="item.value">
                                            {{ item.label }}
                                        </n-radio>
                                    </n-space>
                                </n-radio-group>
                                <div v-if="formDatum.chat_information == 'required'" class="form-tip">
                                    {{ $t("必填：发送聊天内容前必须设置昵称、电话。") }}
                                </div>
                                <div v-else class="form-tip">{{ $t("如果必填，发送聊天前必须设置昵称、电话。") }}</div>
                            </div>
                        </n-form-item>
                        <n-form-item :label="$t('匿名消息')">
                            <div class="select-box">
                                <n-radio-group v-model:value="formDatum.anon_message" name="anon_message">
                                    <n-space>
                                        <n-radio v-for="item in anonMessageList" :key="item.value" :value="item.value">
                                            {{ item.label }}
                                        </n-radio>
                                    </n-space>
                                </n-radio-group>
                                <div v-if="formDatum.anon_message == 'open'" class="form-tip">
                                    {{ $t("允许匿名发送消息给其他成员。") }}
                                </div>
                                <div v-else class="form-tip">{{ $t("禁止匿名发送消息。") }}</div>
                            </div>
                        </n-form-item>
                    </div>
                </div>
                <div class="block-setting-box">
                    <h3>{{ $t("其他设置") }}</h3>
                    <div class="form-box">
                        <n-form-item :label="$t('图片优化')">
                            <div class="select-box">
                                <n-radio-group v-model:value="formDatum.image_compress" name="password_policy">
                                    <n-space>
                                        <n-radio v-for="item in imageCompressList" :key="item.value" :value="item.value">
                                            {{ item.label }}
                                        </n-radio>
                                    </n-space>
                                </n-radio-group>
                                <div class="form-tip">
                                    {{ $t("数码相机4M的图片，优化后仅有700KB左右，而且肉眼基本看不出区别。") }}
                                </div>
                            </div>
                        </n-form-item>
                        <n-form-item :label="$t('保存网络图片')">
                            <div class="select-box">
                                <n-radio-group v-model:value="formDatum.image_save_local" name="password_policy">
                                    <n-space>
                                        <n-radio v-for="item in imageSaveLocalList" :key="item.value" :value="item.value">
                                            {{ item.label }}
                                        </n-radio>
                                    </n-space>
                                </n-radio-group>
                                <div class="form-tip">
                                    {{ $t("是否将消息中的网络图片保存到本地服务器。") }}
                                </div>
                            </div>
                        </n-form-item>
                        <n-form-item :label="$t('是否启动首页')">
                            <div class="select-box">
                                <n-radio-group v-model:value="formDatum.start_home" name="password_policy">
                                    <n-space>
                                        <n-radio v-for="item in startHomelList" :key="item.value" :value="item.value">
                                            {{ item.label }}
                                        </n-radio>
                                    </n-space>
                                </n-radio-group>
                                <div class="form-tip">{{ $t("仅支持网页版。") }}</div>
                                <n-input v-if="formDatum.start_home == 'open'" v-model:value="formDatum.home_footer"
                                    :rows="2" :autosize="{ minRows: 2, maxRows: 8 }" type="textarea"
                                    :placeholder="$t('首页底部：首页底部网站备案号等信息')" />
                            </div>
                        </n-form-item>
                    </div>
                </div>
            </n-form>
        </n-scrollbar>

        <div class="setting-footer">
            <n-button @click="handleSubmit" type="primary">{{ $t("提交") }}</n-button>
            <n-button @click="getSystemSetting">{{ $t("重置") }}</n-button>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref } from "vue"
import { systemSetting } from "@/api/modules/system"
import { useMessage } from "naive-ui"
import { ResultDialog } from "@/api"

const loadIng = ref(false)
const message = useMessage()

const formDatum = ref<any>({
    reg: "",
    reg_invite: "",
    reg_identity: "",
    login_code: "",
    password_policy: "",
    project_invite: "",
    auto_archived: "",
    archived_day: "",
    all_group_mute: "",
    all_group_autoin: "",
    chat_information: "",
    anon_message: "",
    image_compress: "",
    image_save_local: "",
    start_home: "",
    home_footer: "",
})

const regList = [
    { value: "open", label: $t("允许") },
    { value: "invite", label: $t("邀请码") },
    { value: "close", label: $t("禁止") },
]
const regIdentityList = [
    { value: "normal", label: $t("正常帐号") },
    { value: "temp", label: $t("临时帐号") },
]
const loginCodeList = [
    { value: "auto", label: $t("自动") },
    { value: "open", label: $t("开启") },
    { value: "close", label: $t("关闭") },
]
const passwordPolicyList = [
    { value: "simple", label: $t("简单") },
    { value: "complex", label: $t("复杂") },
]
const projectInviteList = [
    { value: "open", label: $t("开启") },
    { value: "close", label: $t("关闭") },
]
const autoArchivedList = [
    { value: "open", label: $t("开启") },
    { value: "close", label: $t("关闭") },
]
const allGroupMuteList = [
    { value: "open", label: $t("开启") },
    { value: "user", label: $t("成员禁言") },
    { value: "all", label: $t("全部禁言") },
]
const allGroupAutoinList = [
    { value: "yes", label: $t("自动") },
    { value: "no", label: $t("关闭") },
]
const chatInformationList = [
    { value: "optional", label: $t("可选") },
    { value: "required", label: $t("必填") },
]
const anonMessageList = [
    { value: "open", label: $t("开启") },
    { value: "close", label: $t("关闭") },
]
const imageCompressList = [
    { value: "open", label: $t("开启") },
    { value: "close", label: $t("关闭") },
]
const imageSaveLocalList = [
    { value: "open", label: $t("开启") },
    { value: "close", label: $t("关闭") },
]
const startHomelList = [
    { value: "open", label: $t("开启") },
    { value: "close", label: $t("关闭") },
]

// 只能用数字
const onlyAllowNumber = (value: string) => !value || /^\d+$/.test(value)

//获取设置
const getSystemSetting = () => {
    loadIng.value = true
    const upData = {
        type: "get"
    }
    systemSetting(upData).then(({ data }) => {
        formDatum.value = data
    })
        .catch(ResultDialog)
        .finally(() => {
            loadIng.value = false
        })
}
getSystemSetting()

// 提交设置
const handleSubmit = () => {
    loadIng.value = true
    const upData = {...formDatum.value}
    upData.type = "save"
    systemSetting(upData).then(({ data, msg }) => {
        formDatum.value = data
        message.success(msg)
    })
        .catch(ResultDialog)
        .finally(() => {
            loadIng.value = false
        })
}
</script>
<style lang="less" scoped>
:deep(.n-form-item-blank) {
    @apply pt-6 items-start !important;
}
</style>
