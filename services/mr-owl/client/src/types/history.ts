export interface HistoryState {
    history: HistoryItem[]
    addHistoryItem: (item: HistoryItem) => void
}

export interface HistoryItem {
    history_id: string
    id: number
    success: boolean
    task_id: string
    task_options: string
    task_result: string
    task_type: string
}
