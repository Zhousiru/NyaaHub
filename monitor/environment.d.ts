declare global {
  namespace NodeJS {
    interface ProcessEnv {
      FETCHER_API: string
    }
  }
}

export {}
