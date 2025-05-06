import { User } from "@/api/interface/user";

export interface GlobalState {
    baseUrl: string
    baseRoute: string
    language: string
    themeName: string
    isElectron: boolean
    timer: object
    okrDetail: any
    addMultipleShow: any
    multipleId: number
    addMultipleData: any
    addMultipleChange: any
    okrEditData: any
    okrEdit: boolean
    doubleSkip: any,
}

export interface UserState {
    info: User.Info,
    isAdmin: string,
    isOkrAdminOwner: string,
    isDepartmentOwner: string
}
