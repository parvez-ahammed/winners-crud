import { Box, Button, HStack } from '@chakra-ui/react';
import './App.css';
import { useColorMode } from './components/ui/color-mode';

function App() {
  
  const {setColorMode} = useColorMode();
  setColorMode('light');

  return (
    <Box width="100vw" height="100vh">
    <HStack>
      <Button>Click me</Button>
      <Button>Click me</Button>
    </HStack>
    </Box>
  )
}

export default App
