import { NewTask } from '@/api/types'
import {
  Button,
  ButtonGroup,
  Checkbox,
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
  onSubmit: (
    data: NewTask,
    downloadPrev: boolean,
    downloadOnly: boolean
  ) => void
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

  const [downloadOnly, setDownloadOnly] = useState<boolean>(false)
  const [downloadPrev, setDownloadPrev] = useState<boolean>(true)

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
    setDownloadOnly(false)
    setDownloadPrev(true)
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
    props.onSubmit(data, downloadPrev, downloadOnly)
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
          ></Input>
          <Checkbox
            my="0.5rem"
            isChecked={downloadOnly}
            onChange={(event) => setDownloadOnly(event.target.checked)}
          >
            Download only
          </Checkbox>
          {downloadOnly ? (
            <>
              <FormLabel>RSS URL</FormLabel>
              <Input
                value={data.config.rss}
                name="rss"
                onChange={(event) => onConfigChange(event)}
              />
            </>
          ) : (
            <>
              <TaskConfigForm
                config={data.config}
                onChange={onConfigChange}
              ></TaskConfigForm>
              <Checkbox
                mt="0.5rem"
                isChecked={downloadPrev}
                onChange={(event) => setDownloadPrev(event.target.checked)}
              >
                Download previous items
              </Checkbox>
            </>
          )}
        </ModalBody>
        <ModalFooter>
          <ButtonGroup>
            <Button colorScheme="blue" onClick={submit}>
              Submit
            </Button>
            <Button onClick={cancel}>Cancel</Button>
          </ButtonGroup>
        </ModalFooter>
      </ModalContent>
    </Modal>
  )
}
