import {StyleProp, Text, View, ViewStyle} from "react-native";
import SsqMetaView from "./Ssq";

const Ticket = ({itemId, data}: { itemId: string, data: any, }) => {

    const viewTap = {
        ssq: <SsqMetaView data={data}/>
        // other 
    }

    return <View>
        {viewTap[itemId]}
    </View>
}

export default Ticket