import {FlatList, Pressable, View} from "react-native";
import React, {useCallback, useEffect} from "react";
import {fetchOrders, pay} from "../../service/api";
import {Avatar, Button, Card, Divider, Text, useTheme} from "react-native-paper";
import {useFocusEffect, useNavigation} from "@react-navigation/native";
import ItemView from "../../components/ItemView";
import {HStack, Stack} from "@react-native-material/core";
import ListView from "../../components/ListView";
import PlanView from "../Plan/View/PlanView";
import Tag from "../../components/Tag";
import {TopUpBuyingTrigger} from "../../triggers/TopUp";

const OrderList = ({category}) => {

    const navigation = useNavigation()

    const {colors} = useTheme()

    return <ListView
        fetch={page => fetchOrders({category, page})}
        renderItem={(x, updateListItem, reload) =>
            <Pressable onPress={() => {
                // @ts-ignore
                navigation.navigate('OrderDetail', {id: x.id})
            }}>
                <Stack bg={colors.background} p={15} spacing={8}>
                    <HStack items="center" justify="between">
                        <HStack items={"center"} spacing={5}>
                            <ItemView item={x.plan?.item} issue={x.plan?.issue}/>
                        </HStack>
                        <Text style={{color: x.status?.color}}>{x.status?.name}</Text>
                    </HStack>
                    <HStack items={"center"} spacing={5}>
                        {x.tags?.map((t, i) => <View key={i}><Tag key={i} color={t.color}>{t.title}</Tag></View>)}
                    </HStack>

                    <PlanView data={x.plan} brief={true}/>
                    <HStack items="center" justify="between" spacing={8}>
                        <Text>{x.created_at}</Text>
                        <HStack items="center">
                            {x.payable && <View>
                                <TopUpBuyingTrigger orderId={x?.id} onConfirmed={() => reload()}/>
                            </View>

                            }
                        </HStack>
                    </HStack>

                </Stack>
            </Pressable>
        }
    />
}

export default OrderList