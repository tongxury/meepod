import {Pressable, StyleProp, View, ViewStyle} from "react-native";
import React, {useContext, useState} from "react";
import {Button, Dialog, Portal, RadioButton, Text, TextInput} from "react-native-paper";
import {Button as RneButton} from "@rneui/themed"
import {AppContext} from "../../../providers/global";
import {Modal, WingBlank} from "@ant-design/react-native";
import {HStack, Stack} from "@react-native-material/core";
import {rejectWithdraw} from "../../../service/api";
import {Withdraw} from "../../../service/typs";

const RejectTrigger = ({id, onConfirm, children, style}: {
    id: string,
    onConfirm?: (newValue: Withdraw) => void,
    children?: React.ReactNode
    style?: StyleProp<ViewStyle>
}) => {

    const [visible, setVisible] = useState(false);

    const [reason, setReason] = useState('');

    const {settingsState: {settings: appSettings}} = useContext<any>(AppContext);

    const confirm = () => {
        rejectWithdraw({id, reason}).then(rsp => {
            onConfirm?.(rsp.data?.data)
            setVisible(false)
        })
    }

    return <View style={style}>
        <RneButton size={'sm'} color={'error'} onPress={() => setVisible(true)}>拒单</RneButton>
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
                <Stack spacing={15}>
                    <RadioButton.Group onValueChange={newValue => setReason(newValue)} value={reason}>
                        {
                            appSettings.rejectWithdrawReasons.map(t =>
                                <HStack key={t} spacing={8} items="center">
                                    <RadioButton value={t}/>
                                    <Text>{t}</Text>
                                </HStack>)
                        }
                    </RadioButton.Group>
                    <TextInput
                        onChangeText={text => setReason(text)}
                        value={reason}
                        outlineStyle={{borderRadius: 10}} mode="outlined" maxLength={150} multiline
                        numberOfLines={4} placeholder="其他原因"/>
                </Stack>
                <HStack items="center">
                    <Button mode="contained" style={{flex: 1}} onPress={() => setVisible(false)}>取消</Button>
                    <WingBlank size="sm"/>
                    <Button onPress={confirm} mode="contained-tonal" disabled={!reason}
                            style={{flex: 2}}>确认</Button>
                </HStack>
            </Stack>

        </Modal>

    </View>
}

export default RejectTrigger