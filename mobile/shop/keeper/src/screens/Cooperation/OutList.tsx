import {FlatList, Pressable, View} from "react-native";
import React, {useCallback, useContext, useState} from "react";
import {useInfiniteScroll, useRequest} from "ahooks";
import {Avatar, Card, useTheme} from "react-native-paper";
import {useFocusEffect, useNavigation} from "@react-navigation/native";
import {Text} from "react-native-paper";
import {HStack, Stack} from "@react-native-material/core";
import {Empty, Footer} from "../../components/ListComponent";
import {fetchCoStores, fetchProxies} from "../../service/api";
import {Button} from "@rneui/themed";
import UserView from "../../components/UserView";
import StatsView from "../../components/StatsView";
import AmountView from "../../components/AmountView";
import ConfirmTrigger from "../../triggers/Confirm";
import TipView from "./Tip";
import CreateTrigger from "./triggers/Create";
import StoreView from "../../components/StoreView";
import UpdateTrigger from "./triggers/Update";
import ItemView from "../../components/ItemView";
import IconText from "../../components/IconText";
import TopUpTrigger from "./triggers/TopUp";
import RecoverTrigger from "./triggers/Recover";
import EndApplyTrigger from "./triggers/EndApply";
import ResumeTrigger from "./triggers/Resume";
import PauseTrigger from "./triggers/Pause";
import {AppContext} from "../../providers/global";
import ListView from "../../components/ListView";

const OutCoStoreList = ({category = 'out'}) => {

    const navigation = useNavigation()

    const [filterValues, setFilterValues] = useState<any>()
    const [refresh, setRefresh] = useState<boolean>()

    const {colors} = useTheme()
    return <View style={{flex: 1}}>
        <HStack items={"center"} justify={"end"} spacing={5} p={10} bg={colors.background}>
            <CreateTrigger style={{flex: 1}} onConfirm={() => setRefresh(!refresh)}/>
            <TipView/>
        </HStack>

        <ListView
            fetch={page => fetchCoStores({...filterValues, category, page})}
            reloadDeps={[filterValues, refresh]}
            renderItem={(x, updateListItem, reload) =>
                // @ts-ignore
                <Pressable onPress={() => undefined}>
                    <Stack bg={colors.background} p={10} spacing={15}>
                        <HStack items={"center"} justify={"between"}>
                            <StoreView data={x.co_store}/>

                            <HStack spacing={5} items={"center"}>
                                <Text>{x.created_at}</Text>
                                <Text style={{color: x.status?.color}}>{x.status?.name}</Text>
                            </HStack>
                        </HStack>

                        <Stack spacing={5}>
                            <Text>我在合作方额的余额:</Text>
                            {/*<AmountView size={"large"} amount={x.balance}/>*/}
                            <HStack spacing={5} items={"end"}>
                                <AmountView size={"large"} amount={x.balance}/>
                                <Text style={{color: colors.primary}} onPress={() =>
                                    // @ts-ignore
                                    navigation.navigate('CoStorePayment', {
                                        storeId: x?.store?.id,
                                        coStoreId: x?.co_store?.id
                                    })}>资金记录</Text>
                            </HStack>
                        </Stack>

                        <Stack spacing={5}>
                            <Text variant={"titleSmall"}>合作彩种</Text>
                            <FlatList numColumns={3} data={x.items} renderItem={({item: x}) => {
                                return <HStack p={3} fill={1}>
                                    <IconText icon={x.item.icon} title={x.item.name}/>
                                </HStack>
                            }}/>
                        </Stack>


                        <HStack justify={"end"} spacing={5}>
                            {x.endapplyable && <EndApplyTrigger id={x.co_store?.id} onConfirm={reload}/>}
                            {x.recoverable && <RecoverTrigger id={x.co_store?.id} onConfirm={reload}/>}
                            {x.resumable && <ResumeTrigger coStoreId={x.co_store?.id} onConfirm={reload}/>}
                            {x.pausable && <PauseTrigger coStoreId={x.co_store?.id} onConfirm={reload}/>}
                            {x.updatable && <UpdateTrigger data={x} onConfirm={reload}/>}
                            {/*{x.topupable && <TopUpTrigger data={x} onConfirm={reload}/>}*/}
                        </HStack>
                    </Stack>
                </Pressable>

            }
        />
    </View>


}

export default OutCoStoreList