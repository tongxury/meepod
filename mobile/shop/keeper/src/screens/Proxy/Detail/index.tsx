import {StyleProp, View, ViewStyle} from "react-native";
import {Stack} from "@react-native-material/core";
import {Appbar, Text} from "react-native-paper";
import React from "react";
import UserList from "./UserList";


const ProxyDetailScreen = ({navigation, route}) => {

    const {id} = route.params;

    return <Stack fill={1} spacing={5}>
        <Appbar.Header>
            <Appbar.BackAction onPress={() => {
                navigation.goBack()
            }}/>
            <Appbar.Content title={<Text variant="titleMedium">推广员详情-{id}</Text>}/>
        </Appbar.Header>
        <Stack fill={true} spacing={5}>
            <UserList proxyId={id}/>
        </Stack>
    </Stack>
}

export default ProxyDetailScreen
