import dayjs from 'dayjs'
import useSWR from 'swr'

import duration from 'dayjs/plugin/duration'
dayjs.extend(duration)

function TorrentTaskEntry({ status }: { status: any }) {
  let humanizedEta = 'N/A'
  if (status.eta > 0) {
    humanizedEta = dayjs
      .duration(status.eta, 'seconds')
      .format('H [hrs] m[mins] ss[secs]')
  }

  let humanizedStatus = 'Unknown'
  switch (status.status) {
    case 0:
      humanizedStatus = 'Stopped'
      break
    case 1:
      humanizedStatus = 'Queued to Verify Local Data'
      break
    case 2:
      humanizedStatus = 'Verifying Local Data'
      break
    case 3:
      humanizedStatus = 'Queued to Download'
      break
    case 4:
      humanizedStatus = 'Downloading'
      break
    case 5:
      humanizedStatus = 'Queued to Seed (Uploading)'
      break
    case 6:
      humanizedStatus = 'Seeding (Uploading)'
      break
  }

  return (
    <div className="bg-white p-6 rounded-3xl">
      <ol>
        <li className="text-xl break-all">{status.name}</li>
        <li className="mt-2 opacity-60">
          Progress: {Number(status.percentDone * 100).toFixed(2)}%
        </li>
        <li className="opacity-60">ETA: {humanizedEta}</li>
        <li className="opacity-60">Peers Connected: {status.peersConnected}</li>
        <li className="opacity-60">Status: {humanizedStatus}</li>
      </ol>
    </div>
  )
}

function TorrentTaskList() {
  const { data, error, isLoading } = useSWR(
    '/api/fetcher/status',
    (...args) => fetch(...args).then((res) => res.json()),
    { refreshInterval: 500 }
  )

  if (isLoading)
    return (
      <div className="flex items-center justify-center min-h-screen">
        Loading...
      </div>
    )
  if (error)
    return (
      <div className="flex items-center justify-center min-h-screen">
        Failed to get fetcher status
      </div>
    )
  if (!data.payload || data.payload?.length === 0)
    return (
      <div className="flex items-center justify-center min-h-screen">
        Idle (*/ω＼*)
      </div>
    )

  return (
    <div className="flex flex-col gap-5 mt-10 max-w-4xl">
      {data.payload?.map((status: any) => (
        <TorrentTaskEntry status={status} key={status.id}></TorrentTaskEntry>
      ))}
    </div>
  )
}

export default TorrentTaskList
