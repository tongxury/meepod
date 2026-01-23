import {
    FlatList,
    SafeAreaView,
    ScrollView,
    StatusBar,
    View,
    StyleSheet,
    Linking,
    ImageBackground,
    Image
} from "react-native";
import React, { useCallback, useContext, useEffect, useState } from "react";
import { Flex, Grid, Toast, WhiteSpace, WingBlank } from '@ant-design/react-native';
import { Avatar, Button, Card, IconButton, useTheme, Text, AnimatedFAB } from "react-native-paper";
import { useFocusEffect } from "@react-navigation/native";
import { HStack, Stack } from "@react-native-material/core";
import { useSafeAreaInsets } from "react-native-safe-area-context";
import { useRequest } from "ahooks";
import { fetchItemStates } from "../../service/api";
import { Chip, FAB, Skeleton } from "@rneui/themed";
import CountDownTimer from "../../components/CountDownTimer";
import Tag from "../../components/Tag";
import { LinearGradient } from "expo-linear-gradient";
import Svg, { Circle, Rect } from 'react-native-svg';
import { lotteryIcons } from '../../utils/lotteryIcons';



export const HomeScreen = ({ navigation }) => {

    const { data, loading, run } = useRequest(fetchItemStates, { manual: true })
    const items = data?.data?.data
    const supported = ['ssq', 'f3d', 'x7c', 'rx9', 'sfc', 'zjc', 'pl3', 'pl5', 'dlt', 'kl8']

    useFocusEffect(useCallback(() => {
        run()
    }, []))


    const { colors } = useTheme()
    const insets = useSafeAreaInsets();

    const itemHeight = 160;


    if (loading && !items) {
        return <Stack style={{
            flex: 1,
            // paddingTop: insets.top
        }}>
            <ImageBackground
                source={require("../../assets/home_bg.png")}
                style={{ height: 240 }}
            />
            <FlatList
                style={{ borderRadius: 10, marginTop: -160, marginHorizontal: 5 }}
                data={[1, 2, 3, 4, 5, 6, 7, 8, 9, 10]}
                numColumns={2}
                renderItem={({ item: x }) => {
                    return <Skeleton animation="wave"
                        style={{ borderRadius: 5, flex: 1, height: itemHeight, margin: 2 }} />
                }} />
        </Stack>
    }

    return (
        <Stack style={{
            flex: 1,
            // paddingTop: insets.top,
        }}>
            <ImageBackground
                source={require("../../assets/home_bg.png")}
                style={{ height: 240 }}
            />
            <FlatList
                style={{ borderRadius: 10, marginTop: -160, marginHorizontal: 5 }}
                data={items}
                numColumns={2}
                renderItem={({ item: x }) => {
                    return <Card
                        style={{ flex: 1, height: itemHeight, borderRadius: 0 }}
                        contentStyle={{ flex: 1 }}
                        onPress={() => {
                            if (supported.includes(x.id)) {
                                if (x.disabled) {
                                    Toast.info(x?.status?.desc)
                                } else {
                                    navigation.navigate('Plan', { id: x.id, name: x.name })
                                }
                            } else {
                                Toast.info('请升级最新版')
                            }
                        }}>
                        <Stack fill={true} spacing={10} p={10} justify={"between"}>
                            <Stack items="center" spacing={10}>
                                <Avatar.Image style={{ backgroundColor: colors.primaryContainer }}
                                    source={lotteryIcons[x.id] || { uri: x.icon }} />
                                <Text variant="titleMedium" style={{ textAlign: 'center' }}>{x.name}</Text>
                            </Stack>
                            <HStack items={"center"} justify={"between"}>
                                <Tag color={x.status?.color}>{x.status?.name}</Tag>
                                {x.extra?.type === 'countdown' && <CountDownTimer color={colors.primary}
                                    timeLeftSecond={x.extra?.value} />}
                                {x.extra.type === 'text' && <Text variant={"titleSmall"}>{x.extra?.value}</Text>}
                            </HStack>
                        </Stack>
                    </Card>
                }}
                keyExtractor={item => item.name}
            />
            <View>
                <FAB
                    visible={true}
                    onPress={() => {
                        navigation.navigate('More')
                    }}
                    // placement="right"
                    style={{
                        position: 'absolute',
                        right: 0,
                        bottom: 20,
                    }}
                    titleStyle={{
                        fontSize: 16
                    }}
                    containerStyle={{
                        borderTopRightRadius: 0,
                        borderBottomRightRadius: 0
                    }}
                    title="更多玩法"
                    icon={false}
                    color={colors.primary}
                />
            </View>

        </Stack>
    );
}

export default HomeScreen

