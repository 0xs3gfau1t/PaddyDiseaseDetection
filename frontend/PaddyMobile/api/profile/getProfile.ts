export async function getLoggedInProfileInfo(token: string) {
  return {
    name: 'Sam',
    email: 'thicc_sam@gmail.com',
    image: undefined,
    verified: false,
    coords: { latitude: 20, longitude: 80 },
  };
}
