import {HStack, Stack} from "@react-native-material/core";
import {Avatar, Text, useTheme} from "react-native-paper";
import {User} from "../../service/typs";
import {StyleProp, ViewStyle} from "react-native";
import {Chip} from "@rneui/themed";
import Tag from "../Tag";

const IconText = ({icon, title, style}: { icon: string, title: string, style?: StyleProp<ViewStyle> }) => {

    const {colors} = useTheme()

    return <HStack style={style} items={"center"} spacing={8}>
        <Avatar.Image size={30} source={{uri: icon}} style={{backgroundColor: colors.primaryContainer}}/>
        <Stack justify={'between'}>
            <Text variant={"titleSmall"}>{title}</Text>
        </Stack>
    </HStack>
}

export default IconText