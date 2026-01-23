import {Appbar, Avatar, Button, Text, useTheme} from "react-native-paper";
import {FlatList, View} from "react-native";
import {useInfiniteScroll, useRequest} from "ahooks";
import {fetchOrderDetail, switchOrder, ticketOrder} from "../../service/api";
import React from "react";
import TicketTrigger from "../../triggers/Order/Ticket";
import {WhiteSpace, WingBlank} from "@ant-design/react-native";
import {TicketView} from "./View/TicketView";
import ImageViewer from "../../components/ImageViewer";
import TicketListView from "./View/TicketListView";
import {HStack, Stack} from "@react-native-material/core";
import SwitchTrigger from "../../triggers/Order/Switch";
import ItemView from "../../components/ItemView";
import {Chip} from "@rneui/themed";
import AmountView from "../../components/AmountView";
import StoreView from "../../components/StoreView";
import Tag from "../../components/Tag";

const OrderDetailScreen = ({navigation, route}) => {

    const {id} = route.params;

    const {data, loading: orderLoading, mutate} = useRequest(() => fetchOrderDetail({id}),)
    const order = data?.data?.data

    const loading = orderLoading && !order

    const onTicket = ({orderId, images}) => {
        ticketOrder({orderId, images}).then(rsp => {
            if (rsp.data?.code == 0) {
                mutate(rsp)
            }
        })
    }


    const onSwitch = ({orderId, toStoreId}) => {
        switchOrder({orderId, toStoreId}).then(rsp => {
            if (rsp.data?.code == 0) {
                mutate(rsp)
            }
        })
    }


    const {colors} = useTheme()

    return <View style={{flex: 1}}>
        <Appbar.Header>
            <Appbar.BackAction onPress={() => {
                navigation.goBack()
            }}/>
            <Appbar.Content title={<Text variant="titleMedium">方案详情-{id}</Text>}/>
        </Appbar.Header>
        <Stack fill={true} spacing={5}>
            <Stack p={10} bg={colors.background} spacing={10}>
                <HStack items="center" justify="between">
                    <ItemView item={order?.plan?.item} issue={order?.plan?.issue} loading={loading}/>
                    <Text>{order?.created_at}</Text>
                </HStack>
                <HStack items={"center"} spacing={5}>
                    {order?.tags?.map((t, i) => <Tag size={"sm"} key={i} title={t.title} color={t.color}/>)}
                </HStack>
                <HStack items={"center"} justify={"between"} p={8}
                        topLeftRadius={5} topRightRadius={5}>
                    <StoreView size={"sm"} data={order?.store}/>
                    {order?.to_store && <Text variant={'labelSmall'}>转给</Text>}
                    {order?.to_store && <StoreView size={"sm"} data={order?.to_store}/>}
                </HStack>
                <HStack items={"center"} spacing={8}>
                    <Text>方案总金额</Text>
                    <AmountView amount={order?.plan?.amount ?? 0} multiple={order?.plan?.multiple}/>
                </HStack>
                <Text style={{textAlign: "center", color: order?.status?.color}}
                      variant="titleMedium">{order?.status?.name}</Text>
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
            <Stack bg={colors.background} fill={1} p={10} spacing={10}>
                <Text variant="titleSmall">投注内容(拆票)</Text>
                <TicketListView itemId={order?.plan?.item?.id} data={order?.plan?.split_tickets}/>
            </Stack>
        </Stack>
        {order?.ticketable &&
            <HStack p={10} spacing={10} bg={colors.background}>

                <TicketTrigger
                    style={{flex: 1}}
                    needUpload={order?.need_upload}
                    onConfirm={(images) =>
                        onTicket({orderId: id, images})
                    }
                />

            </HStack>
        }
    </View>
}

export default OrderDetailScreen