<template>
    <div class="login">
        <div class="wrapper">
            <div class="title">
                {{$t(formData.type === 'reg' ? 'login.reg' : 'login.login')}}
            </div>
            <n-form
                ref="formRef"
                :model="formData"
                :rules="formRules"
                :show-label="false"
                size="large">
                <n-form-item path="email" :label="$t('login.email')">
                    <n-input v-model:value="formData.email" :placeholder="$t('login.enterEmail')"></n-input>
                </n-form-item>
                <n-form-item path="password" :label="$t('login.password')">
                    <n-input v-model:value="formData.password" type="password" :placeholder="$t('login.enterPassword')"></n-input>
                </n-form-item>
                <n-form-item v-if="formData.type === 'reg'" path="password2" :label="$t('login.password2')">
                    <n-input v-model:value="formData.password2" type="password" :placeholder="$t('login.enterPassword2')"></n-input>
                </n-form-item>
                <n-grid :cols="1">
                    <n-grid-item class="buttons">
                        <n-button :loading="loadIng" round type="primary" @click="handleSubmit">
                            {{$t(formData.type === 'reg' ? 'login.reg' : 'login.login')}}
                        </n-button>
                        <n-button :loading="loadIng" round type="default" @click="formData.type=formData.type === 'reg' ? 'login' : 'reg'">
                            {{$t(formData.type === 'reg' ? 'login.login' : 'login.reg')}}
                        </n-button>
                    </n-grid-item>
                </n-grid>
            </n-form>
        </div>
        <div class="policy">{{$t('login.agree')}}</div>
    </div>
</template>

<style lang="less" scoped>
.login {
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;

    .wrapper {
        flex: 1;
        display: flex;
        justify-content: center;
        flex-direction: column;
        width: 90%;
        max-width: 220px;

        .title {
            text-align: center;
            font-size: 24px;
            margin-bottom: 24px;
        }

        .buttons {
            margin-top: 2px;
            display: flex;
            justify-content: center;

            > button + button {
                margin-left: 12px;
            }
        }
    }

    .policy {
        padding: 12px 32px 32px;
    }
}
</style>
<script lang="ts">
import {defineComponent, ref} from "vue";
import {LogoGithub, AddCircleOutline} from "@vicons/ionicons5";
import {FormInst, FormItemRule, FormRules, useMessage} from 'naive-ui'
import {ResultDialog} from "../api";
import utils from "../utils/utils";
import {userLogin, userReg} from "../api/modules/user";
import i18n from "../lang";

export default defineComponent({
    components: {
        LogoGithub,
        AddCircleOutline
    },
    setup() {
        const message = useMessage()
        const loadIng = ref<boolean>(false)
        const formRef = ref<FormInst>()
        const formData = ref({
            type: "login",
            email: "",
            password: "",
            password2: "",
        })
        const formRules: FormRules = {
            email: [
                {
                    validator(rule: FormItemRule, value: string) {
                        if (value) {
                            if (!utils.isEmail(value)) {
                                return new Error(i18n.global.t('login.emailError'))
                            }
                        } else {
                            return new Error(i18n.global.t('login.emailEmpty'))
                        }
                        return true
                    },
                    required: true,
                    trigger: ['input', 'blur']
                }
            ],
            password: [
                {
                    validator(rule: FormItemRule, value: string) {
                        if (value) {
                            if (value.length < 6 || value.length > 20) {
                                return new Error(i18n.global.t('login.passwordLengthError'))
                            }
                        } else {
                            return new Error(i18n.global.t('login.passwordEmpty'))
                        }
                        return true
                    },
                    required: true,
                    trigger: ['input', 'blur']
                }
            ],
            password2: [
                {
                    validator(rule: FormItemRule, value: string) {
                        if (formData.value.type === "reg") {
                            if (value) {
                                if (value !== formData.value.password) {
                                    return new Error(i18n.global.t('login.passwordDiff'))
                                }
                            } else {
                                return new Error(i18n.global.t('login.password2Empty'))
                            }
                        }
                        return true
                    },
                    required: formData.value.type === "reg",
                    trigger: ['input', 'blur']
                }
            ],
        }

        const handleSubmit = (e: MouseEvent) => {
            e.preventDefault()
            formRef.value?.validate((errors) => {
                if (errors) {
                    return;
                }
                //
                if (loadIng.value) {
                    return
                }
                loadIng.value = true;
                (formData.value.type === 'reg' ? userReg : userLogin)(formData.value)
                    .then(({msg}) => {
                        message.success(msg);
                        setTimeout(() => {
                            window.location.href = utils.removeURLParameter(window.location.href, ['result_code', 'result_msg'])
                        }, 300)
                    })
                    .catch(ResultDialog)
                    .finally(() => {
                        loadIng.value = false
                    })
            })
        }
        return {
            i18n,

            loadIng,
            formRef,
            formData,
            formRules,

            handleSubmit,
        }
    }
})
</script>
