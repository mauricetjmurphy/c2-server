import axios from 'axios'
import { LISTENER_URL } from '../../config'
import { Result } from '../../types/result'

export function getResults(): Promise<Result[]> {
    return axios
        .get(`${LISTENER_URL}/results`)
        .then((resp) => {
            return resp.data ?? []
        })
        .catch((error) => {
            throw error
        })
}
