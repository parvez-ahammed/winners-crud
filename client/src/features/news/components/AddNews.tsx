import { VStack, Textarea, Button } from '@chakra-ui/react';

export const AddNews = () => {
  return (
    <VStack w="full" gap={4}>
      <Textarea placeholder="New News..." color="black" textDecor="GrayText" w="full" />
      <Button w="full">Post News</Button>
    </VStack>
  );
};
