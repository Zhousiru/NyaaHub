import dayjs from 'dayjs'
import useSWR from 'swr'
import prettyBytes from 'pretty-bytes'
import { PropsWithChildren } from 'react'

import duration from 'dayjs/plugin/duration'
import utc from 'dayjs/plugin/utc'
dayjs.extend(duration)
dayjs.extend(utc)

function humanizeStatus(statusCode: Number): string {
  let humanized = 'Unknown'

  switch (statusCode) {
    case 0:
      humanized = 'Stopped'
      break
    case 1:
      humanized = 'Queued to Verify Local Data'
      break
    case 2:
      humanized = 'Verifying Local Data'
      break
    case 3:
      humanized = 'Queued to Download'
      break
    case 4:
      humanized = 'Downloading'
      break
    case 5:
      humanized = 'Queued to Seed (Uploading)'
      break
    case 6:
      humanized = 'Seeding (Uploading)'
      break
  }

  return humanized
}

function SimplifyStatus(statusCode: Number): string {
  let simplified = 'Unknown'

  switch (statusCode) {
    case 0:
      simplified = 'Stopped'
      break
    case 1:
    case 2:
    case 3:
    case 4:
      simplified = 'Downloading'
      break
    case 5:
    case 6:
      simplified = 'Uploading'
      break
  }

  return simplified.toUpperCase()
}

function humanizeEta(eta: number): string {
  if (eta < 0) {
    return 'N/A'
  }
  if (eta === 0) {
    return `Soon`
  }

  const etaStr = dayjs.duration(eta, 'seconds').format('H,m,s')
  const [h, m, s] = etaStr.split(',')

  return (h ? `${h} h ` : '') + (m ? `${m} min ` : '') + `${s} s`
}

function humanizeProgress(percentDone: number): string {
  return Number(percentDone * 100).toFixed(2) + '%'
}

function Badge(props: PropsWithChildren<{ className: string }>) {
  return (
    <span
      className={`inline-block text-white rounded-lg py-1 px-2 ${props.className}`}
    >
      {props.children}
    </span>
  )
}

function GetStatusColor(statusCode: number): string {
  switch (statusCode) {
    case 0:
      return 'bg-orange-100'
    case 5:
      return 'bg-green-100'
  }

  return 'bg-blue-100'
}

function TorrentTaskEntry({ status }: { status: any }) {
  const progress = humanizeProgress(status.percentDone)

  function DetailList(props: PropsWithChildren) {
    return <li className={`text-gray-600`}>{props.children}</li>
  }

  return (
    <ol className="bg-white p-6 rounded-3xl overflow-hidden relative [&>li]:relative">
      <div
        className={`absolute inset-y-0 left-0 w-12 transition-all ${GetStatusColor(
          status.status
        )}`}
        style={{ width: progress }}
      ></div>
      <div className="absolute text-4xl bottom-0 right-0 p-5 text-gray-300/50">
        {SimplifyStatus(status.status)}
      </div>
      <li>
        <Badge className="bg-slate-500">{status.collection}</Badge>
      </li>
      <li className="text-xl break-all my-3">{status.name}</li>
      <li className="mb-3">
        <Badge className="bg-teal-700 text-xs">
          ETA: {humanizeEta(status.eta)}
        </Badge>
        <Badge className="bg-lime-700 ml-2 text-xs">
          DOWN: {prettyBytes(status.rateDownload)}/s
        </Badge>
        <Badge className="bg-amber-700 ml-2 text-xs">
          UP: {prettyBytes(status.rateUpload)}/s
        </Badge>
        <Badge className="bg-blue-700 ml-2 text-xs">
          SIZE: {prettyBytes(status.sizeWhenDone / 8)}
        </Badge>
      </li>
      <DetailList>Progress: {progress}</DetailList>
      <DetailList>Peers Connected: {status.peersConnected}</DetailList>
      <DetailList>
        Added Date:{' '}
        {dayjs.utc(status.addedDate).local().format('YYYY-MM-DD HH:mm:ss')}
      </DetailList>
      <DetailList>Status: {humanizeStatus(status.status)}</DetailList>
    </ol>
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
