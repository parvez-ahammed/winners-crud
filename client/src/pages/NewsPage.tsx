import { AddNews } from '@/features/news/components/AddNews';
import NewsList from '@/features/news/components/NewsList';
import { Box, VStack } from '@chakra-ui/react';

export const NewsPage = () => {
  return (
    <Box w="full" h="full" gap={4}>
      <VStack gap={4} width={{ base: '90%', md: '80%', lg: '60%' }} margin="auto">
        <AddNews />
        <NewsList />
      </VStack>
    </Box>
  );
};
