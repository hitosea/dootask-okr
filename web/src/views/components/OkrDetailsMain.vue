<template >
    <div class="flex flex-col h-full md:h-auto md:flex-row md:max-h-[640px] md:min-h-[640px]">
        <div class="md:flex-1 flex flex-col relative md:overflow-hidden bg-white px-16 pt-16 md:pt-0 md:px-0"
            :class="navActive == 0 ? 'navActive' : ''">
            <div
                class="hidden md:flex min-h-[36px] items-center justify-between pb-[15px] border-solid border-0 border-b-[1px] border-[#F2F3F5] relative md:mr-24">
                <div class="flex items-center gap-4">
                    <n-popover v-if="detailData.completed == '0'" placement="bottom" :show="showPopover" trigger="manual" @clickoutside="showPopover = false">
                        <template #trigger>
                            <div @click="showPopover = !showPopover" v-if="detailData.completed == '0'"
                                class="flex items-center justify-center w-[16px] h-[16px] overflow-hidden rounded-full border-[1px] border-solid cursor-pointer"
                                :class="detailData.completed == '0' || detailData.canceled == '0' ? 'border-[#A8ACB6]' : 'border-primary-color bg-primary-color'">
                            </div>
                        </template>
                        <span class="cursor-pointer" @click="handleCancel">
                            {{ detailData.canceled == '0' ? $t('取消目标') : $t('重启目标') }}
                        </span>
                    </n-popover>
                    <div v-if="detailData.completed == '1'"
                        class="flex items-center justify-center w-[16px] h-[16px] overflow-hidden rounded-full border-[1px] border-solid "
                        :class="detailData.completed == '0' ? 'border-[#A8ACB6]' : 'border-primary-color bg-primary-color'">
                        <n-icon v-if="detailData.completed == '1'"
                            :class="detailData.completed == '0' ? 'text-[#A8ACB6]' : ' text-white'" size="14"
                            :component="CheckmarkSharp" />
                    </div>

                    <n-tag v-if="detailData.canceled == '1'">{{ $t('已取消') }} </n-tag>
                    <n-tag v-if="detailData.completed == '1'" type="success">{{ $t('已完成') }}</n-tag>
                    <p class="flex items-center text-text-li opacity-50 text-14">
                        {{ detailData.type == "1" ? $t('承诺型') : $t('挑战性') }} /
                        {{ detailData.ascription == "2" ? $t('个人') : $t('部门') }}
                    </p>
                </div>
                <div class="flex items-center gap-6">
                    <i v-if="detailData.canceled == '0' && detailData.completed == '0' && userInfo.userid == detailData.userid"
                        class="okrfont icon-title text-[#A7ACB6]" @click="handleEdit">&#xe779;</i>

                    <i class="okrfont text-[#FFD023] cursor-pointer text-16" v-if="detailData.is_follow"
                        @click.stop="handleFollowOkr">&#xe683;</i>
                    <i class="okrfont text-[#A7ACB6] cursor-pointer text-16" v-else @click.stop="handleFollowOkr">&#xe679;</i>

                    <div v-if="detailData.score > -1" class="flex items-center cursor-pointer">
                        <img class="mr-8" :src="utils.apiUrl(fenSvg)" />
                        <p class=" text-title-color text-12">{{ detailData.score }}{{ $t('分') }}
                        </p>
                    </div>
                </div>
                <img v-if="detailData.completed == '1'" class="absolute right-24 -bottom-[50px] "
                    src="@/assets/images/icon/complete.png" />
            </div>

            <n-scrollbar ref="scrollbarRef" class="left-scrollbar">
                <div class="md:mr-24">
                    <h3 id="detailTop"
                        class="relative text-text-li md:mt-[24px] break-all text-18 md:text-24 leading-[1.4] font-medium md:min-h-[40px]">
                        {{ detailData.title }}
                        <img v-if="detailData.completed == '1'" class="absolute right-24 top-0 block md:hidden"
                            src="@/assets/images/icon/complete.png" />
                    </h3>
                    <div class="mt-16 md:mt-24 flex flex-col gap-4">
                        <div class="flex items-center">
                            <p class="flex items-center w-[115px]">
                                <i v-if="detailData.ascription == '2'"
                                    class="okrfont icon-item text-[#BBBBBB]">&#xe6e4;</i>
                                <i v-else class="okrfont icon-item text-[#BBBBBB]">&#xe682;</i>

                                <span class="text-[#BBBBBB] text-[14px]"> {{ detailData.ascription == "2" ?
                                    $t('负责人') : $t('部门') }}</span>
                            </p>
                            <p class="flex-1 text-text-li text-14">
                                {{ detailData.alias && detailData.alias.join(',') }}
                            </p>
                        </div>


                        <div class="flex md:items-center">
                            <p class="flex md:items-center w-[115px]">
                                <i class="okrfont icon-item text-[#BBBBBB] -mt-2">&#xe6e8;</i>
                                <span class="text-[#BBBBBB] text-[14px]">{{ $t('起止时间') }}</span>
                            </p>
                            <p class="flex-1 text-text-li text-14 flex  md:items-center flex-col md:flex-row">
                                <span v-if="detailData.start_at">{{ utils.GoDate(detailData.start_at || 0) }} ~ {{
                                    utils.GoDate(detailData.end_at || 0) }}</span>
                                <div class="" v-if="detailData.completed == '0' && detailData.canceled == '0' && detailData.end_at">
                                    <n-tag class="md:ml-4 mt-4 md:mt-0" v-if="within24Hours(detailData.end_at)" type="info"><i
                                            class="okrfont text-14 mr-4">&#xe71d;</i>{{ expiresFormat(detailData.end_at) }}</n-tag>
                                    <n-tag class="ml-4" v-if="isOverdue(detailData)" type="error">{{ $t('超期未完成') }}</n-tag>
                                </div>
                            </p>

                        </div>

                        <div class="flex items-center">
                            <p class="flex items-center w-[115px]">
                                <i class="okrfont icon-item text-[#BBBBBB]">&#xe6ec;</i>
                                <span class="text-[#BBBBBB] text-[14px]">{{ $t('优先级') }}</span>
                            </p>
                            <span class="span h-[16px]" :class="pStatus(detailData.priority)">{{ detailData.priority
                            }}</span>
                        </div>


                        <div class="flex md:hidden items-center" v-if="detailData.score > -1">
                            <p class="flex items-center w-[115px]">
                                <img class="mr-6" :src="utils.apiUrl(fenSvg)" />
                                <span class="text-[#BBBBBB] text-[14px] opacity-50">{{ $t('评分') }}</span>
                            </p>
                            <p class="flex-1 text-text-li text-14">
                                {{ detailData.score }}{{ $t('分') }}
                            </p>
                        </div>
                    </div>

                    <h4 class="hidden md:block text-text-li text-16 font-medium mt-36">KR</h4>

                    <div class="hidden md:flex flex-col mt-20 gap-4 min-h-[60px]">
                        <div class="flex flex-col" v-for="(item, index) in detailData.key_results">
                            <div class="flex items-start">
                                <span class=" text-primary-color text-12 leading-[20px] mr-8 shrink-0">KR{{ index + 1
                                }}</span>
                                <h4 class=" text-text-li text-14 font-normal leading-[20px]">{{ item.title }}</h4>
                            </div>
                            <div class="flex items-center justify-between mt-4">
                                <div class="flex items-center mr-24">
                                    <div class="flex items-start gap-2 max-w-[104px] h-[26px] overflow-hidden">
                                        <div v-if="showUserSelect  && detailData.completed == '0' && detailData.canceled == '0' || item.participant!=''" class="relative">
                                            <div v-if="userInfo.userid != detailData.userid || detailData.canceled =='1'|| detailData.completed =='1' || item.score > -1 || item.superior_score > -1" class="absolute top-0 bottom-0 left-0 right-0 cursor-not-allowed z-[2]"></div>
                                            <UserSelects class=" relative z-[1]" :formkey="index" />
                                        </div>
                                        <n-avatar v-if="!showUserSelect" round :size="20"
                                            src="https://07akioni.oss-cn-beijing.aliyuncs.com/07akioni.jpeg" />
                                    </div>
                                    <div v-if="item.participant.split(',').length > 4"
                                        class="w-[20px] h-[20px] rounded-full bg-primary-color flex items-center justify-center text-white text-12 ">
                                        <span class="scale-[0.8333] origin-center whitespace-nowrap">+{{
                                            item.participant.split(',').length - 4 }}</span>
                                    </div>
                                    <template v-if="!isOverdue(item) || detailData.completed == '1' || detailData.canceled == '1'">
                                        <i class="okrfont mr-4 text-12 ml-24 text-[#A7ABB5]">&#xe6e8;</i>
                                        <p class="flex-1 text-text-li text-12 min-w-[140px] opacity-50 shrink-0 leading-[18px]">{{
                                            utils.GoDate(item.start_at || 0) }} ~{{ utils.GoDate(item.end_at || 0) }}
                                        </p>
                                    </template>
                                    <template v-if="isOverdue(item) && detailData.completed == '0' && detailData.canceled == '0'">
                                        <i class="okrfont mr-4 text-12 ml-24 " :class="isOverdue(item) ?'text-[#ED4014]' : ''"> &#xe6e8;</i>
                                       <p :class="isOverdue(item) ?'text-[#ED4014] text-12' : ''">{{ expiresFormat(item.end_at) }}</p>
                                    </template>
                                </div>
                                <div class="flex items-center  shrink-0 min-w-[200px]">
                                    <div class="flex items-center cursor-pointer min-w-[55px] justify-start flex-1"
                                        @click="handleSchedule(item.id, item.progress, item.progress_status, item.score)">
                                        <n-progress class="-mt-10 mr-[6px]" style="width: 15px; " type="circle"
                                            :show-indicator="false" :offset-degree="180" :stroke-width="15"
                                            :color="colorStatus(item.progress_status)" status="success"
                                            :percentage="item.progress" />
                                        <p class="text-text-li opacity-50 text-12">{{ item.progress }}%</p>
                                    </div>
                                    <div v-if="item.confidence == '0'" class="flex items-center cursor-pointer min-w-[55px] justify-start flex-1"
                                        @click="handleConfidence(item.id, item.confidence, item.score)">
                                        <i class="okrfont mr-6 text-16 text-[#A7ABB5]">&#xe67c;</i>
                                        <p class="text-text-li opacity-50 text-12">{{ $t('信心') }}</p>
                                    </div>
                                    <div v-else class="flex items-center cursor-pointer min-w-[55px] justify-start flex-1"
                                        @click="handleConfidence(item.id, item.confidence, item.score)">
                                        <i class="okrfont mr-6 text-16 text-[#FFA25A]">&#xe674;</i>
                                        <p class="text-text-li opacity-50 text-12">{{ item.confidence }}</p>
                                    </div>
                                    <div v-if="item.kr_score == '0'" class="flex items-center cursor-pointer min-w-[55px] justify-start flex-1"
                                        @click="handleMark(item.id, item.score, item.superior_score, item.progress)">
                                        <i class="okrfont mr-6 text-16 text-[#A7ABB5]">&#xe67d;</i>
                                        <p class="text-text-li opacity-50 text-12">{{ $t('评分') }}</p>
                                    </div>
                                    <div v-else class="flex items-center cursor-pointer min-w-[55px] justify-start flex-1"
                                        @click="handleMark(item.id, item.score, item.superior_score, item.progress)">
                                        <img class="mr-6 -mt-2" :src="utils.apiUrl(fenSvg)" />
                                        <p class="text-text-li opacity-50 text-12">{{ item.kr_score }}{{ $t('分') }}
                                        </p>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>

                    <div class="hidden md:block border-solid border-0 border-b-[1px] border-[#F2F3F5] my-36"></div>

                    <h4 class="hidden md:block text-text-li text-16 font-medium mb-12">
                        {{ $t('对齐目标') }}
                        <i v-if="detailData.canceled == '0' && detailData.completed == '0' && userInfo.userid == detailData.userid"
                            class="okrfont text-16 cursor-pointer text-[#A7ABB5]"
                            @click="() => { selectAlignmentShow = true }">&#xe779;</i>
                    </h4>

                    <div class="pb-[28px] w-full overflow-hidden hidden md:block">
                        <AlignTarget ref="AlignTargetRef" :value="props.show" :id="props.id" :progressShow="true"
                            :userid="detailData.userid"
                            :cancelShow="detailData.canceled == '0' && detailData.completed == '0' && userInfo.userid == detailData.userid"
                            @unalign="handleUnalign" @openDetail="openDetail">
                        </AlignTarget>
                    </div>

                </div>
            </n-scrollbar>
        </div>

        <div
            class="md:min-w-[35.8%] relative flex flex-col flex-1 md:flex-initial border-solid border-0 md:border-l-[2px] border-[#F2F3F5] ">
            <div class="flex items-center justify-between border-solid border-0 border-b-[1px] border-[#F2F3F5] pb-[11px] md:pb-[15px] md:ml-24 min-h-[36px] bg-white"
                :class="navActive == 0 ? 'pt-14 md:pt-0' : 'pt-[32px] md:pt-0'">
                <ul class="flex w-full items-center gap-8 justify-between md:justify-start px-16 md:px-0">
                    <li class="block md:hidden li-nav" :class="navActive == 3 ? 'active' : ''" @click="handleNav(3)">KR</li>
                    <li class="block md:hidden li-nav" :class="navActive == 4 ? 'active' : ''" @click="handleNav(4)">{{
                        $t('对齐') }}
                    </li>
                    <li class="li-nav" :class="navActive == 0 ? 'active' : ''" @click="handleNav(0)">{{ $t('评论') }}
                    </li>
                    <li class="li-nav" :class="navActive == 1 ? 'active' : ''" @click="handleNav(1)">{{ $t('动态') }}
                    </li>
                    <li class="li-nav" :class="navActive == 2 ? 'active' : ''" @click="handleNav(2)">{{ $t('复盘') }}
                    </li>
                </ul>
                <i class="okrfont text-16 cursor-pointer text-[#A7ABB5] hidden md:block n-close" @click="closeModal">&#xe6e5;</i>
            </div>
            <div class="flex-auto relative">
                <div class="md:absolute md:top-[24px] md:bottom-0 md:left-0 md:right-0">
                    <div class="text-center  mt-16 md:mt-0 px-16 md:px-0" v-if="navActive == 3">
                        <div class="flex md:hidden flex-col gap-3 min-h-[60px]">
                            <div class="flex flex-col bg-[#fff] px-16 pt-24 rounded-lg"
                                v-for="(item, index) in detailData.key_results">
                                <div class="flex items-center">
                                    <h4 class=" text-text-li text-14 font-normal text-left break-all">{{ item.title
                                    }}</h4>
                                </div>
                                <div class="flex flex-col justify-between mt-12">
                                    <div class="flex items-center justify-between">
                                        <div  class="flex items-center">
                                            <div class="flex items-start gap-2 max-w-[104px] h-[26px] overflow-hidden">
                                                <div v-if="showUserSelect && detailData.completed == '0' && detailData.canceled == '0'" class="relative">
                                                    <div v-if="userInfo.userid != detailData.userid || detailData.canceled =='1'|| detailData.completed =='1' || item.score > -1 || item.superior_score > -1" class="absolute top-0 bottom-0 left-0 right-0 cursor-not-allowed z-[2]"></div>
                                                    <UserSelects class=" relative z-[1]" :formkey="index" />
                                                </div>
                                                <n-avatar v-if="!showUserSelect" round :size="20"
                                                    src="https://07akioni.oss-cn-beijing.aliyuncs.com/07akioni.jpeg" />
                                            </div>
                                            <div v-if="item.participant.split(',').length > 4"
                                                class="w-[20px] h-[20px] rounded-full bg-primary-color flex items-center justify-center text-white text-12 " :class="item.participant.split(',').length == 4 ? 'ml-6':''">
                                                <span class="scale-[0.8333] origin-center whitespace-nowrap">+{{
                                                    item.participant.split(',').length - 4 }}</span>
                                            </div>
                                         </div>
                                        <div class="flex items-center">
                                            <i class="okrfont text-12 mr-4 text-[#A7ABB5]">&#xe6e8;</i>
                                            <p class="flex-1 text-text-tips text-12  shrink-0">{{ utils.GoDate(item.end_at
                                                || 0) }}
                                            </p>
                                        </div>
                                    </div>
                                    <div
                                        class="flex mt-16 py-8 shrink-0 border-solid border-0 border-t-[1px] border-[#F2F3F5]">
                                        <div class="flex items-center justify-center cursor-pointer flex-1 border-solid border-0 border-r-[1px] border-[#F2F3F5]"
                                            @click="handleSchedule(item.id, item.progress, item.progress_status, item.score)">
                                            <n-progress class="-mt-7 mr-[6px]" style="width: 16px; " type="circle"
                                                :show-indicator="false" :offset-degree="180" :stroke-width="15"
                                                :color="colorStatus(item.progress_status)" status="success"
                                                :percentage="item.progress" />
                                            <p class="text-text-li opacity-50 text-12">{{ item.progress }}%</p>
                                        </div>
                                        <div v-if="item.confidence == '0'"
                                            class="flex flex-1 items-center justify-center cursor-pointer border-solid border-0 border-r-[1px] border-[#F2F3F5]"
                                            @click="handleConfidence(item.id, item.confidence, item.score)">
                                            <i class="okrfont mr-6 text-16 text-[#A7ABB5]">&#xe67c;</i>
                                            <p class="text-text-li opacity-50 text-12">{{ $t('信心') }}</p>
                                        </div>
                                        <div v-else
                                            class="flex flex-1 items-center justify-center cursor-pointer border-solid border-0 border-r-[1px] border-[#F2F3F5]"
                                            @click="handleConfidence(item.id, item.confidence, item.score)">
                                            <i class="okrfont mr-6 text-16 text-[#FFA25A]">&#xe674;</i>
                                            <p class="text-text-li opacity-50 text-12">{{ item.confidence }}</p>
                                        </div>
                                        <div v-if="item.kr_score == '0'"
                                            class="flex flex-1 items-center justify-center cursor-pointer"
                                            @click="handleMark(item.id, item.score, item.superior_score, item.progress)">
                                            <i class="okrfont mr-6 text-16 text-[#A7ABB5]">&#xe67d;</i>
                                            <p class="text-text-li opacity-50 text-12">{{ $t('评分') }}</p>
                                        </div>
                                        <div v-else class="flex flex-1 items-center justify-center cursor-pointer"
                                            @click="handleMark(item.id, item.score, item.superior_score, item.progress)">
                                            <img class="mr-6 -mt-2" :src="utils.apiUrl(fenSvg)" />
                                            <p class="text-text-li opacity-50 text-12">{{ item.kr_score }}{{ $t('分') }}
                                            </p>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div class="overflow-hidden mt-24 md:mt-0 px-16 md:px-0 " v-if="navActive == 4">
                        <AlignTarget ref="AlignTargetRef" :value="props.show" :id="props.id" :progressShow="true"
                            :cancelShow="detailData.canceled == '0' && detailData.completed == '0' && userInfo.userid == detailData.userid"
                            @unalign="handleUnalign" @openDetail="openDetail">
                        </AlignTarget>
                    </div>
                    <div class="text-center flex-1" v-show="navActive == 0">
                        <!-- <n-spin size="small" class="absolute top-0 bottom-0 left-0 right-0"> </n-spin> -->
                        <span v-if="!showDialogWrapper">{{ $t('子应用无法加载') }}</span>
                        <div v-else class=" absolute top-0 md:-top-[20px] -bottom-[16px] left-0 right-0">
                            <DialogWrappers />
                        </div>
                    </div>
                    <n-scrollbar class="mt-16 md:mt-0 px-16 md:px-0" v-if="navActive == 1" :on-scroll="onScrollLogList">
                        <div class="flex text-start mb-[24px] md:pl-24 pr-[10px] " v-for="item in logList"
                            v-if="logList.length">
                            <n-avatar round :size="28" class="mr-8 shrink-0" :src="(item.user_avatar || '').replace(':///', globalStore.baseUrl) " />
                            <div class="flex flex-col gap-3">
                                <p class="text-14 leading-[16px] text-primary-color">{{ item.user_nickname }}<span
                                        class="text-12 text-text-li opacity-60 ml-8">{{ utils.GoDateHMS(item.created_at)
                                        }}</span></p>
                                <h4 class="text-14 leading-[18px] text-title-color font-normal"> <span
                                        class=" font-normal">{{ item.content }}</span></h4>
                            </div>
                        </div>
                        <p v-else class="text-12 mt-20 text-text-tips text-center ">{{ $t('暂无动态') }}</p>
                        <p @click="nextLogList" v-if="logListLastPage > logListPage"
                            class="text-12 mt-20 mb-20 text-text-tips text-center block md:hidden">{{ $t('点击加载更多') }}</p>
                    </n-scrollbar>
                    <n-scrollbar class="mt-16 md:mt-0 px-16 md:px-0" v-if="navActive == 2" >
                        <div class="md:pl-24 pr-[10px]">
                            <p class="cursor-pointer mb-20" v-if="userInfo.userid == detailData.userid"
                                :class="detailData.score < 0 ? 'text-text-tips opacity-50' : 'text-primary-color'"
                                @click="handleAddMultiple"> <i class="okrfont mr-4 text-16 ">&#xe6f2;</i><span
                                    class="text-14">{{ $t('添加复盘') }}</span></p>
                            <div class="flex flex-col gap-3" v-if="replayList.length">
                                <div class="flex items-center justify-between cursor-pointer" v-for="item in replayList"
                                    @click="handleCheckMultiple(item.id)">
                                    <h4 class="text-14 text-text-li font-normal">{{ $t('复盘') }} ({{
                                        utils.GoDate(item.created_at) }})</h4>
                                    <p class="text-12 text-text-li opacity-60">{{ utils.GoDateHMS(item.created_at)
                                    }}
                                    </p>
                                </div>
                            </div>
                            <p @click="nextReplayList" v-if="replayListLastPage > replayListPage"
                                class="text-12 mt-20 mb-20 text-text-tips text-center ">{{ $t('点击加载更多') }}
                            </p>
                        </div>
                    </n-scrollbar>
                </div>
            </div>
            <div v-if="loadIngR"
                class="absolute top-[71px] bottom-[24px] left-0 right-0 z-10 bg-[rgba(255,255,255,0.8)] flex-1 flex justify-center items-center">
                <n-spin size="small" />
            </div>
        </div>
    </div>
    <!-- 对齐目标 -->
    <SelectAlignment :value="selectAlignmentShow" :editData="detailData.align_objective"
        @close="() => { selectAlignmentShow = false }" @submit="submitSelectAlignment"></SelectAlignment>

    <!-- 更新进度   -->
    <DegreeOfCompletion v-model:show="degreeOfCompletionShow" :id="degreeOfCompletionId"
        :progress="degreeOfCompletionProgress" :progress_status="degreeOfCompletionProgressStatus"
        @close="handleCloseDedree">
    </DegreeOfCompletion>

    <!-- 更新信心 -->
    <Confidences :id="confidencesId" :confidence="confidence" v-model:show="confidenceShow" @close="handleCloseConfidenes">
    </Confidences>

    <!-- 更新评分 -->
    <MarkVue v-model:show="markShow" :id="markId" :score="score" :superior_score="superiorScore" :inputShow="inputShow"
        @close="handleCloseMarks">
    </MarkVue>

    <!-- 强提示 -->
    <TipsModal :show="showModal" :content="tipsContent" @close="() => { showModal = false }"></TipsModal>
