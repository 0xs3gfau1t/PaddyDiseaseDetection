import { useEffect, useState } from 'react';
import { fetcher } from '../driver';
import { FetchType } from '@/types/misc';
import endpoints from '@/constants/endpoints';

type ResponseType = {
  diseases: { id: string; identified_as: string; imageLink: string }[];
};

export default function useFetchDiseases({
  page = 0,
  limit = 10,
}: {
  page: number;
  limit: number;
}) {
  const [state, setState] = useState<FetchType<ResponseType>>({
    fetching: true,
    data: null,
  });

  useEffect(() => {
    fetcher({
      params: [
        ['page', page.toString()],
        ['limit', limit.toString()],
      ],
      uri: endpoints.disease,
    })
      .then((r) => {
        if (r.success)
          setState({
            fetching: false,
            data: r.data,
          });
        else throw new Error();
      })
      .catch(() => {
        setState({
          fetching: false,
          data: null,
        });
      });
  }, [page, limit]);

  return state;
}
