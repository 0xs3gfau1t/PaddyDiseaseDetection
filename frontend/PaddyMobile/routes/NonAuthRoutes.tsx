import { createNativeStackNavigator } from "@react-navigation/native-stack";
import pages from "../constants/screens";
import LoginScreen from "../screens/Login";
import SignupScreen from "../screens/Signup";
import LandingScreen from "../screens/Landing";

const Stack = createNativeStackNavigator();
export default function NonAuthRoutes() {
  return (
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
  );
}
