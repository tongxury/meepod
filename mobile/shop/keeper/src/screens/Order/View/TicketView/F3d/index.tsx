import {View} from "react-native";
import React from "react";
import Ball from "../../../../../components/Ball";
import {useTheme, Text,} from "react-native-paper";

const F3dTicketView = ({data}) => {
    const {colors} = useTheme()

    return <View>
        {data.cat == 'result' &&
            <View style={{flexDirection: "row", alignItems: "center", flexWrap: "wrap"}}>
                {data.bai?.map(t => <Ball key={t} style={{margin: 2}} title={t}/>)}
                <View style={{width: 3, height: 10, marginHorizontal: 5, backgroundColor: colors.primary}}/>
                {data.shi?.map(t => <Ball key={t} style={{margin: 2}} title={t}/>)}
                <View style={{width: 3, height: 10, marginHorizontal: 5, backgroundColor: colors.primary}}/>
                {data.gen?.map(t => <Ball key={t} style={{margin: 2}} title={t}/>)}
            </View>
        }
        {data.cat == 'z1' &&
            <View style={{flexDirection: "row", alignItems: "center", flexWrap: "wrap"}}>
                <Text variant={"titleMedium"}>直选</Text>
                <View style={{marginLeft: 5, flexDirection: "row", alignItems: "center", flexWrap: "wrap"}}>
                    {data.bai?.map(t => <Ball key={t} style={{margin: 2}} title={t}/>)}
                    <View style={{width: 3, height: 10, marginHorizontal: 5, backgroundColor: colors.primary}}/>
                    {data.shi?.map(t => <Ball key={t} style={{margin: 2}} title={t}/>)}
                    <View style={{width: 3, height: 10, marginHorizontal: 5, backgroundColor: colors.primary}}/>
                    {data.gen?.map(t => <Ball key={t} style={{margin: 2}} title={t}/>)}
                </View>
            </View>
        }
        {
            data.cat == 'z3' && <View style={{flexDirection: "row", alignItems: "center", flexWrap: "wrap"}}>
                <Text variant={"titleMedium"}>组三</Text>
                <View style={{marginLeft: 5, flexDirection: "row", alignItems: "center", flexWrap: "wrap"}}>
                    {data.ton?.map(t => <Ball key={t} style={{margin: 2}} title={t}/>)}
                </View>
            </View>
        }
        {
            data.cat == 'z6' && <View style={{flexDirection: "row", alignItems: "center", flexWrap: "wrap"}}>
                <Text variant={"titleMedium"}>组六</Text>
                <View style={{flexDirection: "row", alignItems: "center", flexWrap: "wrap"}}>
                    {data.ton?.map(t => <Ball key={t} style={{margin: 2}} title={t}/>)}
                </View>
            </View>
        }

    </View>
}

export default F3dTicketView