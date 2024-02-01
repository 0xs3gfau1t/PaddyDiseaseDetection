import { useEffect, useState } from "react";
import { deleteItemAsync, getItem, setItem } from "expo-secure-store";

const initialStatus = {
  fetching: true,
  token: null as null | string,
};

export default function useAuthContext() {
  const [tokenState, setTokenState] = useState(initialStatus);

  useEffect(() => {
    setTokenState({
      token: getItem("accessToken"),
      fetching: false,
    });
  }, []);

  function logout() {
    deleteItemAsync("accessToken").finally(() => {
      setTokenState({
        ...tokenState,
        token: null,
      });
    });
  }

  function login(token: string) {
    setItem("accessToken", token);
    setTokenState({
      ...tokenState,
      token,
    });
  }
  return { tokenState, logout, login };
}
