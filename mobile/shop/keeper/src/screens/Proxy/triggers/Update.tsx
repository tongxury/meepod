import editor from "../../Profile/Editor";
import {InputItem, Modal} from "@ant-design/react-native";
import {HStack, Stack} from "@react-native-material/core";
import {Button as RneButton, Image, Input} from "@rneui/themed";
import {Button, Text, useTheme} from "react-native-paper";
import {StyleProp, View, ViewStyle, TextInput} from "react-native";
import React, {useContext, useEffect, useState} from "react";
import {useForm, Controller} from "react-hook-form";
import UserView from "../../../components/UserView";
import {addProxy, fetchStoreUsers, updateProxy} from "../../../service/api";
import {Proxy, User} from "../../../service/typs";
import {Selector} from "../../../components/Selector";
import {AppContext} from "../../../providers/global";


const UpdateTrigger = ({data, style, onConfirm}: { data: Proxy,  onConfirm?: () => void, style?: StyleProp<ViewStyle> }) => {

    const [open, setOpen] = useState<boolean>(false)

    const [rewardRate, setRewardRate] = useState<string>()

    const {settingsState: {settings: {proxy}}} = useContext<any>(AppContext);
    const confirm = () => {
        updateProxy(data?.id, rewardRate).then(rsp => {
            if (rsp.data?.code === 0) {
                setOpen(false)
                onConfirm?.()
            }
        })
    }

    useEffect(() => {
        setRewardRate(data?.reward_rate?.toString())
    }, [data])


    const {colors} = useTheme()

    return <View style={style}>
        <RneButton onPress={() => setOpen(true)}>佣金比例</RneButton>
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
                        <UserView data={data?.user}/>
                    </Stack>

                    <Stack spacing={5} fill={1}>
                        <Text variant={"titleMedium"}>佣金百分比</Text>
                        <Text variant={"bodySmall"} style={{color: 'red'}}>调整后历史订单不受影响，只有新的订单会按新的比例计算</Text>
                        <Selector
                            items={proxy?.rewardRateItems}
                            itemStyle={{selected: colors.primary, unselected: colors.primaryContainer}}
                            selectedKeys={[rewardRate]} onSelectChanged={(item, selected) => {
                            if (selected) setRewardRate(item.value)
                        }}/>
                    </Stack>

                </Stack>


                <Button mode="contained" onPress={confirm}>确认</Button>
            </Stack>
        </Modal>
    </View>

}

export default UpdateTrigger