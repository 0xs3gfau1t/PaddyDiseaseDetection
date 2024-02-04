import { LocationType } from '@/types/misc';
import { getCurrentPositionAsync, requestForegroundPermissionsAsync } from 'expo-location';
import { createRef, useEffect, useState } from 'react';
import { Button, StyleSheet, View } from 'react-native';
import MapView, { MapMarker, Marker } from 'react-native-maps';
import { Card } from 'react-native-paper';

export default function LocationPicker({
  onPicked,
  onCancel,
  currentCoords,
}: {
  currentCoords?: LocationType;
  onCancel: () => void;
  onPicked: ({ latitude, longitude }: LocationType) => void;
}) {
  const [pickedLoc, setPickedLoc] = useState<
    | {
        latitude: number;
        longitude: number;
      }
    | undefined
  >(currentCoords);
  const markerRef = createRef<MapMarker>();
  const mapRef = createRef<MapView>();

  useEffect(() => {
    (async () => {
      let { status } = await requestForegroundPermissionsAsync();
      if (status !== 'granted') {
        alert('Permission to access location was denied');
        return;
      }

      const location = await getCurrentPositionAsync();
      setPickedLoc({ latitude: location.coords.latitude, longitude: location.coords.longitude });
    })();
  }, []);

  useEffect(() => {
    if (!pickedLoc) return;
    markerRef.current?.animateMarkerToCoordinate(pickedLoc, 1000);
  }, [pickedLoc]);

  return (
    <Card style={styles.container}>
      <MapView
        style={styles.map}
        region={pickedLoc ? { ...pickedLoc, latitudeDelta: 0.03, longitudeDelta: 0.03 } : undefined}
        ref={mapRef}
      >
        {pickedLoc && <Marker coordinate={pickedLoc} ref={markerRef} />}
      </MapView>
      {pickedLoc && (
        <View style={styles.options}>
          <Button onPress={() => onPicked(pickedLoc)} title='Save' />
          <Button onPress={onCancel} title='Cancel' />
        </View>
      )}
    </Card>
  );
}

const styles = StyleSheet.create({
  container: {
    padding: 10,
  },
  map: {
    width: '100%',
    height: 200,
    borderRadius: 10,
  },
  options: { alignSelf: 'center', paddingTop: 10, flexDirection: 'row', gap: 10 },
});
