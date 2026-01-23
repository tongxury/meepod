import {Chip} from "@rneui/themed";
import {StyleProp, View, ViewStyle} from "react-native";
import {Text} from "react-native-paper";
import React from "react";
import {Stack} from "@react-native-material/core";

const Ball = ({title, color, style}: {
    title: string,
    color?: 'blue' | 'red' | 'dred' | 'darkblue' | 'gold',
    style?: StyleProp<ViewStyle>;
}) => {

    // "#4e059a", selected: "#f53a3a", unselected: '#faa199'
    // selected: "#346dfc", unselected: '#99bbfa'
    const colorTap = {
        red: '#f53a3a',
        blue: '#346dfc',
        dred: '#f53aba',
        darkblue: '#092567',
        gold: '#f6ac36'
    }

    return <Stack center={true}
                  style={Object.assign({
                      borderRadius: 13,
                      width: 26,
                      height: 26,
                      backgroundColor: colorTap[color || 'red']
                  }, style)}>
        <Text style={{color: 'white'}}>{title}</Text>
    </Stack>
}

export default Ball