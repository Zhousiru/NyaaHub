import { NewTask } from '@/api/types'
import {
  Button,
  ButtonGroup,
  FormLabel,
  Input,
  Modal,
  ModalBody,
  ModalContent,
  ModalFooter,
  ModalHeader,
  ModalOverlay,
} from '@chakra-ui/react'
import { ChangeEvent, useEffect, useState } from 'react'
import { TaskConfigForm } from './TaskConfigForm'

export function NewTaskModal(props: {
  isOpen: boolean
  onClose: () => void
  onSubmit: (data: NewTask) => void
}) {
  const [data, setData] = useState<NewTask>({
    collection: '',
    config: {
      rss: '',
      cron: '',
      cronTimeZone: '',
      maxDownload: 30,
      timeout: 30,
    },
  })

  useEffect(() => {
    setData({
      collection: '',
      config: {
        rss: '',
        cron: '',
        cronTimeZone: '',
        maxDownload: 30,
        timeout: 30,
      },
    })
  }, [props.isOpen])

  function onConfigChange(event: ChangeEvent<HTMLInputElement>) {
    setData({
      ...data,
      config: {
        ...data.config,
        [event.target.name]: event.target.value,
      },
    })
  }

  function submit() {
    props.onSubmit(data)
    props.onClose()
  }

  function cancel() {
    props.onClose()
  }

  return (
    <Modal isOpen={props.isOpen} onClose={props.onClose}>
      <ModalOverlay />
      <ModalContent>
        <ModalHeader>New Task</ModalHeader>
        <ModalBody>
          <FormLabel>Collection</FormLabel>
          <Input
            onChange={(event) =>
              setData({ ...data, collection: event.target.value })
            }
            mb={1}
          ></Input>
          <TaskConfigForm
            config={data.config}
            onChange={onConfigChange}
          ></TaskConfigForm>
        </ModalBody>
        <ModalFooter>
          <ButtonGroup>
            <Button colorScheme="blue" onClick={submit}>
              Update
            </Button>
            <Button onClick={cancel}>Cancel</Button>
          </ButtonGroup>
        </ModalFooter>
      </ModalContent>
    </Modal>
  )
}
