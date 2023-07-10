<template ></template>
<script setup lang="ts">
import { useRouter } from "vue-router"
import { UserStore } from "@/store/user"

const userState = UserStore()
const router = useRouter()

const getNeedStartHome = () => {
    if (userState.info.id > 0) {
        router.push("/manage/dashboard")
    } else {
        router.push("/login")
    }
}
getNeedStartHome()

</script>
<style lang="less">

</style>