import {FlatList, View} from "react-native";
import {useInfiniteScroll} from "ahooks";
import {fetchPayments, fetchWithdraws} from "../../service/api";
import React, {memo, useCallback, useEffect, useMemo} from "react";
import {useTheme, Text, Avatar, Button} from "react-native-paper";
import {Empty, Footer} from "../../components/ListComponent";
import {mainBodyHeight} from "../../utils/dimensions";
import {HStack, Stack} from "@react-native-material/core";
import {Chip} from "@rneui/themed";
import Tag from "../../components/Tag";
import UserView from "../../components/UserView";
import MoneyView from "../../components/MoneyView";
import {toDetail} from "../../utils/nav_utils";
import {useNavigation} from "@react-navigation/native";
import ImageViewer from "../../components/ImageViewer";

const PaymentList = ({}) => {

    const navigation = useNavigation()

    const {data: orders, loading, reload, mutate, loadMore} =
        useInfiniteScroll(d => fetchPayments({page: (d?.current ?? 0) + 1})
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
                <Stack spacing={10}>
                    <HStack items="center" justify="between">
                        <HStack items="center" spacing={8}>
                            <MoneyView amount={x.amount}/>
                            <Tag color={x.category?.color}>{x.category?.name}</Tag>
                        </HStack>
                        <HStack items={'center'} spacing={5} justify={"end"}>
                            <Text>{x.created_at}</Text>
                            <Text style={{color: x.status?.color}}>{x.status?.name}</Text>
                        </HStack>
                    </HStack>
                    {x.remark && <Text>{x.remark}</Text>}
                    {x.biz_id &&
                        <HStack items={"center"} justify={"end"}>
                            <Text onPress={() => toDetail(navigation, x.biz_category?.value, x.biz_id)}
                                  style={{color: colors.primary}}>{x.biz_id}</Text>
                        </HStack>
                    }
                </Stack>
            </View>

        }}/>

}

export default PaymentList