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

export type CausesType = {
  name: string;
  image: string;
};
export type SolutionType = {
  name: string;
  photos: string[];
  ingredients: string[];
  description: string;
  id: string;
};

export type UploadDetailType = {
  images: string[];
  severity: string;
  staus: string;
  id: string;
  solutions: SolutionType[];
  identified: {
    name: string;
    id: string;
  };
  causes: CausesType[];
};

export type UploadListItemType = {
  id: string;
  images: string[];
  identified: string | null;
  status: string;
  severity: number;
  name: string[] | null;
};
