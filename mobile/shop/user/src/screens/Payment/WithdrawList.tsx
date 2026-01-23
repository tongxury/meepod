import {FlatList, View} from "react-native";
import {useInfiniteScroll} from "ahooks";
import {fetchPaymentOrders, fetchTopups, fetchWithdraws} from "../../service/api";
import React, {memo, useCallback, useEffect, useMemo} from "react";
import {useTheme, Text, Avatar, Button} from "react-native-paper";
import {Empty, Footer} from "../../components/ListComponent";
import {mainBodyHeight} from "../../utils/dimensions";
import {HStack, Stack} from "@react-native-material/core";
import Tag from "../../components/Tag";
import CancelWithdrawTrigger from "../../triggers/Withdraw/Cancel";
import moment from "moment";
import {updateListItemById} from "../../service/utils";
import MoneyView from "../../components/MoneyView";
import ImageViewer from "../../components/ImageViewer";
import {Button as RneButton} from "@rneui/themed"

const WithdrawList = ({}) => {

    const {data: orders, loading, reload, mutate, loadMore} =
        useInfiniteScroll(d => fetchWithdraws({page: (d?.current ?? 0) + 1})
            .then(rsp => new Promise(resolve => resolve(rsp?.data?.data))),)

    const {colors} = useTheme()

    return <FlatList
        ListEmptyComponent={<Empty height={mainBodyHeight}/>}
        ItemSeparatorComponent={() => <View style={{height: 1.5}}/>}
        ListFooterComponent={<Footer visible={orders?.list?.length > 0} noMore={orders?.no_more}
                                     onPress={loadMore}/>}
        onEndReached={info => loadMore()}
        onRefresh={() => reload()}
        refreshing={loading}
        data={orders?.list ?? []}
        renderItem={({item: x, index: i}) => {
            return <View style={{backgroundColor: colors.background, padding: 15}}>
                <Stack spacing={5}>
                    <HStack items="center" justify="between">
                        <MoneyView amount={x.amount}/>
                        <HStack items={"center"} spacing={5}>
                            <Text>{x.created_at}</Text>
                            <Text style={{color: x.status?.color}}>{x.status?.name}</Text>
                        </HStack>
                    </HStack>
                    {x.remark && <Text>{x.remark}</Text>}
                    {x.image && <ImageViewer size={"small"} images={[{url: x.image}]}/>}
                    <HStack spacing={5} justify={"end"}>
                        {
                            x.cancelable && <CancelWithdrawTrigger
                                id={x.id}
                                onConfirm={(newValue) => updateListItemById(mutate, newValue)}>
                                {/*<RneButton type={'clear'}>撤销</RneButton>*/}
                            </CancelWithdrawTrigger>
                        }
                    </HStack>
                </Stack>
            </View>

        }}/>

}

export default WithdrawList