</template>
<script setup lang="ts">
import { CheckmarkSharp } from '@vicons/ionicons5'
import { getOkrDetail, okrFollow, getLogList, getReplayList, okrCancel, alignUpdate, participantUpdate } from '@/api/modules/okrList'
import AlignTarget from "@/views/components/AlignTarget.vue";
import { ResultDialog } from "@/api"
import utils from '@/utils/utils';
import { useMessage } from "naive-ui"
import SelectAlignment from '@/views/components/SelectAlignment.vue'
import DegreeOfCompletion from '@/views/components/DegreeOfCompletion.vue'
import Confidences from '@/views/components/Confidences.vue';
import MarkVue from '@/views/components/Marks.vue';
import TipsModal from '@/views/components/TipsModal.vue';
import { GlobalStore } from '@/store';
import { useRouter } from 'vue-router'
import { UserStore } from '@/store/user'
import webTs from '@/utils/web';
import fenSvg from '@/assets/images/icon/fen.svg';


const userInfo = UserStore().info
const router = useRouter()
const globalStore = GlobalStore()
const userSelectApps = ref([]);
const navActive = ref(0)
const dialogWrappersApp = ref()
const loadIng = ref(false)
const loadIngR = ref(false)
const showPopover = ref(false)
const detailData = ref<any>({})
const message = useMessage()
const showModal = ref(false)
const tipsContent = ref('')
const scrollbarRef = ref(null)

