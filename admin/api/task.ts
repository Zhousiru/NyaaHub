import { NewTask, Task, TaskConfig } from './types'

const apiUrl =
  typeof window === 'undefined' ? '' : localStorage.getItem('apiUrl') ?? ''
const token =
  typeof window === 'undefined' ? '' : localStorage.getItem('token') ?? ''

const withQuery = (path: string, q?: Record<string, string>) =>
  apiUrl + path + '?' + new URLSearchParams({ token: token, ...q }).toString()

async function listTask(
  limit: number,
  start?: string
): Promise<{
  list: Array<Task>
  hasNext: boolean
}> {
  const res = await fetch(
    withQuery('/listTask', {
      start: start ?? '',
      limit: limit.toString(),
    })
  )

  if (res.status != 200) {
    throw new Error((await res.json()).msg)
  }

  return (await res.json()).payload
}

async function addTask(task: NewTask) {
  task.config.maxDownload = Number(task.config.maxDownload)
  task.config.timeout = Number(task.config.timeout)

  const res = await fetch(withQuery('/addTask'), {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(task),
  })

  if (res.status != 200) {
    throw new Error((await res.json()).msg)
  }
}

async function removeTask(collection: string) {
  const res = await fetch(withQuery('/removeTask', { collection }))

  if (res.status != 200) {
    throw new Error((await res.json()).msg)
  }
}

async function updateTask(collection: string, newConfig: TaskConfig) {
  newConfig.maxDownload = Number(newConfig.maxDownload)
  newConfig.timeout = Number(newConfig.timeout)

  const task: NewTask = {
    collection,
    config: newConfig,
  }

  const res = await fetch(withQuery('/updateTaskConfig'), {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(task),
  })

  if (res.status != 200) {
    throw new Error((await res.json()).msg)
  }
}

export { listTask, addTask, removeTask, updateTask }
