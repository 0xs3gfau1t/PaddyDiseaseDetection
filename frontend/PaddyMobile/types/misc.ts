export type NavProps = {
  navigation: any;
};

export type AuthState = {
  isFetching: boolean;
  token: string | null;
  userData: {
    name: string;
    image: string | undefined;
    email: string;
    verified: boolean;
    coords: { latitude: number; longitude: number };
    location: string;
  } | null;
};

export type FetchType<T> = {
  fetching: boolean;
  data: null | T;
};

export type LocationType = {
  latitude: number;
  longitude: number;
};
