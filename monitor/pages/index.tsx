import Head from 'next/head'
import TorrentTaskList from '../components/TorrentTask'

export default function Home() {
  return (
    <>
      <Head>
        <title>NyaaHub Monitor</title>
        <meta name="robots" content="noindex" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <main className="flex flex-col items-center justify-center">
        <TorrentTaskList></TorrentTaskList>
      </main>
    </>
  )
}