const logListPage = ref(1)
const logListLastPage = ref(99999)
const logList = ref<any>([])

const replayListPage = ref(1)
const replayListLastPage = ref(99999)
const replayList = ref([])
const showDialogWrapper = computed(() => window.Vues?.components?.DialogWrapper ? 1 : 0)
const showUserSelect = computed(() => window.Vues?.components?.UserSelect ? 1 : 0)

const selectAlignmentShow = ref(false)
const AlignTargetRef = ref(null)

const degreeOfCompletionShow = ref(false)
const degreeOfCompletionId = ref(0)
const degreeOfCompletionProgress = ref(0)
const degreeOfCompletionProgressStatus = ref(0)

const confidenceShow = ref(false)
const confidencesId = ref(0)
const confidence = ref(0)

const markShow = ref(false)
const markId = ref(0)
const score = ref(0)
const superiorScore = ref(0)
const inputShow = ref(false)

const nowInterval = ref<any>(null)
const nowTime = ref(0)

const emit = defineEmits(['close', 'edit', 'upData', 'isFollow', 'canceled', 'getList', 'openDetail', 'getDetail'])

const props = defineProps({
    show: {
        type: Boolean,
        default: false,
    },
    id: {
        type: Number,
        default: 0,
    },
})


// 明细
const getDetail = (type) => {
    if (type == 'first') loadIng.value = true
    getOkrDetail({
        id: props.id,
    }).then(({ data }) => {
        detailData.value = type == 'first' ? data : {}
        if( type != 'first'  ){
            nextTick(() => {
                detailData.value = data
                emit('isFollow', detailData.value.is_follow)
                emit('getDetail', detailData.value)
                emit('canceled', detailData.value.canceled)
                if( navActive.value == 2){
                    replayListPage.value = 1
                    replayList.value = []
                    handleGetReplayList()
                }
            })
        }else{
            emit('isFollow', detailData.value.is_follow)
            emit('getDetail', detailData.value)
            emit('canceled', detailData.value.canceled)
            if( navActive.value == 2){
                replayListPage.value = 1
                replayList.value = []
                handleGetReplayList()
            }
        }
    })
    .catch(({ msg }) => {
        message.error(msg)
    })
    .finally(() => {
        if (type == 'first') loadIng.value = false
    })
}

