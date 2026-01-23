import {StyleProp, useWindowDimensions, View, ViewStyle} from "react-native";
import React, {useCallback, useMemo, useState} from "react";
import {Text, useTheme} from 'react-native-paper';
import {mainBodyHeight} from "../../utils/dimensions";
import {TabView, SceneMap, TabBar} from 'react-native-tab-view';


declare type TabItem = {
    key: string,
    title: string,
    component: () => React.ReactNode
}

const Tabs = ({tabs, current, style}: {
    tabs: TabItem[],
    current?: number,
    style?: StyleProp<ViewStyle> | undefined
}) => {

    const {colors} = useTheme()
    const layout = useWindowDimensions();

    const [index, setIndex] = useState(current ?? 0);

    const [routes] = useState(tabs);
    const renderScene = ({route}) => {
        return tabs.filter(t => t.key === route.key)[0].component()
    };

    const {fontScale} = useWindowDimensions()

    return <View style={style}>
        <TabView
            lazy
            navigationState={{index, routes}}
            renderScene={renderScene}
            onIndexChange={setIndex}
            initialLayout={{width: layout.width}}
            renderTabBar={props =>
                <TabBar {...props}
                        indicatorStyle={{backgroundColor: colors.primary}}
                        indicatorContainerStyle={{backgroundColor: colors.background}}
                        activeColor={colors.primary}
                        inactiveColor={colors.onBackground}
                        labelStyle={{fontWeight: "bold"}}

                />
            }
        />
    </View>

}

export default Tabs