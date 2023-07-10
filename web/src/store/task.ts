import { createPinia,defineStore } from 'pinia'
import { task } from './interface/task'
import piniaPersistConfig from "./config/pinia-persist"
export const TaskStore = defineStore({
    id: 'TaskState',
    state: (): task => ({
        cacheProjects: [],
        cacheColumns: [],
        cacheTasks: [],
        cacheProjectParameter: [],
        cacheTaskBrowse: [],
        cacheUserBasic: [],
        cacheDialogs: [],
        messengerSearchKey: {dialog: '', contacts: ''},
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

    },
    persist: piniaPersistConfig('TaskState'),
});
const pinia = createPinia()
export default pinia