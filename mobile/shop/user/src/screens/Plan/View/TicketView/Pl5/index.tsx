import {Text, View} from "react-native";
import React from "react";
import Ball from "../../../../../components/Ball";
import {useTheme} from "react-native-paper";

const X7cTicketView = ({data}) => {
    const {colors} = useTheme()

    const Divider = () => {
        return <View style={{width: 3, height: 10, marginHorizontal: 5, backgroundColor: colors.primary}}/>
    }

    return <View style={{flexDirection: "row", alignItems: "center", flexWrap: "wrap",}}>
        {data.wan?.map(t => <Ball color='red' key={t} style={{margin: 2}} title={t}/>)}
        <Divider/>
        {data.qian?.map(t => <Ball color='red' key={t} style={{margin: 2}} title={t}/>)}
        <Divider/>
        {data.bai?.map(t => <Ball color='red' key={t} style={{margin: 2}} title={t}/>)}
        <Divider/>
        {data.shi?.map(t => <Ball color='red' key={t} style={{margin: 2}} title={t}/>)}
        <Divider/>
        {data.gen?.map(t => <Ball color='red' key={t} style={{margin: 2}} title={t}/>)}
    </View>
}

export default X7cTicketView