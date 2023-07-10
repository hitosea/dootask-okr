/**
 * v-clickoutside
 * 点击外面
 */
import type { Directive, DirectiveBinding } from "vue";
const clickoutside: Directive = {
    mounted(el, binding: DirectiveBinding) {
        const documentHandler = (e) => {
            if (el.contains(e.target)) {
                return false;
            }
            if (binding.value) {
                binding.value(e)
            }
        }
        el.__vueClickOutside__ = documentHandler;
        document.addEventListener('click', documentHandler);
    },
    unmounted(el){
        document.removeEventListener('click', el.__vueClickOutside__);
        delete el.__vueClickOutside__;
    }
};

export default clickoutside;