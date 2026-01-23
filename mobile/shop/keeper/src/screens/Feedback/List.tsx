import {FlatList, Pressable, View} from "react-native";
import React, {useCallback, useEffect} from "react";
import {useInfiniteScroll, useRequest} from "ahooks";
import {fetchFeedbacks, fetchOrders, pay} from "../../service/api";
import {Empty, Footer} from "../../components/ListComponent";
import {Text, useTheme} from "react-native-paper";
import {useFocusEffect, useNavigation} from "@react-navigation/native";

import {HStack, Stack} from "@react-native-material/core";
import {Button} from "@rneui/themed";
import FeedbackTrigger from "./triggers/Add";
import UserView from "../../components/UserView";

const FeedBackList = ({}) => {

    const navigation = useNavigation()

    const {data: orders, loading, loadMore, mutate, reload} =
        useInfiniteScroll(d => fetchFeedbacks({page: (d?.current ?? 0) + 1})
            .then(rsp => new Promise(resolve => resolve(rsp?.data?.data))))

    const {colors} = useTheme()

    return <View style={{flex: 1}}>
        <FlatList
            ListEmptyComponent={<Empty/>}
            ItemSeparatorComponent={() => <View style={{height: 1.5}}/>}
            ListFooterComponent={<Footer visible={orders?.list?.length > 0} noMore={orders?.no_more}
                                         onPress={loadMore}/>}
            onRefresh={() => reload()}
            refreshing={loading}
            onEndReached={info => loadMore()}
            data={orders?.list || []}
            numColumns={1}
            renderItem={({item: x}) =>
                <Stack bg={colors.background} p={15} spacing={10}>
                    <HStack items="center" justify="between">
                        <UserView data={x.user}/>

                        <HStack items="center" spacing={5}>
                            <Text>{x.created_at}</Text>
                            <Text style={{color: x.status?.color}}>{x.status?.name}</Text>
                        </HStack>
                    </HStack>

                    <Text variant={"bodyMedium"}>{x.text}</Text>
                </Stack>
            }
            keyExtractor={x => x.id}
        />
    </View>

}

export default FeedBackList