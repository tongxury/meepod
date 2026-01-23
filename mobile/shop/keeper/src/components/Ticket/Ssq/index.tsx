import {Text, View} from "react-native";
import React from "react";
import Ball from "../../Ball";

const SsqMetaView = ({data}) => {
    return <View style={{flexDirection: "row", alignItems: "center", flexWrap: "wrap",}}>
        {data.red?.map(t => <Ball key={t} style={{margin: 2}} title={t}/>)}
        {data.redD?.map(t => <Ball key={t} style={{margin: 2}} title={t} color='dred'/>)}
        {data.blue?.map(t => <Ball key={t} style={{margin: 2}} title={t} color='blue'/>)}
    </View>
}

export default SsqMetaView