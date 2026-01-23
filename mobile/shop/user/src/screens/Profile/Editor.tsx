import {Platform, View} from "react-native";
import {Appbar, Avatar, Button, List, Text, useTheme} from "react-native-paper";
import React, {useCallback, useContext} from "react";
import {Modal, Toast, WhiteSpace} from "@ant-design/react-native";
import {Image} from "@rneui/themed"
import {fetchUser, updateUser} from "../../service/api";
import {useRequest} from "ahooks";
import {useFocusEffect} from "@react-navigation/native";
import UploaderTrigger from "../../triggers/Uploader";
import {HStack, Stack} from "@react-native-material/core";
import useAuthState from "../../hooks/auth";
import DebugTrigger from "../../triggers/Debug";
import {checkForUpdateAsync, fetchUpdateAsync, reloadAsync} from "expo-updates";
import Constants from "expo-constants";
import ImageViewer from "../../components/ImageViewer";
import useAppUpdates from "../../hooks/app_update";
import {AppContext} from "../../providers/global";
import {Config} from "../../../config";


const ProfileEditorScreen = ({navigation}) => {


    const {colors} = useTheme()
    const {data, loading, runAsync, mutate} = useRequest(fetchUser, {manual: true})
    const user = data?.data?.data

    useFocusEffect(useCallback(() => {
        runAsync({userId: "me"}).then()
    }, []))

    const {logout} = useAuthState()

    const {updateState: {checkToUpdate}} = useContext<any>(AppContext);


    return <View style={{flex: 1}}>
        <Appbar.Header>
            <Appbar.BackAction onPress={() => {
                navigation.goBack()
            }}/>
            <Appbar.Content title={
                <HStack items={'center'} spacing={2}>
                    <Text variant="titleMedium">个人信息</Text>
                    <DebugTrigger/>
                </HStack>
            }/>
        </Appbar.Header>
        <WhiteSpace/>
        <Stack fill={true} bg={colors.background} justify="between">
            <Stack>
                <UploaderTrigger onUploaded={urls => {
                    updateUser({field: 'icon', value: urls[0]}).then(rsp => {
                        if (rsp?.data?.code === 0) {
                            mutate(rsp)
                        }
                    })
                }}>
                    <List.Item
                        title="修改头像"
                        right={props =>
                            <Avatar.Image style={{backgroundColor: colors.secondaryContainer}} size={50}
                                          source={{uri: user?.icon}}/>
                        }
                    />
                </UploaderTrigger>
                <List.Item
                    onPress={() => {
                        Modal.prompt(
                            '修改昵称',
                            '',
                            (value: any) => {
                                updateUser({field: 'nickname', value}).then(rsp => {
                                    if (rsp?.data?.code === 0) {
                                        mutate(rsp)
                                    }
                                })
                            },
                            'default',
                            '',
                            ['输入新昵称'],
                        )
                    }}
                    title="修改昵称"
                    right={props => <Text>{user?.nickname}</Text>}
                />
                <UploaderTrigger onUploaded={urls => {
                    updateUser({field: 'wechat_pay_qrcode', value: urls[0]}).then(rsp => {
                        if (rsp?.data?.code === 0) {
                            mutate(rsp)
                        }
                    })
                }}>
                    <List.Item
                        title="微信收款二维码"
                        right={props =>
                            <ImageViewer size={"medium"} images={[{url: user?.wechat_pay_qrcode}]}/>
                        }
                    />
                </UploaderTrigger>
                <UploaderTrigger onUploaded={urls => {
                    updateUser({field: 'ali_pay_qrcode', value: urls[0]}).then(rsp => {
                        if (rsp?.data?.code === 0) {
                            mutate(rsp)
                        }
                    })
                }}>
                    <List.Item
                        title="支付宝收款二维码"
                        right={props =>
                            <ImageViewer size={"medium"} images={[{url: user?.ali_pay_qrcode}]}/>
                        }
                    />
                </UploaderTrigger>
                {Platform.OS != 'web' &&
                    <List.Item
                        onPress={checkToUpdate}
                        title="检查更新"
                        right={props => <Text>v{Config().version}</Text>}
                    />
                }
            </Stack>

            <Button onPress={() => logout(() => {
                navigation.navigate('Root', {screen: 'Home'})
            })} mode="contained" style={{margin: 10}}>
                退出
            </Button>
        </Stack>

    </View>
}

export default ProfileEditorScreen