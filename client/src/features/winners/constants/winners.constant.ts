import { Winner } from '@/features/winners/interface/winner.interface';

export const dummyWinners: Winner[] = [
  {
    id: '1',
    season: '2024',
    game: 'Chess',
    position: 'Champion',
    teamMember1: 'Alice',
    teamMember2: 'Bob',
  },
  {
    id: '2',
    season: '2024',
    game: 'UNO',
    position: '1st Runners Up',
    teamMember1: 'Bob',
    teamMember2: 'Charlie',
  },
  {
    id: '3',
    season: '2025',
    game: 'FoosBall',
    position: '2nd Runners Up',
    teamMember1: 'Charlie',
    teamMember2: 'Dave',
  },
];
