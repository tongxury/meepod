import {ActivityIndicator, Pressable, View} from "react-native";
import {useRequest} from "ahooks";
import {useFocusEffect} from "@react-navigation/native";
import React, {useCallback, useContext} from "react";
import {Avatar, useTheme, Text, Card} from "react-native-paper";
import {Button, Chip, Image} from "@rneui/themed";
import {HStack, Stack} from "@react-native-material/core";
import {useSafeAreaInsets} from "react-native-safe-area-context";
import IconFontAwesome5 from "react-native-vector-icons/FontAwesome5";
import IconFontAwesome from "react-native-vector-icons/FontAwesome";
import {AppContext} from "../../providers/global";
import {Badge, Toast} from "@ant-design/react-native";
import IconEtp from "react-native-vector-icons/Entypo";
import UploaderTrigger from "../../triggers/Uploader";
import CmrServiceTrigger from "../../triggers/CmrService";
import {fetchMyStore, updateStore} from "../../service/api";
import AmountView from "../../components/AmountView";
import NoticeTrigger from "./triggers/Notice";
import IconMci from "react-native-vector-icons/MaterialCommunityIcons";
import Ionicon from "react-native-vector-icons/Ionicons";
import * as Clipboard from 'expo-clipboard';
import FontAwesome5Icon from "react-native-vector-icons/FontAwesome5";
import {ImageBackground} from "expo-image";

function ProfileScreen({navigation}) {

    const {data, loading, runAsync, mutate} = useRequest(fetchMyStore, {manual: true})
    const result = data?.data
    const store = result?.data

    useFocusEffect(useCallback(() => {
        refresh()
    }, []))

    const refresh = () => {
        runAsync().then()
    }

    const {colors} = useTheme()

    const {top} = useSafeAreaInsets();

    if (loading) {
        return <View></View>
    }

    const {
        counterState: {counter},
        settingsState: {settings},
        updateState: {update, checkToUpdate}
    } = useContext<any>(AppContext);

    const copyToClipboard = (text) => {
        Clipboard.setStringAsync(text).then(result => {
            if (result) {
                Toast.info('邀请码已复制')
            }
        });
    };

    return (
        <Stack style={{flex: 1}} spacing={10}>

            <ImageBackground
                source={require("../../assets/home_bg.png")}
                style={{height: 260, width: '100%', paddingTop: top + 30, paddingBottom: 30}}
            >
                <Stack center={true} spacing={20}>

                    <Pressable onPress={() => navigation.navigate('ProfileEditor')}>
                        <Avatar.Image size={80} style={{elevation: 2, backgroundColor: colors.secondaryContainer}}
                                      source={{uri: store?.icon}}/>
                    </Pressable>
                    <Stack spacing={8} items="center">
                        <HStack items={"center"} spacing={5}>
                            <Text
                                style={{fontSize: 18, fontWeight: "bold", color: colors.onPrimary}}>{store?.name}</Text>
                        </HStack>
                        <HStack items={"center"} spacing={5}>
                            <Text variant="bodyMedium" style={{color: colors.onPrimary}}>邀请码: </Text>
                            <Text variant="titleMedium" style={{color: colors.onPrimary}}
                                  onPress={() => copyToClipboard(store?.id)}>{store?.id}</Text>
                            <Ionicon name='copy' style={{color: colors.onPrimary}}
                                     onPress={() => copyToClipboard(store?.id)}/>
                        </HStack>

                        <HStack items={"center"} spacing={5}>
                            <View><Chip color={store?.member_level?.color}>{store?.member_level?.name}</Chip></View>
                            <View><Chip color={store?.status?.color}>{store?.status?.name}</Chip></View>
                            {/*<Text style={{color: store?.status?.color}}>{store?.status?.name}</Text>*/}
                        </HStack>
                    </Stack>
                </Stack>

            </ImageBackground>


            <HStack p={10} ph={15} bg={colors.background} items="center" justify="between">
                <Stack spacing={5}>
                    <Text variant="bodySmall">余额: </Text>
                    <AmountView size={"large"} amount={store?.balance}/>
                </Stack>
                <HStack items="center" spacing={8}>
                    <Button type={'clear'} onPress={() => navigation.navigate('Payment')}>明细</Button>
                    <Button color={colors.primary} onPress={() => Toast.info('请联系客服充值')}>充值</Button>
                </HStack>
            </HStack>

            <Stack style={{backgroundColor: colors.background}}>
                <HStack p={10} bg={colors.background} items="center" justify={"around"}>
                    {[
                        {
                            title: '基本信息',
                            icon: () => <IconFontAwesome5 color={colors.primary} size={30} name="user-edit"/>,
                            onPress: () => navigation.navigate('ProfileEditor')
                        },
                        {
                            title: '店铺公告',
                            icon: () => <NoticeTrigger notice={store?.notice} onConfirm={refresh}/>,
                        },
                        {
                            title: '联系客服',
                            icon: () => <CmrServiceTrigger>
                                <IconFontAwesome color={colors.primary} size={30} name="wechat"/>
                            </CmrServiceTrigger>,
                            onPress: () => undefined
                        },
                        {
                            title: '用户投诉',
                            icon: () =>
                                <Badge text={counter?.feedback} size={'small'}>
                                    <IconMci color={colors.primary} size={30} name="clipboard-edit"/>
                                </Badge>,
                            onPress: () => navigation.navigate('Feedback')
                        },
                    ].map((t, i) =>
                        <Pressable key={i} onPress={t.onPress}>
                            <Stack spacing={10} items="center" style={{padding: 10}}>
                                {t.icon && t.icon()}
                                {t.title && <Text variant="bodySmall">{t.title}</Text>}
                            </Stack>
                        </Pressable>
                    )}
                </HStack>

                <HStack p={10} bg={colors.background} items="center" justify={"around"}>
                    {[
                        // {
                        //     title: '店铺合作',
                        //     icon: () => <FontAwesome5Icon
                        //         name="hands-helping"
                        //         color={colors.primary} size={30}/>,
                        //     onPress: () => navigation.navigate('Cooperation'),
                        // },
                        {
                            title: '检查更新',
                            icon: () => <Badge dot={update} size={'small'}>
                                <IconMci color={colors.primary} size={30} name="refresh-circle"/>
                            </Badge>,
                            onPress: checkToUpdate
                        },
                        {title: '', icon: () => undefined, onPress: undefined},
                        {title: '', icon: () => undefined, onPress: undefined},
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


        </Stack>
    );
}

export default ProfileScreen