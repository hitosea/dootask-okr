import { User } from "@/api/interface/user";

export interface GlobalState {
    baseUrl: string
    language: string
    themeName: string
    timer: object
    addMultipleShow: any
    multipleId: number
    addMultipleData: any
    addMultipleChange: any
    okrDetailId: any
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
