import {FlatList, Pressable, View} from "react-native";
import React, {useCallback, useState} from "react";
import {Avatar, Card, useTheme} from "react-native-paper";
import {useFocusEffect, useNavigation} from "@react-navigation/native";
import {Text} from "react-native-paper";
import {HStack, Stack} from "@react-native-material/core";
import {Button} from "@rneui/themed";
import ListView from "../../components/ListView";
import {fetchProxyUsers} from "../../service/api";
import UserView from "../../components/UserView";
import StatsView from "../../components/StatsView";

const UserList = ({}) => {

    const navigation = useNavigation()

    const [filterValues, setFilterValues] = useState<any>()

    const {colors} = useTheme()
    return <View style={{flex: 1}}>
        <ListView
            fetch={page => fetchProxyUsers({page})}
            reloadDeps={[filterValues]}
            renderItem={(x, updateListItem, reload) =>
                <Stack bg={colors.background} p={10} spacing={15}>
                    <HStack items={"center"} justify={"between"}>
                        <UserView data={x.user}/>
                        <Text>{x.created_at}</Text>
                    </HStack>
                    <HStack items={"center"} justify={"around"}>
                        <StatsView title="订单数量" value={x.order_count} unit={"单"}/>
                        <StatsView title="订单额度" value={x.order_amount} unit={"元"}/>
                    </HStack>
                </Stack>
            }
        />
    </View>
}

export default UserList