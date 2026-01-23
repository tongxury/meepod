import {HStack, Stack} from "@react-native-material/core";
import {Avatar, Text} from "react-native-paper";
import {Store, User} from "../../service/typs";
import {StyleProp, ViewStyle} from "react-native";
import {Chip} from "@rneui/themed";
import Tag from "../Tag";

const UserView = ({data, style}: { data: User, style?: StyleProp<ViewStyle> }) => {
    return <HStack style={style} items={"center"} spacing={8}>
        <Avatar.Image size={30} source={{uri: data?.icon}}/>
        <Stack justify={'between'}>
            <HStack items={"center"} spacing={5}>
                <Text variant={"titleSmall"}>{data?.nickname}</Text>
                {data?.tags?.map((t, i) => <Tag key={i} title={t.title} color={t.color}/>)}
            </HStack>
        </Stack>
    </HStack>
}

export default UserView