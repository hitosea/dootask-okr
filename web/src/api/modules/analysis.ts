import http from "../index"
import * as Analysis from "../interface/analysis"

export const getAnalyzeDepartment = () => {
    return http.get<Analysis.complete>("okr/analyze/department")
}

//
export const getAnalyzeComplete = (params={}) => {
    return http.get<Analysis.complete>("okr/analyze/complete",params)
}

export const getAnalyzeDeptComplete = (params={}) => {
    return http.get<[Analysis.deptCompletes]>("okr/analyze/dept/complete",params)
}

//
export const getAnalyzeScore = (params={}) => {
    return http.get<Analysis.score>("okr/analyze/score",params)
}

export const getAnalyzeDeptScore = (params={}) => {
    return http.get<[Analysis.deptScore]>("okr/analyze/dept/score",params)
}

//
export const getAnalyzeScoreSate = (params={}) => {
    return http.get<Analysis.complete>("okr/analyze/personnel/score/rate",params)
}

export const getAnalyzeDeptScoreProportion = (params={}) => {
    return http.get<[Analysis.deptScore]>("okr/analyze/dept/score/proportion",params)
}
