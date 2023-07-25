export interface replayList {
    page: number,
    page_size: number,
    count: number,
    data: replayOkr[]
}

export interface replayOkr {
    id: number,
    userid: number,
    okr_id: number,
    okr_userid: number,
    okr_department_id: number,
    okr_title: string,
    okr_progress: number,
    okr_priority: number,
    review: string,
    created_at: string,
    updated_at: string
}