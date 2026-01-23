import {Pressable, StyleProp, ViewStyle} from "react-native";
import React, {useState} from "react";
import {Button} from "react-native-paper";
import {Modal} from "@ant-design/react-native";
import {HStack, Stack} from "@react-native-material/core";

const ConfirmTrigger = ({onConfirm, children, style}: {
    onConfirm?: () => void,
    children: React.ReactNode, style?: StyleProp<ViewStyle>
}) => {

    const [visible, setVisible] = useState(false);


    return <Pressable style={style} onPress={() => setVisible(true)}>
        {children }
        <Modal
            style={{borderTopLeftRadius: 10, borderTopRightRadius: 10}}
            popup
            visible={visible}
            maskClosable={true}
            animationType="slide-up"
            onClose={() => {
                setVisible(false)
            }}>
            <Stack>

                <HStack items={"center"} spacing={10} fill={1}>
                    <Button mode="contained-tonal" style={{flex: 1}} onPress={() => {
                        setVisible(false);
                    }}>取消</Button>

                    <Button onPress={() => {
                        onConfirm?.()
                        setVisible(false)
                    }} mode="contained"
                            style={{flex: 1}}>确认</Button>
                </HStack>
            </Stack>
        </Modal>

    </Pressable>
}

export default ConfirmTrigger