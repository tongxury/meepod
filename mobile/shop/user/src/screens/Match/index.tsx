import {View} from "react-native";
import {mainBodyHeight} from "../../utils/dimensions";
import Tabs from "../../components/Tabs";
import React from "react";
import {useTheme} from "react-native-paper";
import {useSafeAreaInsets} from "react-native-safe-area-context";
import MatchList from "./ZjcList";
import Z14MatchList from "./Z14List";

const MatchScreen = ({navigation}) => {
    const {colors} = useTheme()
    const {top} = useSafeAreaInsets();

    return <View style={{height: mainBodyHeight, paddingTop: top}}>
        <Tabs style={{flex: 1}} tabs={[
            {key: 'z14', title: '足彩', component: () => <Z14MatchList category="z14"/>},
            {key: 'zjc', title: '竞彩足球', component: () => <MatchList category="zjc" issue={''}/>},
        ]}/>
    </View>
}

export default MatchScreen