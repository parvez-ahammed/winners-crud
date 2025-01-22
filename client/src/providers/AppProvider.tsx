import { Provider } from '@/components/ui/provider';
interface AppProviderProps {
  children: React.ReactNode;
}
export const AppProvider = (props: AppProviderProps) => {
  return <Provider>{props.children}</Provider>;
};
