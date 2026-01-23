import {HStack, Stack} from "@react-native-material/core";
import {Avatar, Text} from "react-native-paper";
import React from "react";
import {StyleProp, ViewStyle} from "react-native";


const TitleView = ({icon, title, subTitle, extra, style}: {
    icon: string,
    title: string,
    subTitle?: string,
    extra?: React.ReactNode
    style?: StyleProp<ViewStyle>
}) => {

    return <HStack style={style} items={"center"} justify={"between"}>
        <HStack items={"center"} spacing={8}>
            <Avatar.Image size={40} source={{uri: icon}}/>
            <Stack spacing={3}>
                <Text variant={"titleMedium"}>{title}</Text>
                <Text variant={"labelSmall"}>{subTitle ?? '-'}</Text>
            </Stack>
        </HStack>

        {extra && extra}
    </HStack>
}

export default TitleView