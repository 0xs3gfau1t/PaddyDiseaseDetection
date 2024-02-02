import AuthProvider from "./contexts/auth/auth-provider";
import Root from "./routes";

export default function App() {
  return (
    <AuthProvider>
      <Root />
    </AuthProvider>
  );
}
