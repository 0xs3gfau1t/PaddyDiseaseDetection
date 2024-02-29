import LoggedInTabs from '@/components/LoggedInTabs';
import pages from '@/constants/screens';
import DashboardScreen from '@/screens/withSession/Dashboard';
import DetailScreen from '@/screens/withSession/Detail';
import HeatMap from '@/screens/withSession/HeatMap';
import ProfileScreen from '@/screens/withSession/Profile';
import StatsScreen from '@/screens/withSession/Stats';
import UploadScreen from '@/screens/withSession/Upload';
import { createBottomTabNavigator } from '@react-navigation/bottom-tabs';

const Tab = createBottomTabNavigator();

export default function AuthRoutes() {
  return (
    <Tab.Navigator
      screenOptions={{ headerShown: false }}
      tabBar={(props) => <LoggedInTabs {...props} />}
    >
      <Tab.Screen name={pages.dashboard} component={DashboardScreen} />
      <Tab.Screen name={pages.upload} component={UploadScreen} />
      <Tab.Screen name={pages.live} component={StatsScreen} />
      <Tab.Screen name={pages.profile} component={ProfileScreen} />
      <Tab.Screen name={pages.detail} component={DetailScreen} />
      <Tab.Screen name={pages.heatMap} component={HeatMap} />
    </Tab.Navigator>
  );
}
