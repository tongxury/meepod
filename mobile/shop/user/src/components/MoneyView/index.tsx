import {Text, ProgressBar, useTheme} from "react-native-paper";
import {StyleProp, View, ViewStyle} from "react-native";
import {HStack} from "@react-native-material/core";

const MoneyView = ({amount, size, style}: {
    amount: number,
    size?: 'sm' | 'large'
    style?: StyleProp<ViewStyle> | undefined
}) => {

    const {colors} = useTheme()

    const sty = {
        'sm': {
            amountFontSize: 14,
            multipleFontSize: 8
        },
        'large': {
            amountFontSize: 20,
            multipleFontSize: 13
        }
    }[size ?? 'sm']

    return <View style={style}><Text style={{
        color: amount >= 0 ? 'rgb(11,178,19)' : 'red',
        fontSize: sty.amountFontSize,
        fontWeight: "bold"
    }}>¥ {amount ?? 0} <Text
        style={{fontSize: sty.multipleFontSize}}>元</Text></Text></View>

}

export default MoneyView