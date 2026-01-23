import {FlatList, Pressable, StyleProp, TextStyle, View, ViewStyle} from "react-native";
import React, {DependencyList, useCallback, useEffect, useState} from "react";
import {useInfiniteScroll, useRequest} from "ahooks";
import {useFocusEffect, useNavigation} from "@react-navigation/native";
import {PageData, Result} from "../../service/typs";
import {AxiosResponse} from "axios";
import {Empty, Footer} from "../ListComponent";

const ListView = ({fetch, reloadDeps, renderItem, style}: {
    fetch: (page: number) => Promise<AxiosResponse<Result<PageData<any>>>>
    reloadDeps?: DependencyList;
    renderItem: (item: any, updateListItem: (newItem: any) => void, reload: () => void) => React.ReactElement
    style?: StyleProp<ViewStyle> | undefined;
}) => {

    const {data, loading, mutate, reload, loadMore} =
        useInfiniteScroll(d => fetch((d?.current ?? 0) + 1)
                .then(rsp => new Promise(resolve => resolve(rsp?.data?.data))),
            {isNoMore: d => d?.no_more, manual: true, reloadDeps})


    useFocusEffect(useCallback(() => {
        console.log('ListView useFocusEffect')
        reload()
    }, []))

    const updateListItem = (newItem) => {
        console.log('updateListItem updateListItem', newItem)
        mutate(oldData => {
            const newList = oldData.list.map(u => u.id !== newItem?.id ? u : newItem);
            const newData = {...oldData}
            newData.list = newList
            return newData
        })
    }

    return <FlatList
        refreshing={loading}
        onRefresh={() => reload()}
        style={style}
        ListEmptyComponent={<Empty/>}
        ListFooterComponent={<Footer visible={data?.list?.length > 0} noMore={data?.no_more} onPress={loadMore}/>}
        onEndReached={info => loadMore()}
        data={data?.list || []}
        numColumns={1}
        ItemSeparatorComponent={() => <View style={{height: 1.5}}/>}
        renderItem={({item: x, index}) => renderItem(x, updateListItem, reload)}
        keyExtractor={(x, i) => i.toString()}
    />
}

export default ListView