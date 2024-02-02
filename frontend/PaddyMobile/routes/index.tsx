import { NavigationContainer } from "@react-navigation/native";
import LoadingScreen from "../components/Loading";
import { useAuthContext } from "../contexts/auth/auth-provider";
import AuthRoutes from "./AuthRoutes";
import NonAuthRoutes from "./NonAuthRoutes";

export default function Root() {
  const auth = useAuthContext();
  if (auth.isFetching) return <LoadingScreen />;
  return (
    <NavigationContainer>
      {auth.token ? <AuthRoutes /> : <NonAuthRoutes />}
    </NavigationContainer>
  );
}
