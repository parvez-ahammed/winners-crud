import { Button, HStack, Spacer, Text } from '@chakra-ui/react';
import React from 'react';
import { FaEdit, FaTrash } from 'react-icons/fa';

interface SingleNewsProps {
  newsText: string;
  onEdit: () => void;
  onDelete: () => void;
}

export const SingleNews: React.FC<SingleNewsProps> = ({ newsText, onEdit, onDelete }) => {
  return (
    <HStack gap={4} p={4} bg="CardBg" borderRadius={8} border="1px solid" borderColor="blackAlpha.400" width="100%">
      <Text color="ButtonText" flex="1">
        {newsText}
      </Text>
      <Spacer />
      <Button onClick={onEdit} className="btn-edit">
        <FaEdit />
      </Button>
      <Button onClick={onDelete} className="btn-delete">
        <FaTrash />
      </Button>
    </HStack>
  );
};
