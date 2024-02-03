import { loginPost } from '@/api/auth/signup';
import pages from '@/constants/screens';
import { useAuthContext } from '@/contexts/auth/auth-provider';
import { NavProps } from '@/types/misc';
import { FC, useState } from 'react';
import { ActivityIndicator, Button, Image, StyleSheet, Text, TextInput, View } from 'react-native';

const LoginScreen: FC<NavProps> = ({ navigation }) => {
  const [info, setInfo] = useState({
    email: '',
    password: '',
  });
  const [loggingIn, setLoggingIn] = useState(false);

  const { setToken } = useAuthContext();

  function handleChange(name: string, value: string) {
    setInfo((i) => ({ ...i, [name]: value }));
  }

  async function handleLogin() {
    setLoggingIn(true);
    const { message, accessToken } = await loginPost(info);

    if (accessToken) setToken(accessToken);
    else alert(message);

    setLoggingIn(false);
  }

  return (
    <View style={styles.container}>
      <Text style={styles.heading}>Login</Text>
      <Image source={require('@/assets/icons/tea.png')} style={styles.img} />
      <View style={styles.inputContainer}>
        <TextInput
          onChangeText={(e) => handleChange('email', e)}
          placeholder='Email'
          style={styles.inp}
          keyboardType='email-address'
          autoFocus
        />
        <TextInput
          onChangeText={(e) => handleChange('password', e)}
          placeholder='Password'
          style={styles.inp}
          secureTextEntry
        />
      </View>
      <View
        style={{
          display: 'flex',
          flexDirection: 'row',
          justifyContent: 'center',
        }}
      >
        <Text>Don't have an account? </Text>
        <Text
          style={{ textDecorationLine: 'underline', fontWeight: 'bold' }}
          onPress={() => navigation.navigate(pages.signup)}
        >
          Signup
        </Text>
      </View>
      <View style={{ width: '50%' }}>
        {loggingIn ? (
          <ActivityIndicator />
        ) : (
          <Button title='Login' color={'purple'} onPress={handleLogin} />
        )}
      </View>
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    display: 'flex',
    justifyContent: 'space-around',
    alignItems: 'center',
    height: '100%',
    position: 'relative',
  },
  heading: {
    fontSize: 30,
    fontWeight: 'bold',
    letterSpacing: 5,
  },
  inputContainer: {
    display: 'flex',
    flexDirection: 'column',
    gap: 10,
    width: '60%',
  },
  inp: {
    borderWidth: 1,
    borderRadius: 10,
    borderColor: 'purple',
    paddingVertical: 5,
    textAlign: 'center',
    width: '100%',
  },
  img: {
    position: 'absolute',
    width: '100%',
    opacity: 0.2,
  },
});

export default LoginScreen;
