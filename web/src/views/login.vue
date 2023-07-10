<template>
    <div class="page-login child-view">
        <div class="login-body">
            <div class="login-logo"></div>
            <div class="login-box">
                <div class="login-mode-switch">
                    <div class="login-mode-switch-box">
                        <span class="login-mode-switch-icon" @click="switchLoginMode">
                            <n-tooltip trigger="hover" placement="left">
                                <template #trigger>
                                    <n-input-group>
                                        <svg v-if="loginMode == 'qrcode'" viewBox="0 0 24 24" fill="currentColor"
                                            xmlns="http://www.w3.org/2000/svg" data-icon="PcOutlined">
                                            <path
                                                d="M23 16a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h18a2 2 0 0 1 2 2v12ZM21 4H3v9h18V4ZM3 15v1h18v-1H3Zm3 6a1 1 0 0 1 1-1h10a1 1 0 1 1 0 2H7a1 1 0 0 1-1-1Z"
                                                fill="currentColor"></path>
                                        </svg>
                                        <svg v-else viewBox="0 0 24 24" fill="currentColor"
                                            xmlns="http://www.w3.org/2000/svg" data-icon="QrOutlined">
                                            <path
                                                d="M6.5 7.5a1 1 0 0 1 1-1h1a1 1 0 0 1 1 1v1a1 1 0 0 1-1 1h-1a1 1 0 0 1-1-1v-1Z"
                                                fill="currentColor"></path>
                                            <path
                                                d="M4.5 2.5c-1.1 0-2 .9-2 2v7c0 1.1.9 2 2 2h7c1.1 0 2-.9 2-2v-7c0-1.1-.9-2-2-2h-7Zm0 2h7v7h-7v-7ZM11 16a1 1 0 1 1 2 0 1 1 0 0 1-2 0Zm0 3.5a1 1 0 1 1 2 0v1a1 1 0 1 1-2 0v-1Zm4-7.5a1 1 0 1 1 2 0 1 1 0 0 1-2 0Zm3.5 0a1 1 0 0 1 1-1h1a1 1 0 1 1 0 2h-1a1 1 0 0 1-1-1ZM15 17c0-1.1.9-2 2-2h2.5c1.1 0 2 .9 2 2v2.5c0 1.1-.9 2-2 2H17c-1.1 0-2-.9-2-2V17Zm4.5 0H17v2.5h2.5V17Zm-15-2c-1.1 0-2 .9-2 2v2.5c0 1.1.9 2 2 2H7c1.1 0 2-.9 2-2V17c0-1.1-.9-2-2-2H4.5Zm0 2H7v2.5H4.5V17ZM15 4.5c0-1.1.9-2 2-2h2.5c1.1 0 2 .9 2 2V7c0 1.1-.9 2-2 2H17c-1.1 0-2-.9-2-2V4.5Zm4.5 0H17V7h2.5V4.5Z"
                                                fill="currentColor"></path>
                                        </svg>
                                    </n-input-group>
                                </template>
                                {{ loginMode == "qrcode" ? $t("帐号登录") : $t("扫码登录") }}
                            </n-tooltip>
                        </span>
                    </div>
                </div>
                <h2 class="login-title">{{ loginMode == "qrcode" ? $t("扫码登录") : "Welcome DooTask" }}</h2>
                <p class="login-subtitle">
                    {{ loginMode == "qrcode" ? $t("请使用DooTask移动端扫描二维码。") : $t("输入您的凭证以访问您的帐户。") }}
                </p>
                <transition name="login-mode">
                    <div v-if="loginMode == 'qrcode'" class="login-qrcode" @click="qrcodeRefresh">
                        <vue-qrcode :value="qrcodeUrl" :size="200" :margin="0"></vue-qrcode>
                    </div>
                </transition>
                <transition name="login-mode">
                    <div v-if="loginMode == 'access'" class="login-access">
                        <n-input v-model:value="formData.email" @blur="onBlur" :placeholder="$t('输入您的电子邮箱')" clearable
                            size="large">
                            <template #prefix>
                                <n-icon :component="MailOutline" />
                            </template>
                        </n-input>
                        <n-input type="password" v-model:value="formData.password" :placeholder="$t('输入您的密码')" clearable
                            size="large">
                            <template #prefix>
                                <n-icon :component="LockClosedOutline" />
                            </template>
                        </n-input>
                        <n-input class="code-load-input" v-model:value="code" :placeholder="$t('输入图形验证码')" clearable
                            size="large" v-if="codeNeed">
                            <template #prefix>
                                <n-icon :component="CheckmarkCircleOutline" />
                            </template>
                            <template #suffix>
                                <div class="login-code-end" @click="refreshCode">
                                    <div v-if="codeLoad > 0" class="code-load">
                                        <Loading />
                                    </div>
                                    <span v-else-if="codeUrl === 'error'" class="code-error">{{ $t("加载失败") }}</span>
                                    <img v-else :src="codeUrl" />
                                </div>
                            </template>
                        </n-input>
                        <n-input type="password" v-model:value="formData.confirmPassword" v-if="loginType == 'reg'"
                            :placeholder="$t('输入确认密码')" clearable size="large">
                            <template #prefix>
                                <n-icon :component="LockClosedOutline" />
                            </template>
                        </n-input>
                        <n-input v-model:value="formData.invite" v-if="loginType == 'reg'" :placeholder="$t('请输入注册邀请码')"
                            clearable size="large">
                            <template #prefix> {{ $t("邀请码") }} </template>
                        </n-input>
                        <n-button v-if="loginType == 'login'" :loading="loadIng" @click="handleLogin" type="primary"
                            size="large">{{ $t("登录") }}</n-button>
                        <n-button v-else type="primary" :loading="loadIng" @click="handleReg">{{ $t("注册") }}</n-button>
                        <div class="login-switch">
                            <template v-if="loginType == 'login'">
                                {{ $t("还没有帐号？") }}
                                <a href="javascript:void(0)" @click="changeLoginType"> {{ $t("注册帐号") }}</a>
                            </template>
                            <template v-else>
                                {{ $t("已经有帐号？") }}
                                <a href="javascript:void(0)" @click="changeLoginType"> {{ $t("登录帐号") }}</a>
                            </template>
                        </div>
                    </div>
                </transition>
            </div>
            <div class="login-bottom max-w-90p">
                <n-dropdown :options="options" placement="bottom-start" trigger="click" @select="handleSelect">
                    <p class="login-setting">{{ $t("设置") }} <n-icon size="16" :component="CaretDown" /></p>
                </n-dropdown>
                <div class="login-forgot">{{ $t("忘记密码了？") }}<a href="javascript:void(0)">{{ $t("重置密码") }}</a></div>
            </div>
        </div>
    </div>
