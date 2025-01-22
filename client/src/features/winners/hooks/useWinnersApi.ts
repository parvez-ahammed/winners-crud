import { dummyWinners } from '@/features/winners/constants/winners.constant';
export const useWinnersApi = () => {
  const unoWinners = dummyWinners.filter((winner) => winner.game === 'UNO');
  const chessWinners = dummyWinners.filter((winner) => winner.game === 'Chess');
  const foosBallWinner = dummyWinners.filter((winner) => winner.game === 'FoosBall');
  const tableTennisWinners = dummyWinners.filter((winner) => winner.game === 'Table Tennis');
  return { unoWinners, chessWinners, foosBallWinner, tableTennisWinners };
};
