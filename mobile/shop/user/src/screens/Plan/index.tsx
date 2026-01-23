import {ScrollView, View} from "react-native";
import {Appbar, useTheme, Text} from "react-native-paper";
import PlanEditor from "./Editor";
import {WhiteSpace} from "@ant-design/react-native";

const PlanScreen = ({route, navigation}) => {
    const {id, name} = route.params;

    const {colors} = useTheme()
    return <View style={{flex: 1}}>
        <Appbar.Header>
            <Appbar.BackAction onPress={() => {
                // navigation.navigate("Root", {screen: "Home"})
                navigation.goBack()
            }}/>
            <Appbar.Content title={<Text variant="titleMedium">{name}</Text>}/>
            {/*<Appbar.Action icon="dots-vertical" onPress={() => {*/}
            {/*}}/>*/}
        </Appbar.Header>
        <WhiteSpace size="sm"/>
        <PlanEditor itemId={id}/>
    </View>
}

export default PlanScreen