import { newsData } from '../constants/news-data.constant';

export const useNewsApi = () => {
  return {
    newsData: newsData,
  };
};
