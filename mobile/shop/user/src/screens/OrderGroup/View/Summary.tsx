import {StyleProp, View, ViewStyle} from "react-native";
import {OrderGroup, PageData, Result} from "../../../service/typs";
import {Avatar, Text, useTheme} from "react-native-paper";
import {WhiteSpace, WingBlank} from "@ant-design/react-native";
import {Chip} from "@rneui/themed";
import Progress from "./Progress";
import {HStack, Stack} from "@react-native-material/core";
import ItemView from "../../../components/ItemView";
import AmountView from "../../../components/AmountView";
import group from "../../Group";
import React from "react";

const GroupSummary = ({data, style}: { data: OrderGroup, style?: StyleProp<ViewStyle> | undefined }) => {

    const {colors} = useTheme()
    return <Stack bg={colors.background} p={15} spacing={10} style={style}>
        <HStack items={"center"} justify={"between"}>
            <HStack items={"center"} spacing={5}>
                <ItemView item={data?.plan?.item} issue={data?.plan?.issue}/>
            </HStack>
            <Chip color={data?.status?.color} size="sm" radius={5}>{data?.status?.name}</Chip>
        </HStack>
        <HStack items={"center"} spacing={5}>
            {data?.tags?.map((t, i) => <Chip key={i} color={t.color} size="sm" radius={5}>{t.title}</Chip>)}
        </HStack>
        <Progress data={data}/>
        <HStack items={"center"}>
            <AmountView amount={data?.plan?.amount} multiple={data?.plan?.multiple}/>
        </HStack>
        <Text variant={"bodySmall"}>{data?.remark}</Text>
    </Stack>
}

export default GroupSummary