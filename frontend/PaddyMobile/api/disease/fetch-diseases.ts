import { useContext, useEffect, useState } from 'react';
import { fetcher } from '../driver';
import { FetchType, UploadListItemType } from '@/types/misc';
import endpoints from '@/constants/endpoints';
import { AuthContext } from '@/contexts/auth/auth-provider';
import { STATUS } from '@/constants/misc';

export default function useFetchDiseases({
  page = 0,
  limit = 10,
}: {
  page: number;
  limit: number;
}) {
  const [state, setState] = useState<FetchType<UploadListItemType[]>>({
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

export function useFetchUploaded({ id, item }: { id: string; item: UploadListItemType }) {
  const [itemNew, setItemNew] = useState(item);
  const [tick, setTick] = useState<number | null>(null);

  const { token } = useContext(AuthContext);

  useEffect(() => {
    if (item.status !== STATUS.queued) return;
    const t = setInterval(() => {
      fetcher({
        params: [['itemId', id]],
        uri: endpoints.diseaseStat,
        token: token as string,
      })
        .then((r) => {
          if (r.success) setItemNew((i) => ({ ...i, ...r.data }));
        })
        .catch((e) => {
          console.error(e);
        });
    }, 5000);
    setTick(t);
  }, []);

  useEffect(() => {
    if (tick !== null && itemNew.status !== STATUS.queued) clearInterval(tick);
  }, [itemNew]);

  return itemNew;
}
