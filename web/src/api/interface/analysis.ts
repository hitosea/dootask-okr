/* eslint-disable @typescript-eslint/no-namespace */
export interface complete {
    total: number
    complete: number
}

export interface deptCompletes {
    total: number
    complete: number,
    department_id: number,
    department_name: string
}

export interface score {
    total: number
    complete: number,
    unscored: number,
    zero_to_three: number,
    three_to_seven: number,
    seven_to_ten: number,
}

export interface deptScore {
    department_id: number,
    department_name: string,
    total: number
    complete: number,
    unscored: number,
    zero_to_three: number,
    three_to_seven: number,
    seven_to_ten: number,
    already_reviewed: number,
}