// 关注
const handleFollowOkr = () => {
    loadIng.value = true
    okrFollow({
        id: detailData.value.id,
    }).then(({ msg }) => {
        message.success($t('操作成功'))
        emit('upData', detailData.value.id)
        getDetail('')
    })
        .catch(({ msg }) => {
            message.error(msg)
        })
        .finally(() => {
            loadIng.value = false
        })
}

//
const handleNav = (index) => {

    navActive.value = index
    if (navActive.value == 0 && window.innerWidth < 768) {
        loadDialogWrappers()
    }
    if (navActive.value == 1) {
        logListPage.value = 1
        logList.value = []
        handleGetLogList()
    }
    if (navActive.value == 2) {
        replayListPage.value = 1
        replayList.value = []
        handleGetReplayList()
    }
    if (navActive.value == 3) {
        
        nextTick(()=>{
            loadUserSelects()
        })

    }
}

// 日志下一页
const onScrollLogList = (e) => {
    if (e.target.scrollTop + e.target.offsetHeight + 10 >= e.target.scrollHeight) {
        // 重新请求数据
        if (!loadIngR.value) {
            logListPage.value++
            handleGetLogList()
        }
    }
}

// 日志下一页
const nextLogList = (e) => {
    if (!loadIngR.value) {
        logListPage.value++
        handleGetLogList()
    }
}

