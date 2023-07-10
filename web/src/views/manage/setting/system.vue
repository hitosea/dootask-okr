<template>
    <div class="setting-item">
        <n-tabs class="setting-item-tabs" type="line">
            <n-tab-pane name="系统设置" tab="系统设置">
                <system-setting />
            </n-tab-pane>
            <n-tab-pane name="邮件设置" tab="邮件设置">
                <system-email-setting />
            </n-tab-pane>
        </n-tabs>
    </div>
</template>
<script setup lang="ts">
import SystemSetting from "./components/SystemSetting.vue"
import SystemEmailSetting from "./components/SystemEmailSetting.vue"
</script>
<style lang="less" scoped>
.setting-item-tabs {
    @apply pt-16 px-32 flex-1;

    .n-tab-pane {
        @apply flex flex-1 relative;



        .setting-component-item {
            @apply flex flex-1 flex-col absolute top-0 bottom-0 left-0 right-0;

            :deep(.n-form) {
                @apply py-8 px-12 !important;
            }

            :deep(.block-setting-box) {
                @apply relative pt-44 px-24 pb-4 mt-24 mx-0 mb-12 rounded-lg border-solid border border-border-color;
                transition: box-shadow 0.3s;

                &:hover {
                    @apply shadow-manage-box-Shadow;
                }

                h3 {
                    @apply absolute top-0 left-6 py-4 px-10 rounded bg-white border-solid border border-border-color text-15 font-medium inline-block;
                    transform: translateY(-50%);
                }

                .n-form-item-blank {

                    .select-box {
                        @apply w-full;
                    }

                    .form-tip {
                        @apply text-text-tips py-5;
                    }

                    .form-list {
                        ol {
                            li {
                                @apply list-none;
                                counter-reset: list-1 list-2 list-3 list-4 list-5 list-6 list-7 list-8 list-9;
                                counter-increment: list-0;

                                &:before {
                                    @apply w-auto min-w-16;
                                    content: counter(list-0, decimal) ". ";
                                }
                            }
                        }
                    }
                }
            }

            :deep(.setting-footer) {
                @apply -mx-32;
            }
        }
    }
}
</style>
