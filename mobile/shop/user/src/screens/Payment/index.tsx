import {Appbar, Text} from "react-native-paper";
import React from "react";
import {View} from "react-native";
import Tabs from "../../components/Tabs";
import TopupList from "./TopupList";
import WithdrawList from "./WithdrawList";
import PaymentList from "./PaymentList";
import {mainBodyHeight} from "../../utils/dimensions";
import {useSafeAreaInsets} from "react-native-safe-area-context";

const PaymentScreen = ({navigation, route}) => {

    const {tab} = route.params;
    const {top} = useSafeAreaInsets();


    return <View style={{flex: 1}}>
        <Appbar.Header>
            <Appbar.BackAction onPress={() => {
                navigation.goBack()
            }}/>
            <Appbar.Content title={<Text variant="titleMedium">现金记录</Text>}/>
        </Appbar.Header>
        <Tabs
            style={{flex: 1}}
            tabs={[
                {key: 'pay', title: '消费', component: () => <PaymentList/>},
                {key: 'topUp', title: '充值', component: () => <TopupList/>},
                {key: 'withdraw', title: '提现', component: () => <WithdrawList/>},
            ]}
            current={tab}
        />
    </View>
}

export default PaymentScreen