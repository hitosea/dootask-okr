<template >
    <div>
        我是子应用
    </div>
    <div style="width: 300px;height: 500px; position: relative;">
        <DialogWrappers />
    </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'
import { UserStore } from "@/store/user"

export default defineComponent({
    setup () {
        nextTick(()=>{

            console.log(UserStore().info)

            new window.Vues.Vue({
                el: document.querySelector('DialogWrappers'), 
                store: window.Vues.store,
                render: (h:any) => {
                    return h( window.Vues?.components?.DialogWrapper ,{
                        props:{
                            dialogId: 1
                        }
                    },[ h("div",{slot:"head"}) ])
                },
            });

        })
        return {
            showModal: ref(false)
        }
    }
})


</script>

<style lang="less">

</style>