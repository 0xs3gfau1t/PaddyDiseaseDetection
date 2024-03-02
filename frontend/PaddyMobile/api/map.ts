import { FetchType } from '@/types/misc';
import { useEffect, useState } from 'react';
import { fetcher } from './driver';
import endpoints from '@/constants/endpoints';

export type HeatmapPoint = {
  id: number;
  latitude: number;
  longitude: number;
  weight: number;
};

export default function useMapData() {
  const [state, setState] = useState<FetchType<HeatmapPoint[]>>({
    fetching: true,
    data: null,
  });

  useEffect(() => {
    fetcher({ uri: endpoints.map })
      .then((r) => {
        if (r.success) setState({ fetching: false, data: r.data });
        else throw new Error();
      })
      .catch((e) => {
        console.error(e);
        setState({ fetching: false, data: null });
      });
  }, []);

  return { points: state.data, fetching: state.fetching };
}
