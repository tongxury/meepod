import {FlatList, Pressable, TouchableOpacity, TouchableWithoutFeedback, View} from "react-native";
import {Avatar, Card, Divider, Text, IconButton, useTheme, Button} from "react-native-paper";
import React, {useCallback, useEffect} from "react";
import {useInfiniteScroll, useRequest} from "ahooks";
import {addPlan, deletePlans, fetchPlans} from "../../service/api";

import {useFocusEffect, useNavigation} from "@react-navigation/native";
import {Modal, Tag, WhiteSpace, WingBlank} from "@ant-design/react-native";
import {Empty, Footer} from "../../components/ListComponent";
import TicketListView from "./View/TicketListView";
import ItemView from "../../components/ItemView";
import {HStack, Stack} from "@react-native-material/core";
import {mainBodyHeight} from "../../utils/dimensions";
import {OrderSubmitterTrigger} from "../../triggers/Order/Submit";
import AmountView from "../../components/AmountView";
import ListView from "../../components/ListView";

const PlanList = ({category}) => {

    const navigation = useNavigation()

    const onDelete = (id, reload) => {

        Modal.alert('', '确认删除当前方案吗', [
            {text: '取消', onPress: undefined, style: 'cancel'},
            {
                text: '确认', onPress: () => {
                    deletePlans({id}).then(rsp => {
                        if (rsp.data?.code == 0) {
                            reload()
                        }
                    })
                }
            },
        ])

    }


    const {colors} = useTheme()
    return <ListView
        fetch={page => fetchPlans({category, page})}
        renderItem={(x, updateListItem, reload) =>
            <Stack bg={colors.background} p={10} spacing={10}>
                <HStack items="center" spacing={5}>
                    <ItemView item={x.item} issue={x.issue}/>
                </HStack>
                <AmountView amount={x.amount} multiple={x.multiple}/>
                <TicketListView itemId={x?.item?.id} data={x?.tickets}/>
                <HStack justify="end" items="center">
                    <OrderSubmitterTrigger data={{planId: x?.id}} onSubmitted={(orderId, groupBuy) => {
                    }}>
                        <Button>提交</Button>
                    </OrderSubmitterTrigger>
                    <Text onPress={() => onDelete?.(x?.id, reload)}>删除</Text>
                </HStack>
            </Stack>
        }
    />
}

export default PlanList