import {FlatList, Pressable, View} from "react-native";
import {useInfiniteScroll, useRequest} from "ahooks";
import {fetchOrderGroups, fetchOrders} from "../../service/api";
import React, {useCallback, useEffect} from "react";
import {Empty, Footer} from "../../components/ListComponent";
import GroupSummary from "./View/Summary";
import {useFocusEffect, useNavigation} from "@react-navigation/native";
import {mainBodyHeight} from "../../utils/dimensions";
import ListView from "../../components/ListView";

const OrderGroupList = ({category}) => {

    const navigation = useNavigation()

    return <ListView
        fetch={page => fetchOrderGroups({category, page})}
        renderItem={(x, updateListItem, reload) =>
            <Pressable onPress={() => {
                // @ts-ignore
                navigation.navigate('OrderGroupDetail', {id: x.id})
            }}>
                <GroupSummary data={x}/>
            </Pressable>
        }/>
}

export default OrderGroupList