<template>
    <div class="page-manage child-view">
        <manage-menu></manage-menu>
        <div class="manage-box-main">
            <router-view></router-view>
        </div>
    </div>
</template>
<script lang="ts" setup>
import manageMenu from "./components/menu.vue"

</script>
<style lang="less">
.page-manage {
    @apply flex;

    .manage-box-main {
        @apply flex-1 relative;

        .manage-box-view {
            @apply bottom-0 top-0 left-0 right-0 absolute overflow-auto;
        }
    }
}
</style>
