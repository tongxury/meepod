import {View} from "react-native";
import {Appbar, Text, useTheme} from "react-native-paper";
import {useSafeAreaInsets} from "react-native-safe-area-context";
import {mainBodyHeight} from "../../utils/dimensions";
import Tabs from "../../components/Tabs";
import PlanList from "../Plan/List";
import OrderList from "../Order/List";
import React from "react";
import FollowableOrderList from "../Order/FollowableList";
import OrderGroupList from "../OrderGroup/List";

const MoreScreen = ({navigation}) => {

    const {colors} = useTheme()

    const {top} = useSafeAreaInsets();

    return <View style={{flex: 1}}>
        <Appbar.Header>
            <Appbar.BackAction onPress={() => {
                navigation.goBack()
            }}/>
            <Appbar.Content title={<Text variant="titleMedium">更多玩法</Text>}/>
        </Appbar.Header>

        <Tabs tabs={[
            {key: 'group', title: '店内合买', component: () => <OrderGroupList category="joinable"/>},
            // {key: 'follow', title: '店内跟单', component: () => <FollowableOrderList/>},
        ]}/>
    </View>
}

export default MoreScreen