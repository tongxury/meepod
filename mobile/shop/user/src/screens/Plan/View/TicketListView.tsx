import {FlatList, ScrollView, StyleProp, Text, View, ViewStyle} from "react-native";
import {TicketView} from "./TicketView";
import React from "react";
import {Empty} from "../../../components/ListComponent";

const TicketListView = ({itemId, data, style}: { itemId: string, data: any[], style?: StyleProp<ViewStyle> }) => {

    return <ScrollView style={style}>
        <FlatList
            data={data ?? []}
            numColumns={1}
            scrollEnabled={true}
            ListEmptyComponent={<Empty/>}
            ItemSeparatorComponent={() => <View style={{height: 10}}/>}
            renderItem={({item: x}) => {
                return <TicketView itemId={itemId} data={x}/>
            }}
            keyExtractor={(x, i) => i.toString()}
        />
    </ScrollView>

}

export default TicketListView