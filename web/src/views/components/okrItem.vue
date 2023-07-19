<template >
    <div class="okr-item-main">
        <div class="okr-item-box">
            <n-progress style="width: 52px;" color="var(--primary-color)" indicator-text-color="var(--primary-color)"
                type="circle" :percentage="80" :offset-degree="180" :stroke-width="8">
                <p class="text-primary-color text-14">80<span class="text-12">%</span></p>
            </n-progress>
            <div class="okr-list">
                <div class="okr-title">
                    <div class="okr-title-l">
                        <span class="scale-[0.8333]">P0</span>
                        <h3>完善团队制度，为了更好的服务客户</h3>
                    </div>
                    <div class="okr-title-r">
                        <i class="taskfont okr-title-star">&#xe683;</i>
                        <i class="taskfont okr-title-icon">&#xe671;</i>
                        <p>1/3</p>
                    </div>
                </div>
                <div class="okr-time">
                    <i class="taskfont"> &#xe61a;</i>
                    <p>张三</p>
                    <div class="w-1 bg-[#F2F3F5] mx-12 h-full"></div>
                    <i class="taskfont"> &#xe71d;</i>
                    <p>6h23</p>
                </div>
                <div class="kr-list">
                    <div class="kr-list-item">
                        <span class="bg-[rgba(135,208,104,0.2);] scale-[0.8333]">KR1</span>
                        <p>完善团队制度，为了更好的服务客户，搭建一个团队使用的客户成功系统</p>
                        <div class="kr-list-schedule">
                            <n-progress class="-mt-6 mr-[6px]" style="width: 15px; " type="circle" :show-indicator="false"
                                :offset-degree="180" :stroke-width="15" color="var(--primary-color)" status="success"
                                :percentage="20" />
                            20%
                        </div>
                    </div>

                </div>
                <div class="align-target" @click="handleTarget(1)">
                    对齐目标（5）
                </div>
                <div class="align-target" @click="handleTarget(2)">
                    向上对齐
                </div>
            </div>
        </div>
    </div>
    <AlignTarget :value="alignTargetShow" @close="() => { alignTargetShow = false }"></AlignTarget>
    <SelectAlignment></SelectAlignment>
</template>
<script setup lang="ts">
import AlignTarget from '@/views/components/AlignTarget.vue';
import SelectAlignment from '@/views/components/SelectAlignment.vue'

const alignTargetShow = ref(false)

const handleTarget = (e) => {
    if(e==1){
        alignTargetShow.value = true
    }else{

    }

}

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
                        @apply bg-[#FF7070] text-10 text-white px-6 py-2 rounded-full origin-center flex items-center leading-3;
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

                    .okr-title-icon {
                        @apply text-text-tips mr-4;
                    }

                    p {
                        @apply text-text-tips;
                    }
                }
            }

            .okr-time {
                @apply flex items-center mt-12 text-text-tips;

                i {
                    @apply mr-4;
                }
            }

            .kr-list {
                @apply mt-16 flex flex-col gap-5;

                .kr-list-item {
                    @apply flex items-center;

                    span {
                        @apply text-primary-color text-10 origin-center py-2 px-6 rounded-full flex items-center leading-3;
                    }

                    p {
                        @apply text-title-color text-14 ml-16 truncate;
                    }

                    .kr-list-schedule {
                        @apply ml-auto flex items-center;
                    }
                }
            }

            .align-target {
                @apply mt-16 text-text-tips text-12 cursor-pointer;
            }
        }
    }
}
</style>