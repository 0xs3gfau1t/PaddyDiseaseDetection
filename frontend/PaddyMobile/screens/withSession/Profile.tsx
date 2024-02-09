import LocationPicker from '@/components/profile/LocationPicker';
import { VERIFICATION_EXPIRY_TIME } from '@/constants/misc';
import { AuthContext } from '@/contexts/auth/auth-provider';
import { LocationType } from '@/types/misc';
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
  const { removeToken, userData: userDataContext } = useContext(AuthContext);
  const [editActiveFields, setEditActiveFields] = useState({
    name: false,
    image: false,
    location: false,
    verified: false,
    password: false,
  });
  const [userData, setUserData] = useState(userDataContext);
  const [profileImageSize, _] = useState(Dimensions.get('screen').width * 0.25);
  const [verificationText, setVerificationText] = useState(verificationStat.NOT_SENT);
  const [oldPassword, setOldPassword] = useState('');
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
          <View style={styles.bgContainer} />
        </Card>

        <View>
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
              alignItems: editActiveFields.password ? 'center' : 'stretch',
              width: editActiveFields.password ? '100%' : '50%',
              gap: 10,
              justifyContent: 'center',
            }}
          >
            {editActiveFields.password && (
              <Card
                style={{
                  width: '100%',
                  padding: Dimensions.get('screen').width * 0.02,
                  backgroundColor: '#e87e72',
                }}
              >
                <View style={styles.passwordInputContainer}>
                  <Text style={styles.label}>Old Password</Text>
                  <TextInput onChangeText={(e) => setOldPassword(e)} style={styles.input} />
                </View>
                <View style={styles.passwordInputContainer}>
                  <Text style={styles.label}>New Password</Text>
                  <TextInput onChangeText={(e) => setNewPassword(e)} style={styles.input} />
                </View>
                <View style={styles.passwordInputContainer}>
                  <Text style={styles.label}>Repeat New Password</Text>
                  <TextInput onChangeText={(e) => setReNewPassword(e)} style={styles.input} />
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
                    onPress={() => setEditActiveFields((f) => ({ ...f, password: false }))}
                  />
                </View>
              </Card>
            )}
            {!editActiveFields.password && (
              <Button
                title='Change Password'
                onPress={() => setEditActiveFields((f) => ({ ...f, password: true }))}
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
  },
  passwordInputContainer: {
    flexDirection: 'row',
    marginVertical: 5,
    alignItems: 'center',
  },
});
