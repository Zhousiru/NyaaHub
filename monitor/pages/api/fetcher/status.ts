import axios from 'axios'
import type { NextApiRequest, NextApiResponse } from 'next'

const fetcherApiUrl = process.env.FETCHER_API.split('#')[0]
const fetcherApiToken = process.env.FETCHER_API.split('#')[1]

export default async function handler(
  req: NextApiRequest,
  res: NextApiResponse
) {
  const resp = await axios.get(`${fetcherApiUrl}/status`, {
    params: {
      token: fetcherApiToken,
    },
  })

  res.status(200).json(resp.data)
}
