import {FlatList, ScrollView, Text, View} from "react-native";
import {Empty} from "../../../components/ListComponent";
import {TicketView} from "./TicketView";
import React from "react";
import {WingBlank} from "@ant-design/react-native";
import {Plan} from "../../../service/typs";

const TicketListView = ({itemId, data}: { itemId: string, data: any[] }) => {

    return <ScrollView>
        <FlatList
            data={data ?? []}
            numColumns={1}
            scrollEnabled={true}
            ListEmptyComponent={<Empty/>}
            ItemSeparatorComponent={() => <View style={{height: 20}}/>}
            renderItem={({item: x}) => {
                return <TicketView itemId={itemId} data={x}/>
            }}
            keyExtractor={(x, i) => i.toString()}
        />
    </ScrollView>

}

export default TicketListView