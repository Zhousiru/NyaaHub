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
  useToast,
} from '@chakra-ui/react'
import { useEffect, useState } from 'react'

interface Setting {
  apiUrl: string
  token: string
}

export function SettingModal(props: { isOpen: boolean; onClose: () => void }) {
  const [setting, setSetting] = useState<Setting>({ apiUrl: '', token: '' })

  useEffect(() => {
    setSetting({
      apiUrl: localStorage.getItem('apiUrl') ?? '',
      token: localStorage.getItem('token') ?? '',
    })
  }, [props.isOpen])

  const toast = useToast()

  function update() {
    localStorage.setItem('apiUrl', setting.apiUrl)
    localStorage.setItem('token', setting.token)
    props.onClose()
    toast({
      title: 'Updated',
      status: 'success',
    })
  }

  function cancel() {
    props.onClose()
  }

  return (
    <Modal isOpen={props.isOpen} onClose={props.onClose}>
      <ModalOverlay />
      <ModalContent>
        <ModalHeader>Setting</ModalHeader>
        <ModalBody>
          <FormLabel>Scheduler API URL</FormLabel>
          <Input
            value={setting.apiUrl}
            onChange={(event) =>
              setSetting({ ...setting, apiUrl: event.target.value })
            }
          />
          <FormLabel mt={1}>API token</FormLabel>
          <Input
            value={setting.token}
            onChange={(event) =>
              setSetting({ ...setting, token: event.target.value })
            }
          />
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
