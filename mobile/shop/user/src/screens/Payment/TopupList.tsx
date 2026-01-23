import {FlatList, View} from "react-native";
import {fetchPaymentOrders, fetchTopups} from "../../service/api";
import React, {memo, useCallback, useEffect, useMemo} from "react";
import {useTheme, Text, Avatar, Button} from "react-native-paper";
import {HStack, Stack} from "@react-native-material/core";
import {CancelTopupTrigger, TopUpTrigger} from "../../triggers/TopUp";
import {PayTrigger} from "../../triggers/Pay";
import CountDownTimer from "../../components/CountDownTimer";
import MoneyView from "../../components/MoneyView";
import Tag from "../../components/Tag";
import ListView from "../../components/ListView";

const TopupList = ({}) => {

    const {colors} = useTheme()

    return <ListView
        fetch={page => fetchTopups({page})}
        renderItem={(x, updateListItem, reload) =>
            <View style={{backgroundColor: colors.background, padding: 8}}>
                <Stack spacing={8}>
                    <HStack items="center" justify="between">
                        <HStack items="center" spacing={5}>
                            <MoneyView amount={x.amount}/>
                            <Tag color={x.category?.color}>{x.category?.name}</Tag>
                        </HStack>
                        <HStack items={'center'} justify={"end"} spacing={5}>
                            <Text>{x.created_at}</Text>
                            <Text style={{color: x.status?.color}}>{x.status?.name}</Text>
                        </HStack>
                    </HStack>
                    <HStack items="center" justify="end">

                        <HStack spacing={5}>
                            {
                                x.payable && <CountDownTimer color={colors.primary} timeLeftSecond={x.time_left}/>
                            }

                            {
                                x.cancelable && <CancelTopupTrigger
                                    id={x.id}
                                    onConfirm={(newValue) => updateListItem(newValue)}>
                                </CancelTopupTrigger>
                            }
                            {/*{*/}
                            {/*    x.payable && <TopUpTrigger topup={x}>*/}
                            {/*        <Button>支付</Button>*/}
                            {/*    </TopUpTrigger>*/}
                            {/*}*/}
                        </HStack>
                    </HStack>
                </Stack>
            </View>
        }
    />
}

export default TopupList