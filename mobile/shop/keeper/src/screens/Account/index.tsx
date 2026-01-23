import React, {useState} from "react";
import {SafeAreaView, Text, View} from "react-native";
import AccountList from "./AccountList";
import WithdrawList from "./WithdrawList";
import Tabs from "../../components/Tabs";
import {useSafeAreaInsets} from "react-native-safe-area-context";
import {useTheme} from "react-native-paper";
import RewardList from "./RewardList";


const AccountScreen = ({navigation}) => {

    const tabs = [
        {key: 'account', title: '账本', component: () => <AccountList/>},
        {key: 'withdraw', title: '提现', component: () => <WithdrawList/>},
        {key: 'reward', title: '兑奖', component: () => <RewardList/>},
    ]

    const insets = useSafeAreaInsets();

    const {colors} = useTheme()

    return <SafeAreaView style={{flex: 1, paddingTop: insets.top}}>
        <Tabs style={{flex: 1}} tabs={tabs}/>
    </SafeAreaView>
}

export default AccountScreen