</template>

<script lang="ts" setup>
import { ref, computed } from "vue"
import { userLogin, userReg, codeImg, needCode, getQrcodeStatus } from "@/api/modules/user"
import { ResultDialog } from "@/api"
import { useMessage } from "naive-ui"
import utils from "@/utils/utils"
import webTs from "@/utils/web"
import { UserStore } from "@/store/user"
import { useRouter } from "vue-router"
import { MailOutline, LockClosedOutline, CaretDown, CheckmarkCircleOutline } from "@vicons/ionicons5"
import { GlobalStore } from "@/store"
import VueQrcode from "qrcode.vue"

const globalStore = GlobalStore()
const router = useRouter()
const message = useMessage()
const loadIng = ref<boolean>(false)
const code = ref("")
const codeUrl = ref("")
const codeLoad = ref(0)
const userState = UserStore()
const loginMode = ref("access")
const qrcodeVal = ref("")
const codeNeed = ref(false)
const qrcodeRefresh = () => { }
const codeId = ref("")
const qrcodeTimer = ref<any>()
const loginType = ref<String>("login")
const formData = ref({
    email: "",
    password: "",
    confirmPassword: "",
    invite: "",
})
// 设置菜单选项
const options = [
    {
        label: $t("主题皮肤"),
        key: "theme",
        children: [
            { label: $t("跟随系统"), key: "flow" },
            { label: $t("明亮"), key: "light" },
            { label: $t("暗黑"), key: "dark" },
        ],
    },
    {
        label: $t("语言"),
        key: "language",
        children: [
            { label: $t("简体中文"), key: "zh" },
            { label: $t("English"), key: "en" },
        ],
    },
]

// 登录
const handleLogin = () => {
    if (formData.value.email == "") return message.info($t("请填写邮箱"))
    if (!utils.isEmail(formData.value.email)) return message.info($t("请填写正确邮箱"))
    if (formData.value.password == "") return message.info($t("请填写密码"))
    const upData = {
        email: formData.value.email,
        password: formData.value.password,
        code_id: codeId.value,
        code: code.value,
    }
    loadIng.value = true
    userLogin(upData)
        .then(({ data, msg }) => {
            userState.info = data
            message.success(msg)
            router.push("/manage/setting/personal")
        })
        .catch(({ code, msg, data }) => {
            ResultDialog({ code, msg, data })
            if (data.code == "need") {
                onBlur()
            }
        })
        .finally(() => {
            loadIng.value = false
        })
}

// 注册
const handleReg = () => {
    if (formData.value.email == "") return message.info($t("请填写邮箱"))
    if (!utils.isEmail(formData.value.email)) return message.info($t("请填写正确邮箱"))
    if (formData.value.password == "") return message.info($t("请填写密码"))
    if (formData.value.confirmPassword == "") return message.info($t("请再次确认密码"))
    if (formData.value.confirmPassword != formData.value.password) return message.info($t("两次填写的密码不符"))
    loadIng.value = true
    const upData = {
        email: formData.value.email,
        password: formData.value.password,
    }
    userReg(upData)
        .then(({ msg }) => {
            message.success(msg)
        })
        .catch(ResultDialog)
        .finally(() => {
            loadIng.value = false
        })
}

