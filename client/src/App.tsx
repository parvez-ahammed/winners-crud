import { Box, VStack } from '@chakra-ui/react';
import './App.css';
import { AddNews } from './components/custom/AddNews';
import { useColorMode } from './components/ui/color-mode';
import NewsList from './components/custom/NewsList';
interface NewsItem {
  id: number;
  title: string;
  description: string;
}

// eslint-disable-next-line react-refresh/only-export-components
export const newsData: NewsItem[] = [
  { id: 1, title: 'News Title 1', description: 'Description for news 1' },
  { id: 2, title: 'News Title 2', description: 'Description for news 2' },
  { id: 3, title: 'News Title 3', description: 'Description for news 3' },
];
function App() {
  const { setColorMode } = useColorMode();
  setColorMode('light');

  return (
    <Box w="full" h="full" gap={4}>
      <VStack gap={4} width={{ base: '90%', md: '80%', lg: '60%' }} margin="auto">
        <AddNews />
        <NewsList />
      </VStack>
    </Box>
  );
}

export default App;
