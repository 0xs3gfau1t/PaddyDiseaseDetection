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
  id: string;
  name: string;
  photos: string[];
  ingredients: string[];
  description: string;
};

export type ROI = {
  box: number[];
  conf: number;
  classId: number;
  name: string;
  color: string;
};

export type UploadDetailType = {
  id: string;
  name: string[];
  images: string[];
  severity: number;
  staus: string;
  identified: {
    id: string;
    name: string;
    solutions: SolutionType[];
    causes: CausesType[];
  }[];
  roi: string;
};

export type UploadListItemType = {
  id: string;
  name: string[];
  severity: number;
  status: string;
  images: string[];
};

export type DashboardDataType = {
  userName: string;
  userSubmissions: number;
  userDiseaseDetected: number;
  creditsRemaining: number;
  areaSubmissions: number;
  areaDiseaseDetected: number;
  expertsOnline: number;
  expertsTotal: number;
};
