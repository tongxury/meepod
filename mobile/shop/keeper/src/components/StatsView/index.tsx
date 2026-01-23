import {StyleProp, ViewStyle} from "react-native";
import {HStack, Stack} from "@react-native-material/core";
import {Text, useTheme} from "react-native-paper";

const StatsView = ({title, value, unit, style}: {
    title: string,
    value: number,
    unit?: string,
    style?: StyleProp<ViewStyle>
}) => {

    const {colors} = useTheme()

    return <Stack style={style} spacing={5} center={true}>
        <Text variant={"labelMedium"} style={{fontSize: 10}}>
            {title}
        </Text>

        <HStack items={"center"} spacing={5}>
            <Text variant={"titleLarge"} allowFontScaling={false}
                  style={{fontSize: 20, color: colors.primary, fontWeight: "bold"}}>
                {value ?? 0}
            </Text>

            {unit && <Text variant={"labelSmall"} style={{fontSize: 10}}>
                {unit}
            </Text>}

        </HStack>


    </Stack>
};

export default StatsView
