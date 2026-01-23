import {FlatList, Pressable, View} from "react-native";
import React, {useCallback, useContext, useState} from "react";
import {useInfiniteScroll, useRequest} from "ahooks";
import {Appbar, Avatar, Banner, Card, useTheme} from "react-native-paper";
import {useFocusEffect, useNavigation} from "@react-navigation/native";
import {Text} from "react-native-paper";
import {HStack, Stack} from "@react-native-material/core";
import {Empty, Footer} from "../../components/ListComponent";
import {fetchCoStorePayments,} from "../../service/api";
import {Selector} from "../../components/Selector";
import StoreView from "../../components/StoreView";
import MoneyView from "../../components/MoneyView";
import {AppContext} from "../../providers/global";
import {WhiteSpace} from "@ant-design/react-native";
import List from "./List";

const PaymentScreen = ({}) => {

    const navigation = useNavigation()

    const {colors} = useTheme()

    return <Stack style={{flex: 1}}>
        <Appbar.Header>
            <Appbar.BackAction onPress={() => {
                navigation.goBack()
            }}/>
            <Appbar.Content title={<Text variant="titleMedium">流水明细</Text>}/>
        </Appbar.Header>
        <WhiteSpace/>
        <List/>
    </Stack>


}

export default PaymentScreen