import {Stack} from "@react-native-material/core";
import {StyleProp, Text, ViewStyle} from "react-native"
import * as React from "react";


const Tag = ({children, color, style}: {
    color?: string,
    children: React.ReactNode,
    style?: StyleProp<ViewStyle>
}) => {

    return <Stack style={style} center={true} pv={1} ph={5} radius={3} bg={color}>
        <Text style={{color: 'white', fontSize: 10}} numberOfLines={1}>
            {children}
        </Text>
    </Stack>

}

export default Tag