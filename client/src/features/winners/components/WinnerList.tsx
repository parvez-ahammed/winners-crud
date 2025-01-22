import React from 'react';

import { SimpleGrid } from '@chakra-ui/react';
import { SingleNews } from '@/features/winners/components/SingleWinner';
import { newsData } from '@/features/winners/constants/news-data.constant';

const NewsList: React.FC = () => {
  return (
    <SimpleGrid columns={{ sm: 1, md: 1, lg: 1 }} gap={5} width={'100%'}>
      {newsData.map((news) => (
        <SingleNews newsText={news.title} onEdit={() => {}} onDelete={() => {}} />
      ))}
    </SimpleGrid>
  );
};

export default NewsList;
