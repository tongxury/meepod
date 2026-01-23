import {FlatList, Text, View} from "react-native";
import {Avatar, Button, Card} from "react-native-paper";
import React from "react";


export declare type SelectorItem = {
    label: string,
    value: any
}

export const SelectorV2 = ({items, columns, selectedKeys, onSelectChanged, itemStyle}: {
    items: SelectorItem[],
    columns?: number,
    selectedKeys: any[],
    itemStyle?: {
        selected: string,
        unselected: string,
    }
    onSelectChanged: (item: SelectorItem, selected: boolean) => void,
}) => {

    return <FlatList
        style={{width: '100%'}}
        scrollEnabled={false}
        data={items}
        numColumns={columns}
        renderItem={({item}) => {
            const isSelected = (selectedKeys || []).includes(item.value)

            return <Button
                onPress={() => onSelectChanged(item, !isSelected)}
                key={item.value}
                style={{margin: 3, backgroundColor: isSelected ? itemStyle?.selected : itemStyle?.unselected, flex: 1}}
                mode="contained">
                <Text style={{color: "white"}}>{item.label}</Text>
            </Button>

        }}
        keyExtractor={item => item.value}
    />
}
