import LocationPicker from '@/components/profile/LocationPicker';
import { VERIFICATION_EXPIRY_TIME } from '@/constants/misc';
import { AuthContext } from '@/contexts/auth/auth-provider';
import { LocationType } from '@/types/misc';
import { useContext, useState } from 'react';
import {
  ActivityIndicator,
  Alert,
  Button,
  Dimensions,
  Image,
  Pressable,
  StyleSheet,
  Text,
  TextInput,
  View,
} from 'react-native';
import OutsidePressHandler from 'react-native-outside-press';
import { Card } from 'react-native-paper';
import AntDesign from 'react-native-vector-icons/AntDesign';
import Octicons from 'react-native-vector-icons/Octicons';

const verificationStat = {
  NOT_SENT: 'Verify',
  SENT: 'Sent',
  SENDING: 'Sending',
};

export default function ProfileScreen() {
  const { removeToken, userData: userDataContext } = useContext(AuthContext);
  const [editActiveFields, setEditActiveFields] = useState({
    name: false,
    image: false,
    location: true,
    verified: false,
    password: false,
  });
  const [userData, setUserData] = useState(userDataContext);
  const [profileImageSize, _] = useState(Dimensions.get('screen').width * 0.25);
  const [verificationText, setVerificationText] = useState(verificationStat.NOT_SENT);
  const [newPassword, setNewPassword] = useState('');
  const [reNewPassword, setReNewPassword] = useState('');

  function handleLocationPicked({ latitude, longitude }: LocationType) {
    console.log('Setting your new location to ', latitude, longitude);
    setEditActiveFields((f) => ({ ...f, location: false }));
  }

  function handleDelete() {
    Alert.alert('Confirm delete account?', 'Deleting your account is permanent and irreversible', [
      {
        text: 'Confirm',
        onPress: () => alert('Account Deleted'),
      },
      { text: 'Cancel', onPress: () => {} },
    ]);
  }

  function handlePasswordChange() {
    if (newPassword !== reNewPassword) alert("Passwords don't match");
    console.log('Invoking change password endpoint');
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
    <View style={styles.container}>
      <View style={styles.imageContainer}>
        <Image
          source={require('@/assets/icons/tea.png')}
          style={{ ...styles.profilePic, width: profileImageSize, height: profileImageSize }}
        />
        <View style={{ flexDirection: 'row', gap: 10, alignItems: 'center' }}>
          {userData?.verified ? (
            <Octicons name='verified' size={20} style={{ color: 'aqua' }} />
          ) : (
            <Octicons name='unverified' size={20} style={{ color: 'orange' }} />
          )}
          {editActiveFields.name ? (
            <OutsidePressHandler
              onOutsidePress={() => setEditActiveFields((f) => ({ ...f, name: false }))}
            >
              <TextInput value={userData?.name} onChangeText={(e) => {}} autoFocus />
            </OutsidePressHandler>
          ) : (
            <Text style={styles.name}>{userData?.name}</Text>
          )}
          {!editActiveFields.name && <AntDesign name='edit' size={20} />}
        </View>
      </View>

      <View>
        <Card>
          {editActiveFields.location ? (
            <LocationPicker
              onPicked={handleLocationPicked}
              onCancel={() => setEditActiveFields((f) => ({ ...f, location: false }))}
              currentCoords={userData ? userData.coords : undefined}
            />
          ) : (
            <EditableField
              label='Location:'
              value='Pokhara'
              onEditActive={() => setEditActiveFields((f) => ({ ...f, location: true }))}
            />
          )}

          <EditableField label='Email:' value={userData?.email} disable />
        </Card>
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
        style={{ borderTopWidth: 1, borderStyle: 'dashed', marginTop: 10, paddingTop: 10, gap: 10 }}
      >
        <Text style={{ fontSize: 18, fontWeight: 'bold' }}>Danger Zone</Text>
        <View
          style={{
            flexDirection: 'row',
            flexWrap: 'wrap',
            gap: 10,
            justifyContent: 'center',
          }}
        >
          {!editActiveFields.password ? (
            <Button
              title='Change Password'
              onPress={() => setEditActiveFields((f) => ({ ...f, password: true }))}
            />
          ) : (
            <Button title='Save' onPress={handlePasswordChange} />
          )}
          <Button title='Delete Account' onPress={handleDelete} />
        </View>
      </View>
    </View>
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
    paddingTop: Dimensions.get('screen').height * 0.04,
    paddingHorizontal: Dimensions.get('screen').width * 0.02,
    gap: 10,
  },
  imageContainer: {
    backgroundColor: 'green',
    justifyContent: 'center',
    alignItems: 'center',
    borderRadius: 20,
    paddingVertical: 10,
    position: 'relative',
    gap: 10,
  },
  profilePic: { borderWidth: 2, borderRadius: 20 },
  name: { fontSize: 20, fontWeight: 'bold' },
  componentInfo: {
    justifyContent: 'space-between',
    paddingVertical: 10,
    paddingHorizontal: 20,
    flexDirection: 'row',
  },
});
