import {Appbar, Text, useTheme} from "react-native-paper";
import React from "react";
import {View} from "react-native";
import {Stack} from "@react-native-material/core";
import PaymentList from "./List";

const CoStorePaymentScreen = ({navigation, route}) => {

    const {storeId, coStoreId} = route.params;

    const {colors} = useTheme()

    return <View style={{flex: 1}}>
        <Appbar.Header>
            <Appbar.BackAction onPress={() => {
                navigation.goBack()
            }}/>
            <Appbar.Content title={<Text variant="titleMedium">资金记录</Text>}/>
        </Appbar.Header>
        <Stack fill={true}>
            <PaymentList storeId={storeId} coStoreId={coStoreId}/>
        </Stack>
    </View>
}


export default CoStorePaymentScreen