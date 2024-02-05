import { Pressable, StyleSheet, Text, View } from 'react-native';
import { Card } from 'react-native-paper';
import AntDesign from 'react-native-vector-icons/AntDesign';

export default function WelcomeCard({
  userName,
  userSubmissions,
  userDiseaseDetected,
  creditsRemaining,
}: {
  userName: string;
  userSubmissions: number;
  userDiseaseDetected: number;
  creditsRemaining: number;
}) {
  return (
    <Card style={styles.cardContainer}>
      <View style={styles.leftBg} />
      <Text style={styles.heading}>Welcome, {userName}</Text>
      <View style={styles.detailContainer}>
        <View style={styles.detail}>
          <Text style={styles.detailText}>{userSubmissions}</Text>
          <Text>submissions</Text>
        </View>
        <View style={styles.detail}>
          <Text style={styles.detailText}>{userDiseaseDetected} </Text>
          <Text>detections</Text>
        </View>
        <View style={styles.detail}>
          <Text style={styles.detailText}>{creditsRemaining} </Text>
          <Text>credits remaining</Text>
          <Pressable>
            <View style={styles.buy}>
              <AntDesign name='shoppingcart' size={20} />
              <Text>Add More</Text>
            </View>
          </Pressable>
        </View>
      </View>
    </Card>
  );
}

const styles = StyleSheet.create({
  cardContainer: {
    alignItems: 'flex-start',
    justifyContent: 'center',
    position: 'relative',
    backgroundColor: '#892bccaa',
  },
  leftBg: {
    backgroundColor: '#6ce0bbaa',
    borderRadius: 15,
    borderTopRightRadius: 200,
    borderBottomRightRadius: 0,
    width: '70%',
    position: 'absolute',
    height: '100%',
  },
  heading: {
    fontSize: 23,
    fontWeight: 'bold',
    padding: 10,
  },
  detailContainer: { padding: 10, paddingTop: 0 },
  detail: {
    flexDirection: 'row',
    gap: 20,
    alignItems: 'center',
  },
  detailText: {
    fontSize: 22,
    width: 30,
    textAlign: 'right',
    fontWeight: '600',
  },
  buy: {
    flexDirection: 'row',
    gap: 5,
    padding: 5,
    borderRadius: 10,
    backgroundColor: '#8b3bc4',
  },
});
