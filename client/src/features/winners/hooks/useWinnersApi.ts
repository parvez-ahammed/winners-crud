import { winnersApi } from '../api/winners.api';

export const useWinnersApi = () => {
  const winners = winnersApi();
  const unoWinners = winners.filter((winner) => winner.game === 'UNO');
  const chessWinners = winners.filter((winner) => winner.game === 'Chess');
  const foosBallWinner = winners.filter((winner) => winner.game === 'FoosBall');
  const tableTennisWinners = winners.filter((winner) => winner.game === 'Table Tennis');
  return { unoWinners, chessWinners, foosBallWinner, tableTennisWinners };
};
