import {ActivityIndicator, FlatList, Pressable, StyleProp, View, ViewStyle} from "react-native";
import React, {useContext, useState} from "react";
import {Avatar, Button, Dialog, Portal, RadioButton, Text, TextInput} from "react-native-paper";
import {AppContext} from "../../../providers/global";
import {Modal, WhiteSpace, WingBlank} from "@ant-design/react-native";
import {decrease} from "../../../service/api";
import {Account} from "../../../service/typs";
import {Stack} from "@react-native-material/core";
import {Button as RneButton} from "@rneui/themed"

const DecrTrigger = ({id, onConfirm, children, style}: {
    id: string,
    onConfirm?: (newValue: Account) => void,
    children?: React.ReactNode,
    style?: StyleProp<ViewStyle>
}) => {

    const [visible, setVisible] = useState(false);

    const [amount, setAmount] = useState<string>('0')
    const [remark, setRemark] = useState<string>()

    const {settingsState: {settings: appSettings}} = useContext<any>(AppContext);

    const confirm = () => {
        decrease({id, amount: parseInt(amount), remark}).then(rsp => {
            onConfirm?.(rsp.data?.data)
            setVisible(false)
        })
    }

    return <View style={style}>
        <RneButton size={'sm'} onPress={() => setVisible(true)} color={'error'}>扣减</RneButton>
        <Modal
            style={{borderTopLeftRadius: 10, borderTopRightRadius: 10}}
            popup
            visible={visible}
            maskClosable={true}
            animationType="slide-up"
            onClose={() => {
                setVisible(false)
            }}>
            <Stack p={20} spacing={15}>
                <Stack spacing={5}>
                    <Text variant="titleMedium">扣减金额</Text>
                    <TextInput
                        mode="outlined"
                        dense
                        outlineStyle={{borderRadius: 5}}
                        value={amount}
                        onChangeText={(value: any) => {
                            if (!isNaN(Number(value))) {
                                setAmount(value)
                            }
                        }}
                        placeholder="请输入金额">
                    </TextInput>
                </Stack>

                <Stack spacing={5}>
                    <Text variant="titleMedium">备注</Text>
                    <TextInput
                        mode="outlined"
                        dense
                        outlineStyle={{borderRadius: 5}}
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

    </View>
}

export default DecrTrigger