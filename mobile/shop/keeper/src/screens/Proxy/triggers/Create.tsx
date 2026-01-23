import editor from "../../Profile/Editor";
import {InputItem, Modal} from "@ant-design/react-native";
import {HStack, Stack} from "@react-native-material/core";
import {Button as RneButton, Image, Input} from "@rneui/themed";
import {Button, Text, useTheme} from "react-native-paper";
import {StyleProp, View, ViewStyle, TextInput} from "react-native";
import React, {useContext, useEffect, useState} from "react";
import {useForm, Controller} from "react-hook-form";
import UserView from "../../../components/UserView";
import {useRequest} from "ahooks";
import {addProxy, fetchStoreUsers} from "../../../service/api";
// import AntDesignIcon from "react-native-vector-icons/AntDesign";
import {AntDesign as AntDesignIcon} from "@expo/vector-icons";
import {Selector} from "../../../components/Selector";
import {AppContext} from "../../../providers/global";


const CreateTrigger = ({onConfirm, style}: { onConfirm?: () => void, style?: StyleProp<ViewStyle> }) => {

    const [open, setOpen] = useState<boolean>(false)

    const [phone, setPhone] = useState<string>()
    const [rewardRate, setRewardRate] = useState<string>('0.02')

    const {data, run} = useRequest(fetchStoreUsers, {manual: true, debounceWait: 1000})
    const users = data?.data?.data?.list

    const {settingsState: {settings: {proxy}}} = useContext<any>(AppContext);


    useEffect(() => {
        if (phone) {
            run({phone})
        }
    }, [phone])

    const confirm = () => {

        addProxy({userId: users?.[0]?.user?.id, rewardRate}).then(rsp => {
            if (rsp.data?.code === 0) {
                setOpen(false)
                onConfirm?.()
            }
        })
    }


    const {colors} = useTheme()

    return <View style={style}>
        <RneButton icon={<AntDesignIcon name={'plus'} color={colors.onPrimary}/>}
                   onPress={() => setOpen(true)}>添加推广员</RneButton>
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
                            placeholder={"必须为店内用户"}/>
                        {users?.length> 0 && <UserView data={users?.[0]?.user}/>}
                    </Stack>

                    <Stack spacing={5} fill={1}>
                        <Text variant={"titleMedium"}>佣金百分比</Text>

                        <Selector
                            items={proxy?.rewardRateItems}
                            itemStyle={{selected: colors.primary, unselected: colors.primaryContainer}}
                            selectedKeys={[rewardRate]} onSelectChanged={(item, selected) => {
                            if (selected) setRewardRate(item.value)
                        }}/>

                    </Stack>

                </Stack>

                <Button mode="contained" disabled={!(users?.length > 0)} onPress={confirm}>确认</Button>
            </Stack>
        </Modal>
    </View>

}

export default CreateTrigger