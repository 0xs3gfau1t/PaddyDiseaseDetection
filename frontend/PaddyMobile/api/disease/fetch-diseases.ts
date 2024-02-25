import { useContext, useEffect, useState } from 'react';
import { fetcher } from '../driver';
import { FetchType } from '@/types/misc';
import endpoints from '@/constants/endpoints';
import { AuthContext } from '@/contexts/auth/auth-provider';
import { STATUS } from '@/constants/misc';

type ResponseType = {
  id: string;
  name: string;
  severity: number;
  status: string;
  image: string;
}[];

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

  const { token } = useContext(AuthContext);

  const triggerFetch = () =>
    fetcher({
      params: [
        ['page', page.toString()],
        ['limit', limit.toString()],
      ],
      uri: endpoints.diseases,
      token: token as string,
    })
      .then((r) => {
        if (r.success)
          setState({
            fetching: false,
            data: r.data,
          });
        else throw new Error();
      })
      .catch((e) => {
        console.error(e);
        setState({
          fetching: false,
          data: null,
        });
      });

  useEffect(() => {
    triggerFetch();
  }, [page, limit]);

  return { state, triggerFetch };
}

export function useFetchUploaded({ id, item }: { id: string; item: any }) {
  const [itemNew, setItemNew] = useState(item);
  const [tick, setTick] = useState(0);

  const { token } = useContext(AuthContext);

  useEffect(() => {
    if (item.status !== STATUS.queued) return;
    fetcher({
      params: [['itemId', id]],
      uri: endpoints.disease,
      token: token as string,
    })
      .then((r) => {
        if (r.success) setItemNew(r.data);
        if (r.data.status === STATUS.queued) setTimeout(() => setTick(tick + 1), 1000);
        else throw new Error();
      })
      .catch((e) => {
        console.error(e);
        setItemNew(item);
      });
  }, [tick]);

  return itemNew;
}
