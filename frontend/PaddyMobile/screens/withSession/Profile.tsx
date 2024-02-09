import { changePassword, deleteAccount, editProfile } from '@/api/profile';
import LocationPicker from '@/components/profile/LocationPicker';
import { VERIFICATION_EXPIRY_TIME } from '@/constants/misc';
import { AuthContext } from '@/contexts/auth/auth-provider';
import { LocationType } from '@/types/misc';
import { reverseGeocodeAsync } from 'expo-location';
import { useContext, useMemo, useState } from 'react';
import {
  ActivityIndicator,
  Alert,
  Button,
  Dimensions,
  Image,
  Pressable,
  ScrollView,
  StyleSheet,
  Text,
  TextInput,
  View,
} from 'react-native';
import OutsidePressHandler from 'react-native-outside-press';
import { Card } from 'react-native-paper';
import AntDesign from 'react-native-vector-icons/AntDesign';
import Octicons from 'react-native-vector-icons/Octicons';

const defaultProfilePic = require('@/assets/icons/tea.png');

const verificationStat = {
  NOT_SENT: 'Verify',
  SENT: 'Sent',
  SENDING: 'Sending',
};

export default function ProfileScreen() {
  const { removeToken, userData: userDataContext, token } = useContext(AuthContext);
  const [editActiveFields, setEditActiveFields] = useState({
    name: 0,
    image: 0,
    location: 0,
    verified: 0,
    password: 0,
  });
  const [userData, setUserData] = useState(
    userDataContext || {
      name: 'N/A',
      image: undefined,
      verified: false,
      email: 'N/A',
      coords: { latitude: 0, longitude: 0 },
      location: 'N/A',
    }
  );
  const [profileImageSize, _] = useState(Dimensions.get('screen').width * 0.25);
  const [verificationText, setVerificationText] = useState(verificationStat.NOT_SENT);
  const [oldPassword, setOldPassword] = useState('');
  const [newPassword, setNewPassword] = useState('');
  const [reNewPassword, setReNewPassword] = useState('');

  async function handleLocationPicked({ latitude, longitude }: LocationType) {
    const location = await reverseGeocodeAsync({ latitude, longitude });
    setEditActiveFields((f) => ({ ...f, location: 1 }));
    if (!token) return;
    const { success } = await editProfile({
      coords: { latitude, longitude },
      location: location[0].city || location[0].subregion || location[0].region || undefined,
      token,
    });
    if (success) alert('Location updated successfully');
    else alert('Failed updating location');
    setEditActiveFields((f) => ({ ...f, location: 0 }));
  }

  async function handleNameChangeSubmit() {
    setEditActiveFields((f) => ({ ...f, name: 1 }));
    if (!token || !userData?.name) return;
    await editProfile({ name: userData.name, token });
    setEditActiveFields((f) => ({ ...f, name: 0 }));
  }

  function handleDelete() {
    if (!token) return;
    Alert.alert('Confirm delete account?', 'Deleting your account is permanent and irreversible', [
      {
        text: 'Confirm',
        onPress: () => {
          deleteAccount({ token }).then((r) => {
            if (r.success) {
              alert('Account deleted');
              removeToken();
            } else alert("Couldn't delete account");
          });
        },
      },
      { text: 'Cancel', onPress: () => {} },
    ]);
  }

  const renderImage = useMemo(() => {
    if (userData?.image)
      return (
        <Image
          source={{ uri: userData.image }}
          style={{ ...styles.image, width: profileImageSize, height: profileImageSize }}
        />
      );
    else
      return (
        <Image
          source={defaultProfilePic}
          style={{ ...styles.image, width: profileImageSize, height: profileImageSize }}
        />
      );
  }, [userData, defaultProfilePic, profileImageSize]);

  async function handlePasswordChange() {
    setEditActiveFields((f) => ({ ...f, password: 1 }));
    if (!token) return;
    if (newPassword !== reNewPassword) alert("Passwords don't match");
    const { success } = await changePassword({ oldPassword, newPassword, token });
    if (success) alert('Updated new password');
    else alert("Couldn't change password");
    setEditActiveFields((f) => ({ ...f, password: 0 }));
  }

  async function handleVerificationRequest() {
    setVerificationText(verificationStat.SENDING);

    // Make a fetch request

    //setVerificationText(verificationStat.SENT);

    setTimeout(() => {
      setVerificationText(verificationStat.NOT_SENT);
    }, VERIFICATION_EXPIRY_TIME);
  }

  return (
    <ScrollView>
      <View style={styles.container}>
        <Card style={styles.imageContainer}>
          {renderImage}
          <View
            style={{ flexDirection: 'row', gap: 10, alignItems: 'center', alignSelf: 'center' }}
          >
            {userData?.verified ? (
              <Octicons name='verified' size={20} style={{ color: 'aqua' }} />
            ) : (
              <Octicons name='unverified' size={20} style={{ color: 'orange' }} />
            )}
            {editActiveFields.name < 0 ? (
              <OutsidePressHandler
                onOutsidePress={() => setEditActiveFields((f) => ({ ...f, name: 0 }))}
                style={{ flexDirection: 'row', gap: 10 }}
              >
                <TextInput
                  value={userData?.name}
                  onChangeText={(e) => setUserData((d) => ({ ...d, name: e }))}
                  autoFocus
                />
                <Button title='Save' onPress={handleNameChangeSubmit} />
              </OutsidePressHandler>
            ) : (
              <>
                <Text style={styles.name}>{userData?.name}</Text>
                <Pressable onPress={() => setEditActiveFields((f) => ({ ...f, name: -1 }))}>
                  <AntDesign name='edit' size={20} />
                </Pressable>
              </>
            )}
          </View>
          <View style={styles.bgContainer} />
        </Card>

        <View>
          {editActiveFields.location < 0 ? (
            <LocationPicker
              onPicked={handleLocationPicked}
              onCancel={() => setEditActiveFields((f) => ({ ...f, location: 0 }))}
              currentCoords={userData ? userData.coords : undefined}
            />
          ) : (
            <EditableField
              label='Location:'
              value={userData.location}
              onEditActive={() => setEditActiveFields((f) => ({ ...f, location: -1 }))}
            />
          )}

          <EditableField label='Email:' value={userData?.email} disable />
        </View>
        {!userData?.verified && (
          <Pressable onPress={handleVerificationRequest}>
            <View
              style={{
                paddingHorizontal: 10,
                paddingVertical: 5,
                borderRadius: 10,
                borderWidth: 1,
                borderColor: 'rose',
                alignSelf: 'center',
              }}
            >
              {verificationText === verificationStat.SENDING ? (
                <ActivityIndicator />
              ) : (
                <Text>{verificationText}</Text>
              )}
            </View>
          </Pressable>
        )}
        <View
          style={{
            borderTopWidth: 1,
            borderStyle: 'dashed',
            marginTop: 10,
            paddingTop: 10,
            gap: 10,
          }}
        >
          <Text style={{ fontSize: 18, fontWeight: 'bold' }}>Danger Zone</Text>
          <View
            style={{
              alignSelf: 'center',
              alignItems: editActiveFields.password < 0 ? 'center' : 'stretch',
              width: editActiveFields.password ? '100%' : '50%',
              gap: 10,
              justifyContent: 'center',
            }}
          >
            {editActiveFields.password < 0 && (
              <Card
                style={{
                  width: '100%',
                  padding: Dimensions.get('screen').width * 0.02,
                  backgroundColor: '#e87e72',
                }}
              >
                <View style={styles.passwordInputContainer}>
                  <Text style={styles.label}>Old Password</Text>
                  <TextInput
                    onChangeText={(e) => setOldPassword(e)}
                    style={styles.input}
                    secureTextEntry
                  />
                </View>
                <View style={styles.passwordInputContainer}>
                  <Text style={styles.label}>New Password</Text>
                  <TextInput
                    onChangeText={(e) => setNewPassword(e)}
                    style={styles.input}
                    secureTextEntry
                  />
                </View>
                <View style={styles.passwordInputContainer}>
                  <Text style={styles.label}>Repeat New Password</Text>
                  <TextInput
                    onChangeText={(e) => setReNewPassword(e)}
                    style={styles.input}
                    secureTextEntry
                  />
                </View>
                <View
                  style={{
                    flexDirection: 'row',
                    alignItems: 'center',
                    justifyContent: 'center',
                    gap: 10,
                  }}
                >
                  <Button title='Save' onPress={handlePasswordChange} />
                  <Button
                    title='Cancel'
                    onPress={() => setEditActiveFields((f) => ({ ...f, password: 0 }))}
                  />
                </View>
              </Card>
            )}
            {!editActiveFields.password && (
              <Button
                title='Change Password'
                onPress={() => setEditActiveFields((f) => ({ ...f, password: -1 }))}
              />
            )}
            <Button title='Logout' onPress={removeToken} />
            <Button title='Delete Account' onPress={handleDelete} />
          </View>
        </View>
      </View>
    </ScrollView>
  );
}

