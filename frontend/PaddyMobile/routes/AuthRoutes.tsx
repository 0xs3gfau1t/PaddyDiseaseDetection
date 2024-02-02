import { createNativeStackNavigator } from "@react-navigation/native-stack";
import ProfileScreen from "../screens/Profile";
import pages from "../constants/screens";

const Stack = createNativeStackNavigator();

export default function AuthRoutes() {
  return (
    <Stack.Navigator>
      <Stack.Screen name={pages.profile} component={ProfileScreen} />
    </Stack.Navigator>
  );
}
