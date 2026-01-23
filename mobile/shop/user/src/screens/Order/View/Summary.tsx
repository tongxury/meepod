import {StyleProp, ViewStyle, Text} from "react-native";
import React from "react";
import {View} from "react-native";
import {Chip} from "@rneui/themed";
import {Avatar, useTheme} from "react-native-paper";

const OrderSummary = ({data}) => {

    const {colors} = useTheme()
    return <View>
        <View style={{flexDirection: "row"}}>
            <Chip titleStyle={{marginHorizontal: 8}} type="outline" size="sm"
                  icon={<Avatar.Image  style={{backgroundColor: colors.background}}  size={20} source={{uri: data?.plan?.item?.icon}}/>}
                  title={data?.plan?.item?.name}></Chip>
        </View>

    </View>
}

export default OrderSummary