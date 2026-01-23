import {Platform, Pressable, View} from "react-native";
import {Appbar, Avatar, Button, List, Text, useTheme} from "react-native-paper";
import React, {useCallback} from "react";
import {Modal, Toast, WhiteSpace} from "@ant-design/react-native";
import {useRequest} from "ahooks";
import {useFocusEffect} from "@react-navigation/native";
import {fetchMyStore, updateStore} from "../../service/api";
import UploaderTrigger from "../../triggers/Uploader";
import {HStack, Stack} from "@react-native-material/core";
import useAuthState from "../../hooks/auth";
import ItemSettingsTrigger from "../../triggers/ItemSettings";
import DebugTrigger from "./triggers/Debug";
import Constants from "expo-constants";
import {checkForUpdateAsync, fetchUpdateAsync, reloadAsync} from "expo-updates";
import {Config} from "../../../config";

const ProfileEditorScreen = ({navigation}) => {

    const {data, loading, run, mutate} = useRequest(fetchMyStore, {manual: true})
    const result = data?.data
    const store = result?.data

    useFocusEffect(useCallback(() => {
        run()
    }, []))

    const {logout} = useAuthState()

    const {colors} = useTheme()

    const checkUpdate = () => {
        checkForUpdateAsync().then(rsp => {
            if (rsp.isAvailable) {
                fetchUpdateAsync().then(rsp => {
                    reloadAsync().then(() => {
                        Toast.info('已更新为最新版本')
                    })
                })
            } else {
                Toast.info('已是最新版本')
            }
        });
    }

    return <Stack style={{flex: 1}}>
        <Appbar.Header>
            <Appbar.BackAction onPress={() => {
                navigation.goBack()
            }}/>
            <Appbar.Content title={
                <HStack items={'center'} spacing={2}>
                    <Text variant="titleMedium">店铺设置</Text>
                    <DebugTrigger/>
                </HStack>
            }/>

        </Appbar.Header>
        <WhiteSpace/>
        <Stack fill={true} bg={colors.background} justify="between">
            <Stack>
                <UploaderTrigger onUploaded={urls => {
                    updateStore({field: 'icon', value: urls[0].key}).then(rsp => {
                        if (rsp?.data?.code === 0) {
                            mutate(rsp)
                        }
                    })
                }}>
                    <List.Item
                        title="修改头像"
                        right={props =>
                            <Avatar.Image style={{backgroundColor: colors.secondaryContainer}} size={50}
                                          source={{uri: store?.icon}}/>}
                    />
                </UploaderTrigger>
                <ItemSettingsTrigger values={store?.selected_item_ids} onConfirm={run}>
                    <List.Item
                        title="设置彩种"
                        right={props => <Text variant={"titleMedium"}
                                              style={{color: colors.primary}}>已设置{store?.selected_item_ids?.length ?? 0}个</Text>}
                    />
                </ItemSettingsTrigger>
                {Platform.OS != 'web' &&
                    <List.Item
                        onPress={checkUpdate}
                        title="检查更新"
                        right={props => <Text>v{Config().version}</Text>}
                    />
                }
            </Stack>

            <Button style={{margin: 10}} onPress={() => logout(() => {
                navigation.navigate('Root')
            })} mode="contained">退出</Button>
        </Stack>
    </Stack>
}

export default ProfileEditorScreen