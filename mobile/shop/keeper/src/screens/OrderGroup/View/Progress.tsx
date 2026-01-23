import {StyleProp, Text, View, ViewStyle} from "react-native";
import {ProgressBar} from "react-native-paper";
import {WhiteSpace, WingBlank} from "@ant-design/react-native";
import {OrderGroup} from "../../../service/typs";

const Progress = ({data, style}: {
    data: OrderGroup,
    style?: StyleProp<ViewStyle> | undefined
}) => {

    return <View style={style}>
        <ProgressBar progress={data?.volume_ordered / data?.volume}/>
        <WhiteSpace/>
        <View style={{flexDirection: "row", alignItems: "center", justifyContent: "space-between"}}>
            <Text>合买进度<Text>{(data?.volume_ordered / data?.volume * 100).toFixed(0)}%</Text></Text>

            <Text>已参与<Text>{data?.joiner_count}人/{data?.volume_ordered}份</Text><WingBlank size="sm"/>
                余<Text>{data?.volume - data?.volume_ordered}份</Text></Text>
        </View>
    </View>
}

export default Progress