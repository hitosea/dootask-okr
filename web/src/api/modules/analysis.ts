import http from "../index"
import * as Analysis from "../interface/analysis"

// 
export const getAnalyzeComplete = () => {
    return http.get<Analysis.complete>("okr/analyze/complete")
}

export const getAnalyzeDeptComplete = () => {
    return http.get<[Analysis.deptCompletes]>("okr/analyze/dept/complete")
}

// 
export const getAnalyzeScore = () => {
    return http.get<Analysis.score>("okr/analyze/score")
}

export const getAnalyzeDeptScore = () => {
    return http.get<[Analysis.deptScore]>("okr/analyze/dept/score")
}

// 
export const getAnalyzeScoreSate = () => {
    return http.get<Analysis.complete>("okr/analyze/personnel/score/rate")
}

export const getAnalyzeDeptScoreProportion = () => {
    return http.get<[Analysis.deptScore]>("okr/analyze/dept/score/proportion")
}