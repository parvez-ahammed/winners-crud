import React from 'react';

import { Table, Box, Text, HStack } from '@chakra-ui/react';
import { NativeSelectField, NativeSelectRoot } from '@/components/ui/native-select';
import { dummyWinners } from '../constants/winners.constant';

const NewsList: React.FC = () => {
  const gameName = 'UNO';
  return (
    <Box w="full" border={1} p={4} borderRadius="lg" borderWidth={2} boxShadow="sm">
      <HStack justify={'space-between'} mb={4}>
        <Text fontSize={'2xl'}>Hello World</Text>
        <NativeSelectRoot w="auto">
          <NativeSelectField placeholder="Select Season" name="season" items={['2024', '2025']} />
        </NativeSelectRoot>
      </HStack>
      <Table.Root size="sm">
        <Table.Header>
          <Table.Row>
            <Table.ColumnHeader>Position</Table.ColumnHeader>
            {gameName === 'UNO' || gameName === 'Chess' ? (
              <Table.ColumnHeader>Name</Table.ColumnHeader>
            ) : (
              <Table.ColumnHeader>Team Member 1</Table.ColumnHeader>
            )}
            {gameName !== 'UNO' && gameName !== 'Chess' && <Table.ColumnHeader>Team Member 2</Table.ColumnHeader>}
          </Table.Row>
        </Table.Header>
        <Table.Body>
          {dummyWinners.map((winners) => (
            <Table.Row key={winners.id}>
              <Table.Cell>{winners.position}</Table.Cell>
              <Table.Cell>{winners.teamMember1}</Table.Cell>
            </Table.Row>
          ))}
        </Table.Body>
      </Table.Root>
    </Box>
  );
};

export default NewsList;
