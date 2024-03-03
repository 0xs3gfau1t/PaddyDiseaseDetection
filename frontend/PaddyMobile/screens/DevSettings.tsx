import { Text, TextInput, View } from 'react-native';
import { useContext } from 'react';
import { AuthContext } from '@/contexts/auth/auth-provider';

export default function DevSettings() {
    const { apiUrl, setApiUrl } = useContext(AuthContext);

    return (
        <View>
            <Text>Backend Url</Text>
            <TextInput value={apiUrl} onChangeText={(e) => setApiUrl(e)} />
        </View>
    );
}
