import {View} from "react-native";
import {Button, Button as RneButton} from "@rneui/themed";
import React, {useState} from "react";
import {Modal} from "@ant-design/react-native";
import {StyleProp, ViewStyle} from "react-native";
import {TextInput, useTheme} from "react-native-paper";
import {Stack} from "@react-native-material/core";
import IconMci from "react-native-vector-icons/MaterialCommunityIcons";
import {addFeedback} from "../../../service/api";


const FeedbackTrigger = ({onConfirm, style}: {
    onConfirm?: () => void,
    style?: StyleProp<ViewStyle>
}) => {

    const [open, setOpen] = useState<boolean>(false)

    const [text, setText] = useState<string>()

    const {colors} = useTheme()

    const confirm = () => {
        addFeedback({text: text}).then(rsp => {
            setOpen(false)
            onConfirm?.()
        })
    }

    return <View style={style}>
        <Button onPress={() => setOpen(true)} style={{width: '100%'}}>提交反馈</Button>
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
                <TextInput
                    value={text}
                    onChangeText={text => setText(text)}
                    maxLength={100}
                    outlineStyle={{borderRadius: 8}} multiline
                    numberOfLines={6} dense mode={"flat"}
                    underlineStyle={{width: 0}}>

                </TextInput>
                <Button onPress={confirm}>确认</Button>
            </Stack>
        </Modal>
    </View>
}

export default FeedbackTrigger