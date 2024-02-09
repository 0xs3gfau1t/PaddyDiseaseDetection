import { FileSystemUploadType, uploadAsync } from 'expo-file-system';

export const fetcher = async ({
  params,
  uri,
  token,
}: {
  params?: [string, string][];
  uri: string;
  token?: string;
}) => {
  const paramsJoined = params
    ? params.reduce((prev, current) => prev + '&' + current[0] + '=' + current[1], '')
    : '';
  return fetch(`${uri}?${paramsJoined}`, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  })
    .then(async (r) => {
      if (!r.ok) throw r;
      try {
        const res = await r.json();
        return { success: true, message: res.message as string, data: res.data };
      } catch (e) {
        return { success: true, message: null, data: null };
      }
    })
    .catch(async (e) => {
      try {
        const res = await e.json();

        return { success: false, message: res.message as string, data: res.data };
      } catch (e) {
        return { success: false, message: null, data: null };
      }
    });
};

export const poster = async ({
  data,
  uri,
  token,
}: {
  data: FormData;
  uri: string;
  token?: string;
}) => {
  return fetch(uri, {
    method: 'POST',
    body: data,
    headers: {
      Authorization: `Bearer ${token}`,
    },
  })
    .then(async (r) => {
      if (!r.ok) throw r;
      try {
        const res = await r.json();
        return { success: true, message: res.message as string, data: res.data };
      } catch (e) {
        return { success: true, message: null, data: null };
      }
    })
    .catch(async (e) => {
      try {
        const res = await e.json();

        return { success: false, message: res.message as string };
      } catch (e) {
        return { success: false, message: null, data: null };
      }
    });
};

export const patcher = async ({
  data,
  uri,
  token,
}: {
  data: FormData;
  uri: string;
  token: string;
}) => {
  return fetch(uri, {
    method: 'PATCH',
    body: data,
    headers: {
      Authorization: `Bearer ${token}`,
    },
  })
    .then(async (r) => {
      if (!r.ok) throw r;
      try {
        const res = await r.json();
        return { success: true, message: res.message as string, data: res.data };
      } catch (e) {
        return { success: true, message: null, data: null };
      }
    })
    .catch(async (e) => {
      try {
        const res = await e.json();

        return { success: false, message: res.message as string };
      } catch (e) {
        return { success: false, message: null, data: null };
      }
    });
};

export const executioner = async ({
  data,
  uri,
  token,
}: {
  data: FormData | null;
  uri: string;
  token: string;
}) => {
  return fetch(uri, {
    method: 'DELETE',
    body: data,
    headers: {
      Authorization: `Bearer ${token}`,
    },
  })
    .then(async (r) => {
      if (!r.ok) throw r;
      return { success: true, message: null, data: null };
    })
    .catch(async (e) => {
      try {
        const res = await e.json();

        return { success: false, message: res.message as string };
      } catch (e) {
        return { success: false, message: null, data: null };
      }
    });
};

export const uploader = async ({
  fileUri,
  fieldName,
  uri,
  token,
}: {
  fileUri: string;
  fieldName: string;
  uri: string;
  token?: string;
}) => {
  return uploadAsync(uri, fileUri, {
    httpMethod: 'POST',
    uploadType: FileSystemUploadType.MULTIPART,
    fieldName,
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
};
