import { defineStore } from 'pinia';
import { UserState } from './interface';
import piniaPersistConfig from "./config/pinia-persist";
import { getUserInfo } from "../api/modules/user";
import utils from "../utils/utils";
export const UserStore = defineStore({
    id: 'UserState',
    state: (): UserState => ({
        info: {
            id: 0,
            email: "",
            name: "",
            token: "",
            avatar: "",
            created_at: "",
            updated_at: "",
            userIsAdmin: false,
            identity:[],
            userid:0,
        },
    }),
    actions: {
        refresh() {
            return new Promise((resolve, reject) => {
                getUserInfo().then(({ data }) => {
                    if (utils.isEmpty(data.name)) {
                        data.name = data.email
                    }
                    this.info = data
                    resolve(data)
                })
                .catch(err => {
                    reject(err)
                })
            })
        },
        setUserInfo(info: any = null) {
            this.info = info || {
                id: 0,
                email: "",
                name: "",
                token: "",
                avatar: "",
                created_at: "",
                updated_at: "",
            };
        },
        setToken(token: any = '') {
            this.info.token = token;
        }
    },
    persist: piniaPersistConfig('UserState'),
});
