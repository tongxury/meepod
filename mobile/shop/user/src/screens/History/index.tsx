import {View} from "react-native";
import React, {useState} from "react";
import OrderList from "../Order/List";
import OrderGroupList from "../OrderGroup/List";
import PlanList from "../Plan/List";
import {Text, useTheme} from 'react-native-paper';
import {mainBodyHeight} from "../../utils/dimensions";
import {useSafeAreaInsets} from "react-native-safe-area-context";
import Tabs from "../../components/Tabs";

const HistoryScreen = ({navigation}) => {

    const {colors} = useTheme()
    const {top} = useSafeAreaInsets();


    return <View style={{height: mainBodyHeight, paddingTop: top}}>
        <Tabs style={{flex: 1}} tabs={[
            {key: 'order', title: '订单', component: () => <OrderList category="order"/>},
            {key: 'my', title: '合买', component: () => <OrderGroupList category="my"/>},
            {key: 'saved', title: '方案', component: () => <PlanList category="saved"/>}
        ]}/>
    </View>
}

export default HistoryScreen