import {FlatList, Pressable, View} from "react-native";
import React, {useCallback, useEffect, useState} from "react";
import {useInfiniteScroll, useRequest} from "ahooks";
import {fetchMatches, fetchMatchFilters, fetchOrders, pay} from "../../service/api";
import {Empty, Footer} from "../../components/ListComponent";
import {Text, useTheme} from "react-native-paper";
import {useFocusEffect, useNavigation} from "@react-navigation/native";

import {mainBodyHeight} from "../../utils/dimensions";
import {HStack, Stack} from "@react-native-material/core";
import {Button, Button as RneButton, ButtonGroup, Chip} from "@rneui/themed";
import {ButtonSelector} from "../../components/ButtonSelector";
import ListView from "../../components/ListView";

const MatchList = ({category, issue}) => {

    const navigation = useNavigation()

    const {colors} = useTheme()

    return <View style={{flex: 1}}>

        <ListView
            fetch={page => fetchMatches({category, issue: issue ?? '', page})}
            reloadDeps={[issue]}
            renderItem={(x, updateListItem, reload) =>
                <Stack p={10} style={{flex: 1}} spacing={8} bg={colors.background}>
                    <HStack items="center" spacing={5}>
                        <View><Chip buttonStyle={{backgroundColor: colors.primary}}
                                    titleStyle={{color: colors.onPrimary}}
                                    radius={0}>{x.league}</Chip></View>
                        <HStack items="center" justify={"between"} fill={true} spacing={10}>
                            <Text variant="titleMedium" numberOfLines={1}
                                  ellipsizeMode="tail">{x.home_team_tag}{x.home_team}</Text>
                            <Text variant="titleMedium">VS</Text>
                            <Text variant="titleMedium"
                                  numberOfLines={1}>{x.guest_team_tag}{x.guest_team}</Text>
                        </HStack>
                    </HStack>


                    <ButtonGroup

                        containerStyle={{height: 30}}
                        buttons={x.odds?.items?.map(op => ({
                            element: () => op.result === x.result?.value ?
                                <Text style={{
                                    color: colors.primary,
                                    fontWeight: "bold"
                                }}>{op.name} {op.value ?? ''}</Text> :
                                <Text>{op.name} {op.value ?? ''}</Text>
                        }))}
                        // selectedIndex={1}
                        disabled={true}
                        // disabledSelectedStyle={{backgroundColor: colors.primary,}}
                        // disabledSelectedTextStyle={{color: colors.onPrimary,}}
                        // onPress={(value) => {
                        //     setSelectedIndex(value);
                        // }}
                        // containerStyle={{ marginBottom: 20 }}
                    />


                    <HStack items={"center"} justify={"between"} spacing={5}>
                        <Text variant={"bodySmall"}>{x.start_at}</Text>
                        <HStack spacing={5}>
                            <Text variant={"bodySmall"}>{x.result?.goals}</Text>
                            <Text variant={"bodySmall"} style={{color: x.status?.color}}>{x.status?.name}</Text>
                        </HStack>
                    </HStack>

                </Stack>

            }
        />
    </View>

}

export default MatchList