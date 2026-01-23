import {InputItem, Modal, WhiteSpace} from "@ant-design/react-native";
import {View} from "react-native";
import {Avatar, Button, Card, Text, IconButton, TextInput, useTheme} from "react-native-paper";
import React, {useEffect, useMemo, useState} from "react";
import {useRequest} from "ahooks";
import {fetchStoreInfo} from "../../../service/api";
import {Stack} from "@react-native-material/core";

const StoreSwitcher = ({open, onClose, onSwitch}: {
    open: boolean,
    onClose: () => void,
    onSwitch: (storeId, proxyId) => void
}) => {

    const [storeInput, setStoreInput] = useState<string>('')
    const [proxyId, setProxyId] = useState<string>('')

    const {data, run: fetch} = useRequest(fetchStoreInfo, {
        manual: true,
        debounceWait: 500
    })

    const storeInfo = data?.data?.data


    useEffect(() => {
        if (storeInput) {
            fetch(storeInput)
        }
    }, [storeInput])

    const {colors} = useTheme()

    return <Modal
        popup
        animationType="slide-up"
        onClose={onClose}
        style={{borderTopLeftRadius: 10, borderTopRightRadius: 10}}
        maskClosable
        visible={open}
        bodyStyle={{paddingVertical: 30, paddingHorizontal: 15}}
    >
        <Stack items="center" spacing={30} style={{flex: 1}}>
            <Stack items="center" spacing={8}>
                <Text variant="titleMedium" style={{fontSize: 20}}>选择店铺</Text>
                <Text variant="bodySmall" style={{color: colors.primary}}>邀请码请与彩店店主或代理联系获取</Text>
            </Stack>
            <Stack spacing={20} style={{paddingHorizontal: 50}}>
                <TextInput
                    outlineStyle={{borderRadius: 10}}
                    mode="outlined"
                    placeholder="请输入店铺邀请码"
                    value={storeInput}
                    onChangeText={text => setStoreInput(text)}
                />

                {
                    storeInfo &&
                    <Card.Title
                        title={storeInfo?.name}
                        // subtitle="Card Subtitle"
                        left={(props) => <Avatar.Image style={{backgroundColor: colors.background}}  {...props}
                                                       source={{uri: storeInfo?.icon}}/>}
                        // right={(props) => <IconButton {...props} icon="dots-vertical" onPress={() => {}} />}
                    />

                }
                <TextInput
                    outlineStyle={{borderRadius: 10}}
                    mode="outlined"
                    placeholder="请输入推荐人邀请码"
                    value={proxyId}
                    onChangeText={text => setProxyId(text)}
                />
                <Button disabled={!storeInfo} onPress={() => onSwitch?.(storeInput, proxyId)}
                        mode="contained">确认选择</Button>
            </Stack>

        </Stack>
    </Modal>
}
export default StoreSwitcher