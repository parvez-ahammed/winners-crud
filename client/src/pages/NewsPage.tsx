import { WinnersList } from '@/features/winners/components/WinnerList';
import { useWinnersApi } from '@/features/winners/hooks/useWinnersApi';
import { Box, Button, VStack } from '@chakra-ui/react';

export const NewsPage = () => {
  const { unoWinners, chessWinners, foosBallWinner, tableTennisWinners } = useWinnersApi();
  return (
    <Box w="full" h="full" gap={4}>
      <VStack gap={4} width={{ base: '90%', md: '80%', lg: '60%' }} margin="auto">
        {/* Button to navigate to /winners/add */}
        <Box w="full" textAlign="right">
          <Button colorScheme="black" w="full" onClick={() => (window.location.href = '/winners/add')}>
            Add Winner
          </Button>
        </Box>
        {/* Winners list */}
        <WinnersList gameName="FoosBall" winnerData={foosBallWinner} />
        <WinnersList gameName="UNO" winnerData={unoWinners} />
        <WinnersList gameName="Chess" winnerData={chessWinners} />
        <WinnersList gameName="Table Tennis" winnerData={tableTennisWinners} />
      </VStack>
    </Box>
  );
};
