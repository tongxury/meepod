import {FlatList, Pressable, View} from "react-native";
import React, {useCallback, useContext, useState} from "react";
import {useInfiniteScroll, useRequest} from "ahooks";
import {Avatar, Banner, Card, useTheme} from "react-native-paper";
import {useFocusEffect, useNavigation} from "@react-navigation/native";
import {Text} from "react-native-paper";
import {HStack, Stack} from "@react-native-material/core";
import {Empty, Footer} from "../../../components/ListComponent";
import {fetchCoStorePayments,} from "../../../service/api";
import {Selector} from "../../../components/Selector";
import StoreView from "../../../components/StoreView";
import MoneyView from "../../../components/MoneyView";
import {AppContext} from "../../../providers/global";
import {toDetail} from "../../../utils/nav_utils";
import ListView from "../../../components/ListView";

const PaymentList = ({storeId, coStoreId}) => {

    const navigation = useNavigation()

    const [month, setMonth] = useState<string>('')
    const [cat, setCat] = useState<string>('')

    const {settingsState: {settings: {filters}}} = useContext<any>(AppContext);

    const {colors} = useTheme()

    return <View style={{flex: 1}}>

        <Stack p={10} spacing={8} bg={colors.background}>
            <Selector items={filters?.month ?? []} selectedKeys={[month]}
                      itemStyle={{selected: colors.primary, unselected: colors.primaryContainer}}
                      onSelectChanged={(item, selected) => {
                          setMonth(selected ? item.value : '')
                      }}/>

            <Selector items={filters?.coStorePaymentCats ?? []} selectedKeys={[cat]}
                      itemStyle={{selected: colors.primary, unselected: colors.primaryContainer}}
                      onSelectChanged={(item, selected) => {
                          setCat(selected ? item.value : '')
                      }}/>

        </Stack>

        <ListView
            fetch={page => fetchCoStorePayments({storeId, coStoreId, month, cat, page})}
            reloadDeps={[month, cat]}
            renderItem={(x, updateListItem, reload) =>
                <Pressable onPress={() => undefined}>
                    <Stack bg={colors.background} p={10} spacing={10}>
                        <HStack items={"center"} justify={"between"}>
                            <HStack items={"center"} spacing={5}>
                                <Text>转出方:</Text>
                                <StoreView data={x.store}/>
                            </HStack>
                            <HStack items={"center"} spacing={5}>
                                <Text>转入方:</Text>
                                <StoreView data={x.co_store}/>
                            </HStack>
                        </HStack>

                        <HStack items={"center"} justify={"between"}>
                            <HStack items={"center"} spacing={20}>
                                <Text style={{color: x.category?.color}}>{x.category?.name}</Text>
                                <MoneyView amount={x.amount}/>
                            </HStack>
                            <HStack spacing={5} items={"center"}>
                                <Text>{x.created_at}</Text>
                                <Text style={{color: x.status?.color}}>{x.status?.name}</Text>
                            </HStack>
                        </HStack>
                        {x.biz_id && <Text
                            style={{color: colors.primary}}
                            // @ts-ignore
                            onPress={() => toDetail(navigation, x.biz_category?.value, x.biz_id)}>订单明细</Text>
                        }
                    </Stack>
                </Pressable>

            }
        />
    </View>


}

export default PaymentList