// 日志列表
const handleGetLogList = () => {
    if (logListLastPage.value >= logListPage.value) {
        loadIngR.value = true
        getLogList({
            id: detailData.value.id,
            page: logListPage.value,
            page_size: 10,
        }).then(({ data }) => {
            data.data.map(item => {
                if (item.content.includes('创建OKR')) {
                    item.content = $t('创建OKR')
                }
                if (item.content.includes('修改O目标标题')) {
                    item.content = $t('修改O目标标题') + ": " + item.records.title_change[0] + ' => ' + item.records.title_change[1]
                }
                if (item.content.includes('修改O目标周期')) {
                    item.content = $t('修改O目标周期') + ": " + item.records.time_change[0] + ' => ' + item.records.time_change[1]
                }
                if (item.content.includes('修改O目标状态')) {
                    item.content = $t('修改O目标状态') + ": " + $t(item.records.status_change[0]) + ' => ' + $t(item.records.status_change[1])
                }
                if (item.content.includes('修改对齐目标')) {
                    item.content = $t('修改对齐目标')
                }
                if (item.content.includes('取消对齐目标')) {
                    let parent_title = ''
                    item.records.parent_title ? parent_title = ' => ' + item.records.parent_title : ''
                    item.content = $t('取消对齐目标') + ": " + item.records.title + parent_title
                }
                if (item.content.includes('修改KR标题')) {
                    item.content = $t('修改KR标题') + ": " + item.records.title_change[0] + ' => ' + item.records.title_change[1]
                }
                if (item.content.includes('修改KR周期')) {
                    item.content = $t('修改KR周期') + ": " + item.records.title + ' [ ' + item.records.time_change[0] + ' => ' + item.records.time_change[1] + ' ]'
                }
                if (item.content.includes('修改KR参与人')) {
                    item.content = $t('修改KR参与人') + ": " + item.records?.title
                }
                if (item.content.includes('修改KR进度')) {
                    item.content = $t('修改KR进度') + ": " + item.records.title + ' [ ' + item.records.progress_change[0] + '%' + ' => ' + item.records.progress_change[1] + '%' + ' ]'
                }
                if (item.content.includes('修改KR状态')) {
                    item.content = $t('修改KR状态') + ": " + item.records.title + ' [ ' + $t(item.records.progress_status_change[0]) + ' => ' + $t(item.records.progress_status_change[1]) + ' ]'
                }
                if (item.content.includes('修改KR信心指数')) {
                    item.content = $t('修改KR信心指数') + ": " + item.records.title + ' [ ' + item.records.confidence_change[0] + ' => ' + item.records.confidence_change[1] + ' ]'
                }
                if (item.content.includes('责任人打分')) {
                    item.content = $t('责任人打分') + ": " + item.records.title
                }
                if (item.content.includes('上级打分')) {
                    item.content = $t('上级打分') + ": " + item.records.title
                }
                if (item.content.includes('添加KR')) {
                    item.content = $t('添加KR') + ": " + item.records.title
                }
                if (item.content.includes('删除KR')) {
                    item.content = $t('删除KR') + ": " + item.records.title
                }
                logList.value.push(item)
            })


            logListLastPage.value = data.last_page
        })
            .catch(({ msg }) => {

            })
            .finally(() => {
                loadIngR.value = false
            })
    }
}

