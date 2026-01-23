import {FlatList, ImageBackground, Pressable, View} from "react-native";
import {Avatar, Button, Card, IconButton, Text, useTheme} from "react-native-paper";
import {useRequest} from "ahooks";
import {fetchUserAccount, fetchUserProfile} from "../../service/api";
import {useFocusEffect} from "@react-navigation/native";
import React, {useCallback, useContext} from "react";
import IconMci from "react-native-vector-icons/MaterialCommunityIcons";
import IconAntd from "react-native-vector-icons/AntDesign";
import IconFontAwesome5 from "react-native-vector-icons/FontAwesome5";
import {HStack, Stack} from "@react-native-material/core";
import WithdrawTrigger from "../../triggers/Withdraw";
import {TopUpTrigger, TopUpTriggerV2} from "../../triggers/TopUp";
import {useSafeAreaInsets} from "react-native-safe-area-context";
import {checkForUpdateAsync, fetchUpdateAsync, reloadAsync} from "expo-updates";
import {Badge, Toast} from "@ant-design/react-native";
import Constants from "expo-constants";
import {Skeleton} from "@rneui/themed";
import useAppUpdates from "../../hooks/app_update";
import {AppContext} from "../../providers/global";

const ProfileScreen = ({navigation}) => {

    const {data, loading: profileLoading, runAsync} = useRequest(fetchUserProfile, {manual: true})
    const profile = data?.data?.data

    const {
        data: accountData,
        loading: accountLoading,
        runAsync: runAsyncAccount
    } = useRequest(fetchUserAccount, {manual: true})
    const account = accountData?.data?.data


    useFocusEffect(useCallback(() => {
        fetch()
    }, []))

    const fetch = () => {
        runAsync({userId: "me"}).then()
        runAsyncAccount({userId: "me"}).then()
    }

    const {colors} = useTheme()

    const toProfileEditor = () => {
        navigation.navigate('ProfileEditor')
    }

    const {top} = useSafeAreaInsets();

    const {updateState: {update, checkToUpdate}} = useContext<any>(AppContext);


    if ((profileLoading || accountLoading) && !profile) {
        return <Stack style={{flex: 1}} spacing={10}>
            <Stack bg={colors.background} pt={top + 30} pb={30} items="center" spacing={20}>
                <Skeleton animation={'wave'} width={80} height={80} circle/>
                <Stack spacing={5} items="center">
                    <Skeleton width={120} height={50}/>
                </Stack>
            </Stack>
            <HStack p={20} bg={colors.background} items="center" justify="between">
                <Stack spacing={10}>
                    <Skeleton width={50} height={16}/>
                    <Skeleton width={80} height={20}/>
                </Stack>

                <HStack items={"center"} spacing={10}>
                    <Skeleton width={60} height={25}/>
                    <Skeleton width={60} height={25}/>
                </HStack>
            </HStack>
            <Stack p={20} style={{flex: 1, backgroundColor: colors.background}} spacing={20}>
                <HStack spacing={20} justify={"between"} items={"center"}>
                    <Skeleton style={{flex: 1}} height={60}/>
                    <Skeleton style={{flex: 1}} height={60}/>
                    <Skeleton style={{flex: 1}} height={60}/>
                    <Skeleton style={{flex: 1}} height={60}/>
                </HStack>
            </Stack>
        </Stack>
    }


    return (
        <Stack style={{flex: 1}} spacing={10}>

            <ImageBackground
                source={require("../../assets/home_bg.png")}
                style={{height: 240, width: '100%', paddingTop: top + 30, paddingBottom: 30}}
            >
                <Stack fill={true} center={true} spacing={20}>
                    <Pressable onPress={toProfileEditor}>
                        <Avatar.Image size={80} style={{elevation: 2, backgroundColor: colors.secondaryContainer}}
                                      source={{uri: profile?.user?.icon}}/>
                    </Pressable>
                    <Stack spacing={5} items="center">
                        <Text style={{
                            fontSize: 18,
                            color: colors.onPrimary,
                            fontWeight: "bold"
                        }}>{profile?.user?.nickname}</Text>
                        <Text variant="bodySmall" style={{color: colors.onPrimary,}}>{profile?.user?.phone}</Text>
                    </Stack>
                </Stack>
            </ImageBackground>

            <HStack p={20} bg={colors.background} items="center" justify="between">
                <Stack spacing={8}>
                    <Text variant="bodySmall">账户余额</Text>
                    <Text><Text style={{color: colors.primary, fontWeight: "bold"}}
                                variant="titleLarge">{account?.balance ?? 0} </Text>元</Text>
                </Stack>
                <HStack items="center" spacing={8}>
                    <View><WithdrawTrigger balance={account?.balance} onConfirmed={() => fetch()}>
                        <Button mode="contained-tonal" labelStyle={{margin: 3}}>提现</Button></WithdrawTrigger></View>
                    <View><TopUpTriggerV2 onConfirmed={() => fetch()}>
                        <Button labelStyle={{margin: 3}} mode="contained">充值</Button></TopUpTriggerV2></View>
                </HStack>
            </HStack>

            <Stack style={{flex: 1, backgroundColor: colors.background}}>
                <HStack p={8} bg={colors.background} items="center" justify="between">
                    {[
                        {
                            title: '消费记录',
                            icon: () => <IconAntd color={colors.primary} size={30} name="pay-circle1"/>,
                            onPress: () => navigation.navigate('Payment', {tab: 0})
                        },
                        {
                            title: '充值记录',
                            icon: () => <IconFontAwesome5 color={colors.primary} size={30} name="money-bill"/>,
                            onPress: () => navigation.navigate('Payment', {tab: 1})
                        },
                        {
                            title: '提现记录',
                            icon: () => <IconFontAwesome5 color={colors.primary} size={30} name="file-invoice-dollar"/>,
                            onPress: () => navigation.navigate('Payment', {tab: 2})
                        },
                        {
                            title: '个人信息',
                            icon: () => <IconFontAwesome5 color={colors.primary} size={30} name="user-edit"/>,
                            onPress: toProfileEditor
                        },


                    ].map((t, i) =>
                        <Pressable key={i} onPress={t.onPress} style={{flex: 1}}>
                            <Stack spacing={10} items="center" p={5}>
                                {t.icon && t.icon()}
                                {t.title && <Text variant="bodySmall">{t.title}</Text>}
                            </Stack>
                        </Pressable>
                    )}
                </HStack>
                <HStack p={8} bg={colors.background} items="center" justify="between">
                    {[

                        {
                            title: '合作推广',
                            icon: () => <IconFontAwesome5 color={colors.primary} size={30} name="hands-helping"/>,
                            onPress: () => navigation.navigate('Proxy')
                        },

                        {
                            title: '投诉反馈',
                            icon: () => <IconMci color={colors.primary} size={30} name="clipboard-edit"/>,
                            onPress: () => navigation.navigate('Feedback')
                        },
                        {
                            title: '检查更新',
                            icon: () => <Badge dot={update} size={'small'}>
                                <IconMci color={colors.primary} size={30} name="refresh-circle"/>
                            </Badge>,
                            onPress: checkToUpdate
                        },
                        {title: '', icon: () => undefined, onPress: undefined},
                    ].map((t, i) =>
                        <Pressable key={i} onPress={t.onPress} style={{flex: 1}}>
                            <Stack spacing={10} items="center" p={5}>
                                {t.icon && t.icon()}
                                <Text variant="bodySmall">{t.title}</Text>
                            </Stack>
                        </Pressable>
                    )}
                </HStack>
            </Stack>

            <View></View>
        </Stack>
    );
}

export default ProfileScreen