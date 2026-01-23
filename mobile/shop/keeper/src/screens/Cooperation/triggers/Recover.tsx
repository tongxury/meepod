import {Pressable, StyleProp, View, ViewStyle} from "react-native";
import React, {useState} from "react";
import {Modal} from "@ant-design/react-native";
import {HStack, Stack} from "@react-native-material/core";
import {Button as RneButton} from "@rneui/themed";
import {Button} from "react-native-paper";
import {deleteProxy, recoverCoStore, recoverProxy} from "../../../service/api";

const RecoverTrigger = ({id, onConfirm, style}: {
    id: string,
    onConfirm?: () => void,
    style?: StyleProp<ViewStyle>
}) => {

    const [visible, setVisible] = useState(false);

    const confirm = () => {
        recoverCoStore(id).then(rsp => {
            setVisible(false)
            onConfirm?.()
        })
    }

    return <View style={style}>
        <RneButton onPress={() => setVisible(true)} >恢复</RneButton>
        <Modal
            style={{borderTopLeftRadius: 10, borderTopRightRadius: 10}}
            popup
            visible={visible}
            maskClosable={true}
            animationType="slide-up"
            onClose={() => {
                setVisible(false)
            }}>
            <Stack p={20}>
                <HStack items={"center"} spacing={10} fill={1}>
                    <Button style={{flex: 1}} mode={"contained"} onPress={() => {
                        setVisible(false);
                    }}>取消</Button>

                    <Button mode={"contained-tonal"} onPress={confirm}
                            style={{flex: 1}}>确认</Button>
                </HStack>
            </Stack>
        </Modal>

    </View>
}

export default RecoverTrigger