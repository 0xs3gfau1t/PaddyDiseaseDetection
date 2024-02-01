import { NavigationContainer } from "@react-navigation/native";
import { createNativeStackNavigator } from "@react-navigation/native-stack";
import LandingScreen from "./screens/Landing";
import pages from "./constants/screens";
import SignupScreen from "./screens/Signup";
import LoginScreen from "./screens/Login";
import ProfileScreen from "./screens/Profile";
import useAuthContext from "./hooks/authContext";
import LoadingScreen from "./components/Loading";

const Stack = createNativeStackNavigator();

export default function App() {
  const { tokenState } = useAuthContext();
  if (tokenState.fetching) return <LoadingScreen />;
  else if (!tokenState.token)
    return (
      <NavigationContainer>
        <Stack.Navigator>
          <Stack.Screen
            name={pages.landing}
            component={LandingScreen}
            options={{ headerShown: false }}
          />
          <Stack.Screen
            name={pages.signup}
            component={SignupScreen}
            options={{ headerShown: false }}
          />
          <Stack.Screen
            name={pages.login}
            component={LoginScreen}
            options={{ headerShown: false }}
          />
        </Stack.Navigator>
      </NavigationContainer>
    );

  return (
    <NavigationContainer>
      <Stack.Navigator>
        <Stack.Screen name={pages.profile} component={ProfileScreen} />
      </Stack.Navigator>
    </NavigationContainer>
  );
}
