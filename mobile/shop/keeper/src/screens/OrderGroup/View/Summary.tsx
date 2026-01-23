import {StyleProp, View, ViewStyle} from "react-native";
import {OrderGroup, PageData, Result} from "../../../service/typs";
import {Avatar, Text, useTheme} from "react-native-paper";
import {Chip} from "@rneui/themed";
import Progress from "./Progress";
import {HStack, Stack} from "@react-native-material/core";
import ItemView from "../../../components/ItemView";
import AmountView from "../../../components/AmountView";
import React from "react";

const GroupSummary = ({data, style}: { data: OrderGroup, style?: StyleProp<ViewStyle> | undefined }) => {

    const {colors} = useTheme()
    return <Stack bg={colors.background} spacing={15} style={style}>
        <Stack spacing={8}>
            <HStack items={"center"} justify={"between"}>
                <ItemView item={data?.plan?.item} issue={data?.plan?.issue}/>
                <Chip color={data?.status?.color} size="sm" radius={5}>{data?.status?.name}</Chip>
            </HStack>
            <HStack items={"center"} spacing={5}>
                {data?.tags?.map((t, i) => <Chip color={t.color} size="sm" radius={5}>{t.title}</Chip>)}
            </HStack>
            <HStack items={"center"} justify={"between"}>
                <HStack items={"center"} spacing={5}>
                    <Avatar.Image size={20} source={{uri: data?.store?.icon}}/>
                    <Text variant={"titleSmall"}>{data?.store?.name}</Text>
                </HStack>
                {data?.to_store && <Text variant={'labelSmall'}>转给</Text>}
                {data?.to_store &&
                    <HStack items={"center"} spacing={5}>
                        <Avatar.Image size={20} source={{uri: data?.to_store?.icon}}/>
                        <Text variant={"titleSmall"}>{data?.to_store?.name}</Text>
                    </HStack>
                }
            </HStack>
        </Stack>

        <Progress data={data}/>
        <HStack items={"center"}>
            <AmountView amount={data?.plan?.amount} multiple={data?.plan?.multiple}/>
        </HStack>
        <Text variant={"bodySmall"}>{data?.remark}</Text>
    </Stack>
}

export default GroupSummary