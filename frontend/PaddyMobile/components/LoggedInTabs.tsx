import { pagesIcons } from '@/constants/TabIcons';
import pages from '@/constants/screens';
import { useMemo } from 'react';
import { View, TouchableOpacity } from 'react-native';
import { Card } from 'react-native-paper';

export default function LoggedInTabs({ state, descriptors, navigation }: any) {
  const viewableTabs = useMemo(
    () => state.routes.filter((r: any) => r.name !== pages.detail),
    [state.routes]
  );
  return (
    <View style={{ flexDirection: 'row', justifyContent: 'center' }}>
      {viewableTabs.map((route: any, index: number) => {
        const { options } = descriptors[route.key];
        const label =
          options.tabBarLabel !== undefined
            ? options.tabBarLabel
            : options.title !== undefined
              ? options.title
              : route.name;

        const isFocused = state.index === index;

        const onPress = () => {
          const event = navigation.emit({
            type: 'tabPress',
            target: route.key,
            canPreventDefault: true,
          });

          if (!isFocused && !event.defaultPrevented) {
            navigation.navigate(route.name, route.params);
          }
        };

        const onLongPress = () => {
          navigation.emit({
            type: 'tabLongPress',
            target: route.key,
          });
        };

        return (
          <TouchableOpacity
            accessibilityRole='button'
            accessibilityState={isFocused ? { selected: true } : {}}
            accessibilityLabel={options.tabBarAccessibilityLabel}
            testID={options.tabBarTestID}
            onPress={onPress}
            onLongPress={onLongPress}
            style={{ padding: 5 }}
            key={index}
          >
            <Card style={{ padding: 10 }}>{pagesIcons[label]}</Card>
          </TouchableOpacity>
        );
      })}
    </View>
  );
}
