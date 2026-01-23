import {View} from "react-native";
import React, {useCallback, useState} from "react";
import {useInfiniteScroll, useRequest} from "ahooks";
import {fetchAccounts, fetchAccountSummary, topUp} from "../../service/api";
import {useFocusEffect} from "@react-navigation/native";
import {Avatar, Card, Text, useTheme} from "react-native-paper";
import {WhiteSpace, WingBlank} from "@ant-design/react-native";
import {Account} from "../../service/typs";
import {HStack, Stack} from "@react-native-material/core";
import DecrTrigger from "../../triggers/Account/Decrease";
import AmountView from "../../components/AmountView";
import ListView from "../../components/ListView";
import {Button} from "@rneui/themed";
import UserView from "../../components/UserView";
import UserTrigger from "../../triggers/User";

const AccountList = ({}) => {

    const {
        data: accountSummaryResult,
        run: runFetchAccountSummary
    } = useRequest(fetchAccountSummary, {manual: true})
    const summary = accountSummaryResult?.data?.data

    useFocusEffect(useCallback(() => {
        runFetchAccountSummary({})
    }, []))


    const onConfirm = (newValue: Account, updateListItem: any) => {
        updateListItem(newValue)
        runFetchAccountSummary({})
    }

    const {colors} = useTheme()

    return <View style={{flex: 1}}>
        <View style={{
            padding: 10,
            flexDirection: "row",
            alignItems: "center",
            backgroundColor: colors.background
        }}>
            <Text>总人数: <Text variant="titleMedium">{summary?.user_count}人</Text></Text>
            <WingBlank size="sm"/>
            <Text>总余额: <Text variant="titleMedium">¥ {summary?.total_balance}元</Text></Text>
        </View>
        <ListView
            fetch={page => fetchAccounts({page})}
            renderItem={(x, updateListItem) =>
                <Stack bg={colors.background} p={15} spacing={5}>
                    <UserTrigger data={x.user}/>
                    <HStack spacing={5} items={"center"}>
                        <Text>账本余额：</Text>
                        <AmountView size={"large"} amount={x.balance}/>
                    </HStack>
                    <HStack items="center" justify="end" spacing={5}>
                        {x.decrable &&
                            <DecrTrigger id={x.id} onConfirm={(newValue) => onConfirm(newValue, updateListItem)}/>}

                        {/*<TopUpTrigger id={x.id} onConfirm={onConfirm}>*/}
                        {/*    <Button mode="contained" labelStyle={{margin: 5}}>充值</Button>*/}
                        {/*</TopUpTrigger>*/}
                    </HStack>
                </Stack>
            }
        />
    </View>
}

export default AccountList