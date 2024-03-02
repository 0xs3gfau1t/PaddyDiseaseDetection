import { useMemo } from 'react';
import { StyleSheet, Text, View } from 'react-native';

export default function AreaCard({
  areaSubmissions,
  areaDiseaseDetected,
}: {
  areaSubmissions: number;
  areaDiseaseDetected: number;
}) {
  const renderSeverity = useMemo(() => {
    const calculatedSeverity = areaDiseaseDetected / areaSubmissions;

    if (calculatedSeverity > 0.6)
      return <Text style={{ ...styles.commonSevere, ...styles.severe }}>Severe</Text>;
    else return <Text style={{ ...styles.commonSevere, ...styles.nonSevere }}>Not Severe</Text>;
  }, [areaDiseaseDetected, areaSubmissions]);

  return (
    <View style={styles.container}>
      <Text style={styles.heading}>In your area</Text>
      <View style={styles.detail}>
        <Text style={styles.detailText}>{areaSubmissions}</Text>
        <Text>submissions</Text>
      </View>
      <View style={styles.detail}>
        <Text style={styles.detailText}>{areaDiseaseDetected} </Text>
        <Text>detections</Text>
      </View>
      {renderSeverity}
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    padding: 10,
    justifyContent: 'flex-end',
    paddingVertical: 20,
    backgroundColor: '#ddd4',
    borderTopRightRadius: 10,
    borderTopLeftRadius: 10,
  },
  heading: { fontWeight: 'bold', fontSize: 20, paddingBottom: 10 },
  commonSevere: {
    borderWidth: 1,
    alignSelf: 'flex-end',
    padding: 5,
    paddingHorizontal: 10,
    borderRadius: 10,
  },
  severe: { borderColor: 'red' },
  nonSevere: { borderColor: 'orange' },
  detail: {
    flexDirection: 'row',
    gap: 20,
    alignItems: 'center',
  },
  detailText: {
    fontSize: 22,
    width: 50,
    textAlign: 'right',
    fontWeight: '600',
  },
});
