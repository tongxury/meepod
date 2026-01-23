import {Pressable, StyleProp, TextStyle, View, ViewStyle} from "react-native";
import React, {useCallback, useEffect, useState} from "react";
import {fetchOrders, fetchRewards, rejectReward, sendReward} from "../../service/api";
import {useTheme} from "react-native-paper";
import {useFocusEffect, useNavigation} from "@react-navigation/native";
import {Text} from "react-native-paper";
import {Modal, WingBlank} from "@ant-design/react-native";
import {Button, Chip} from "@rneui/themed";
import {HStack, Stack} from "@react-native-material/core";
import TitleView from "../../components/TitleView";
import AmountView from "../../components/AmountView";
import RejectTrigger from "../../triggers/reward/Reject";
import ListView from "../../components/ListView";
import UserTrigger from "../../triggers/User";
import {toDetail} from "../../utils/nav_utils";

const RewardList = () => {

    const navigation = useNavigation()

    const {colors} = useTheme()
    const onReward = (rewardId, updateListItem) => {
        Modal.alert('', '请确认已向用户发放奖金', [
            {text: '取消', onPress: undefined, style: 'cancel'},
            {
                text: '确认', onPress: () => {
                    sendReward({rewardId}).then(rsp => {
                        // updateList(mutate, rsp, rewardId)
                        updateListItem(rsp?.data?.data)
                    })
                }
            },
        ])
    }

    const onReject = (rewardId, reason, updateListItem) => {
        rejectReward({rewardId, reason}).then(rsp => {
            // updateList(mutate, rsp, rewardId)
            updateListItem(rsp?.data?.data)
        })
    }

    return <ListView
        fetch={page => fetchRewards({page})}
        renderItem={(x, updateListItem) =>
            <Pressable onPress={() => undefined}>
                <Stack p={10} bg={colors.background} spacing={10}>
                    <HStack items={"center"} justify={"between"}>
                        <UserTrigger data={x.user}/>
                        <HStack spacing={5} items={"center"}>
                            <Text>{x.created_at}</Text>
                            <Text style={{color: x.status?.color}}>{x.status?.name}</Text>
                        </HStack>
                    </HStack>

                    <HStack items={"center"} spacing={5}>
                        <Text>中奖金额</Text>
                        <AmountView amount={x.amount}/>
                    </HStack>

                    <HStack items={"center"} justify={"between"}>
                        <Text onPress={() => toDetail(navigation, x.biz_category?.value, x.biz_id)}
                              style={{color: colors.primary}}>中奖方案</Text>

                        <HStack items={"center"} spacing={5}>
                            {x.rejectable && <RejectTrigger onConfirm={id => onReject(x.id, id, updateListItem)}/>}
                            {x.rewardable && <Button onPress={() => onReward(x.id, updateListItem)}>兑奖</Button>}
                        </HStack>

                    </HStack>
                </Stack>

            </Pressable>

        }
    />
}

export default RewardList