import endpoints from '@/constants/endpoints';
import { executioner, patcher, poster } from './driver';

export function editProfile({
  token,
  location,
  coords,
  name,
}: {
  token: string;
  location?: string;
  coords?: { latitude: number; longitude: number };
  name?: string;
}) {
  const formData = new FormData();
  if (coords) {
    formData.append('latitude', coords.latitude.toString());
    formData.append('longitude', coords.longitude.toString());
  }
  if (location) formData.append('location', location);
  if (name) formData.append('name', name);
  return patcher({ uri: endpoints.profile.root, token, data: formData });
}

export function deleteAccount({ token }: { token: string }) {
  return executioner({ uri: endpoints.profile.root, data: null, token });
}

export function changePassword({
  token,
  oldPassword,
  newPassword,
}: {
  token: string;
  oldPassword: string;
  newPassword: string;
}) {
  const formData = new FormData();
  formData.append('oldPassword', oldPassword);
  formData.append('newPassword', newPassword);

  return poster({ uri: endpoints.profile.changePassword, data: formData, token });
}
