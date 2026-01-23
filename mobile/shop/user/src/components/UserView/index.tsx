import {HStack, Stack} from "@react-native-material/core";
import {Avatar, Text} from "react-native-paper";
import {Store, User} from "../../service/typs";
import {StyleProp, ViewStyle} from "react-native";
import {Chip} from "@rneui/themed";
import Tag from "../Tag";

const UserView = ({data, size, style}: { data: User, size?: 'sm' | 'md', style?: StyleProp<ViewStyle> }) => {

    const sty = {
        'sm': {
            avatarSize: 20,
        },
        'md': {
            avatarSize: 30,
        }
    }[size ?? 'sm']

    return <HStack style={style} items={"center"} spacing={8}>
        <Avatar.Image size={sty.avatarSize} source={{uri: data?.icon}}/>
        <Stack justify={'between'}>
            <HStack items={"center"} spacing={5}>
                <Text variant={"titleSmall"}>{data?.phone}</Text>
                {data?.tags?.map((t, i) => <Tag key={i} color={t.color}>{t.title}</Tag>)}
            </HStack>
        </Stack>
    </HStack>
}

export default UserView