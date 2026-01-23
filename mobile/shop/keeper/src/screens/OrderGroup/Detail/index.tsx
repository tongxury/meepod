import {View} from "react-native";
import {useRequest} from "ahooks";
import {
    fetchOrderGroupDetail,
    switchOrder,
    switchOrderGroup,
    ticketOrder,
    ticketOrderGroup
} from "../../../service/api";
import React, {useState} from "react";
import {Appbar, Text, Button, useTheme} from "react-native-paper";
import GroupSummary from "../View/Summary";
import Tabs from "../../../components/Tabs";
import {HStack, Stack} from "@react-native-material/core";
import TicketListView from "../../Order/View/TicketListView";
import TicketTrigger from "../../../triggers/Order/Ticket";
import {TicketView} from "../../Order/View/TicketView";
import ImageViewer from "../../../components/ImageViewer";
import DetailTrigger from "./Detail";
import SwitchTrigger from "../../../triggers/Order/Switch";

const GroupDetailScreen = ({route, navigation}) => {
    const { id, biz_category} = route.params;

    const {data, mutate, loading, refresh} = useRequest(() => fetchOrderGroupDetail({id, biz_category}),)
    const order = data?.data?.data

    const {colors} = useTheme()

    const onTicket = ({orderId, images}) => {
        ticketOrderGroup({orderId, images}).then(rsp => {
            if (rsp.data?.code == 0) {
                mutate(rsp)
            }
        })
    }

    const onSwitch = ({orderId, toStoreId}) => {
        switchOrderGroup({orderId, toStoreId}).then(rsp => {
            if (rsp.data?.code == 0) {
                mutate(rsp)
            }
        })
    }


    return <View style={{flex: 1}}>
        <Appbar.Header>
            <Appbar.BackAction onPress={() => {
                navigation.goBack()
            }}/>
            <Appbar.Content title={<Text variant={"titleMedium"}>合买订单-{order?.id}</Text>}/>
        </Appbar.Header>
        <Stack fill={1} spacing={5}>
            <Stack bg={colors.background} p={10} spacing={10}>
                <GroupSummary data={order}/>
                <DetailTrigger orderId={order?.id}/>
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
            <Stack bg={colors.background} p={10} spacing={10}>
                <Text variant="titleSmall">投注内容(拆票)</Text>
                {/*<TicketListView itemId={order?.plan?.item?.id} data={order?.plan?.split_tickets}/>*/}
                <TicketListView itemId={order?.plan?.item?.id} data={order?.plan?.split_tickets}/>
            </Stack>
        </Stack>
        <View>
            {
                order?.ticketable && <HStack p={10} spacing={10} bg={colors.background}>
                    <TicketTrigger
                        style={{flex: 1}}
                        needUpload={order?.need_upload}
                        onConfirm={(images) =>
                            onTicket({orderId: order?.id, images})
                        }
                    />

                </HStack>
            }


        </View>
    </View>
}


export default GroupDetailScreen