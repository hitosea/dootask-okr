import http from "../index"
import * as Replay from "../interface/replay"

export const getReplayList = () => {
    return http.get<Replay.replayList>("okr/replay/list")
}