import editor from "../../Profile/Editor";
import {InputItem, Modal} from "@ant-design/react-native";
import {HStack, Stack} from "@react-native-material/core";
import {Button as RneButton, Image, Input} from "@rneui/themed";
import {Button, Text, useTheme} from "react-native-paper";
import {StyleProp, View, ViewStyle, TextInput} from "react-native";
import React, {useEffect, useState} from "react";
import {useForm, Controller} from "react-hook-form";
import UserView from "../../../components/UserView";
import {useRequest} from "ahooks";
import {addProxy, addProxyUser, fetchStoreUsers} from "../../../service/api";
import {Proxy} from "../../../service/typs";


const AddUserTrigger = ({proxy, style}: { proxy: Proxy, style?: StyleProp<ViewStyle> }) => {

    const [open, setOpen] = useState<boolean>(false)

    const [phone, setPhone] = useState<string>()
    const [rewardRate, setRewardRate] = useState<number>(1)

    const {data, run} = useRequest(fetchStoreUsers, {manual: true})
    const users = data?.data?.data?.list


    useEffect(() => {
        run({phone})
    }, [phone])

    const onConfirm = () => {
        addProxyUser(proxy?.id, users?.[0]?.user?.id).then(rsp => {
            if (rsp.data?.code === 0) {
                setOpen(false)
            }
        })
    }

    const {colors} = useTheme()

    return <View style={style}>
        <RneButton onPress={() => setOpen(true)}>添加用户</RneButton>
        <Modal
            style={{borderTopLeftRadius: 10, borderTopRightRadius: 10}}
            popup
            visible={open}
            maskClosable={true}
            animationType="slide-up"
            onClose={() => {
                setOpen(false)
            }}>
            <Stack p={20} spacing={30}>
                <Stack spacing={20}>

                    <Stack spacing={8} fill={1}>
                        <Text variant={"titleMedium"}>推广员</Text>
                        <UserView data={proxy?.user}/>
                    </Stack>

                    <Stack spacing={8} fill={1}>
                        <Text variant={"titleMedium"}>手机号查询</Text>
                        <TextInput
                            value={phone}
                            onChangeText={text => {
                                setPhone(text)
                            }}
                            style={{
                                minHeight: 40,
                                flex: 1,
                                paddingHorizontal: 10,
                                backgroundColor: colors.primaryContainer,
                                borderRadius: 5
                            }}
                            placeholder={"用户手机号"}/>
                        {users?.length > 0 && <UserView data={users?.[0]?.user}/>}
                    </Stack>

                </Stack>


                <Button mode="contained" disabled={users?.length === 0} onPress={onConfirm}>确认</Button>
            </Stack>
        </Modal>
    </View>

}

export default AddUserTrigger