import React from 'react';

import { NativeSelectField, NativeSelectRoot } from '@/components/ui/native-select';
import { Box, HStack, Table, Text, Image } from '@chakra-ui/react';
import { Winner } from '../interface/winner.interface';
import { isSinglePersonGame } from '@/helpers/utility';
import nodata from '@/assets/images/no-data.svg';
import { FaEdit, FaTrash } from 'react-icons/fa';

interface WinnersListProps {
  gameName: string;
  winnerData: Winner[];
}
export const WinnersList: React.FC<WinnersListProps> = ({ gameName, winnerData }) => {
  return (
    <Box w="full" border={1} p={4} borderRadius="lg" borderWidth={2} boxShadow="lg">
      <HStack justify={'space-between'} mb={4}>
        <Text fontSize={'2xl'}>{gameName}</Text>
        <NativeSelectRoot w="auto">
          <NativeSelectField placeholder="Select Season" name="season" items={['2024', '2025']} />
        </NativeSelectRoot>
      </HStack>
      <Table.Root size="sm" variant={'line'} interactive>
        <Table.Header>
          <Table.Row>
            <Table.ColumnHeader width="20%" textAlign="center">
              Position
            </Table.ColumnHeader>
            {isSinglePersonGame(gameName) ? (
              <Table.ColumnHeader width="70%" textAlign="center">
                Name
              </Table.ColumnHeader>
            ) : (
              <Table.ColumnHeader width="35%" textAlign="center">
                Team Member 1
              </Table.ColumnHeader>
            )}
            {isSinglePersonGame(gameName) ? undefined : (
              <Table.Cell width="35%" textAlign="center">
                Team Member 2
              </Table.Cell>
            )}
            <Table.ColumnHeader width="20%">Action</Table.ColumnHeader>
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
              <Table.Cell textAlign="center">{winners.position}</Table.Cell>
              <Table.Cell width="35%" textAlign="center">
                {winners.teamMember1}
              </Table.Cell>
              {isSinglePersonGame(gameName) ? undefined : (
                <Table.Cell width="35%" textAlign="center">
                  {winners.teamMember2}
                </Table.Cell>
              )}

              <Table.Cell width="20%" textAlign={'center'}>
                <HStack gap={4}>
                  <FaEdit />
                  <FaTrash />
                </HStack>
              </Table.Cell>
            </Table.Row>
          ))}
        </Table.Body>
      </Table.Root>
    </Box>
  );
};
