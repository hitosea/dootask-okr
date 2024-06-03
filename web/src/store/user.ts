import { computed } from "vue";
import { defineStore } from 'pinia';
import { UserState } from './interface';
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
            identity: [],
            userid: 0,
            department_owner: null,
            department: null,
            okr_admin_owner: null,
        },
        isAdmin: '',
        isOkrAdminOwner: '',
        isDepartmentOwner: ''
    }),
    actions: {
        auth() {
            return {
                isAdmin: computed(() => {
                    return this.info?.identity?.indexOf('admin') !== -1
                }),
                isOkrAdminOwner: computed(() => {
                    return this.info?.okr_admin_owner
                }),
                isDepartmentOwner: computed(() => {
                    return this.info?.department_owner
                })
            }
        },
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
    }
});
