import http from "../index"
import * as Icreated from "../interface/icreated"

export const getNumberofOkrIcreatedPending = () => {
    return http.get<Icreated.okrCompletedAndUncompleted>("okr/statistics/completes")
}

export const getOkrDegreeOfCompletionAndScore = () => {
    return http.get<Icreated.okrDegreeOfCompletionAndScore>("okr/statistics/overall")
}