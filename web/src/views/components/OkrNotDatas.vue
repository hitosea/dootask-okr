<template>
    <div class=" w-full h-full flex mt-[10%] items-center flex-col">
        <div>
            <img class="mr-6 -mt-2" src="@/assets/images/icon/notData.svg" />
        </div>
        <div class="text-[#515A6E] text-14 text-opacity-50 text-center">        
            <slot name="content"></slot>
        </div>
    </div>
</template>
