// 通用返回
export interface Result {
    code: number
    msg: string,
}

// 通用返回数据
export interface ResultData<T = any> extends Result {
    data?: T
}

// 数据库基础字段
export interface DatabaseBase {
    id: number
    created_at: number
    updated_at: number
}

// 分页请求
export interface PageReq<T = any> {
    count?: number
    data?: T
    page?: number
    page_size?: number
    last_page?: number
}

// 分页返回
export interface Page {
    page: number
    page_size: number

    next_page: number
    prev_page: number
    page_count: number

    total: number
}
//
export interface replayData {
    okr_title?:any
    kr_history?:any
    okr_progress?:any
    review?:any
    problem?:any
}
