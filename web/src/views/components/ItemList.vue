<template >
    <n-select v-model:value="value" :placeholder="$t('请选择项目')" filterable :options="itemOptions" :on-focus="openList" />
</template>
<script setup lang="ts">
import { getProjectList } from '@/api/modules/created'

const value = ref('')
const itemOptions = ref([])

const page = ref(1)
const loadIng = ref(false)

const props = defineProps({
    edit: {
        type: Boolean,
        default: false,
    },
})


const openList = () => {
    getItem()
}

const getItem = () => {
    const data = {
        page: page.value,
        page_size: 10,
    }
    loadIng.value = true
    getProjectList(data).then(({ data }) => {
        itemOptions.value = []
        data.map(item => {
            itemOptions.value.push({
                label: item.name,
                value: item.id,
            })
        })
        loadIng.value = false
    })
}

watch(()=>props.edit,(newValue)=>{
    if(newValue){
        getItem()
    }
},{immediate:true} )
</script>
<style lang="less"></style>
