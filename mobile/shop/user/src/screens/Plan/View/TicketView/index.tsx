import {StyleProp, View, ViewStyle} from "react-native";
import SsqTicketView from "./Ssq";
import DltTicketView from "./Dlt";
import F3dTicketView from "./F3d";
import X7cTicketView from "./X7c";
import Z14TicketView from "./Z14";
import ZjcTicketView from "./Zjc";
import Pl5TicketView from "./Pl5";

export const TicketView = ({itemId, data, style}: { itemId: string, data: any, style?: StyleProp<ViewStyle> }) => {

    const viewTap = {
        ssq: <SsqTicketView data={data}/>,
        dlt: <DltTicketView data={data}/>,
        f3d: <F3dTicketView data={data}/>,
        x7c: <X7cTicketView data={data}/>,
        rx9: <Z14TicketView data={data}/>,
        sfc: <Z14TicketView data={data}/>,
        zjc: <ZjcTicketView data={data}/>,
        pl3: <F3dTicketView data={data}/>,
        pl5: <Pl5TicketView data={data}/>,
    }

    return <View style={style}>
        {viewTap[itemId]}
    </View>
}
