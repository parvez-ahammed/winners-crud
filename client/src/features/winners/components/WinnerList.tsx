import React from 'react';

import { NativeSelectField, NativeSelectRoot } from '@/components/ui/native-select';
import { Box, HStack, Table, Text, Image } from '@chakra-ui/react';
import { Winner } from '../interface/winner.interface';
import { isSinglePersonGame } from '@/helpers/utility';
import nodata from '@/assets/images/no-data.svg';

interface WinnersListProps {
  gameName: string;
  winnerData: Winner[];
}
export const WinnersList: React.FC<WinnersListProps> = ({ gameName, winnerData }) => {
  return (
    <Box w="full" border={1} p={4} borderRadius="lg" borderWidth={2} boxShadow="sm">
      <HStack justify={'space-between'} mb={4}>
        <Text fontSize={'2xl'}>{gameName}</Text>
        <NativeSelectRoot w="auto">
          <NativeSelectField placeholder="Select Season" name="season" items={['2024', '2025']} />
        </NativeSelectRoot>
      </HStack>
      <Table.Root size="sm">
        <Table.Header>
          <Table.Row>
            <Table.ColumnHeader>Position</Table.ColumnHeader>
            {isSinglePersonGame(gameName) ? <Table.ColumnHeader>Name</Table.ColumnHeader> : <Table.ColumnHeader>Team Member 1</Table.ColumnHeader>}
            {isSinglePersonGame(gameName) ? undefined : <Table.Cell>Team Member 2</Table.Cell>}
          </Table.Row>
        </Table.Header>
        <Table.Body>
          {winnerData.length === 0 && (
            <Table.Row>
              <Table.Cell colSpan={isSinglePersonGame(gameName) ? 2 : 3} textAlign="center">
                <Box display="flex" justifyContent="center" alignItems="center" flexDirection="column">
                  <Image src={nodata} alt="No data" height="200px" />
                  <Text color={'gray.400'}>No winners to display</Text>
                </Box>
              </Table.Cell>
            </Table.Row>
          )}
          {winnerData.map((winners) => (
            <Table.Row key={winners.id}>
              <Table.Cell>{winners.position}</Table.Cell>
              <Table.Cell>{winners.teamMember1}</Table.Cell>
              {isSinglePersonGame(gameName) ? undefined : <Table.Cell>{winners.teamMember2}</Table.Cell>}
            </Table.Row>
          ))}
        </Table.Body>
      </Table.Root>
    </Box>
  );
};