//复盘下一页
// const onScrollReplayList = (e) => {
//     if (e.target.scrollTop + e.target.offsetHeight + 10 >= e.target.scrollHeight) {
//         // 重新请求数据
//         if (!loadIng.value) {
//             replayListPage.value ++
//             handleGetReplayList()
//         }
//     }
// }

const nextReplayList = (e) => {
    // 重新请求数据
    if (!loadIng.value) {
        replayListPage.value ++
        handleGetReplayList()
    }
}

// 复盘列表
const handleGetReplayList = () => {
    if (replayListLastPage.value >= replayListPage.value) {
        loadIngR.value = true
        getReplayList({
            id: detailData.value.id,
            page: replayListPage.value,
            page_size: 10,
        }).then(({ data }) => {
            data.data.map(item => {
                replayList.value.push(item)
            })
            replayListLastPage.value = data.last_page
        })
            .catch(({ msg }) => {

            })
            .finally(() => {
                loadIngR.value = false
            })
    }
}

//编辑
const handleEdit = () => {
    if (window.innerWidth < 768) {
        router.push({
            path: globalStore.baseRoute + '/addOkr',
        })
        globalStore.$patch((state) => {
            state.okrEditData = detailData.value
            state.okrEdit = true
        })
    }
    else {
        emit('edit', utils.cloneJSON(detailData.value))
    }
}

const closeModal = () => {
    emit('close')
}

//打开进度
const handleSchedule = (id, progress, progress_status, score) => {
    if (userInfo.userid != detailData.value.userid) {
        tipsContent.value = $t('仅限负责人操作')
        showModal.value = true
        return
    }
    if (score > -1) return message.error($t('KR已评分无法操作'))
    if (detailData.value.completed == '1') return message.error($t('O目标已完成无法操作'))
    if (detailData.value.canceled == '1') return message.error($t('O目标已取消无法操作'))
    degreeOfCompletionId.value = id
    degreeOfCompletionProgress.value = progress
    degreeOfCompletionProgressStatus.value = progress_status
    degreeOfCompletionShow.value = true
}

//关闭进度
const handleCloseDedree = (type) => {
    if (type == 1) {
        getDetail('')
        AlignTargetRef.value.getList()
        emit('upData', detailData.value.id)
    }
    degreeOfCompletionId.value = 0
    degreeOfCompletionProgress.value = 0
    degreeOfCompletionProgressStatus.value = 0
    degreeOfCompletionShow.value = false
}

