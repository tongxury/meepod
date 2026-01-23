import {FlatList, ScrollView, StyleProp, Text, View, ViewStyle} from "react-native";
import React from "react";
import {Plan} from "../../../service/typs";
import TicketListView from "./TicketListView";
import {Stack} from "@react-native-material/core";
import AmountView from "../../../components/AmountView";

const PlanView = ({data, brief, style}: { data: Plan, brief?: boolean, style?: StyleProp<ViewStyle> }) => {

    if (data?.tickets?.length > 1 && brief) {
        return <Stack style={style} spacing={5}>
            <AmountView amount={data.amount} multiple={data?.multiple}/>
            <TicketListView itemId={data?.item?.id} data={data?.tickets?.slice(0, 1)}/>
            <Text style={{fontSize: 10}}>... 更多</Text>
        </Stack>
    } else {
        return <Stack style={style} spacing={5}>
            <AmountView amount={data.amount} multiple={data?.multiple}/>
            <TicketListView itemId={data?.item?.id} data={data?.tickets}/>
        </Stack>
    }


}

export default PlanView