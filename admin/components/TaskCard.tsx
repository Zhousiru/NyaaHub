import { getLog, removeTask, updateTask } from '@/api'
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
import { LogModal } from './LogModal'
dayjs.extend(utc)

export function TaskCard(props: { data: Task }) {
  const lastUpdate = dayjs
    .utc(props.data.lastUpdate)
    .local()
    .format('YYYY-MM-DD HH:mm:ss')

  const [deleted, setDeleted] = useState<boolean>(false)
  const [data, setData] = useState<Task>(props.data)
  const [log, setLog] = useState<string>('')

  const {
    isOpen: isEditOpen,
    onOpen: onEditOpen,
    onClose: onEditClose,
  } = useDisclosure()
  const {
    isOpen: isLogOpen,
    onOpen: onLogOpen,
    onClose: onLogClose,
  } = useDisclosure()

  const toast = useToast()

  function deleteTask() {
    const loadingToast = toast({ title: 'Loading...', status: 'loading' })
    removeTask(props.data.collection)
      .then(() => {
        setDeleted(true)
        toast({
          title: 'Deleted',
          status: 'success',
        })
      })
      .catch((err) => {
        toast({
          title: 'Failed to delete task',
          description: err.message,
          status: 'error',
        })
      })
      .finally(() => {
        toast.close(loadingToast)
      })
  }

  function onEditUpdate(config: TaskConfig) {
    const loadingToast = toast({ title: 'Loading...', status: 'loading' })
    updateTask(props.data.collection, config)
      .then(() => {
        toast({
          title: 'Updated',
          status: 'success',
        })
      })
      .catch((err) => {
        toast({
          title: 'Failed to update task',
          description: err.message,
          status: 'error',
        })
      })
      .finally(() => {
        toast.close(loadingToast)
      })
    setData({
      ...data,
      config,
    })
  }

  function viewLog() {
    const loadingToast = toast({ title: 'Loading...', status: 'loading' })
    getLog(props.data.collection)
      .then((data) => {
        setLog(data)
      })
      .catch((err) => {
        toast({
          title: 'Failed to load log',
          description: err.message,
          status: 'error',
        })
      })
      .finally(() => {
        toast.close(loadingToast)
      })
    onLogOpen()
  }

  if (deleted) {
    return <></>
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
              {data.download} / {data.config.maxDownload}
            </ListItem>
            <ListItem>
              <Text as="b">Timeout: </Text>
              {data.config.timeout} day(s)
            </ListItem>
          </List>
        </CardBody>
        <CardFooter>
          <ButtonGroup display={{ base: 'none', md: 'flex' }}>
            <Button onClick={onEditOpen} leftIcon={<EditIcon></EditIcon>}>
              Edit
            </Button>
            <Button leftIcon={<ViewIcon></ViewIcon>} onClick={viewLog}>
              View Log
            </Button>
            <Button
              colorScheme="red"
              variant="outline"
              onClick={deleteTask}
              leftIcon={<DeleteIcon></DeleteIcon>}
            >
              Delete
            </Button>
          </ButtonGroup>
          <ButtonGroup display={{ base: 'flex', md: 'none' }}>
            <Button onClick={onEditOpen}>
              <EditIcon></EditIcon>
            </Button>
            <Button onClick={onLogOpen}>
              <ViewIcon></ViewIcon>
            </Button>
            <Button colorScheme="red" variant="outline" onClick={deleteTask}>
              <DeleteIcon></DeleteIcon>
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
      <LogModal isOpen={isLogOpen} onClose={onLogClose} log={log}></LogModal>
    </>
  )
}
