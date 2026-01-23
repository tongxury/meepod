import {Text, View} from "react-native";
import React from "react";
import Ball from "../../../../../components/Ball";
import {Flex} from "@ant-design/react-native";

const SsqMetaView = ({data}) => {
    return <View style={{flexDirection: "row", alignItems: "center", flexWrap: "wrap"}}>
        {data.red?.map(t => <Ball key={t} style={{margin: 2}} title={t}/>)}
        {data.redD?.map(t => <Ball key={t} style={{margin: 2}} title={t} color='dred'/>)}
        {data.blue?.map(t => <Ball key={t} style={{margin: 2}} title={t} color='blue'/>)}
    </View>
    // return <Flex align="center" wrap="wrap" style={{flex: 1}}>
    //     {data.red?.map(t => <Ball key={t} style={{margin: 2}} title={t}/>)}
    //     {data.redD?.map(t => <Ball key={t} style={{margin: 2}} title={t} color='dred'/>)}
    //     {data.blue?.map(t => <Ball key={t} style={{margin: 2}} title={t} color='blue'/>)}
    // </Flex>
}

export default SsqMetaView