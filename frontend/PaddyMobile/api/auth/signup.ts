import { TOKEN_HOLDER } from '@/constants/auth';
import endpoints from '@/constants/endpoints';
import * as SecureStore from 'expo-secure-store';

export type SignUpProps = {
  email: string;
  password: string;
  name: string;
  location: string;
};

export type LoginProps = {
  email: string;
  password: string;
};

export async function signUpPost(info: SignUpProps) {
  try {
    const formData = new FormData();
    Object.entries(info).forEach(([key, val]) => {
      formData.append(key, val);
    });

    const res = await fetch(endpoints.auth.signup, {
      method: 'POST',
      body: formData,
    });

    if (!res.ok) throw await res.json();

    return { good: true, message: 'Signed up successfully' };
  } catch (e) {
    console.error(e);
    return { good: false, message: "Couldn't signup" };
  }
}

export async function loginPost(info: LoginProps) {
  try {
    const formData = new FormData();
    Object.entries(info).forEach(([key, val]) => {
      formData.append(key, val);
    });

    const res = await fetch(endpoints.auth.login, {
      method: 'POST',
      body: formData,
    });

    if (!res.ok) throw await res.json();

    const { accessToken }: { accessToken: string } = await res.json();

    return { good: true, message: 'Logged in successfully', accessToken };
  } catch (e) {
    console.error(e);
    return { good: false, message: "Couldn't login" };
  }
}

export async function logout() {
  try {
    await SecureStore.deleteItemAsync(TOKEN_HOLDER);

    return { good: true, message: 'Logged Out' };
  } catch (e) {
    console.error(e);
    return { good: false, message: "Couldn't logout" };
  }
}
