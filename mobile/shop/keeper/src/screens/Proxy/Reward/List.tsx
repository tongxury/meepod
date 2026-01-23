import {FlatList, Pressable, View} from "react-native";
import React, {useCallback, useContext, useState} from "react";
import {useInfiniteScroll, useRequest} from "ahooks";
import {Avatar, Banner, Card, useTheme} from "react-native-paper";
import {useFocusEffect, useNavigation} from "@react-navigation/native";
import {Text} from "react-native-paper";
import {HStack, Stack} from "@react-native-material/core";
import {Empty, Footer} from "../../../components/ListComponent";
import {fetchProxyRewards} from "../../../service/api";
import {Button, Chip} from "@rneui/themed";
import UserView from "../../../components/UserView";
import StatsView from "../../../components/StatsView";
import AmountView from "../../../components/AmountView";
import TipView from "./Tip";
import {Selector} from "../../../components/Selector";
import PayTrigger from "./triggers/Pay";
import {AppContext} from "../../../providers/global";
import ListView from "../../../components/ListView";

const List = ({}) => {

    const navigation = useNavigation()

    const [month, setMonth] = useState<string>('')
    const [cat, setCat] = useState<string>('')

    const {settingsState: {settings: {filters}}} = useContext<any>(AppContext);

    const {colors} = useTheme()

    return <View style={{flex: 1}}>

        <Stack p={10} spacing={8} bg={colors.background}>
            <TipView/>
            <Selector items={filters?.month ?? []} selectedKeys={[month]}
                      itemStyle={{selected: colors.primary, unselected: colors.primaryContainer}}
                      onSelectChanged={(item, selected) => {
                          setMonth(selected ? item.value : '')
                      }}/>

            <Selector items={filters?.proxyRewardCats ?? []} selectedKeys={[cat]}
                      itemStyle={{selected: colors.primary, unselected: colors.primaryContainer}}
                      onSelectChanged={(item, selected) => {
                          setCat(selected ? item.value : '')
                      }}/>
        </Stack>

        <ListView
            fetch={page => fetchProxyRewards({month, cat, page})}
            reloadDeps={[month, cat]}
            renderItem={(x, updateListItem, reload) =>
                <Pressable onPress={() => undefined}>
                    <Stack bg={colors.background} p={10} spacing={10}>
                        <HStack items={"center"} justify={"between"}>
                            <UserView data={x.user}/>

                            <HStack spacing={5} items={"center"}>
                                <Text>{x.month}</Text>
                                <Chip color={x.status?.color}>{x.status?.name}</Chip>
                            </HStack>
                        </HStack>
                        <HStack items={"center"} spacing={20}>
                            <AmountView amount={x.reward_amount} size={"large"}/>
                            <Text><Text style={{
                                fontSize: 20,
                                fontWeight: 'bold',
                                color: colors.primary
                            }}>{x.reward_rate * 100}</Text> %</Text>
                        </HStack>
                        <HStack items={"center"} justify={"between"}>
                            <StatsView title="推广用户" value={x.user_count} unit={"人"}/>
                            <StatsView title="订单数" value={x.order_count} unit={"单"}/>
                            <StatsView title="订单总额" value={x.order_amount} unit={"元"}/>
                            {/*<StatsView title="佣金比例" value={3000} unit={"%"}/>*/}
                        </HStack>
                        <HStack justify={"end"} spacing={5}>
                            {x.payable && <PayTrigger id={x.id} onConfirm={reload}/>}
                        </HStack>
                    </Stack>
                </Pressable>
            }
        />
    </View>


}

export default List