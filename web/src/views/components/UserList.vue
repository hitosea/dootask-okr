<template >
    <n-select v-model:value="value" :placeholder="$t('请选择项目')" filterable multiple :options="options"
        :render-label="renderLabel" :render-tag="renderMultipleSelectTag" :on-focus="openList" />
</template>
<script setup lang="ts">
import { getUserList } from '@/api/modules/created'
import {
    NAvatar,
    NText,
    NTag,
    SelectRenderTag,
    SelectRenderLabel
} from 'naive-ui'

const value = ref('')
const options = ref([])

const page = ref(1)
const loadIng = ref(false)

const renderLabel: SelectRenderLabel = (option) => {
    return h(
        'div',
        {
            style: {
                display: 'flex',
                alignItems: 'center'
            }
        },
        [
            h(NAvatar, {
                src: option.userimg as string,
                round: true,
                size: 'small'
            }),
            h(
                'div',
                {
                    style: {
                        marginLeft: '12px',
                        padding: '4px 0'
                    }
                },
                [
                    h('div', null, [option.label as string]),
                    h(
                        NText,
                        { depth: 3, tag: 'div' },
                    )
                ]
            )
        ]
    )
}

const renderMultipleSelectTag: SelectRenderTag = ({
    option,
    handleClose
}) => {
    return h(
        NTag,
        {
            style: {
                padding: '0 6px 0 4px'
            },
            round: true,
            closable: true,
            onClose: (e) => {
                e.stopPropagation()
                handleClose()
            }
        },
        {
            default: () =>
                h(
                    'div',
                    {
                        style: {
                            display: 'flex',
                            alignItems: 'center'
                        }
                    },
                    [
                        h(NAvatar, {
                            src: option.userimg as string,
                            round: true,
                            size: 22,
                            style: {
                                marginRight: '4px'
                            }
                        }),
                        option.label as string
                    ]
                )
        }
    )
}

const openList = () => {
    getItem()
}

const getItem = () => {
    const data = {
        page: page.value,
        page_size: 10,
    }
    loadIng.value = true
    getUserList(data).then(({ data }) => {
        options.value = []
        data.map(item => {
            options.value.push({
                label: item.nickname,
                value: item.userid,
                userimg:item.userimg,
            })
        })
        loadIng.value = false
    })
}


</script>
<style lang="less"></style>