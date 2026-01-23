import {HStack, Stack} from "@react-native-material/core";
import {Avatar, Text, useTheme} from "react-native-paper";
import {Store, User} from "../../service/typs";
import {StyleProp, ViewStyle} from "react-native";
import {Chip} from "@rneui/themed";
import Tag from "../Tag";

const StoreView = ({data, size, style}: { data: Store, size?: 'sm'| 'lg', style?: StyleProp<ViewStyle> }) => {
    const {colors} = useTheme()


    const sty = {
        'sm': {
            iconSize: 20,
            titleSize: 14,
        },
        'lg': {
            iconSize: 30,
            titleSize: 16,
        }
    }[size ?? 'sm']


    return <HStack style={style} items={"center"} spacing={8}>
        <Avatar.Image size={sty.iconSize} source={{uri: data?.icon}} style={{backgroundColor: colors.primaryContainer}}/>
        <Stack justify={'between'}>
            <HStack items={"center"} spacing={5}>
                <Text variant={"titleSmall"} style={{fontSize: sty.titleSize}}>{data?.name}</Text>
                {data?.tags?.map((t, i) => <Tag key={i} title={t.title} color={t.color}/>)}
            </HStack>
        </Stack>
    </HStack>
}

export default StoreView