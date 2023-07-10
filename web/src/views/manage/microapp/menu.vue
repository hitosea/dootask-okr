<template>
    <div class="page-manage child-view">
        <div class="manage-box-main">
            <n-spin size="small" :show="showSpin" class="h-full">
                <div v-if="errorText" class="flex items-center justify-center h-screen">
                    <p class="text-center text-text-tips">{{errorText}}</p>
                </div>
                <micro-app v-if="microAppUrl" class="h-full"
                    :name="microAppName" 
                    :url="microAppUrl" 
                    baseRoute="/" 
                    keep-alive 
                    inline 
                    disableSandbox 
                    :data='microAppData'
                    @created='handleCreate'
                    @beforemount='handleBeforeMount'
                    @mounted='handleMount'
                    @unmount='handleUnmount'
                    @error='handleError'
                    @datachange='handleDataChange'
                ></micro-app>
                <template #description>
                    {{ $t("玩命加载中") }}
                </template>
            </n-spin>
        </div>
    </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue';
import { useRoute } from 'vue-router';
import { GlobalStore } from "@/store"
import { EventCenterForMicroApp } from '@micro-zoe/micro-app'

const globalStore = GlobalStore()
const route = useRoute();
const getPluginMenu = computed(() => globalStore.pluginConfig.menu || [])
const microAppUrl = ref("");
const microAppName = ref("");
const microAppData = ref({msg: '来自基座的数据'});
const showSpin = ref(false);
const errorText = ref("");

// 监听路由变更，更新子应用地址
watch(() => route.path,(newPath, oldPath) => { 
    microAppUrl.value = ""
    const match = newPath.match(/\d+/);            
    const menuId = match ? parseInt(match[0]) : 0;     
    getPluginMenu.value.forEach((item:any) => {
        if( item.id == menuId){
            microAppName.value = item.code+"-" +item.id
            microAppUrl.value = item.path
            // @ts-ignore
            window.eventCenterForAppNameVite = new EventCenterForMicroApp(microAppName.value)
        }
    });
},{ immediate: true });

// 子应用创建了
const handleCreate = (e:any) => {
    console.log("子应用创建了",e)
}

// 子应用即将被渲染
const handleBeforeMount = (e:any) => {
    showSpin.value = true
    errorText.value = ""
    console.log("子应用即将被渲染",e)
}

// 子应用已经渲染完成
const handleMount = (e:any) => {
    showSpin.value = false
    console.log("子应用已经渲染完成",e)
    setTimeout(() => {
        // @ts-ignore
        microAppData.value = {msg: '来自基座的新数据'}
    }, 2000)
}

// 子应用卸载了
const handleUnmount = (e:any) => {
    console.log("子应用卸载了",e)
}

// 子应用加载出错了
const handleError = (e:any) => {
    showSpin.value = false
    errorText.value = $t('子应用加载出错了') +" : " + e.detail.error
}

// 监听子应用数据变更
const handleDataChange = (e:any) => {
    console.log('来自子应用 child-vite 的数据:', e.detail.data)
}
</script>
