import {SafeAreaView, StyleProp, View, ViewStyle} from "react-native";
import {useSafeAreaInsets} from "react-native-safe-area-context";
import {Appbar, Text, useTheme} from "react-native-paper";
import Tabs from "../../components/Tabs";
import React from "react";
import PaymentList from "./Payment";
import OutCoStoreList from "./OutList";
import InCoStoreList from "./InList";
import {Stack} from "@react-native-material/core";
import {WhiteSpace} from "@ant-design/react-native";

const CooperationScreen = ({navigation}) => {
    const insets = useSafeAreaInsets();

    const {colors} = useTheme()

    // return <SafeAreaView style={{flex: 1, paddingTop: insets.top}}>
    //     <Tabs style={{flex: 1}} tabs={[
    //         {key: '1', title: '合作转出', component: () => <OutCoStoreList/>},
    //         {key: '2', title: '合作转入', component: () => <InCoStoreList/>},
    //         // {key: '3', title: '账单明细', component: () => <PaymentList/>},
    //     ]}/>
    // </SafeAreaView>

    return <View style={{flex: 1}}>
        <Appbar.Header>
            <Appbar.BackAction onPress={() => {
                navigation.goBack()
            }}/>
            <Appbar.Content title={<Text variant="titleMedium">店铺合作</Text>}/>
        </Appbar.Header>
        <Stack fill={true}>
            <Tabs style={{flex: 1}} tabs={[
                {key: '1', title: '合作转出', component: () => <OutCoStoreList/>},
                {key: '2', title: '合作转入', component: () => <InCoStoreList/>},
                // {key: '3', title: '账单明细', component: () => <PaymentList/>},
            ]}/>
        </Stack>
    </View>
};

export default CooperationScreen
