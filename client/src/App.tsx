import { BrowserRouter, Route, Routes } from 'react-router-dom';

import './App.css';
import { useColorMode } from './components/ui/color-mode';
import { AddWinner } from './features/winners/components/AddWinner';
import { WinnersPage } from './pages/NewsPage';
import { AppProvider } from './providers/AppProvider';
import { RedirectToWinners } from './pages/RedirectToWinners';

function App() {
  const { setColorMode } = useColorMode();
  setColorMode('light');

  return (
    <AppProvider>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<RedirectToWinners />} />
          <Route path="/winners" element={<WinnersPage />} />
          <Route path="/winners/add" element={<AddWinner />} />
        </Routes>
      </BrowserRouter>
    </AppProvider>
  );
}

export default App;
