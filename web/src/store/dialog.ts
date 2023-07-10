import {defineStore } from 'pinia'
import { diaLog } from './interface/dialog'
import piniaPersistConfig from "./config/pinia-persist"
export const DiaLogStore = defineStore({
    id: 'diaLogState',
    state: (): diaLog => ({
        loadDialogs: 0,
        dialogId: 0,
        dialogTodos: [],
        cacheDialogs: [],
        appNotificationPermission: true,
        audioPlaying: null,
        loads: [],
        dialogMsgs: [],
    }),
    actions: {
        // setUserInfo (info:any=null) {
        //     this.info = info || {
        //         id: null,
        //         email: "",
        //         name: "",
        //         token: "",
        //         avatar: "",
        //         created_at: "",
        //         updated_at: "",
        //     };
        // }
    },
    getters: {
        isLoad: (state) => (key) => {
            const load = state.loads.find(item => item.key === key)
            return load && load.num > 0
        }
    },
    persist: piniaPersistConfig('diaLogState'),
});