function EditableField({
  label,
  value,
  disable,
  onEditActive,
}: {
  label: string;
  value?: string;
  disable?: boolean;
  onEditActive?: () => void;
}) {
  return (
    <View style={styles.componentInfo}>
      <View style={{ flexDirection: 'row', gap: 10 }}>
        <Text style={{ width: Dimensions.get('screen').width * 0.2 }}>{label}</Text>
        <Text>{value}</Text>
      </View>
      {!disable && (
        <Pressable onPress={onEditActive}>
          <AntDesign name='edit' size={20} />
        </Pressable>
      )}
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    paddingVertical: Dimensions.get('screen').height * 0.04,
    paddingHorizontal: Dimensions.get('screen').width * 0.02,
    gap: 10,
  },
  imageContainer: {
    borderRadius: 20,
    paddingBottom: 10,
    position: 'relative',
    gap: 10,
    backgroundColor: '#ddda',
  },
  image: {
    borderWidth: 2,
    borderRadius: 20,
    alignSelf: 'center',
    marginTop: 10,
  },
  bgContainer: {
    position: 'absolute',
    height: '60%',
    backgroundColor: '#892bccaa',
    width: '100%',
    zIndex: -1,
    borderRadius: 20,
  },
  profilePic: { borderWidth: 2, borderRadius: 20 },
  name: { fontSize: 20, fontWeight: 'bold' },
  componentInfo: {
    justifyContent: 'space-between',
    paddingVertical: 10,
    paddingHorizontal: 20,
    flexDirection: 'row',
  },
  label: {
    width: Dimensions.get('screen').width * 0.4,
  },
  input: {
    borderWidth: 1,
    width: Dimensions.get('screen').width * 0.5,
    borderRadius: 10,
    borderColor: '#000',
    textAlign: 'center',
  },
  passwordInputContainer: {
    flexDirection: 'row',
    marginVertical: 5,
    alignItems: 'center',
  },
});
