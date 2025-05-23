import axios from 'axios'
import { LISTENER_URL } from '../../config'
import { Task } from '../../types/task'

export function getTasks(): Promise<Task[]> {
    return axios
        .get(`${LISTENER_URL}/tasks`)
        .then((resp) => {
            return resp.data ?? []
        })
        .catch((error) => {
            throw error
        })
}

export function sendTask(task: Task): Promise<Task> {
    return axios
        .post(`${LISTENER_URL}/task`, task)
        .then((resp) => {
            return resp.data ?? {}
        })
        .catch((error) => {
            throw error
        })
}
