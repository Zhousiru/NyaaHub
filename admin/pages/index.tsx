import {
  Card,
  CardBody,
  Center,
  Text,
  HStack,
  VStack,
  Box,
  Button,
  ButtonGroup,
  useDisclosure,
  useToast,
  Toast,
} from '@chakra-ui/react'

import dayjs from 'dayjs'
import utc from 'dayjs/plugin/utc'
import { NewTask, Task } from '@/api/types'
import { TaskCard } from '@/components/TaskCard'
import { NewTaskModal } from '@/components/NewTaskModal'
import { AddIcon, SettingsIcon } from '@chakra-ui/icons'
import { SettingModal } from '@/components/SettingModal'
import Head from 'next/head'
import { useEffect, useState } from 'react'
import { addTask, listTask } from '@/api/task'
dayjs.extend(utc)

export default function Home() {
  const toast = useToast()

  const [taskList, setTaskList] = useState<Array<Task>>()
  const [hasNext, sethasNext] = useState<boolean>(false)

  function refresh() {
    const loadingToast = toast({ title: 'Loading...', status: 'loading' })
    listTask(5)
      .then((el) => {
        sethasNext(el.hasNext)
        setTaskList(el.list)
      })
      .catch((err) => {
        toast({
          title: 'Failed to load data',
          description: err.message,
          status: 'error',
        })
      })
      .finally(() => {
        toast.close(loadingToast)
      })
  }

  useEffect(refresh, [toast])

  function loadMore() {
    if (!taskList) {
      return
    }

    const loadingToast = toast({ title: 'Loading...', status: 'loading' })
    listTask(5, taskList[taskList.length - 1].collection)
      .then((el) => {
        sethasNext(el.hasNext)
        setTaskList([...taskList, ...el.list])
      })
      .catch((err) => {
        toast({
          title: 'Failed to load data',
          description: err.message,
          status: 'error',
        })
      })
      .finally(() => {
        toast.close(loadingToast)
      })
  }

  const {
    isOpen: isNewOpen,
    onOpen: onNewOpen,
    onClose: onNewClose,
  } = useDisclosure()

  const {
    isOpen: isSettingOpen,
    onOpen: onSettingOpen,
    onClose: onSettingClose,
  } = useDisclosure()

  function newTask(data: NewTask) {
    const loadingToast = toast({ title: 'Loading...', status: 'loading' })
    addTask(data)
      .then(() => {
        toast({
          title: 'Added',
          status: 'success',
        })
      })
      .catch((err) => {
        toast({
          title: 'Failed to add task',
          description: err.message,
          status: 'error',
        })
      })
      .finally(() => {
        toast.close(loadingToast)
        refresh()
      })
  }

  return (
    <>
      <Head>
        <title>Scheduler Admin</title>
      </Head>

      <NewTaskModal
        isOpen={isNewOpen}
        onClose={onNewClose}
        onSubmit={newTask}
      ></NewTaskModal>
      <SettingModal
        isOpen={isSettingOpen}
        onClose={onSettingClose}
      ></SettingModal>

      <Center>
        <VStack
          spacing={4}
          my="10vh"
          width="clamp(350px, 80vw, 900px)"
          align="stretch"
        >
          <Card>
            <CardBody>
              <HStack>
                <Text fontWeight={100} fontSize="1.5rem" flex="1">
                  Scheduler Admin
                </Text>
                <Box>
                  <ButtonGroup>
                    <Button onClick={onSettingOpen}>
                      <SettingsIcon></SettingsIcon>
                    </Button>
                    <Button
                      colorScheme="blue"
                      onClick={onNewOpen}
                      leftIcon={<AddIcon></AddIcon>}
                      display={{ base: 'none', md: 'flex' }}
                    >
                      New Task
                    </Button>
                    <Button
                      colorScheme="blue"
                      onClick={onNewOpen}
                      display={{ base: 'flex', md: 'none' }}
                    >
                      <AddIcon></AddIcon>
                    </Button>
                  </ButtonGroup>
                </Box>
              </HStack>
            </CardBody>
          </Card>
          {taskList?.map((el) => (
            <TaskCard data={el} key={el.collection}></TaskCard>
          ))}
          {hasNext ? (
            <Button h="3rem" onClick={loadMore}>
              Load More
            </Button>
          ) : (
            <Center h="3rem" color="blackAlpha.300">
              — EOF —
            </Center>
          )}
        </VStack>
      </Center>
    </>
  )
}
