import { Pressable, View} from "react-native";
import React, {useCallback, useContext, useState} from "react";
import {useInfiniteScroll, useRequest} from "ahooks";
import {Avatar, Banner, Card, useTheme} from "react-native-paper";
import {useFocusEffect, useNavigation} from "@react-navigation/native";
import {Text} from "react-native-paper";
import {HStack, Stack} from "@react-native-material/core";
import {Empty, Footer} from "../../components/ListComponent";
import {fetchCoStorePayments, fetchStorePayments,} from "../../service/api";
import {Selector} from "../../components/Selector";
import StoreView from "../../components/StoreView";
import MoneyView from "../../components/MoneyView";
import {AppContext} from "../../providers/global";
import {toDetail} from "../../utils/nav_utils";
import ListView from "../../components/ListView";

const PaymentList = ({}) => {

    const navigation = useNavigation()

    const [month, setMonth] = useState<string>('')

    const {settingsState: {settings: {filters}},} = useContext<any>(AppContext);

    const {colors} = useTheme()

    return <View style={{flex: 1}}>
        <Stack p={10} spacing={8} bg={colors.background}>
            <Selector items={filters?.month ?? []} selectedKeys={[month]}
                      itemStyle={{selected: colors.primary, unselected: colors.primaryContainer}}
                      onSelectChanged={(item, selected) => {
                          setMonth(selected ? item.value : '')
                      }}/>

        </Stack>

        <ListView
            fetch={page =>  fetchStorePayments({month, page})}
            reloadDeps={[month]}
            renderItem={(x, updateListItem) =>
                <Pressable onPress={() => undefined}>
                    <Stack bg={colors.background} p={10} spacing={10}>
                        <HStack items={"center"} justify={"between"}>
                            {/*<StoreView data={x.co_store}/>*/}
                            <HStack items={"center"} spacing={20}>
                                <Text style={{color: x.category?.color}}>{x.category?.name}</Text>
                                <MoneyView amount={x.amount}/>
                            </HStack>

                            <HStack spacing={5} items={"center"}>
                                <Text>{x.created_at}</Text>
                                <Text style={{color: x.status?.color}}>{x.status?.name}</Text>
                            </HStack>
                        </HStack>
                        <HStack justify={'end'}>
                            <Text style={{color: colors.primary}}
                                // @ts-ignore
                                  onPress={() => toDetail(navigation, x.biz_category?.value, x.biz_id)}>订单明细</Text>
                        </HStack>
                    </Stack>
                </Pressable>
        }
        />

    </View>


}

export default PaymentList