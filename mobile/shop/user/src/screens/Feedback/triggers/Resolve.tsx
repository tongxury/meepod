import {View} from "react-native";
import {Button, Button as RneButton} from "@rneui/themed";
import React, {useState} from "react";
import {Modal} from "@ant-design/react-native";
import {StyleProp, ViewStyle} from "react-native";
import {TextInput, useTheme} from "react-native-paper";
import {Stack} from "@react-native-material/core";
import {addFeedback, resolve} from "../../../service/api";
import {Feedback} from "../../../service/typs";


const ResolveTrigger = ({id, onConfirm, style}: {
    id: string,
    onConfirm?: (newValue: Feedback) => void,
    style?: StyleProp<ViewStyle>
}) => {

    const [open, setOpen] = useState<boolean>(false)

    const {colors} = useTheme()

    const confirm = () => {
        resolve({id}).then(rsp => {
            if (rsp?.data?.code === 0) {
                setOpen(false)
                onConfirm?.(rsp?.data?.data)
            }

        })
    }

    return <View style={style}>
        <Button size={'sm'} onPress={() => setOpen(true)} >已解决</Button>
        <Modal
            style={{borderTopLeftRadius: 10, borderTopRightRadius: 10}}
            popup
            visible={open}
            maskClosable={true}
            animationType="slide-up"
            onClose={() => {
                setOpen(false)
            }}>
            <Stack spacing={20} p={15}>
                <Button onPress={confirm}>确认</Button>
            </Stack>
        </Modal>
    </View>
}

export default ResolveTrigger