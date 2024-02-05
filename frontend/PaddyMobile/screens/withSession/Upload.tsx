import useFetchDiseases from '@/api/disease/fetch-diseases';
import UploadNew from '@/components/UploadNew';
import { useMemo, useState } from 'react';
import {
  ActivityIndicator,
  Dimensions,
  Image,
  ScrollView,
  StyleSheet,
  Text,
  View,
} from 'react-native';
import { Card } from 'react-native-paper';
import AntDesign from 'react-native-vector-icons/AntDesign';

export default function UploadScreen() {
  const [page, _] = useState(0);
  const [limit, __] = useState(5);
  const { fetching, data } = useFetchDiseases({ page, limit });

  function triggerRefresh() {
    console.log('List triggered');
  }

  const renderStats = useMemo(() => {
    if (fetching)
      return (
        <View
          style={{
            flexDirection: 'row',
            width: '100%',
            gap: 20,
            alignItems: 'center',
          }}
        >
          <ActivityIndicator size={50} />
          <Text>Fetching history</Text>
        </View>
      );
    else
      return (
        <View style={{ gap: 5, flexDirection: 'row', alignItems: 'center' }}>
          <AntDesign name='question' size={40} />
          <Text>No history found</Text>
        </View>
      );
  }, [fetching, data]);

  return (
    <View style={styles.container}>
      <ScrollView style={{ width: '100%', padding: 10, borderRadius: 20 }}>
        <Text
          style={{
            fontSize: 20,
            fontWeight: 'bold',
            paddingBottom: 10,
          }}
        >
          Upload History
        </Text>
        <Card style={styles.history}>
          {data
            ? data.diseases.map((item) => (
                <View style={styles.historyItem}>
                  <Text key={item.id}>Identified as: {item.identified_as}</Text>
                </View>
              ))
            : renderStats}
        </Card>
      </ScrollView>
      <UploadNew onUpload={triggerRefresh} />

      <Image source={require('@/assets/icons/tea.png')} style={styles.imgBg} />
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    height: '100%',
    flexDirection: 'column',
    paddingVertical: 40,
    gap: 10,
    alignItems: 'center',
  },
  imgBg: {
    position: 'absolute',
    width: Dimensions.get('screen').width,
    opacity: 0.1,
  },
  history: {
    width: '95%',
    gap: 2,
    alignSelf: 'center',
  },
  historyItem: {
    padding: 5,
  },
});
