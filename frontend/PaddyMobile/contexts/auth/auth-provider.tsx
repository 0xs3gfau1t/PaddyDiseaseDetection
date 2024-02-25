import { ReactNode, createContext, useEffect, useState } from 'react';
import { deleteItemAsync, getItem, setItem } from 'expo-secure-store';
import { AuthContextType } from '@/types/contexts/auth';
import { AuthState } from '@/types/misc';
import { TOKEN_HOLDER } from '@/constants/auth';
import { fetcher } from '@/api/driver';
import endpoints from '@/constants/endpoints';

export const AuthContext = createContext<AuthContextType>({
  isFetching: true,
  token: null,
  userData: null,
  setToken: (_token) => {},
  removeToken: () => {},
});

export default function AuthProvider({ children }: { children: ReactNode }) {
  const [authState, setAuthState] = useState<AuthState>({
    isFetching: true,
    token: null,
    userData: null,
  });

  useEffect(() => {
    const token = getItem(TOKEN_HOLDER);
    if (token) setToken(token);
    else setAuthState({ isFetching: false, token, userData: null });
  }, []);

  function setToken(token: string) {
    setItem(TOKEN_HOLDER, token);
    fetcher({ uri: endpoints.profile.root, token }).then((r) => {
      setAuthState({
        userData: r.data,
        isFetching: false,
        token,
      });
    });
  }

  async function removeToken() {
    await deleteItemAsync(TOKEN_HOLDER);
    setAuthState({ isFetching: false, token: null, userData: null });
  }
  return (
    <AuthContext.Provider value={{ ...authState, setToken, removeToken }}>
      {children}
    </AuthContext.Provider>
  );
}
