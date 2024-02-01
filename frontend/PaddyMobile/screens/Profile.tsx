import { useState } from "react";
import { Button, StyleSheet, Text, View } from "react-native";
import useAuthContext from "../hooks/authContext";

type ProfileProps = {
  loading: boolean;
  profileInfo: { name: string; email: string } | null;
};

const ProfileScreen = () => {
  const { logout } = useAuthContext();
  const [profileData, setProfileData] = useState<ProfileProps>({
    loading: false,
    profileInfo: { name: "Sam", email: "thicc_sam@gmail.com" },
  });

  return (
    <View style={styles.container}>
      <Text>Hello, {profileData.profileInfo?.name}</Text>
      <Button onPress={logout} title="Logout" />
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    height: "100%",
    alignItems: "center",
    justifyContent: "center",
    gap: 50,
  },
});

export default ProfileScreen;
