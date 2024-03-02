import { DashboardDataType } from '@/types/misc';
import { useIsFocused } from '@react-navigation/native';
import { useEffect, useState } from 'react';
import { fetcher } from './driver';
import endpoints from '@/constants/endpoints';

export function useGetDashboard(token: string | null) {
  const [dashboardData, setDashboardData] = useState({
    data: null as DashboardDataType | null,
    loading: true,
  });
  const isFocused = useIsFocused();

  useEffect(() => {
    if (!isFocused || !token) return;
    fetcher({ token, uri: endpoints.dashboard })
      .then((r) => {
        if (r.success) setDashboardData({ data: r.data, loading: false });
        else throw new Error("Couldn't fetch dashboard data");
      })
      .catch((e) => {
        console.error(e);
        setDashboardData({ data: null, loading: false });
      });
  }, [isFocused]);

  return dashboardData;
}
