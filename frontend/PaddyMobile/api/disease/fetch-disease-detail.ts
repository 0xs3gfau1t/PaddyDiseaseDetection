import { useContext, useEffect, useState } from 'react';
import { fetcher } from '../driver';
import endpoints from '@/constants/endpoints';
import { AuthContext } from '@/contexts/auth/auth-provider';
import { UploadDetailType } from '@/types/misc';

export default function useFetchDiseaseDetail({ id }: { id: string }) {
  const [detail, setDetail] = useState<null | UploadDetailType>(null);
  const [fetching, setFetching] = useState(true);
  const { token } = useContext(AuthContext);

  useEffect(() => {
    if (!token) return alert('Unauthorized action');
    fetcher({ uri: endpoints.disease, token, params: [['itemId', id]] })
      .then((r) => {
        if (!r.success) throw r.message;
        setDetail(r.data);
      })
      .catch((e: any) => {
        if (typeof e === 'string') alert(e);
        console.error(e);
      })
      .finally(() => setFetching(false));
  }, [id]);

  return { detail, fetching };
}
