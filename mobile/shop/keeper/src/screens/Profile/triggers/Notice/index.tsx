import {View} from "react-native";
import {Button, Button as RneButton} from "@rneui/themed";
import {AntDesign as AntDesignIcon} from "@expo/vector-icons";
import React, {useState} from "react";
import {Modal} from "@ant-design/react-native";
import {StyleProp, ViewStyle} from "react-native";
import {TextInput, useTheme} from "react-native-paper";
import IconFontAwesome from "react-native-vector-icons/FontAwesome";
import {Stack} from "@react-native-material/core";
import {updateStoreNotice} from "../../../../service/api";


const NoticeTrigger = ({notice, onConfirm, style}: { notice: string, onConfirm?: () => void, style?: StyleProp<ViewStyle> }) => {

    const [open, setOpen] = useState<boolean>(false)

    const [text, setText] = useState<string>(notice)

    const {colors} = useTheme()

    const confirm = () => {
        updateStoreNotice({notice: text}).then(rsp => {
            setOpen(false)
            onConfirm?.()
        })
    }

    return <View style={style}>
        <IconFontAwesome onPress={() => setOpen(true)} color={colors.primary} size={30} name="bell"/>
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

export default NoticeTrigger