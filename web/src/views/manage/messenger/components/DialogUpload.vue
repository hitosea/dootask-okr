<template>
    <n-upload name="files" ref="upload" :action="actionUrl" :headers="headers" :data="params" multiple
        :show-file-list="false" @on-finish="handleSuccess" :on-before-upload="beforeUpload">
    </n-upload>
</template>

<script setup lang="ts">
import { defineProps, defineEmits } from "vue"
import { DiaLogStore } from '@/store/dialog';
import { UserStore } from "@/store/user";
import webTs from '@/utils/web';
import utils from "@/utils/utils";
import { UploadFileInfo, useNotification } from "naive-ui"

const cacheDialogs = DiaLogStore().cacheDialogs
const actionUrl = ref(webTs.apiUrl('../api/v1/system/upload/file'))
const userToken = UserStore().info.token
const upload = ref(null)
const uploadFormat = ref([])
const notification = useNotification()

const emit = defineEmits(['on-progress', 'on-success', 'on-error'])

const props = defineProps({
    dialogId: {
        type: Number,
        default: 0
    },
    maxSize: {
        type: Number,
        default: 1024000
    }
})


const headers = computed(() => {
    return {
        fd: utils.getSessionStorageString("userWsFd"),
        token: userToken,
    }
})

const params = computed(() => {
    return {
        dialog_id: props.dialogId,
        reply_id: dialogData.value.extra_quote_id || 0,
    }
})

const dialogData = computed(() => {
    return cacheDialogs.find(({ id }) => id == props.dialogId) || {};
})


// naive ui 没有上传时得用上传之前代替
// const handleProgress = (event, file) => {
//     //上传时
//     if (file.tempId === undefined) {
//         if (getCurrentInstance().type.name === 'DialogWrapper') {
//             file.tempId = getCurrentInstance().getTempId()
//         } else {
//             file.tempId = utils.randNum(1000000000, 9999999999)
//         }
//         emit('on-progress', file)
//     }
// }

const beforeUpload = (file) => {
    //上传之前
    if (file.tempId === undefined) {
        if (getCurrentInstance().type.name === 'DialogWrapper') {
            file.tempId = utils.randNum(1000000000, 9999999999)
        } else {
            file.tempId = utils.randNum(1000000000, 9999999999)
        }
        emit('on-progress', file)
    }

}



const handleSuccess = (file: UploadFileInfo, event: ProgressEvent) => {
    //上传完成
    // if (event.ret === 1) {
    //     file.data = event.data;
    //     emit('on-success', file);
    //     if (event.target.task_id) {
    //         // this.$store.dispatch("getTaskFiles", res.data.task_id) 重写接口
    //     }
    // } else {
    //     notification['warning']({
    //         content: $t('发送失败'),
    //         meta: '文件 ' + file.name + ' 发送失败，' + event.msg,
    //         duration: 2500,
    //         keepAliveOnHover: true
    //     })
    //     emit('on-error', file);
    //     upload.value.fileList.pop();
    // }
    if(file.type){
        handleFormatError
    }
    if(file.file.size){
        handleMaxSize
    }
}

const handleFormatError = (file) => {
    //上传类型错误
    notification['warning']({
        content: $t('文件格式不正确'),
        meta: $t('文件') + file.name + $t(' 格式不正确，仅支持发送：') + uploadFormat.value.join(','),
        duration: 2500,
        keepAliveOnHover: true
    })
}

const handleMaxSize = (file) => {
    //上传大小错误
    notification['warning']({
        content: $t('超出文件大小限制'),
        meta: $t('文件') + file.name + $t(' 太大，不能发送超过') + utils.bytesToSize(props.maxSize * 1024) + '。',
        duration: 2500,
        keepAliveOnHover: true
    })
}

const handleClick = () => {
    //手动上传
    upload.value.openOpenFileDialog()
}

const upLoad = (file) => {
    //手动传file
    upload.value.submit(file);
}


</script>
