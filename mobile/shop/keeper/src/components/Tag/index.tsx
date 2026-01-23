import {Stack} from "@react-native-material/core";
import {Text} from "react-native-paper"
import {StyleProp, ViewStyle} from "react-native";


const Tag = ({title, size, color, textColor, style}: {
    title: string,
    size?: 'sm' | 'lg',
    color: string,
    textColor?: string,
    style?: StyleProp<ViewStyle>
}) => {

    const fontSize = {
        'sm': 8,
        'large': 16,
    }

    return <Stack center={true} pv={1} ph={5} radius={3} bg={color}>
        <Text style={{color: textColor ?? 'white', fontSize: fontSize[size ?? 'sm']}} numberOfLines={1}>
            {title}
        </Text>
    </Stack>

}

export default Tag