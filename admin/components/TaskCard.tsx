import { Task, TaskConfig } from '@/api/types'
import {
  DeleteIcon,
  EditIcon,
  ExternalLinkIcon,
  ViewIcon,
} from '@chakra-ui/icons'
import {
  Card,
  CardHeader,
  Heading,
  CardBody,
  List,
  ListItem,
  Code,
  CardFooter,
  ButtonGroup,
  Button,
  Text,
  Link,
  useDisclosure,
  useToast,
} from '@chakra-ui/react'

import dayjs from 'dayjs'
import utc from 'dayjs/plugin/utc'
import { useState } from 'react'
import { EditModal } from './EditModal'
dayjs.extend(utc)

export function TaskCard(props: { data: Task }) {
  const lastUpdate = dayjs
    .utc(props.data.lastUpdate)
    .local()
    .format('YYYY-MM-DD HH:mm:ss')

  const [data, setData] = useState(props.data)
  const {
    isOpen: isEditOpen,
    onOpen: onEditOpen,
    onClose: onEditClose,
  } = useDisclosure()
  const toast = useToast()

  function deleteTask() {
    console.log('debug: delete', props.data.collection)
    toast({
      title: 'Deleted',
      status: 'success',
    })
  }

  function onEditUpdate(config: TaskConfig) {
    setData({
      ...data,
      config,
    })
    console.log('debug: update', props.data.collection)
    toast({
      title: 'Updated',
      status: 'success',
    })
  }

  return (
    <>
      <Card>
        <CardHeader>
          <Heading size="md" wordBreak="break-all">
            {data.collection}
          </Heading>
          <Text color="blackAlpha.600">Last updated on {lastUpdate}</Text>
        </CardHeader>
        <CardBody>
          <List>
            <ListItem>
              <Text as="b">RSS URL: </Text>
              <Link href={data.config.rss} isExternal>
                Open
                <ExternalLinkIcon mx="2px" />
              </Link>
            </ListItem>
            <ListItem>
              <Text as="b">Cron: </Text>
              <Code>{data.config.cron}</Code> ({data.config.cronTimeZone})
            </ListItem>
            <ListItem>
              <Text as="b">Download: </Text>
              {data.downloaded} / {data.config.maxDownload}
            </ListItem>
            <ListItem>
              <Text as="b">Timeout: </Text>
              {data.config.timeout} day(s)
            </ListItem>
          </List>
        </CardBody>
        <CardFooter>
          <ButtonGroup>
            <Button onClick={onEditOpen} leftIcon={<EditIcon></EditIcon>}>
              Edit
            </Button>
            <Button leftIcon={<ViewIcon></ViewIcon>}>View Log</Button>
            <Button
              colorScheme="red"
              variant="outline"
              onClick={deleteTask}
              ml={2}
              leftIcon={<DeleteIcon></DeleteIcon>}
            >
              Delete
            </Button>
          </ButtonGroup>
        </CardFooter>
      </Card>
      <EditModal
        isOpen={isEditOpen}
        onClose={onEditClose}
        config={data.config}
        onUpdate={onEditUpdate}
      ></EditModal>
    </>
  )
}
