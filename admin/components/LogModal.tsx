import {
  Button,
  ButtonGroup,
  Modal,
  ModalBody,
  ModalContent,
  ModalFooter,
  ModalHeader,
  ModalOverlay,
  Textarea,
} from '@chakra-ui/react'

export function LogModal(props: {
  isOpen: boolean
  onClose: () => void
  log: string
}) {
  return (
    <Modal isOpen={props.isOpen} onClose={props.onClose}>
      <ModalOverlay />
      <ModalContent maxWidth={900}>
        <ModalHeader>View Log</ModalHeader>
        <ModalBody>
          <Textarea
            value={props.log}
            onChange={(event) => {
              event.preventDefault()
            }}
            h="50vh"
          ></Textarea>
        </ModalBody>
        <ModalFooter>
          <ButtonGroup>
            <Button colorScheme="blue" onClick={props.onClose}>
              Close
            </Button>
          </ButtonGroup>
        </ModalFooter>
      </ModalContent>
    </Modal>
  )
}
