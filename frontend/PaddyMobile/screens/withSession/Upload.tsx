import { StyleSheet, Text, View } from 'react-native';

export default function UploadScreen() {
  return (
    <View style={styles.container}>
      <Text>Hello, Upload</Text>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    height: '100%',
    alignItems: 'center',
    justifyContent: 'center',
    gap: 50,
  },
});
