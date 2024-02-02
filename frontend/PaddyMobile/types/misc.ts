export type NavProps = {
  navigation: any;
};

export type AuthState = {
  isFetching: boolean;
  token: string | null;
  userData: {
    name: string;
    image: string;
    email: string;
  } | null;
};
