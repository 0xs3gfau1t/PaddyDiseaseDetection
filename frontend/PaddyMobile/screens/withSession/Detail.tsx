import useFetchDiseaseDetail from '@/api/disease/fetch-disease-detail';
import { CausesType, SolutionType } from '@/types/misc';
import { useRoute } from '@react-navigation/native';
import { useState } from 'react';
import {
  ActivityIndicator,
  Dimensions,
  Image,
  Pressable,
  ScrollView,
  StyleSheet,
  Text,
  View,
} from 'react-native';
import { Card } from 'react-native-paper';
import Carousel from 'react-native-reanimated-carousel';
import AntDesign from 'react-native-vector-icons/AntDesign';
import Entypo from 'react-native-vector-icons/Entypo';

const _mockCauses = [
  {
    name: 'Bacteria',
    image:
      'https://fastly.picsum.photos/id/620/200/300.jpg?hmac=ZLg-jrbFo8ASzAzQlxN4yMTNJJBpZtcpDXfwBxAvcT4',
  },
];

function CarouselRenderItem({ item }: { item: string }) {
  return <Image source={{ uri: item }} {...styles.imgDimensions} style={{ alignSelf: 'center' }} />;
}

function CarouselRenderCauseItem({ item }: { item: { name: string; image: string } }) {
  return (
    <View style={{ flex: 1 }}>
      <Image source={{ uri: item.image }} style={{ flex: 1, width: '100%', height: '100%' }} />
      <Text style={{ textAlign: 'center' }}>{item.name}</Text>
    </View>
  );
}

export default function DetailScreen() {
  const { params } = useRoute<any>();
  const { detail, fetching } = useFetchDiseaseDetail({ id: params?.id });

  let renderContent = <ActivityIndicator />;
  if (!fetching && !detail) renderContent = <Text>Error while fetching</Text>;
  else if (detail)
    renderContent = (
      <ScrollView>
        <Carousel
          loop
          width={styles.imgDimensions.width / 0.7}
          height={styles.imgDimensions.height}
          data={detail.images}
          renderItem={CarouselRenderItem}
        />
        <Text style={styles.diseaseName}>{detail.identified.name}</Text>
        <SolutionsView solutions={detail.solutions} causes={detail.causes} />
      </ScrollView>
    );

  return <View style={styles.container}>{renderContent}</View>;
}

const styles = StyleSheet.create({
  container: {
    paddingVertical: Dimensions.get('screen').height * 0.04,
    paddingHorizontal: Dimensions.get('screen').width * 0.02,
    height: '100%',
    gap: 10,
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
  solutionsContainer: { marginHorizontal: 1 },
});

function SolutionsView({ solutions, causes }: { solutions: SolutionType[]; causes: CausesType[] }) {
  return (
    <View style={styles.solutionsContainer}>
      <Card style={{ alignItems: 'center', paddingBottom: 10 }}>
        <Text style={{ paddingVertical: 3, fontSize: 20, fontWeight: 'bold', textAlign: 'center' }}>
          Causes
        </Text>
        <Carousel
          loop
          width={styles.imgDimensions.width}
          height={styles.imgDimensions.height}
          data={_mockCauses}
          renderItem={CarouselRenderCauseItem}
        />
      </Card>
      <View style={{ alignItems: 'center', paddingVertical: 10, gap: 5 }}>
        <Text style={{ paddingVertical: 5, fontSize: 20, fontWeight: 'bold', textAlign: 'center' }}>
          Solutions
        </Text>

        {solutions.map((s) => (
          <SolutionItemView detail={s} />
        ))}
      </View>
    </View>
  );
}

function SolutionItemView({ detail }: { detail: SolutionType }) {
  const [isExpanded, setIsExpanded] = useState(false);

  return (
    <Card style={{ padding: 15 }}>
      <Pressable
        style={{
          flexDirection: 'row',
          justifyContent: 'space-between',
          width: '100%',
          alignItems: 'center',
        }}
        onPress={() => setIsExpanded(!isExpanded)}
      >
        <Text style={{ fontWeight: 'bold' }}>{detail.name}</Text>
        <AntDesign name={isExpanded ? 'up' : 'down'} />
      </Pressable>
      {isExpanded && (
        <View style={{ paddingVertical: 10, gap: 10, alignItems: 'center' }}>
          <Carousel
            loop
            width={styles.imgDimensions.width}
            height={styles.imgDimensions.height}
            data={[_mockCauses[0].image] || detail.photos}
            renderItem={CarouselRenderItem}
            style={{ paddingVertical: 10 }}
          />
          <View style={{ width: '100%' }}>
            <Text style={{ fontWeight: 'bold', fontSize: 15 }}>Ingredients</Text>
            <View style={{ flexWrap: 'wrap', gap: 5, flexDirection: 'row' }}>
              {detail.ingredients.map((i, _) => (
                <View key={_} style={{ flexDirection: 'row', alignItems: 'center' }}>
                  <Entypo name='dot-single' />
                  <Text>{i}</Text>
                </View>
              ))}
            </View>
          </View>
          <View
            style={{
              flexDirection: 'row',
              alignItems: 'center',
              gap: 5,
              borderRadius: 20,
              backgroundColor: '#aaaa',
              paddingHorizontal: 10,
              paddingVertical: 5,
              alignSelf: 'flex-end',
            }}
          >
            <Text>Buy now</Text>
            <AntDesign name='shoppingcart' />
          </View>
          <Text style={{ fontSize: 15, fontWeight: 'bold' }}>How to use?</Text>
          <Text>{detail.description}</Text>
        </View>
      )}
    </Card>
  );
}
