import SsqTicketView from "./Ssq";
import DltTicketView from "./Dlt";
import F3dTicketView from "./F3d";
import X7cTicketView from "./X7c";
import Z14TicketView from "./Z14";
import ZjcTicketView from "./Zjc";
import Pl5TicketView from "./Pl5";

export const TicketView = ({itemId, data}: { itemId: string, data: any, }) => {

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

    return viewTap[itemId]

    // return <View style={{flex: 1}}>
    //     {viewTap[itemId]}
    // </View>
}
