import { NavigationContainer } from '@react-navigation/native';
import AuthRoutes from './AuthRoutes';
import NonAuthRoutes from './NonAuthRoutes';
import { AuthContext } from '@/contexts/auth/auth-provider';
import LoadingScreen from '@/components/Loading';
import { EventProvider } from 'react-native-outside-press';
import { useContext } from 'react';

export default function Root() {
  const { isFetching, token } = useContext(AuthContext);
  if (isFetching) return <LoadingScreen />;
  return (
    <EventProvider>
      <NavigationContainer>{token ? <AuthRoutes /> : <NonAuthRoutes />}</NavigationContainer>
    </EventProvider>
  );
}
