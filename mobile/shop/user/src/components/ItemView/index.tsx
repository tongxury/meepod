import { StyleProp, View, ViewStyle } from "react-native";
import { Avatar, Text, useTheme } from "react-native-paper";
import { WingBlank } from "@ant-design/react-native";
import React from "react";
import { HStack, Stack } from "@react-native-material/core";
import { Issue, Item } from "../../service/typs";
import { Skeleton } from "@rneui/themed";
import { lotteryIcons } from '../../utils/lotteryIcons';


const ItemView = ({ item, issue, loading, style }: {
    item: Item,
    issue?: Issue,
    loading?: boolean,
    style?: StyleProp<ViewStyle> | undefined
}) => {

    const { colors } = useTheme()

    if (loading) {
        return <HStack items={"center"} spacing={5} style={style}>

            <Skeleton circle width={40} height={40} />
            <Stack spacing={6}>
                <Skeleton width={60} height={15} />
                <Skeleton width={100} height={10} />
            </Stack>
        </HStack>
    }


    return <HStack items={"center"} spacing={5} style={style}>
        <Avatar.Image style={{ backgroundColor: colors.background }} size={40}
            source={lotteryIcons[item?.id] || { uri: item?.icon }} />
        <Stack spacing={6}>
            <Text variant="titleSmall">{item?.name}</Text>
            <Text variant="labelSmall">{issue ? `第${issue.index}期` : `-`}</Text>
        </Stack>
    </HStack>
}

export default ItemView