// 变更登录类型
const changeLoginType = () => {
    loginType.value == "login" ? (loginType.value = "reg") : (loginType.value = "login")
    if (loginType.value == "reg") {
        codeNeed.value = false
    } else {
        onBlur()
    }
}

// 设置
const handleSelect = (key: string | number) => {
    switch (key) {
        case "dark":
            globalStore.setThemeName("dark")
            break
        case "light":
            globalStore.setThemeName("light")
            break
        case "en":
            globalStore.setLanguage("en")
            break
        case "zh":
            globalStore.setLanguage("zh")
            break
    }
}

// 切换登录模式
const switchLoginMode = () => {
    if (loginMode.value === "qrcode") {
        loginMode.value = "access"
        onBlur()
        clearInterval(qrcodeTimer.value)
    } else {
        loginMode.value = "qrcode"
        qrcodeVal.value = utils.randomString(32)
        qrcodeTimer.value = setInterval(qrcodeStatus, 2000)
    }
}

// 获取二维码的值



const qrcodeUrl = computed(() => {
    return webTs.apiUrl("../login?qrcode=" + qrcodeVal.value)
})

// 监听二维码状态
const qrcodeStatus = () => {
    const upData = {
        code: qrcodeVal.value,
    }
    getQrcodeStatus(upData).then(({ data, msg }) => {
        if (data) {
            userState.info = data
            message.success(msg)
            router.push("/manage/setting/personal")
        }
    })
}

// 判断要不要验证码
const onBlur = () => {
    const upData = {
        email: formData.value.email,
    }
    needCode(upData)
        .then(({ data }) => {
            codeNeed.value = data
            if (codeNeed.value) {
                refreshCode()
            }
        })
        .catch(ResultDialog)
}

// 刷新验证码
const refreshCode = () => {
    codeImg()
        .then(({ data }) => {
            codeUrl.value = data.image_path
            codeId.value = data.captcha_id
        })
        .catch(() => {
            codeUrl.value = "error"
        })
}
</script>

<style lang="less">
.page-login {
    @apply bg-bg-login flex items-center;

    .login-body {
        @apply flex items-center flex-col max-h-full overflow-hidden py-32 w-full;

        .login-logo {
            @apply block w-84 h-84 bg-logo;
        }

        .login-box {
            @apply bg-bg-login-box rounded-2xl mt-36 w-400 max-w-90p shadow-login-box-Shadow relative;

            .login-mode-switch {
                @apply absolute top-1 right-1 z-10 rounded-lg overflow-hidden;

                .login-mode-switch-box {
                    @apply w-80 h-80 cursor-pointer overflow-hidden bg-primary-color-80;
                    transition: background-color 0.3s;
                    transform: translate(40px, -40px) rotate(45deg);

                    &:hover {
                        @apply bg-primary-color;
                    }

                    .login-mode-switch-icon {
                        @apply absolute text-32 w-50 h-50 bottom-negative-20 left-4 flex items-start justify-start text-white;
                        transform: rotate(-45deg);

                        svg {
                            @apply w-8 h-8 ml-13 mt-3;
                        }
                    }
                }
            }

            .login-title {
                @apply text-24 font-semibold text-center mt-46;
            }

            .login-subtitle {
                @apply text-14 text-text-tips text-center mt-12 px-12;
            }

            .login-qrcode {
                @apply flex items-center justify-center m-auto my-50;
            }

            .login-access {
                @apply mt-30 mx-40 mb-32;

                .n-input {
                    @apply mt-24;
                    transition: all 0s;
                }

                .code-load-input {
                    .n-input-wrapper {
                        @apply pr-0;
                    }

                    .login-code-end {
                        @apply h-38 overflow-hidden cursor-pointer ml-1;

                        .code-load,
                        .code-error {
                            @apply h-full flex items-center justify-center w-5 mx-20;
                        }

                        .code-error {
                            @apply w-auto text-14 opacity-80;
                        }

                        img {
                            @apply h-full min-w-16;
                        }
                    }
                }

                .n-button {
                    @apply mt-24 w-full;
                }

                .login-switch {
                    @apply mt-24 text-text-tips;

                    a {
                        @apply text-primary-color;
                        text-decoration: none;
                    }
                }
            }
        }

        .login-bottom {
            @apply flex items-center justify-between mt-24 w-388;

            .login-setting {
                @apply flex items-center cursor-pointer;
            }

            .login-forgot {
                @apply text-text-tips;

                a {
                    @apply text-primary-color;
                    text-decoration: none;
                }
            }
        }
    }
}

/*登录右侧滑入*/
.login-mode-enter-active {
    transition: all 0.3s ease;
}

.login-mode-leave-active {
    position: absolute;
    z-index: -1;
    display: none;
}

.login-mode-enter,
.login-mode-leave-to {
    transform: translate(100%, 0);
    opacity: 0;
}

input:-webkit-autofill {
    -webkit-box-shadow: 0 0 0px 1000px white inset;
}

.dark input:-webkit-autofill {
    -webkit-box-shadow: 0 0 0px 1000px #2b2b2b inset;
    -webkit-text-fill-color: #ffffff;
}
</style>
