import { uploader } from '@/api/driver';
import endpoints from '@/constants/endpoints';
import { AuthContext } from '@/contexts/auth/auth-provider';
import { MediaTypeOptions, launchCameraAsync, launchImageLibraryAsync } from 'expo-image-picker';
import { useContext, useState } from 'react';
import { ActivityIndicator, Pressable, StyleSheet, Text, View } from 'react-native';
import AntDesign from 'react-native-vector-icons/AntDesign';
import Entypo from 'react-native-vector-icons/Entypo';

export default function UploadNew({ onUpload }: { onUpload: () => void }) {
  const [isSending, setIsSending] = useState(false);
  const { token } = useContext(AuthContext);

  async function handleUpload(image?: string) {
    if (!image) return alert('Select an image first');
    if (!token) return alert('Not authorized.');
    setIsSending(true);

    try {
      const resp = await uploader({
        fileUri: image,
        fieldName: 'images',
        uri: endpoints.uploadImage,
        token,
      });
      if (resp.status !== 200) throw resp;
      try {
        const respMessage = JSON.parse(resp.body);
        alert(respMessage.message);
      } catch (e) {}
    } catch (e) {
      alert("Couldn't upload image");
    }
    setIsSending(false);
    onUpload();
  }

  async function pickImage(cameraMode?: boolean) {
    let resultFxn = cameraMode ? launchCameraAsync : launchImageLibraryAsync;

    const result = await resultFxn({
      mediaTypes: MediaTypeOptions.All,
      allowsEditing: true,
      aspect: [4, 3] as any,
      quality: 1,
    });

    if (!result.canceled) {
      const uri = result.assets[0].uri;
      handleUpload(uri);
    }
  }

  if (isSending)
    return (
      <View style={styles.pickerContainer}>
        <ActivityIndicator size={40} />
        <Text>Uploading Image</Text>
      </View>
    );

  return (
    <View style={styles.pickerContainer}>
      <AntDesign name='clouduploado' size={40} />
      <Text>Upload</Text>
      <Pressable onPress={() => pickImage()} style={styles.picker}>
        <Entypo name='images' size={30} />
      </Pressable>
      <Pressable onPress={() => pickImage(true)} style={styles.picker}>
        <AntDesign name='camera' size={30} />
      </Pressable>
    </View>
  );
}

const styles = StyleSheet.create({
  optionContainer: {
    flexDirection: 'row',
    padding: 10,
    borderWidth: 1,
    borderBottomWidth: 0,
    borderRadius: 10,
    gap: 10,
    width: '100%',
    justifyContent: 'center',
  },
  pickerContainer: {
    flexDirection: 'row',
    alignItems: 'center',
    gap: 10,
    borderWidth: 1,
    borderRadius: 10,
    justifyContent: 'center',
    padding: 10,
    borderStyle: 'dashed',
    zIndex: 0.5,
  },
  picker: {
    borderWidth: 1,
    padding: 10,
    borderRadius: 10,
    zIndex: 0.5,
  },
  pickedImage: {
    width: 100,
    height: 'auto',
    aspectRatio: 1,
  },
});
