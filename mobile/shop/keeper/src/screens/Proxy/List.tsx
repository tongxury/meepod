import {FlatList, Pressable, View} from "react-native";
import React, {useCallback, useContext, useRef, useState} from "react";
import {Avatar, Card, useTheme} from "react-native-paper";
import {useFocusEffect, useNavigation} from "@react-navigation/native";
import {Text} from "react-native-paper";
import {HStack, Stack} from "@react-native-material/core";
import {fetchProxies} from "../../service/api";
import CreateTrigger from "./triggers/Create";
import StatsView from "../../components/StatsView";
import AmountView from "../../components/AmountView";
import DeleteTrigger from "./triggers/Delete";
import RecoverTrigger from "./triggers/Recover";
import UpdateTrigger from "./triggers/Update";
import AddUserTrigger from "./triggers/AddUser";
import ListView from "../../components/ListView";
import UserTrigger from "../../triggers/User";

const ProxyList = ({}) => {

    const navigation = useNavigation()

    const [filterValues, setFilterValues] = useState<any>()
    const [refresh, setRefresh] = useState<boolean>(false)

    const {colors} = useTheme()
    return <View style={{flex: 1}}>
        <HStack items={"center"} justify={"end"} spacing={5} p={10} bg={colors.background}>
            <CreateTrigger style={{flex: 1}} onConfirm={() => setRefresh(!refresh)}/>
        </HStack>

        <ListView
            fetch={page => fetchProxies({...filterValues, page})}
            reloadDeps={[filterValues, refresh]}
            renderItem={(x, updateListItem, reload) =>
                <Pressable onPress={() => undefined}>
                    <Stack bg={colors.background} p={10} spacing={15}>
                        <HStack items={"center"} justify={"between"}>
                            <UserTrigger data={x.user}/>

                            <HStack spacing={5} items={"center"}>
                                <Text>{x.created_at}</Text>
                                <Text style={{color: x.status?.color}}>{x.status?.name}</Text>
                            </HStack>
                        </HStack>

                        <HStack items={"center"} spacing={20}>
                            <HStack items={"center"} spacing={5}>
                                <Text>邀请码</Text>
                                <Text variant={"titleMedium"} style={{fontWeight: 'bold'}}>{x.id}</Text>
                            </HStack>
                            <AmountView amount={x.reward_amount} size={"large"}/>
                            <Text><Text style={{
                                fontSize: 20,
                                fontWeight: 'bold',
                                color: colors.primary
                            }}>{(x.reward_rate * 100).toFixed(1)}</Text> %</Text>
                        </HStack>
                        <HStack items={"center"} justify={"between"}>
                            <StatsView style={{flex: 1}} title="用户数(下过单)" value={x.user_count} unit={"人"}/>
                            <StatsView style={{flex: 1}} title="订单数" value={x.order_count} unit={"单"}/>
                            <StatsView style={{flex: 1}} title="订单总额" value={x.order_amount} unit={"元"}/>
                            {/*<StatsView title="佣金比例" value={3000} unit={"%"}/>*/}
                        </HStack>
                        <HStack justify={"end"} spacing={5}>
                            {x.deletable && <DeleteTrigger id={x.id} onConfirm={reload}/>}
                            {x.recoverable && <RecoverTrigger id={x.id} onConfirm={reload}/>}
                            {x.updatable && <UpdateTrigger data={x} onConfirm={reload}/>}
                            {x.addable && <AddUserTrigger proxy={x}/>}
                        </HStack>
                    </Stack>
                </Pressable>
            }
        />
    </View>
}

export default ProxyList