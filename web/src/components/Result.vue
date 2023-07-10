<template>
    <div class="result">
        <n-result
            :status="status"
            :title="title"
            :description="desc"
            @click="onClick">
            <template #footer>
                <n-button @click.stop="goHome">{{i18n.global.t('result.home')}}</n-button>
            </template>
        </n-result>
    </div>
</template>

<style scoped lang="less">
.result {
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    padding: 0 24px;
    display: flex;
    align-items: center;
    justify-content: center;
}
</style>
<script lang="ts" setup>
import {ref} from 'vue'
import { useDialog } from 'naive-ui'
import utils from '../utils/utils'
import result from "../utils/result";
import i18n from "../lang";
import {CODE} from "../api/constant";
const dialog = useDialog()

let resultCode: any = result.code()
const resultMsg = result.msg()
const resultData = result.data()

if (resultCode === CODE.StatusBadRequest) {
    resultCode = "info"
}

const status = ref(utils.inArray(resultCode, ["500", "error", "info", "success", "warning", "404", "403", "418"]) ? resultCode : "info")
const title = ref(resultMsg.length <= 10 ? resultMsg : '')
const desc = ref(resultMsg.length > 10 ? resultMsg : '')
const goHome = () => {
    window.location.href = "/"
}
const onClick = () => {
    if (resultData === '') {
        return
    }
    dialog.warning({
        title: i18n.global.t('commons.dialog.info'),
        content: resultData,
        positiveText: i18n.global.t('commons.dialog.okText'),
    })
}
</script>
