import {Text, Image, StyleProp, ImageStyle, ViewStyle, View} from "react-native";
import {Stack} from "@react-native-material/core";
import {Button, useTheme} from "react-native-paper";

export const Empty = ({height, buttonText, onPress}: {
    height?: number,
    buttonText?: string,
    onPress?: () => void
}) => {

    const {colors} = useTheme()

    return <Stack spacing={10} center={true}
                  style={{height: height, minHeight: 300, backgroundColor: colors.background}}>
        <Image source={require('../../assets/no_data.png')} style={{width: 60, height: 60}}/>
        <Text>暂无数据</Text>
        {buttonText && <Button onPress={onPress}>{buttonText}</Button>}
    </Stack>
}

export const Footer = ({noMore, onPress, visible}: { noMore?: boolean, onPress?: () => void, visible: boolean }) => {
    const {colors} = useTheme()

    if (!visible) {
        return <View></View>
    }

    return <View
        style={{
            marginTop: 1,
            paddingTop: 10,
            flexDirection: "row",
            justifyContent: "center",
            height: 100,
            // backgroundColor: colors.background
        }}>


        {noMore ? <Text>暂无更多</Text> : <Text onPress={onPress}>加载中</Text>}

    </View>
}
