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
import {addCoStoreTopUp, addProxy, addProxyUser, fetchStoreUsers} from "../../../service/api";
import {CoStore, Proxy} from "../../../service/typs";
import StoreView from "../../../components/StoreView";


const TopUpTrigger = ({data, onConfirm, style}: {
    data: CoStore,
    onConfirm?: () => void,
    style?: StyleProp<ViewStyle>
}) => {

    const [open, setOpen] = useState<boolean>(false)

    const [amount, setAmount] = useState<string>()
    const confirm = () => {
        addCoStoreTopUp({storeId: data?.store?.id, amount: parseInt(amount)}).then(rsp => {
            if (rsp.data?.code === 0) {
                setOpen(false)
                onConfirm?.()
            }
        })
    }

    const {colors} = useTheme()

    return <View style={style}>
        <RneButton onPress={() => setOpen(true)}>充值</RneButton>
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
                <Stack spacing={8} fill={1}>
                    <Text variant={"titleMedium"}>添加合作店铺在我店的余额</Text>
                    <Text variant={"bodySmall"} style={{color: 'red'}}>请确认合作店铺已经全额转账</Text>

                    <StoreView data={data?.store}/>
                </Stack>

                <Stack spacing={8} fill={1}>
                    <Text variant={"titleMedium"}>金额</Text>
                    <TextInput onChangeText={text => setAmount(text)} value={amount} style={{
                        height: 40,
                        backgroundColor: colors.primaryContainer,
                        borderRadius: 5,
                        paddingHorizontal: 10
                    }}/>
                </Stack>

                <Button mode="contained" disabled={!(parseInt(amount) > 0)} onPress={confirm}>确认</Button>
            </Stack>
        </Modal>
    </View>

}

export default TopUpTrigger