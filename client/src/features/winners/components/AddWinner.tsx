import React, { useState, ChangeEvent, FormEvent } from 'react';
import { Box, Button, Input, Stack } from '@chakra-ui/react';
import { Field } from '@/components/ui/field';
import { Fieldset } from '@chakra-ui/react';
import { NativeSelectField, NativeSelectRoot } from '@/components/ui/native-select';
import { Winner } from '@/features/winners/interface/winner.interface';
import { dummyWinners } from '../constants/winners.constant';

export const AddWinner: React.FC = () => {
  const [selectedGame, setSelectedGame] = useState<string>('');
  const [formData, setFormData] = useState<Winner>({
    season: '',
    game: '',
    position: '',
    teamMember1: '',
    teamMember2: '',
  });

  const handleGameChange = (value: string) => {
    setSelectedGame(value);
    setFormData((prev) => ({
      ...prev,
      game: value,
      name: '',
      teamMember1: '',
      teamMember2: '',
    }));
  };

  const handleInputChange = (e: ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData((prev) => ({ ...prev, [name]: value }));
  };

  const handleSubmit = (e: FormEvent) => {
    e.preventDefault();
    console.log('Form Data:', formData);
    formData.id = (dummyWinners.length + 1).toString();
    dummyWinners.push(formData);
    console.log('Winners:', dummyWinners);

    setFormData({
      season: '',
      game: '',
      position: '',
      teamMember1: '',
      teamMember2: '',
    });

    setSelectedGame('');

    alert('Winner added successfully');
  };

  return (
    <Box w="full" border={1} p={4} borderRadius="lg" borderWidth={2} boxShadow="lg">
      <form onSubmit={handleSubmit} style={{ width: '100%' }}>
        <Fieldset.Root size="lg" maxW="full" w="full">
          <Stack>
            <Fieldset.Legend color={'black'}>Winners Details</Fieldset.Legend>
            <Fieldset.HelperText>Please provide details about the winners</Fieldset.HelperText>
          </Stack>

          <Fieldset.Content>
            <Field label="Season" required>
              <NativeSelectRoot>
                <NativeSelectField
                  placeholder="Select Season"
                  name="season"
                  items={['2024', '2025']}
                  onChange={(e) => setFormData((prev) => ({ ...prev, season: e.target.value }))}
                />
              </NativeSelectRoot>
            </Field>
            <Field label="Game name" required>
              <NativeSelectRoot>
                <NativeSelectField
                  name="game"
                  placeholder="Select Game"
                  items={['Chess', 'UNO', 'FoosBall', 'Table Tennis']}
                  onChange={(e) => handleGameChange(e.target.value)}
                />
              </NativeSelectRoot>
            </Field>
            <Field label="Position" required>
              <NativeSelectRoot>
                <NativeSelectField
                  placeholder="Select Position"
                  name="position"
                  items={['Champion', '1st Runners Up', '2nd Runners Up']}
                  onChange={(e: ChangeEvent<HTMLSelectElement>) => setFormData((prev) => ({ ...prev, position: e.target.value }))}
                />
              </NativeSelectRoot>
            </Field>

            {['Chess', 'UNO'].includes(selectedGame) ? (
              <Field label="Name" required>
                <Input name="name" value={formData.teamMember1} onChange={handleInputChange} />
              </Field>
            ) : (
              <Stack direction="row" gap={4}>
                <Field label="Team Member 1" required>
                  <Input name="teamMember1" value={formData.teamMember1} onChange={handleInputChange} />
                </Field>
                <Field label="Team Member 2" required>
                  <Input name="teamMember2" value={formData.teamMember2} onChange={handleInputChange} />
                </Field>
              </Stack>
            )}
          </Fieldset.Content>

          <Button type="submit" alignSelf="flex-end" color={'white'} bg={'black'} _hover={{ bg: 'blue.500' }}>
            Submit
          </Button>
        </Fieldset.Root>
      </form>
    </Box>
  );
};
