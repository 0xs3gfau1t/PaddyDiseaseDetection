import { useFetchUploaded } from '@/api/disease/fetch-diseases';
import { STATUS } from '@/constants/misc';
import pages from '@/constants/screens';
import { UploadListItemType } from '@/types/misc';
import { Dimensions, Image, Pressable, StyleSheet, Text, View } from 'react-native';
import { ActivityIndicator, Card } from 'react-native-paper';
import MaterialCommunityIcons from 'react-native-vector-icons/MaterialCommunityIcons';

export default function UploadItem({
  item,
  navigation,
}: {
  item: UploadListItemType;
  navigation: any;
}) {
  const itemNew = useFetchUploaded({ id: item.id, item: item });

  let renderStatIcon = <ActivityIndicator />;
  if (itemNew.status === STATUS.processed)
    renderStatIcon = <MaterialCommunityIcons name='check' size={20} style={{ color: 'green' }} />;
  else if (itemNew.status === STATUS.queued)
    renderStatIcon = <MaterialCommunityIcons name='timer' size={20} style={{ color: 'purple' }} />;

  return (
    <Pressable onPress={() => navigation.navigate(pages.detail, { id: item.id })}>
      <Card
        style={{
          marginVertical: 5,
          width: '90%',
          alignSelf: 'center',
          backgroundColor: item.name && item.name.length > 0 ? '#f885' : '#afa5',
        }}
      >
        <View style={styles.historyItem}>
          <Image
            source={item.images ? { uri: item.images[0] } : require('@/assets/icons/tea.png')}
            style={styles.diseaseImage}
          />
          <Text style={{ width: 160 }}>
            {itemNew.name ? formatName(itemNew.name) : 'Processing...'}
          </Text>
          <Text style={{ width: 20 }}>{item.severity}</Text>
          {renderStatIcon}
        </View>
      </Card>
    </Pressable>
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

function formatName(name: string[]) {
  return name.length ? name.join(', ') : 'No Disease Detected';
}
