<template >
    <n-upload :action="upUrl" :default-file-list="previewFileList" :list-type="props.listType" @preview="handlePreview"
        :headers="uploadHeaders" @before-upload="beforeUpload" @finish="handleFinish" @error="handleError" method="POST"
        :max="props.maxNumber" :multiple="props.multiple" />
    <n-modal v-model:show="showModal" preset="card" style="width: 600px" :title="$t('查看')">
        <img :src="previewImageUrl" style="width: 100%" />
    </n-modal>
</template>
<script setup lang="ts">
import { ref, computed } from "vue"
import { UploadFileInfo, useMessage } from "naive-ui"
import webTs from "@/utils/web";
import utils from "@/utils/utils";
import { UserStore } from "@/store/user"

const userInfo = UserStore()
const previewImageUrl = ref("")
const showModal = ref(false)
const message = useMessage()

const props = defineProps({
    maxNumber: {
        type: Number,
        default: 3,
    },
    multiple: {
        type: Boolean,
        default: false,
    },
    value: {
    },
    listType: {
        type: String,
        default: 'image-card',
    }
})

//action上传api
const upUrl = computed(() => {
    return webTs.apiUrl("../api/v1/system/upload/image")
})

// 头部参数
const uploadHeaders = computed<Record<string, string>>(() => {
    return {
        fd: utils.getSessionStorageString("userWsFd").toString(),
        token: userInfo.info.token,
    }
})

//上传前判断
const beforeUpload = async (data: {
    file: UploadFileInfo
    fileList: UploadFileInfo[]
}) => {
    if (data.file.file?.type !== 'image/png' && data.file.file?.type !== 'image/jpg' && data.file.file?.type !== 'image/gif' && data.file.file?.type !== 'image/jpeg') {
        message.error($t('只能上传图片文件，请重新上传'))
        return false
    }
    return true
}

const handleFinish = ( options: { file: UploadFileInfo, event?: ProgressEvent }) =>{
    const res = JSON.parse((options.event?.target as XMLHttpRequest).response)
    if (res.code == '200') {
        options.file.url = res.data.url
    }
    return options.file
}



//上传错误
const handleError = () => {
    message.error($t('上传失败！'))
}

const previewFileList = ref<UploadFileInfo[]>([

])

const handlePreview = (file: UploadFileInfo) => {
    const { url } = file
    previewImageUrl.value = url as string
    showModal.value = true
}
</script>
<style lang="less"></style>