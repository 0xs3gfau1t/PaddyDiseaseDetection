import useFetchDiseaseDetail from '@/api/disease/fetch-disease-detail';
import { useRoute } from '@react-navigation/native';
import { useMemo } from 'react';
import {
  ActivityIndicator,
  Dimensions,
  Image,
  ScrollView,
  StyleSheet,
  Text,
  View,
} from 'react-native';
import Carousel from 'react-native-reanimated-carousel';

function CarouselRenderItem({ item }: { item: string }) {
  return <Image source={{ uri: item }} {...styles.imgDimensions} style={{ alignSelf: 'center' }} />;
}

export default function DetailScreen() {
  const { params } = useRoute<any>();
  const { detail, fetching } = useFetchDiseaseDetail({ id: params?.id });
  const width = useMemo(() => Dimensions.get('screen').width, []);

  if (fetching) return <ActivityIndicator />;

  if (!detail) return <Text>Error while fetching</Text>;

  return (
    <View style={styles.container}>
      <Text style={styles.diseaseName}>Detected: {detail.name}</Text>
      <ScrollView contentContainerStyle={{ flex: 1 }}>
        <Carousel
          loop
          width={width}
          data={detail.images}
          scrollAnimationDuration={500}
          renderItem={CarouselRenderItem}
        />
        {/*Causes*/}
      </ScrollView>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    paddingVertical: Dimensions.get('screen').height * 0.04,
    paddingHorizontal: Dimensions.get('screen').width * 0.02,
    gap: 10,
    height: '100%',
  },
  imgDimensions: {
    width: Dimensions.get('screen').width * 0.7,
    height: Dimensions.get('screen').height * 0.3,
  },
  diseaseName: {
    fontSize: 20,
    alignSelf: 'center',
    paddingVertical: 10,
  },
});
