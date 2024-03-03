import { AuthState } from '../misc';

export type AuthContextType = {
  setToken: (token: string) => void;
  removeToken: () => void;
  apiUrl: string;
  setApiUrl: (url: string) => void;
} & AuthState;
