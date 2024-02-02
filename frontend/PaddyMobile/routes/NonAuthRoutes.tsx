import { createNativeStackNavigator } from '@react-navigation/native-stack';
import pages from '../constants/screens';
import LandingScreen from '@/screens/withoutSession/Landing';
import SignupScreen from '@/screens/withoutSession/Signup';
import LoginScreen from '@/screens/withoutSession/Login';

const Stack = createNativeStackNavigator();
export default function NonAuthRoutes() {
  return (
    <Stack.Navigator>
      <Stack.Screen
        name={pages.landing}
        component={LandingScreen}
        options={{ headerShown: false }}
      />
      <Stack.Screen name={pages.signup} component={SignupScreen} options={{ headerShown: false }} />
      <Stack.Screen name={pages.login} component={LoginScreen} options={{ headerShown: false }} />
    </Stack.Navigator>
  );
}
