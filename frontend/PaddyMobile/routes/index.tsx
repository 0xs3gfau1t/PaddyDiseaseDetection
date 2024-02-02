import { NavigationContainer } from '@react-navigation/native';
import AuthRoutes from './AuthRoutes';
import NonAuthRoutes from './NonAuthRoutes';
import { useAuthContext } from '@/contexts/auth/auth-provider';
import LoadingScreen from '@/components/Loading';

export default function Root() {
  const auth = useAuthContext();
  if (auth.isFetching) return <LoadingScreen />;
  return (
    <NavigationContainer>{auth.token ? <AuthRoutes /> : <NonAuthRoutes />}</NavigationContainer>
  );
}
