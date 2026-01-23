import {ActivityIndicator, FlatList, Pressable, StyleProp, View, ViewStyle} from "react-native";
import React, {useContext, useEffect, useState} from "react";
import {Avatar, Button, Dialog, Portal, RadioButton, Text, TextInput} from "react-native-paper";
import {ImagePicker, InputItem, Modal, TextareaItem, WhiteSpace, WingBlank} from "@ant-design/react-native";
import {HStack, Stack} from "@react-native-material/core";
import {StoreUser, User} from "../../service/typs";
import UserView from "../../components/UserView";
import {useRequest} from "ahooks";
import {fetchStoreUser, updateStoreUser} from "../../service/api";
import {set} from "react-hook-form";
import Tag from "../../components/Tag";
import {Divider} from "@rneui/themed";
import ImageViewer from "../../components/ImageViewer";


const UserTrigger = ({data, style}: { data: User, style?: StyleProp<ViewStyle> }) => {

    const [visible, setVisible] = useState(false);

    const [storeUser, setStoreUser] = useState<StoreUser>()

    const {runAsync} = useRequest(fetchStoreUser, {manual: true})

    useEffect(() => {
        if (visible) {
            runAsync({userId: data?.id}).then(rsp => {
                setStoreUser(rsp?.data?.data)
            })
        } else {
            setStoreUser(undefined)
        }
    }, [visible])

    const changeRemark = () => {
        Modal.prompt(
            '修改备注',
            '',
            (value: any) => {
                updateStoreUser({userId: data?.id, field: 'remark', value}).then(rsp => {
                    if (rsp?.data?.code === 0) {
                        setStoreUser(rsp?.data?.data)
                    }
                })
            },
            'default',
            '',
            ['输入新备注'],
        )
    }

    return <View style={style}>
        <Pressable onPress={() => setVisible(true)}>
            <UserView data={data}/>
        </Pressable>
        <Modal
            style={{borderTopLeftRadius: 10, borderTopRightRadius: 10}}
            popup
            visible={visible && !!storeUser}
            maskClosable={true}
            animationType="slide-up"
            onClose={() => {
                setVisible(false)
            }}>
            <Stack p={20} spacing={10}>
                <Stack fill={1} items={"center"} spacing={10}>
                    <Avatar.Image source={{uri: storeUser?.user?.icon}}/>
                    <Text variant={"titleMedium"}>{storeUser?.remark ?? storeUser?.user?.nickname}</Text>
                    {storeUser?.user?.tags && <HStack items={"center"} spacing={5}>
                        {storeUser?.user?.tags?.map((t, i) => <Tag key={i} title={t.title} color={t.color}/>)}
                    </HStack>}
                </Stack>
                <Stack spacing={5}>
                    <HStack justify={"between"}>
                        <Text variant={"bodySmall"}>手机号: </Text>
                        <Text variant={"bodySmall"}>{storeUser?.user?.phone} </Text>
                    </HStack>

                    <HStack justify={"between"} items={"center"}>
                        <Text variant={"bodySmall"}>微信收款二维码: </Text>
                        <ImageViewer size={"small"} images={[{url: storeUser?.user?.wechat_pay_qrcode}]}/>
                    </HStack>
                    <HStack justify={"between"} items={"center"}>
                        <Text variant={"bodySmall"}>支付宝收款二维码: </Text>
                        <ImageViewer size={"small"} images={[{url: storeUser?.user?.ali_pay_qrcode}]}/>
                    </HStack>
                </Stack>
                <HStack items={'center'} justify={"around"}>
                    <Button onPress={changeRemark}>设置备注</Button>
                </HStack>
            </Stack>
        </Modal>

    </View>
}

export default UserTrigger