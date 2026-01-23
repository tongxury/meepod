import {FlatList, Pressable, View} from "react-native";
import {useInfiniteScroll, useRequest} from "ahooks";
import React, {useCallback, useEffect} from "react";
import {useFocusEffect, useNavigation} from "@react-navigation/native";
import {fetchOrderGroupOrders} from "../../../service/api";
import {Avatar, Text, useTheme} from "react-native-paper";
import {Chip} from "@rneui/themed";
import moment from "moment";
import {Empty, Footer} from "../../../components/ListComponent";
import {mainBodyHeight} from "../../../utils/dimensions";
import {HStack, Stack} from "@react-native-material/core";
import Tag from "../../../components/Tag";

const ShareList = ({groupId}) => {

    // const navigation = useNavigation()

    const {data: orders, loading, loadMore, mutate, reload} =
        useInfiniteScroll(d => fetchOrderGroupOrders({groupId, page: (d?.current ?? 0) + 1})
            .then(rsp => new Promise(resolve => resolve(rsp?.data?.data))))

    const {colors} = useTheme()

    return <FlatList
        ListEmptyComponent={<Empty />}
        ItemSeparatorComponent={() => <View style={{height: 1.5}}/>}
        ListFooterComponent={<Footer visible={orders?.list?.length > 0} noMore={orders?.no_more}
                                     onPress={loadMore}/>}
        onRefresh={() => reload()}
        refreshing={loading}
        onEndReached={info => loadMore()}
        data={orders?.list || []}
        numColumns={1}
        renderItem={({item: x}) =>
            <HStack items={"center"} justify={"between"} p={10} bg={colors.background}>
                <HStack items={"center"} spacing={8}>
                    <Avatar.Image style={{backgroundColor: colors.background}} size={40}
                                  source={{uri: x.user?.icon}}/>
                    <Stack spacing={10}>
                        <HStack items={"center"} spacing={5}>
                            <Text variant={"titleMedium"} style={{fontWeight: "bold",}}>{x.user?.phone}</Text>
                            <HStack items={"center"}>
                                {x.user?.tags?.map((t, i) => <Tag color={t.color} key={i}>{t.title}</Tag>)}
                            </HStack>
                        </HStack>
                        <Text variant={"labelSmall"}>{x.created_at}</Text>
                    </Stack>
                </HStack>
                <Text>{x.volume}份</Text>
            </HStack>
        }
        keyExtractor={
            x => x?.user?.id
        }
    />
}

export default ShareList