//打开信心
const handleConfidence = (id, confidences, score) => {
    if (userInfo.userid != detailData.value.userid) {
        tipsContent.value = $t('仅限负责人操作')
        showModal.value = true
        return
    }
    if (score > -1) return message.error($t('KR已评分无法操作'))
    if (detailData.value.completed == '1') return message.error($t('O目标已完成无法操作'))
    if (detailData.value.canceled == '1') return message.error($t('O目标已取消无法操作'))
    confidencesId.value = id
    confidence.value = confidences
    confidenceShow.value = true
}

//关闭信心
const handleCloseConfidenes = (type) => {
    confidencesId.value = 0
    confidence.value = 0
    confidenceShow.value = false
    if (type == 1) {
        getDetail('')
    }
}

//打开评分
const handleMark = (id, scores, superior_score, progress) => {
    if (userInfo.userid == detailData.value.userid || (detailData.value.superior_user?.indexOf(userInfo.userid) != -1 && detailData.value.superior_user?.indexOf(userInfo.userid) != undefined)) {
        if (detailData.value.canceled == '1') return message.error($t('O目标已取消无法操作'))
        if (progress < 100) return message.error($t('KR进度尚未达到100%'))
        if (scores < 0 && detailData.value.superior_user > 0 && detailData.value.superior_user?.indexOf(userInfo.userid) != -1) return message.error($t('负责人未评分'))
        markId.value = id
        score.value = scores
        superiorScore.value = superior_score
        if (userInfo.userid == detailData.value.userid && scores < 0) {
            inputShow.value = true
        }
        if (userInfo.userid == detailData.value.userid && scores > -1) {
            inputShow.value = false
        }
        if (detailData.value.superior_user?.indexOf(userInfo.userid) != -1 && superior_score < 0) {
            inputShow.value = true
        }
        if (detailData.value.superior_user?.indexOf(userInfo.userid) != -1 && superior_score > -1) {
            inputShow.value = false
        }
        if (scores > -1 && superior_score > -1) {
            inputShow.value = false
        }
        markShow.value = true
    } else {
        message.error($t('仅负责人和负责人上级可以操作！'))
    }
}
//关闭评分
const handleCloseMarks = (type) => {
    markId.value = 0
    score.value = 0
    superiorScore.value = 0
    markShow.value = false
    if (type == 1) {
        getDetail('')
    }
}

// 取消O
const handleCancel = () => {
    if (userInfo.userid != detailData.value.userid) {
        tipsContent.value = $t('仅限负责人操作')
        showModal.value = true
        showPopover.value = false
        return
    }
    loadIng.value = true
    okrCancel({
        id: detailData.value.id,
    }).then(({ msg }) => {
        message.success($t('修改成功'))
        emit('upData', detailData.value.id)
        getDetail('')
    })
        .catch(ResultDialog)
        .finally(() => {
            showPopover.value = false
            loadIng.value = false
        })
}

// 取消
const handleUnalign = () => {
    emit('upData', detailData.value.id)
    getDetail('')
}

// 点击对齐目标跳转
const openDetail = (id, userid) => {
    emit('openDetail', id, userid)
    userSelectApps.value.forEach(app => {
        let dom = document.createElement("UserSelects")
        dom.setAttribute('formkey',app._vnode.data.formkey)
        app.$el.replaceWith(dom);
        app.$destroy()
    })
    nextTick(() => {
        scrollbarRef.value.scrollTo({top: 0})
        getDetail('')
        AlignTargetRef.value.getList()
        loadDialogWrappers();
        if (window.innerWidth < 768) {
            navActive.value = 3
        }
        else {
            navActive.value = 0
        }
    })
}

// 加载聊天组件
const loadDialogWrappers = () => {
    dialogWrappersApp.value && (dialogWrappersApp.value.$children[0].allMsgs = [])
    nextTick(() => {
        if (!window.Vues) return false;
        if(dialogWrappersApp.value){
            dialogWrappersApp.value.$el.replaceWith(document.createElement("DialogWrappers"));
            dialogWrappersApp.value.$destroy()
        }
        dialogWrappersApp.value = new window.Vues.Vue({
            el: document.querySelector('DialogWrappers'),
            store: window.Vues.store,
            render: (h: any) => {
                return h(window.Vues?.components?.DialogWrapper, {
                    props: {
                        dialogId: detailData.value.dialog_id
                    }
                }, [h("div", { slot: "head" })])
            }
        });
    })
}

//添加复盘
const handleAddMultiple = () => {
    if (userInfo.userid != detailData.value.userid) {
        tipsContent.value = $t('仅限负责人操作')
        showModal.value = true
        return
    }
    if (detailData.value.score < 0) return message.error($t('KR评分未完成'))
    if (window.innerWidth < 768) {
        router.push({
            path: globalStore.baseRoute + '/addMultiple',
        })
        globalStore.$patch((state) => {
            state.addMultipleData = detailData.value
            state.multipleId = 0
            state.doubleSkip = true
        })
    }
    else {
        closeModal()
        globalStore.$patch((state) => {
            state.addMultipleData = detailData.value
            state.addMultipleShow = true
        })
    }
}


//查看复盘
const handleCheckMultiple = (id) => {
    if (window.innerWidth < 768) {
        router.push({
            path: globalStore.baseRoute + '/addMultiple',
        })
        globalStore.$patch((state) => {
            state.addMultipleData = detailData.value
            state.multipleId = id
            state.doubleSkip = true
        })
    }
    else {
        closeModal()
        globalStore.$patch((state) => {
            state.multipleId = id
            state.addMultipleShow = true
        })
    }
}


