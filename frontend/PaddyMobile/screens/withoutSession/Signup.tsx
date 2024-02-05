import { signUpPost } from '@/api/auth/signup';
import pages from '@/constants/screens';
import { NavProps } from '@/types/misc';
import { FC, useState } from 'react';
import { Button, Image, StyleSheet, Text, TextInput, View } from 'react-native';

const SignupScreen: FC<NavProps> = ({ navigation }) => {
  const [info, setInfo] = useState({
    name: '',
    location: '',
    email: '',
    password: '',
    rePassword: '',
  });

  function handleChange(name: string, value: string) {
    setInfo((i) => ({ ...i, [name]: value }));
  }

  async function handleSignUp() {
    if (info.password !== info.rePassword) return alert("Passwords doesn't match");
    else if (info.password.length < 0) return alert('Passwords must be 5 or more characters');

    const { message } = await signUpPost(info);

    alert(message);
  }

  return (
    <View style={styles.container}>
      <Text style={styles.heading}>Create a new account</Text>
      <Image source={require('@/assets/icons/tea.png')} style={styles.img} />
      <View style={styles.inputContainer}>
        <TextInput
          onChangeText={(e) => handleChange('name', e)}
          placeholder='Name'
          style={styles.inp}
          autoFocus
        />
        <TextInput
          onChangeText={(e) => handleChange('location', e)}
          placeholder='Location'
          style={styles.inp}
        />
        <TextInput
          onChangeText={(e) => handleChange('email', e)}
          placeholder='Email'
          style={styles.inp}
          keyboardType='email-address'
        />
        <TextInput
          onChangeText={(e) => handleChange('password', e)}
          placeholder='Password'
          style={styles.inp}
          secureTextEntry
        />
        <TextInput
          onChangeText={(e) => handleChange('rePassword', e)}
          placeholder='Retype Password'
          style={styles.inp}
          secureTextEntry
        />
        <View
          style={{
            display: 'flex',
            flexDirection: 'row',
            justifyContent: 'center',
          }}
        >
          <Text>Already have an account? </Text>
          <Text
            style={{ textDecorationLine: 'underline', fontWeight: 'bold' }}
            onPress={() => navigation.navigate(pages.login)}
          >
            Login
          </Text>
        </View>
      </View>
      <View style={{ width: '50%' }}>
        <Button title='Signup' color={'purple'} onPress={handleSignUp} />
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
  },
  heading: {
    fontSize: 30,
    fontWeight: 'bold',
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

export default SignupScreen;
