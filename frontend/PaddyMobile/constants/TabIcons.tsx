import pages from './screens';
import AntDesign from 'react-native-vector-icons/AntDesign';
import MaterialCommunityIcons from 'react-native-vector-icons/MaterialCommunityIcons';

export const pagesIcons = {
  [pages.profile]: <AntDesign name='user' size={40} />,
  [pages.dashboard]: <AntDesign name='home' size={40} />,
  [pages.upload]: <AntDesign name='upload' size={40} />,
  [pages.live]: <AntDesign name='find' size={40} />,
  [pages.heatMap]: <MaterialCommunityIcons name='google-maps' size={40} />
};
