import useFetchDiseases from '@/api/disease/fetch-diseases';
import UploadNew from '@/components/UploadNew';
import UploadItem from '@/components/upload/UploadItem';
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

export default function UploadScreen({ navigation }: any) {
  const [page, _] = useState(0);
  const [limit, __] = useState(5);
  const {
    state: { fetching, data },
    triggerFetch,
  } = useFetchDiseases({ page, limit });

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
        <View
          style={{ gap: 5, flexDirection: 'row', alignItems: 'center', justifyContent: 'center' }}
        >
          <Text>No History</Text>
          <AntDesign name='question' size={40} />
        </View>
      );
  }, [fetching, data]);

  return (
    <View style={styles.container}>
      <Text
        style={{
          fontSize: 20,
          fontWeight: 'bold',
          paddingBottom: 10,
        }}
      >
        Upload History
      </Text>
      <ScrollView style={{ width: '100%', padding: 10, borderRadius: 20 }}>
        {data
          ? data.map((item) => <UploadItem item={item} key={item.id} navigation={navigation} />)
          : renderStats}
      </ScrollView>
      <UploadNew onUpload={triggerFetch} />

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
    zIndex: -1,
  },
});
