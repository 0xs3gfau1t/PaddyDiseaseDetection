import { useFetchUploaded } from '@/api/disease/fetch-diseases';
import { STATUS } from '@/constants/misc';
import { Dimensions, Image, StyleSheet, Text, View } from 'react-native';
import { ActivityIndicator, Card } from 'react-native-paper';
import MaterialCommunityIcons from 'react-native-vector-icons/MaterialCommunityIcons';

export default function UploadItem({ item }: any) {
  const itemNew = useFetchUploaded({ id: item.id, item: item });

  let renderStatIcon = <ActivityIndicator />;
  if (itemNew.status === STATUS.processed)
    renderStatIcon = <MaterialCommunityIcons name='check' size={20} style={{ color: 'green' }} />;
  else if (itemNew.status === STATUS.queued)
    renderStatIcon = <MaterialCommunityIcons name='timer' size={20} style={{ color: 'purple' }} />;

  return (
    <Card style={{ marginVertical: 5, width: '90%', alignSelf: 'center' }}>
      <View style={styles.historyItem}>
        <Image
          source={item.images ? { uri: item.images } : require('@/assets/icons/tea.png')}
          style={styles.diseaseImage}
        />
        <Text style={{ width: 160 }}>{itemNew.name}</Text>
        <Text style={{ width: 20 }}>{itemNew.severity}</Text>
        {renderStatIcon}
      </View>
    </Card>
  );
}

const styles = StyleSheet.create({
  historyItem: {
    paddingHorizontal: 10,
    paddingVertical: 10,
    flexDirection: 'row',
    gap: 5,
    alignItems: 'center',
    justifyContent: 'space-between',
  },
  diseaseImage: {
    width: Dimensions.get('screen').width * 0.1,
    height: Dimensions.get('screen').width * 0.1,
    borderRadius: 10,
  },
  status: {
    paddingVertical: 2,
    paddingHorizontal: 5,
    borderRadius: 10,
    borderWidth: 1,
  },
});
