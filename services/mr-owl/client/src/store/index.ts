import { create } from 'zustand'
import { HistoryState, HistoryItem } from '../types/history'
import { Result, ResultState } from '../types/result'
import { Task, TaskState } from '../types/task'
import { MenuState } from '../types/menu'

interface StoreState extends HistoryState, ResultState, TaskState, MenuState {}

const useStore = create<StoreState>((set) => ({
    selectedMenuItem: 1,
    setSelectedMenuItem: (selectedItem: number) =>
        set(() => ({
            selectedMenuItem: selectedItem,
        })),

    history: [],
    addHistoryItem: (item: HistoryItem) =>
        set((state) => ({
            history: [...state.history, item],
        })),

    results: [],
    addResults: (results: Result[]) =>
        set((state) => ({
            results: [...state.results, ...results],
        })),
    addResultItem: (item: Result) =>
        set((state) => ({
            results: [...state.results, item],
        })),

    tasks: [],
    addTasks: (tasks: Task[]) =>
        set((state) => ({
            tasks: [...state.tasks, ...tasks],
        })),
    addTaskItem: (item: Task) =>
        set((state) => ({
            tasks: [...state.tasks, item],
        })),
}))

export default useStore
