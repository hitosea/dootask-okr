// import directives
import { App } from "vue";
import longpress from "./modules/longpress";
import touchmouse from "./modules/touchmouse";
import clickoutside from "./modules/clickoutside";
import TransferDom from "./modules/transfer-dom";
 
const directivesList: any = {
    // Custom directives
    longpress,
    touchmouse,
    clickoutside,
    TransferDom,
};
 
const directives = {
    install: function (app: App<Element>) {
        Object.keys(directivesList).forEach(key => {
            // 注册自定义指令
            app.directive(key, directivesList[key]);
        });
    }
};
 
export default directives;