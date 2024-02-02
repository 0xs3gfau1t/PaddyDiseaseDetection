import { AuthState } from "../misc";

export type AuthContextType = {
  setToken: (token: string) => void;
  removeToken: () => void;
} & AuthState;
