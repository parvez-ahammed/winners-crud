import './App.css';
import { useColorMode } from './components/ui/color-mode';
import { NewsPage } from './pages/NewsPage';
import { AppProvider } from './providers/AppProvider';

function App() {
  const { setColorMode } = useColorMode();
  setColorMode('light');

  return (
    <AppProvider>
      <NewsPage />
    </AppProvider>
  );
}

export default App;
