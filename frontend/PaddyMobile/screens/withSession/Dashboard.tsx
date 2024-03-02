import AreaCard from '@/components/dashboard/Area';
import OptionCard from '@/components/dashboard/Option';
import WelcomeCard from '@/components/dashboard/WelcomeCard';
import { ScrollView, StyleSheet, Text, View } from 'react-native';
import FontAwesome6 from 'react-native-vector-icons/FontAwesome6';
import Entypo from 'react-native-vector-icons/Entypo';
import { useGetDashboard } from '@/api/dashboard';
import { useContext } from 'react';
import { AuthContext } from '@/contexts/auth/auth-provider';
import { ActivityIndicator } from 'react-native-paper';

export default function DashboardScreen() {
  const { token, removeToken } = useContext(AuthContext);
  const { data, loading } = useGetDashboard(token);

  if (loading)
    return (
      <View
        style={{ width: '100%', height: '100%', justifyContent: 'center', alignItems: 'center' }}
      >
        <Text>Loading Dashboard Data...</Text>
        <ActivityIndicator />
      </View>
    );

  if (!data) {
    removeToken();
    return (
      <View
        style={{ width: '100%', height: '100%', justifyContent: 'center', alignItems: 'center' }}
      >
        <Text>Oops something went wrong...</Text>
      </View>
    );
  }

  return (
    <ScrollView>
      <View style={styles.container}>
        <WelcomeCard {...data} />
        <AreaCard {...data} />
        <View>
          <Text style={{ fontSize: 20, fontWeight: 'bold' }}>Services</Text>
          <View style={styles.optionContainer}>
            <OptionCard
              icon={<FontAwesome6 name='user-doctor' size={40} />}
              text={`${data.expertsOnline}/${data.expertsTotal} online`}
            />
            <OptionCard icon={<FontAwesome6 name='shop' size={40} />} text='Store' />
            <OptionCard icon={<Entypo name='video' size={40} />} text='Reels' />
          </View>
        </View>
      </View>
    </ScrollView>
  );
}

const styles = StyleSheet.create({
  container: {
    paddingTop: 50,
    paddingHorizontal: 20,
    gap: 10,
  },
  optionContainer: {
    flex: 1,
    flexDirection: 'row',
    flexWrap: 'wrap',
    paddingVertical: 10,
    gap: 5,
  },
});
