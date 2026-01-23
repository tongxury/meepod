import {FlatList, Pressable, View} from "react-native";
import React, {useCallback, useState} from "react";
import {fetchOrders, rejectOrder, acceptOrder, payOrder, switchOrder} from "../../../service/api";
import {Avatar, Card, useTheme} from "react-native-paper";
import {useFocusEffect, useNavigation} from "@react-navigation/native";
import FilterView from "./FilterView";
import RejectTrigger from "../../../triggers/Order/Reject";
import AcceptTrigger from "../../../triggers/Order/Accept";
import {Text} from "react-native-paper";
import PayTrigger from "../../../triggers/Order/Pay";
import {HStack, Stack} from "@react-native-material/core";
import ItemView from "../../../components/ItemView";
import {Button, Chip, Divider} from "@rneui/themed";
import SwitchTrigger from "../../../triggers/Order/Switch";
import Tag from "../../../components/Tag";
import ListView from "../../../components/ListView";
import PlanView from "../View/PlanView";
import UserTrigger from "../../../triggers/User";

const OrderList = ({}) => {

    const navigation = useNavigation()

    const [filterValues, setFilterValues] = useState<any>()

    const onReject = (orderId, reasonId, updateListItem) => {
        rejectOrder({orderId, reasonId}).then(rsp => {
            updateListItem(rsp?.data?.data)
        })
    }

    const onAccept = (orderId, updateListItem) => {
        acceptOrder({orderId}).then(rsp => {
            updateListItem(rsp?.data?.data)
        })
    }


    const onPay = (orderId, updateListItem) => {
        payOrder({orderId}).then(rsp => {
            updateListItem(rsp?.data?.data)
        })
    }


    const onSwitch = (orderId, toStoreId, updateListItem) => {
        switchOrder({orderId, toStoreId}).then(rsp => {
            updateListItem(rsp?.data?.data)

        })
    }


    const openDetail = (id) => {
        // @ts-ignore
        navigation.navigate('OrderDetail', {id})
    }

    const {colors} = useTheme()
    return <View style={{flex: 1}}>
        <FilterView onValueChange={setFilterValues}/>
        <ListView
            fetch={page => fetchOrders({...filterValues, page})}
            reloadDeps={[filterValues]}
            renderItem={(x, updateListItem) =>
                <Pressable onPress={() => openDetail(x.id)}>
                    <Stack bg={colors.background}>
                        <HStack items={"center"} justify={"between"} bg={colors.primaryContainer} p={8}
                                topLeftRadius={5} topRightRadius={5}>
                            <HStack items={"center"} spacing={5}>
                                <Avatar.Image size={20} source={{uri: x.store?.icon}}/>
                                <Text variant={"titleSmall"}>{x.store?.name}</Text>
                            </HStack>
                            {x.to_store && <Text variant={'labelSmall'}>转给</Text>}
                            {x.to_store &&
                                <HStack items={"center"} spacing={5}>
                                    <Avatar.Image size={20} source={{uri: x.to_store?.icon}}/>
                                    <Text variant={"titleSmall"}>{x.to_store?.name}</Text>
                                </HStack>
                            }
                        </HStack>

                        <Stack p={10} spacing={10}>
                            <HStack items="center" justify="between">
                                <HStack items={'center'} spacing={5}>
                                    <ItemView item={x?.plan?.item} issue={x?.plan?.issue}/>
                                </HStack>
                                <HStack items={"center"} spacing={8}>
                                    <Text style={{fontSize: 10}}>{x.created_at}</Text>
                                    <Text style={{color: x.status?.color}}>{x.status?.name}</Text>
                                </HStack>
                            </HStack>

                            <HStack items={"center"}>
                                {x?.tags?.map((t, i) => <Tag size={"sm"} key={i} title={t.title} color={t.color}/>)}
                            </HStack>
                            <PlanView data={x.plan} brief={true}/>

                            <HStack items="center" justify={"between"}>
                                <HStack items={"center"} spacing={10}>
                                    {/*<UserView data={x.user}/>*/}
                                    <UserTrigger data={x.user}/>
                                </HStack>
                                <HStack items="center" spacing={8} justify={"end"}>
                                    {x.keeper_payable &&
                                        <PayTrigger onConfirm={() => onPay(x.id, updateListItem)}>
                                            <Button>结账</Button>
                                        </PayTrigger>
                                    }
                                    {x.rejectable &&
                                        <RejectTrigger
                                            onConfirm={reasonId => onReject(x.id, reasonId, updateListItem)}/>}
                                    {x.acceptable &&
                                        <AcceptTrigger onConfirm={() => onAccept(x.id, updateListItem)}/>
                                    }
                                    {x.switchable &&
                                        <SwitchTrigger amount={x?.amount} itemId={x?.plan?.item?.id}
                                                       onConfirm={storeId =>
                                                           onSwitch(x?.id, storeId, updateListItem)}/>}
                                    {x.ticketable &&
                                        <Button onPress={() => openDetail(x.id)} color={'success'}>去出票</Button>
                                    }
                                </HStack>
                            </HStack>
                        </Stack>
                    </Stack>
                </Pressable>
            }/>
    </View>
}

export default OrderList