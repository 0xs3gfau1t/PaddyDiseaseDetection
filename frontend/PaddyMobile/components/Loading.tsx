import { Image, Text, View } from 'react-native';

export default function LoadingScreen() {
  return (
    <View style={{ height: '100%', alignItems: 'center', justifyContent: 'center' }}>
      <Image
        source={require('../assets/illustrations/process.gif')}
        style={{ width: '100%', height: '50%' }}
      />
      <Text>Loading...</Text>
    </View>
  );
}
