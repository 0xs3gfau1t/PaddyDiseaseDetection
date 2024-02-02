import { ReactNode, createContext, useContext, useEffect, useState } from 'react';
import { deleteItemAsync, getItem, setItem } from 'expo-secure-store';
import { AuthContextType } from '@/types/contexts/auth';
import { AuthState } from '@/types/misc';
import { TOKEN_HOLDER } from '@/constants/auth';
import { getLoggedInProfileInfo } from '@/api/profile/getProfile';

export const AuthContext = createContext({
  isFetching: true,
  token: null,
  userData: null,
} as AuthContextType);

export default function AuthProvider({ children }: { children: ReactNode }) {
  const [authState, setAuthState] = useState<AuthState>({} as AuthContextType);
  const [accessToken, setAccessToken] = useState<string | null>(null);

  useEffect(() => {
    const token = getItem(TOKEN_HOLDER);
    if (token) {
      getLoggedInProfileInfo(token).then((info) => {
        setAuthState({
          userData: info,
          isFetching: false,
          token,
        });
      });
    } else {
      setAuthState({ isFetching: false, token, userData: null });
    }
  }, [accessToken]);

  function setToken(token: string) {
    setItem(TOKEN_HOLDER, token);
    setAccessToken(token);
  }

  async function removeToken() {
    await deleteItemAsync(TOKEN_HOLDER);
    setAccessToken(null);
  }
  return (
    <AuthContext.Provider value={{ ...authState, setToken, removeToken }}>
      {children}
    </AuthContext.Provider>
  );
}

export const useAuthContext = () => {
  const context = useContext(AuthContext);

  return context;
};
