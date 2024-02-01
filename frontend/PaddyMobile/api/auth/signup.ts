import endpoints from "../../constants/endpoints";

export type SignUpProps = {
  email: string;
  password: string;
  name: string;
  location: string;
};

export default async function signUpPost(info: SignUpProps) {
  try {
    const formData = new FormData();
    Object.entries(info).forEach(([key, val]) => {
      formData.append(key, val);
    });

    const res = await fetch(endpoints.auth.signup, {
      method: "POST",
      body: formData,
    });

    if (!res.ok) throw await res.json();

    return { good: true, message: "Signed up successfully" };
  } catch (e) {
    console.error(e);
    return { good: false, message: "Couldn't signup" };
  }
}
