import { StyleSheet, Text, View } from 'react-native';
import { Card } from 'react-native-paper';

export default function OptionCard({ text, icon }: { text: string; icon: any }) {
  return (
    <Card style={styles.container}>
      <View style={{ alignItems: 'center', paddingBottom: 5 }}>{icon}</View>
      <Text style={{ textAlign: 'center' }}>{text}</Text>
    </Card>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'center',
    padding: 10,
    gap: 5,
  },
});
