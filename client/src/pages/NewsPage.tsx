import { AddNews } from '@/features/winners/components/AddWinner';
import WinnersList from '@/features/winners/components/WinnerList';
import { Box, VStack } from '@chakra-ui/react';

export const NewsPage = () => {
  return (
    <Box w="full" h="full" gap={4}>
      <VStack gap={4} width={{ base: '90%', md: '80%', lg: '60%' }} margin="auto">
        <Box w="full" border={1} p={4} mb={4} borderRadius="lg" borderWidth={2} boxShadow="sm">
          <AddNews />
        </Box>
        <WinnersList />
      </VStack>
    </Box>
  );
};
