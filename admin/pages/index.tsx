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
} from '@chakra-ui/react'

import dayjs from 'dayjs'
import utc from 'dayjs/plugin/utc'
import { NewTask, Task } from '@/api/types'
import { TaskCard } from '@/components/TaskCard'
import { NewTaskModal } from '@/components/NewTaskModal'
import { AddIcon, SettingsIcon } from '@chakra-ui/icons'
import { SettingModal } from '@/components/SettingModal'
import Head from 'next/head'
dayjs.extend(utc)

export default function Home() {
  const mockData: Array<Task> = [
    {
      collection: '不當哥哥了！Test 1',
      config: {
        rss: 'https://nyaa.si/?page=rss&q=[NC-Raws]%20%E4%B8%8D%E7%95%B6%E5%93%A5%E5%93%A5%E4%BA%86%EF%BC%81%20Baha',
        cron: '35 23 * * 4',
        cronTimeZone: 'TZ2',
        maxDownload: 30,
        timeout: 30,
      },
      downloaded: 0,
      lastUpdate: '2023-01-17 11:23:32.8408691+00:00',
    },
    {
      collection: '不當哥哥了！Test 2',
      config: {
        rss: 'https://nyaa.si/?page=rss&q=[NC-Raws]%20%E4%B8%8D%E7%95%B6%E5%93%A5%E5%93%A5%E4%BA%86%EF%BC%81%20Baha',
        cron: '35 23 * * 4',
        cronTimeZone: 'TZ1',
        maxDownload: 30,
        timeout: 30,
      },
      downloaded: 0,
      lastUpdate: '2023-01-17 11:23:32.8408691+00:00',
    },
  ]

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

  const toast = useToast()

  function newTask(data: NewTask) {
    console.log('debug new:', data)
    toast({
      title: 'Added',
      status: 'success',
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
          {mockData.map((el) => (
            <TaskCard data={el} key={el.collection}></TaskCard>
          ))}
        </VStack>
      </Center>
    </>
  )
}
