<template >
    <div id="kkkkk">
        <DialogWrapper :dialogId="1111"  is-messenger/>
    </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'

export default defineComponent({
    setup () {

        setTimeout(function () {
            // new Vue({
            //     el: '#kkkkk',
            //     components: {
            //         // 'dialogwrapper' : window.DialogWrapper,
            //     }
            // })

            new window.Vue({
                el: document.querySelector('DialogWrapper'),
                render: h => h(window.DialogWrapper),
            });
        },100)

        return {
            showModal: ref(false)
        }
    }
})


</script>

<style lang="less">

</style>