// 加载选择用户组件
const loadUserSelects = () => {
    nextTick(() => {
        if (!window.Vues) return false;
        document.querySelectorAll('userselects').forEach(e => {
            let index = e.getAttribute('formkey')
            let item = detailData.value.key_results[e.getAttribute('formkey')];
            let app = new window.Vues.Vue({
                el: e,
                store: window.Vues.store,
                render: (h: any) => {
                    return h(window.Vues?.components?.UserSelect, {
                        class: "okr-user-selects",
                        formkey: index,
                        props: {
                            value: ((item.participant).split(',').map(h => Number(h) )).filter(value => value !== 0),
                            title: $t('选择参与人'),
                            border: false,
                            avatarSize: 20,
                            addIcon: ((item.participant).split(',').map(h => Number(h) )).filter(value => value !== 0).length == 0 ,
                        },
                        on: {
                            "on-show-change": (show: any, values: any) => {
                                if (!show) {
                                    if (item.participant != values.join(',')) {
                                        item.participant = values.join(',');
                                        participantChange(item, e.getAttribute('formkey'))
                                    }
                                }
                            }
                        }
                    })
                },
            });
            userSelectApps.value.push(app)
        })
    })
}


//更新参与人
const participantChange = (item, index) => {
    if (detailData.value.key_results[index].score > -1 || detailData.value.key_results[index].superior_score > -1) {
        return
    }
    const upData = {
        participant: item.participant,
        id: item.id,
    }
    loadIng.value = true
    participantUpdate(upData)
        .then(({ msg }) => {
            message.success($t('修改成功'))
            getDetail('')
        })
        .catch(({ msg }) => {
            message.error(msg)
        })
        .finally(() => {
            loadIng.value = false
        })
}


//超期判断
const within24Hours = (date) => {
    let time = utils.GoDateHMS(date)
    return Number(utils.Date(time, true)) - nowTime.value < 86400
}

const expiresFormat = (date) => {
    return webTs.countDownFormat(date, nowTime.value)
}


const isOverdue = (item) => {
    let time = utils.GoDateHMS(item.end_at)
    return Number(utils.Date(time, true)) < nowTime.value;
}


//添加对齐目标
const submitSelectAlignment = (e) => {
    const upData = {
        align_objective: e.join(','),
        id: detailData.value.id,
    }
    loadIng.value = true
    alignUpdate(upData)
        .then(({ msg }) => {
            message.success($t('修改成功'))
            emit('upData', detailData.value.id)
            getDetail('')
            AlignTargetRef.value.getList()
        })
        .catch(({ msg }) => {
            message.error(msg)
        })
        .finally(() => {
            loadIng.value = false
        })
}

//颜色判断
const colorStatus = (color) => {
    let result = ''
    if (color == 1 || color == 0) {
        result = '#8BCF70'
    }
    if (color == 2) {
        result = '#FFA25A'
    }
    if (color == 3) {
        result = '#FF7070'
    }
    return result
}


// 卸载
window.addEventListener('apps-unmount', function () {
    closeDrawer()
})

onBeforeUnmount(() => {
    closeDrawer()
})

// 关闭
const closeDrawer = () => {
    dialogWrappersApp.value && (dialogWrappersApp.value.$children[0].allMsgs = []);
    nextTick(() => {
        dialogWrappersApp.value && dialogWrappersApp.value.$destroy() && (dialogWrappersApp.value = null);
        navActive.value = 0
        detailData.value = {};
        loadIng.value = true;
        userSelectApps.value.forEach(app => app.$destroy())
    })

}

const pStatus = (p) => {
    return p == 'P0' ? 'span-1' : p == 'P1' ? 'span-2' : 'span-3'
}

watch(() => props.show, (newValue) => {
    if (newValue) {
        getDetail('first')
    }
}, { immediate: true })


watch(() => detailData.value.dialog_id, (newValue) => {
    if (newValue) {
        loadUserSelects()
        if (window.innerWidth >= 768) {
            loadDialogWrappers()
        }
    }
}, { immediate: true })



onMounted(() => {
    nextTick(() => {
    if (window.innerWidth < 768 && !globalStore.doubleSkip) {
        navActive.value = 3
    }
    if(window.innerWidth < 768 && globalStore.doubleSkip){
        globalStore.$patch((state) => {
            state.doubleSkip = false
            navActive.value = 2
        })
    }
})
    nowInterval.value = setInterval(() => {
        nowTime.value = utils.Time();
    }, 1000);
})

onUnmounted(() => {
    clearInterval(nowInterval.value);
})

defineExpose({
    getDetail,
    closeDrawer,
    handleEdit,
    handleCancel,
    handleFollowOkr
})
</script>
<style lang="less" scoped>
.navActive {
    @apply hidden md:flex;
}

.icon-title {
    @apply text-16 cursor-pointer;
}

.icon-item {
    @apply text-16 mr-8;
}

.span {
    @apply text-12 font-medium text-white px-6 py-4 rounded-full flex items-center leading-3 shrink-0;
}

.span-1 {
    @apply bg-[#FF7070];
}

.span-2 {
    @apply bg-[#FC984B];
}

.span-3 {
    @apply bg-[#72A1F7];
}


.li-nav {
    @apply list-none cursor-pointer text-text-li text-14 px-4 relative;
}

.li-nav.active {
    @apply text-primary-color md:text-text-li md:text-16 opacity-100 relative font-medium;
}

.li-nav.active::before {
    @apply w-full bg-primary-color absolute left-0 right-0 -bottom-12 block md:hidden;
    height: 2px;
    content: ' ';
}

:deep(.n-button--error-type) {
    @apply bg-primary-color;
}

:deep(.left-scrollbar) {
    .n-scrollbar-rail {
        @apply right-0;
    }
}</style>
