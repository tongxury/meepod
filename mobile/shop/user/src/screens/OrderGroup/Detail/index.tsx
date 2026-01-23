import {View} from "react-native";
import {useRequest} from "ahooks";
import {fetchOrderGroupDetail} from "../../../service/api";
import React, {useState} from "react";
import {Appbar, Text, Button, useTheme} from "react-native-paper";
import GroupSummary from "../View/Summary";
import {Stack} from "@react-native-material/core";
import {TicketView} from "../../Plan/View/TicketView";
import TicketListView from "../../Plan/View/TicketListView";
import ImageViewer from "../../../components/ImageViewer";
import DetailTrigger from "./Detail";
import JoinGroupTrigger from "../../../triggers/OrderGroup/Join";

const GroupDetailScreen = ({route, navigation}) => {
    const {id, biz_category} = route.params;

    const {data, loading, refresh} = useRequest(() => fetchOrderGroupDetail({id, biz_category}),)
    const order = data?.data?.data

    const {colors} = useTheme()

    return <View style={{flex: 1}}>
        <Appbar.Header>
            <Appbar.BackAction onPress={() => {
                navigation.goBack()
            }}/>
            <Appbar.Content title={<Text variant={"titleMedium"}>合买订单-{order?.id}</Text>}/>
        </Appbar.Header>
        <Stack fill={1} spacing={5}>
            <Stack bg={colors.background} >
                <GroupSummary data={order}/>
                <DetailTrigger orderId={order?.id}/>
            </Stack>
            <Stack bg={colors.background} p={10} spacing={10}>
                <Text variant="titleSmall">投注内容</Text>
                <TicketListView itemId={order?.plan?.item?.id} data={order?.plan?.tickets}/>
            </Stack>
            {order?.prized && order?.plan?.issue?.result &&
                <Stack bg={colors.background} p={10} spacing={10}>
                    <Text variant="titleSmall">开奖结果:</Text>
                    <TicketView itemId={order?.plan?.item?.id} data={order?.plan?.issue?.result}/>
                </Stack>
            }
            {order?.ticket_images?.length > 0 &&
                <Stack bg={colors.background} p={10} spacing={10}>
                    <Text variant="titleSmall">投注票样</Text>
                    <ImageViewer images={order?.ticket_images?.map(t => {
                        return {url: t}
                    })}/>
                </Stack>
            }

        </Stack>
        <View style={{padding: 15}}>
            {
                order?.joinable && <JoinGroupTrigger
                    data={{groupId: order?.id, volumeLeft: order?.volume - order?.volume_ordered, floor: order?.floor}}
                    onSubmitted={(orderId, groupId) => {
                        refresh()
                    }}>
                    <Button mode="contained" style={{flex: 1}}>参与合买</Button>
                </JoinGroupTrigger>
            }


        </View>
    </View>
}


export default GroupDetailScreen