import { newsData } from '../constants/winners.constant';

export const useNewsApi = () => {
  return {
    newsData: newsData,
  };
};
