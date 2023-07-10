<template>
    <div class="common-circle" :data-id="percent" :style="style">
        <svg viewBox="0 0 28 28">
            <g fill="none" fill-rule="evenodd">
                <path class="common-circle-path" d="M-500-100h997V48h-997z"></path>
                <g fill-rule="nonzero">
                    <path class="common-circle-g-path-ring" stroke-width="3"
                        d="M14 25.5c6.351 0 11.5-5.149 11.5-11.5S20.351 2.5 14 2.5 2.5 7.649 2.5 14 7.649 25.5 14 25.5z">
                    </path>
                    <path class="common-circle-g-path-core" :d="arc(args)" />
                </g>
            </g>
        </svg>
    </div>
</template>

<script setup lang="ts">
import { defineProps, computed ,ref } from "vue"

const sizeReturn =ref('')

const props = defineProps({
    percent: {
        type: Number,
        default: 0
    },
    size: {
        type: Number,
        default: 120
    },
})


const style = computed(() => {
    const sizeClone = props.size
    if (isNumeric(props.size)) {
        sizeReturn.value = sizeClone.toString() + 'px';
    }
    return {
        width: sizeReturn.value,
        height: sizeReturn.value,
    }
})

const args = computed(() => {
    let end = Math.min(360, 360 / 100 * props.percent);
    if (end == 360) {
        end = 0;
    } else if (end == 0) {
        end = 360;
    }
    return {
        x: 14,
        y: 14,
        r: 14,
        start: 360,
        end: end,
    }
})



const isNumeric = (n) => {
    return n !== '' && !isNaN(parseFloat(n)) && isFinite(n);
}

const point = (x, y, r, angel) => {
    return [
        (x + Math.sin(angel) * r).toFixed(2),
        (y - Math.cos(angel) * r).toFixed(2),
    ]
}

const full = (x, y, R, r) => {
    if (r <= 0) {
        return `M ${x - R} ${y} A ${R} ${R} 0 1 1 ${x + R} ${y} A ${R} ${R} 1 1 1 ${x - R} ${y} Z`;
    }
    return `M ${x - R} ${y} A ${R} ${R} 0 1 1 ${x + R} ${y} A ${R} ${R} 1 1 1 ${x - R} ${y} M ${x - r} ${y} A ${r} ${r} 0 1 1 ${x + r} ${y} A ${r} ${r} 1 1 1 ${x - r} ${y} Z`;
}

const part = (x, y, R, r, start, end) => {
    const [s, e] = [(start / 360) * 2 * Math.PI, (end / 360) * 2 * Math.PI];
    const P = [
        point(x, y, r, s),
        point(x, y, R, s),
        point(x, y, R, e),
        point(x, y, r, e),
    ];
    const flag = e - s > Math.PI ? '1' : '0';
    return `M ${P[0][0]} ${P[0][1]} L ${P[1][0]} ${P[1][1]} A ${R} ${R} 0 ${flag} 1 ${P[2][0]} ${P[2][1]} L ${P[3][0]} ${P[3][1]} A ${r} ${r}  0 ${flag} 0 ${P[0][0]} ${P[0][1]} Z`;
}

const arc = (opts) => {
    const { x = 0, y = 0 } = opts;
    let { R = 0, r = 0, start, end, } = opts;

    [R, r] = [Math.max(R, r), Math.min(R, r)];
    if (R <= 0) return '';
    if (start !== +start || end !== +end) return full(x, y, R, r);
    if (Math.abs(start - end) < 0.000001) return '';
    if (Math.abs(start - end) % 360 < 0.000001) return full(x, y, R, r);

    [start, end] = [start % 360, end % 360];

    if (start > end) end += 360;
    return part(x, y, R, r, start, end);
}


</script>
<style lang="less">
.common-circle {
    @apply rounded-full;
    .common-circle-path {
        @apply fill-transparent;
    }

    .common-circle-g-path-ring {
        @apply stroke-text-li;
    }

    .common-circle-g-path-core {
        @apply fill-text-li;
        transform: scale(0.56);
        transform-origin: 50%;
    }
}
</style>