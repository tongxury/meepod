import {FlatList, Pressable, View} from "react-native";
import React, {useCallback, useEffect} from "react";
import {useInfiniteScroll, useRequest} from "ahooks";
import {fetchOrders, pay} from "../../service/api";
import {Empty, Footer} from "../../components/ListComponent";
import {Avatar, Button, Card, Divider, Text, useTheme} from "react-native-paper";
import {useFocusEffect, useNavigation} from "@react-navigation/native";

import ItemView from "../../components/ItemView";
import {mainBodyHeight} from "../../utils/dimensions";
import {HStack, Stack} from "@react-native-material/core";
import {Chip} from "@rneui/themed";
import FollowTrigger from "../../triggers/Order/Follow";

const FollowableOrderList = () => {

    const navigation = useNavigation()

    const {data: orders, loading, loadMore, mutate, reload} =
        useInfiniteScroll(d => fetchOrders({category: "followable", page: (d?.current ?? 0) + 1})
            .then(rsp => new Promise(resolve => resolve(rsp?.data?.data))))

    const {colors} = useTheme()


    return <View style={{flex: 1}}>
        <FlatList
            // scrollEnabled={true}
            ListEmptyComponent={<Empty height={mainBodyHeight} buttonText="去创建" onPress={() => {
                // @ts-ignore
                navigation.navigate('Root', {screen: 'Home'})
            }}/>}
            ItemSeparatorComponent={() => <View style={{height: 1.5}}/>}
            ListFooterComponent={<Footer visible={orders?.list?.length > 0} noMore={orders?.no_more}
                                         onPress={loadMore}/>}
            onRefresh={() => reload()}
            refreshing={loading}
            onEndReached={info => loadMore()}
            data={orders?.list || []}
            numColumns={1}
            renderItem={({item: x}) =>
                <Pressable onPress={() => {
                    // @ts-ignore
                    navigation.navigate('OrderDetail', {id: x.id})
                }}>
                    <Stack bg={colors.background} p={15} spacing={10}>
                        <HStack items="center" justify="between">
                            <HStack items={"center"} spacing={5}>
                                <View><ItemView item={x.plan?.item} issue={x.plan?.issue}/></View>
                                <HStack items={"center"} spacing={5}>
                                    {x.tags?.map((t, i) => <View key={i}><Chip size={"sm"} key={i}
                                                                       color={t.color}>{t.title}</Chip></View>)}
                                </HStack>
                            </HStack>
                            <Text style={{color: x.status?.color}}>{x.status?.name}</Text>
                        </HStack>

                        <Text>编号{x.id}</Text>
                        <Text variant="titleMedium" style={{color: colors.primary}}>{x.amount}<Text>元</Text></Text>
                        <HStack items="center" justify="between" spacing={8}>
                            <Text>{x.created_at}</Text>
                            <HStack items="center">
                                {x.followable && <View>
                                    {/*@ts-ignore*/}
                                    <FollowTrigger id={x.id} onConfirm={(id: string) =>  navigation.navigate('OrderDetail', {id})}/>
                                </View>

                                }
                            </HStack>
                        </HStack>

                    </Stack>
                </Pressable>
            }
            keyExtractor={x => x.id}
        />
    </View>

}

export default FollowableOrderList