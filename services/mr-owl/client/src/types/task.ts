export interface TaskState {
    tasks: Task[]
    addTaskItem: (item: Task) => void
}

export interface Task {
    command: string
    id: number
    task_id: string
    task_type: string
}
