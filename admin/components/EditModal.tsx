import { TaskConfig } from '@/api/types'
import {
  Button,
  ButtonGroup,
  Modal,
  ModalBody,
  ModalContent,
  ModalFooter,
  ModalHeader,
  ModalOverlay,
} from '@chakra-ui/react'
import { ChangeEvent, useEffect, useState } from 'react'
import { TaskConfigForm } from './TaskConfigForm'

export function EditModal(props: {
  isOpen: boolean
  onClose: () => void
  config: TaskConfig
  onUpdate: (config: TaskConfig) => void
}) {
  const [config, setConfig] = useState<TaskConfig>(props.config)

  useEffect(() => {
    setConfig(props.config)
  }, [props.config, props.isOpen])

  function onChange(event: ChangeEvent<HTMLInputElement>) {
    setConfig({
      ...config,
      [event.target.name]: event.target.value,
    })
  }

  function update() {
    props.onUpdate(config)
    props.onClose()
  }

  function cancel() {
    props.onClose()
  }

  return (
    <Modal isOpen={props.isOpen} onClose={props.onClose}>
      <ModalOverlay />
      <ModalContent>
        <ModalHeader>Edit Task Config</ModalHeader>
        <ModalBody>
          <TaskConfigForm config={config} onChange={onChange}></TaskConfigForm>
        </ModalBody>
        <ModalFooter>
          <ButtonGroup>
            <Button colorScheme="blue" onClick={update}>
              Update
            </Button>
            <Button onClick={cancel}>Cancel</Button>
          </ButtonGroup>
        </ModalFooter>
      </ModalContent>
    </Modal>
  )
}
