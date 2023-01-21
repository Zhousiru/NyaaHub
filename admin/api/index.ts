import { NewTask, Task, TaskConfig } from './types'
import dayjs from 'dayjs'
import utc from 'dayjs/plugin/utc'
dayjs.extend(utc)

const apiUrl = () => localStorage.getItem('apiUrl') ?? ''
const token = () => localStorage.getItem('token') ?? ''

const withQuery = (path: string, q?: Record<string, string>) =>
  apiUrl() +
  path +
  '?' +
  new URLSearchParams({ token: token(), ...q }).toString()

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

async function addTask(
  task: NewTask,
  downloadPrev: boolean,
  downloadOnly: boolean = false
) {
  task.config.maxDownload = Number(task.config.maxDownload)
  task.config.timeout = Number(task.config.timeout)

  const res = await fetch(
    withQuery('/addTask', {
      prev: String(downloadPrev),
      downloadOnly: String(downloadOnly),
    }),
    {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(task),
    }
  )

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

async function getLog(collection: string): Promise<string> {
  // TODO: Declare a log interface

  const res = await fetch(withQuery('/getLog', { collection }))

  if (res.status != 200) {
    throw new Error((await res.json()).msg)
  }

  const payload = (await res.json()).payload

  let ret = ''
  for (const log of payload) {
    const logTime = dayjs
      .utc(log.time)
      .local()
      .format('YYYY-MM-DD HH:mm:ss.SSS')

    ret += `[${logTime}][${log.type}] ${log.msg}\n`
  }

  return ret
}

export { listTask, addTask, removeTask, updateTask, getLog }
