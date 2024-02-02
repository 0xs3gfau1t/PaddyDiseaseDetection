import { Button, StyleSheet, Text, View } from 'react-native';
import { useAuthContext } from '../contexts/auth/auth-provider';

const ProfileScreen = () => {
  const { removeToken, userData } = useAuthContext();

  return (
    <View style={styles.container}>
      <Text>Hello, {userData?.name}</Text>
      <Button onPress={removeToken} title='Logout' />
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    height: '100%',
    alignItems: 'center',
    justifyContent: 'center',
    gap: 50,
  },
});

export default ProfileScreen;
