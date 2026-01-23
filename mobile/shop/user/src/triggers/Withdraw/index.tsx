import {Modal} from "@ant-design/react-native";
import React, {useState} from "react";
import {Pressable, View} from "react-native";
import {Button, Chip, Text, TextInput, useTheme} from "react-native-paper";
import {styles} from "../../utils/styles";
import {HStack, Stack} from "@react-native-material/core";
import {addWithdrawOrder} from "../../service/api";

const WithdrawTrigger = ({balance, onConfirmed, children}: {
    balance: number,
    onConfirmed?: (amount: number) => void,
    children: React.ReactNode,
}) => {

    const [open, setOpen] = useState<boolean>(false)
    const {colors} = useTheme()
    const [amount, setAmount] = useState<number>()

    const onChange = (text: string) => {

        if (!text || Number(text)) {
            setAmount(isNaN(parseInt(text)) ? 0 : parseInt(text))
        }
    }

    const onConfirm = () => {
        addWithdrawOrder({amount}).then(rsp => {
            setOpen(false)
            onConfirmed?.(amount)
        })
    }

    return <Pressable onPress={() => setOpen(true)}>
        {children}
        <Modal
            popup
            style={{...styles.popup}}
            visible={open}
            maskClosable={true}
            animationType="slide-up"
            onClose={() => setOpen(false)}
        >
            <Stack spacing={30}>
                <Stack spacing={10}>
                    <Text variant="titleMedium" style={{fontWeight: "bold", textAlign: "center"}}>提现</Text>

                    <Text variant="titleMedium">当前总余额 <Text
                        style={{color: colors.primary, fontWeight: 'bold'}}>{balance}元</Text></Text>

                    <TextInput value={amount?.toString() ?? '0'} placeholder="请输入金额" dense mode="outlined"
                               outlineStyle={{borderRadius: 10}} onChangeText={onChange}></TextInput>
                    <HStack items="center" spacing={5}>
                        <Chip onPress={() => onChange(`${balance}`)}>全部</Chip>
                        <Chip onPress={() => onChange(`${balance / 2}`)}>一半</Chip>
                        <Chip disabled={balance < 500} onPress={() => onChange(`${balance - 500}`)}>留500</Chip>
                        <Chip disabled={balance < 100} onPress={() => onChange(`${balance - 100}`)}>留100</Chip>
                    </HStack>
                </Stack>

                <HStack items="center" spacing={10}>
                    <Button mode="contained-tonal" style={{flex: 1}} onPress={() => setOpen(false)}>取消</Button>
                    <Button disabled={amount > balance || amount <= 0} mode="contained" style={{flex: 1}}
                            onPress={onConfirm}>确认</Button>
                </HStack>
            </Stack>
        </Modal>
    </Pressable>

}

export default WithdrawTrigger