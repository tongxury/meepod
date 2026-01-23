import {Text, ProgressBar, useTheme} from "react-native-paper";
import {StyleProp, View, ViewStyle} from "react-native";
import {HStack} from "@react-native-material/core";

const AmountView = ({amount, multiple, size, style}: {
    amount: number,
    multiple?: number,
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


    const Amount = () => {
        if (multiple) {
            return <View style={style}><Text style={{
                fontSize: sty.amountFontSize,
                color: colors.primary,
                fontWeight: "bold"
            }}>¥ {amount * multiple ?? 0}元 <Text
                style={{fontSize: sty.multipleFontSize}}>[{multiple ?? 1}倍]</Text></Text></View>
        } else {
            return <View style={style}><Text style={{
                color: colors.primary,
                fontSize: sty.amountFontSize,
                fontWeight: "bold"
            }}>¥ {amount ?? 0} <Text
                style={{fontSize: sty.multipleFontSize}}>元</Text></Text></View>
        }
    }

    return <HStack>
        <Amount/>
    </HStack>

}

export default AmountView