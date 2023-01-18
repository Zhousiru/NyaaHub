import { TaskConfig } from '@/api/types'
import { FormLabel, Input, HStack } from '@chakra-ui/react'
import { ChangeEvent } from 'react'

export function TaskConfigForm(props: {
  config: TaskConfig
  onChange: (e: ChangeEvent<HTMLInputElement>) => void
}) {
  return (
    <>
      <FormLabel>RSS URL</FormLabel>
      <Input
        value={props.config.rss}
        name="rss"
        onChange={(event) => props.onChange(event)}
      />
      <FormLabel mt="1">Cron and time zone</FormLabel>
      <HStack>
        <Input
          placeholder="Cron expression"
          value={props.config.cron}
          name="cron"
          onChange={(event) => props.onChange(event)}
        />
        <Input
          w="60%"
          placeholder="Time zone"
          value={props.config.cronTimeZone}
          name="cronTimeZone"
          onChange={(event) => props.onChange(event)}
        />
      </HStack>
      <FormLabel mt="1">Max download</FormLabel>
      <Input
        type="number"
        value={props.config.maxDownload}
        name="maxDownload"
        onChange={(event) => props.onChange(event)}
      />
      <FormLabel mt="1">Timeout (days)</FormLabel>
      <Input
        type="number"
        value={props.config.timeout}
        name="timeout"
        onChange={(event) => props.onChange(event)}
      />
    </>
  )
}
