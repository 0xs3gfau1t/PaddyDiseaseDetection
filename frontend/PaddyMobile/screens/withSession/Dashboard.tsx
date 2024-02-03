import AreaCard from '@/components/dashboard/Area';
import OptionCard from '@/components/dashboard/Option';
import WelcomeCard from '@/components/dashboard/WelcomeCard';
import { ScrollView, StyleSheet, Text, View } from 'react-native';
import FontAwesome6 from 'react-native-vector-icons/FontAwesome6';
import Entypo from 'react-native-vector-icons/Entypo';

const mockedData = {
  userName: 'ThiccTsunade',
  userSubmissions: 10,
  userDiseaseDetected: 2,
  creditsRemaining: 10,
  areaSubmissions: 102,
  areaDiseaseDetected: 20,
  expertsOnline: 10,
  expertsTotal: 20,
};

export default function DashboardScreen() {
  return (
    <ScrollView>
      <View style={styles.container}>
        <WelcomeCard {...mockedData} />
        <AreaCard {...mockedData} />
        <View>
          <Text style={{ fontSize: 20, fontWeight: 'bold' }}>Services</Text>
          <View style={styles.optionContainer}>
            <OptionCard
              icon={<FontAwesome6 name='user-doctor' size={40} />}
              text={`${mockedData.expertsOnline}/${mockedData.expertsTotal} online`}
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
