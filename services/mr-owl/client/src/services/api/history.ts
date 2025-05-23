import axios from 'axios'
import { HistoryItem } from '../../types/history'
import { LISTENER_URL } from '../../config'

export function getHistory(): Promise<HistoryItem[]> {
    return axios
        .get(`${LISTENER_URL}/history`)
        .then((resp) => {
            return resp.data ?? []
        })
        .catch((error) => {
            throw error
        })
}
