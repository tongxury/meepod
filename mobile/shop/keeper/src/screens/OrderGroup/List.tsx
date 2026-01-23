import {Pressable, View} from "react-native";
import {useInfiniteScroll, useRequest} from "ahooks";
import {
    acceptOrder,
    acceptOrderGroup,
    fetchOrderGroups,
    fetchOrders,
    rejectOrder,
    rejectOrderGroup, switchOrderGroup
} from "../../service/api";
import React, {useCallback, useEffect} from "react";
import {Empty, Footer} from "../../components/ListComponent";
import GroupSummary from "./View/Summary";
import {useFocusEffect, useNavigation} from "@react-navigation/native";
import {mainBodyHeight} from "../../utils/dimensions";
import {HStack, Stack} from "@react-native-material/core";
import {Appbar, Text, Button, useTheme, Avatar} from "react-native-paper";
import RejectTrigger from "../../triggers/Order/Reject";
import AcceptTrigger from "../../triggers/Order/Accept";
import {updateList} from "../../service/utils";
import {Button as RneButton} from "@rneui/themed";
import {color} from "@rneui/base";
import SwitchTrigger from "../../triggers/Order/Switch";
import ListView from "../../components/ListView";


const OrderGroupList = () => {

    const navigation = useNavigation()

    const onReject = (orderId, reasonId, updateListItem) => {
        rejectOrderGroup({orderId, reasonId}).then(rsp => {
            updateListItem(rsp?.data?.data)
        })
    }

    const onAccept = (orderId, updateListItem) => {
        acceptOrderGroup({orderId}).then(rsp => {
            updateListItem(rsp?.data?.data)
        })
    }
    const onSwitch = (orderId, toStoreId, updateListItem) => {
        switchOrderGroup({orderId, toStoreId}).then(rsp => {
            updateListItem(rsp?.data?.data)
        })
    }

    const openDetail = (id) => {
        // @ts-ignore
        navigation.navigate('OrderGroupDetail', {id})
    }

    const {colors} = useTheme()

    return <ListView
        fetch={page => fetchOrderGroups({page})}
        renderItem={(x, updateListItem) =>
            <Pressable onPress={() => openDetail(x.id)}>
                <Stack bg={colors.background} spacing={10} p={10}>
                    <GroupSummary data={x}/>
                    <HStack items={"center"} spacing={5} justify={"end"} ph={10}>
                        {x.rejectable &&
                            <RejectTrigger onConfirm={reasonId => onReject(x.id, reasonId, updateListItem)}/>
                        }
                        {x.acceptable &&
                            <AcceptTrigger onConfirm={() => onAccept(x.id, updateListItem)}/>
                        }
                        {x.switchable &&
                            <SwitchTrigger amount={x?.amount} itemId={x?.plan?.item?.id} onConfirm={storeId =>
                                onSwitch(x?.id, storeId, updateListItem)}/>}
                        {x.ticketable &&
                            <Button onPress={() => openDetail(x.id)} mode="contained"
                                    labelStyle={{margin: 5}}>去出票</Button>
                        }
                    </HStack>

                </Stack>
            </Pressable>
        }
    />

}

export default OrderGroupList