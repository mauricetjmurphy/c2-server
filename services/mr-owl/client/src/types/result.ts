export interface ResultState {
    results: Result[]
    addResultItem: (item: Result) => void
}

export interface Result {
    command: string
    contents: string
    id: number
    result_id: string
    success: boolean
    task_id: string
}
