import http from "../index"
import * as Icreated from "../interface/icreated"

export const getNumberofOkrIcreatedPending = () => {
    return http.get<Icreated.okrCompletedAndUncompleted>("okr/statistics/all")
}

export const getDegreeOfCompletionAndScore = () => {
    return http.get<Icreated.degreeOfCompletionAndScore>("okr/overall")
}