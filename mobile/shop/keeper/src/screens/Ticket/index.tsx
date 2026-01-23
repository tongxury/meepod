import {SafeAreaView, View} from "react-native";
import Tabs from "../../components/Tabs";
import {useSafeAreaInsets} from "react-native-safe-area-context";
import {useTheme} from "react-native-paper";
import OrderList from "../Order/List";
import OrderGroupList from "../OrderGroup/List";

const TicketScreen = () => {
    const insets = useSafeAreaInsets();

    const {colors} = useTheme()

    return <SafeAreaView style={{flex: 1, paddingTop: insets.top}}>
        <Tabs style={{flex: 1}} tabs={[
            {key: 'plan', title: '方案', component: () => <OrderList/>},
            {key: 'orderGroup', title: '合买', component: () => <OrderGroupList />},
            // {key: 'reward', title: '兑奖', component: () => <RewardList/>}
        ]}/>
    </SafeAreaView>
}

export default TicketScreen