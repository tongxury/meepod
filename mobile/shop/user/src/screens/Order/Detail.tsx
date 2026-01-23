import {View, Modal as RNModel} from "react-native";
import {Appbar, Avatar, Button, Card, Divider, Menu, Text, useTheme} from "react-native-paper";
import {useRequest} from "ahooks";
import {deleteOrder as delete_, deletePlans, fetchOrderDetail, pay} from "../../service/api";
import React, {useEffect} from "react";
import {TicketView} from "../Plan/View/TicketView";
import ImageViewer from "../../components/ImageViewer";
import TicketListView from "../Plan/View/TicketListView";
import {Stack, HStack} from "@react-native-material/core";
import CancelTrigger from "../../triggers/Order/Cancel";
import {Button as RneButton, Chip} from "@rneui/themed";
import ItemView from "../../components/ItemView";
import AmountView from "../../components/AmountView";
import {TopUpBuyingTrigger, TopUpTrigger} from "../../triggers/TopUp";
// import {Stack, HStack} from "../../components/Layout/StackLayout"

const OrderDetailScreen = ({route, navigation}) => {
    const {id} = route.params;

    const {data, loading: orderLoading, run, mutate} = useRequest(() => fetchOrderDetail({id}), {manual: true})
    const order = data?.data?.data

    const {colors} = useTheme()

    useEffect(() => {
        run()
    }, [])

    const Body = () => {

        const loading = orderLoading && !order

        if (loading && !order) {
            return <Stack fill={1}>
                <Stack fill={true} spacing={5}>
                    <Stack p={10} bg={colors.background} spacing={15}>
                        <HStack items="center" justify="between">
                            <ItemView item={order?.plan?.item} issue={order?.plan?.issue} loading={loading}/>
                            {order?.cancelable &&
                                <CancelTrigger id={order?.id} onConfirm={newValue => mutate(newValue)}/>}
                        </HStack>
                        <HStack items={"center"} spacing={5}>
                            {order?.tags?.map((t, i) => <View><Chip size={"sm"} key={i}
                                                                    color={t.color}>{t.title}</Chip></View>)}
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
                            <TicketView itemId={order?.plan?.item?.id}
                                        data={order?.plan?.issue?.result}/>
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
                        <Text variant="titleSmall">投注内容</Text>
                        <TicketListView itemId={order?.plan?.item?.id} data={order?.plan?.tickets}/>
                    </Stack>
                </Stack>
            </Stack>
        }

        return <Stack fill={1}>
            <Stack fill={true} spacing={5}>
                <Stack p={10} bg={colors.background} spacing={15}>
                    <HStack items="center" justify="between">
                        <ItemView item={order?.plan?.item} issue={order?.plan?.issue}/>
                        {order?.cancelable && <CancelTrigger id={order?.id} onConfirm={newValue => mutate(newValue)}/>}
                    </HStack>
                    <HStack items={"center"} spacing={5}>
                        {order?.tags?.map((t, i) => <View><Chip size={"sm"} key={i}
                                                                color={t.color}>{t.title}</Chip></View>)}
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
                        <TicketView itemId={order?.plan?.item?.id}
                                    data={order?.plan?.issue?.result}/>
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
                    <Text variant="titleSmall">投注内容</Text>
                    <TicketListView itemId={order?.plan?.item?.id} data={order?.plan?.tickets}/>
                </Stack>
            </Stack>
            <HStack p={10} spacing={10} justify={"around"}>
                {order?.status?.value == "submitted" &&
                    <TopUpBuyingTrigger orderId={order?.id} onConfirmed={() => run()}/>
                }
                {order?.status?.value !== "submitted" &&
                    <Button style={{flex: 1}} mode="contained" onPress={() => {
                        navigation.replace('Plan', {id: order?.plan?.item?.id, name: order?.plan?.item?.name})
                    }}>再来一单</Button>
                }
            </HStack>
        </Stack>
    }

    return <View style={{flex: 1}}>
        <Appbar.Header>
            <Appbar.BackAction onPress={() => {
                navigation.goBack()
            }}/>
            <Appbar.Content title={<Text variant="titleMedium">方案详情-{id}</Text>}/>
        </Appbar.Header>
        <Body/>
    </View>
}

export default OrderDetailScreen