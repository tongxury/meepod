import {ActivityIndicator, FlatList, Pressable, View} from "react-native";
import React, {useContext, useState} from "react";
import {Avatar, Button, Dialog, Portal, RadioButton, Text, TextInput} from "react-native-paper";
import {AppContext} from "../../../providers/global";
import {ImagePicker, InputItem, Modal, TextareaItem, WhiteSpace, WingBlank} from "@ant-design/react-native";
import {topUp} from "../../../service/api";
import {Account} from "../../../service/typs";
import {Stack} from "@react-native-material/core";


const TopUpTrigger = ({id, onConfirm, children}: {
    id: string,
    onConfirm?: (newValue: Account) => void,
    children: React.ReactNode
}) => {

    const [visible, setVisible] = useState(false);

    const [amount, setAmount] = useState<string>('0')
    const [remark, setRemark] = useState<string>()

    const {settingsState: {settings: appSettings}} = useContext<any>(AppContext);

    const confirm = () => {
        topUp({id, amount: parseInt(amount), remark}).then(rsp => {
            onConfirm?.(rsp.data?.data)
            setVisible(false)
        })
    }

    return <Pressable onPress={() => setVisible(true)}>
        {children}
        <Modal
            style={{borderTopLeftRadius: 10, borderTopRightRadius: 10}}
            popup
            visible={visible}
            maskClosable={true}
            animationType="slide-up"
            onClose={() => {
                setVisible(false)
            }}>
            <Stack p={20} spacing={30}>
                <Stack spacing={8}>
                    <Text variant="titleMedium">充值金额</Text>
                    <TextInput
                        mode="outlined"
                        outlineStyle={{borderRadius: 10}}
                        value={amount}
                        onChangeText={(value: any) => {
                            if (!isNaN(Number(value))) {
                                setAmount(value)
                            }
                        }}
                        placeholder="请输入金额">
                    </TextInput>
                </Stack>

                <Stack spacing={8}>
                    <Text variant="titleMedium">备注</Text>
                    <TextInput
                        mode="outlined"
                        outlineStyle={{borderRadius: 10}}
                        multiline
                        placeholder="输入备注"
                        value={remark}
                        onChangeText={val => setRemark(val)}
                        maxLength={100}
                        numberOfLines={4}
                    />
                </Stack>

                <View style={{flexDirection: "row", alignItems: "center"}}>
                    <Button mode="contained" style={{flex: 1}}>取消</Button>
                    <WingBlank size="sm"/>
                    <Button onPress={confirm} mode="contained-tonal"
                            style={{flex: 2}}>确认</Button>
                </View>
            </Stack>
        </Modal>

    </Pressable>
}

export default TopUpTrigger