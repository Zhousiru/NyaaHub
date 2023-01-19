interface Task {
  collection: string
  config: TaskConfig
  download: number
  lastUpdate: string
}

interface NewTask {
  collection: string
  config: TaskConfig
}

interface TaskConfig {
  rss: string
  cron: string
  cronTimeZone: string
  maxDownload: number
  timeout: number
}

export type { Task, NewTask, TaskConfig }
