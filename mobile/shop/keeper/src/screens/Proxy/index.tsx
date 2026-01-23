import AccountList from "../Account/AccountList";
import WithdrawList from "../Account/WithdrawList";
import {useSafeAreaInsets} from "react-native-safe-area-context";
import {useTheme} from "react-native-paper";
import {SafeAreaView} from "react-native";
import Tabs from "../../components/Tabs";
import React from "react";
import ProxyList from "./List";
import List from "./Reward/List";


const ProxyScreen = ({navigation}) => {

    const tabs = [
        {key: '1', title: '推广员', component: () => <ProxyList/>},
        {key: '2', title: '佣金结算', component: () => <List />},
    ]

    const insets = useSafeAreaInsets();

    const {colors} = useTheme()

    return <SafeAreaView style={{flex: 1, paddingTop: insets.top}}>
        <Tabs style={{flex: 1}} tabs={tabs}/>
    </SafeAreaView>
}

export default ProxyScreen