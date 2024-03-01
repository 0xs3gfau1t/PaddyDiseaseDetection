import useMapData from '@/api/map';
import { View, StyleSheet } from 'react-native';
import MapView, { Heatmap as GoogleHeatMap, PROVIDER_GOOGLE } from 'react-native-maps';

export default function HeatMap() {
  const { points } = useMapData();

  return (
    <View style={styles.container}>
      <MapView
        initialRegion={{
          latitude: 28.227674230390186,
          longitude: 83.91140561833058,
          latitudeDelta: 8,
          longitudeDelta: 8,
        }}
        provider={PROVIDER_GOOGLE}
        style={{ flex: 1 }}
          mapType="standard"
        
          nts && points.length > 0 && (
            gleHeatMap
            points={points}
              ient={{
              colors: ['#E50000', '#F29305', '#EEC20B', '#BBCF4C', '#79BC6A'],
              startPoints: [0.01, 0.25, 0.5, 0.75, 1],
                colorMapSize: 256,
              }}
            ></GoogleHeatMap>
          )}
        </MapView>
      </View>
    );
}

  t styles = StyleSheet.create({
    ainer: {
      flex: 1,
  },
     {
    width: '100%',
      height: '100%',
    },
});
