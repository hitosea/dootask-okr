<template >
    <div class="okr-item-main">
        <div class="okr-item-box">
            <n-progress style="width: 52px;" color="var(--primary-color)" indicator-text-color="var(--primary-color)"
                type="circle" :percentage="80" :offset-degree="180">
                <p class="text-primary-color text-14">80<span class="text-12">%</span></p>
            </n-progress>
            <div class="okr-list">
                <div class="okr-title">
                    <div class="okr-title-l">
                        <span>P0</span>
                        <h3>完善团队制度，为了更好的服务客户</h3>
                    </div>
                    <div class="okr-title-r">
                        <n-icon class="okr-title-star" size="16" :component="Star20Filled" />
                        <i class="taskfont">&#xe648;</i>
                        <p>1/3</p>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>
<script setup lang="ts">
import { Star20Filled, Star20Regular } from '@vicons/fluent'
</script>
<style lang="less">
.okr-item-main {
    @apply flex flex-col gap-6;

    .okr-item-box {
        @apply px-24 py-32 bg-white rounded-lg flex gap-4;

        .okr-list {
            @apply flex flex-col flex-1;

            .okr-title {
                @apply flex justify-between items-center;

                .okr-title-l {
                    @apply flex items-center gap-2;

                    span {
                        @apply bg-[#FF7070] text-10 text-white px-6 py-2 rounded-full scale-[0.8333] origin-center flex items-center leading-3;
                    }

                    h3 {
                        @apply text-title-color text-14 font-normal;
                    }
                }

                .okr-title-r {
                    @apply flex items-center;

                    .okr-title-star {
                        @apply text-[#FFD023] pr-16;
                    }
                    p{
                        @apply text-text-tips;
                    }
                }
            }
        }
    }
}</style>