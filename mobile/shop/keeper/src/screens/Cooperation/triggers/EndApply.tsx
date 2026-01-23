import {Pressable, StyleProp, View, ViewStyle} from "react-native";
import React, {useState} from "react";
import {Modal} from "@ant-design/react-native";
import {HStack, Stack} from "@react-native-material/core";
import {Button as RneButton} from "@rneui/themed";
import {Text} from "react-native-paper";
import {Button} from "react-native-paper";
import {applyEndCoStore, deleteProxy} from "../../../service/api";

const DeleteTrigger = ({id, onConfirm, style}: {
    id: string,
    onConfirm?: () => void,
    style?: StyleProp<ViewStyle>
}) => {

    const [visible, setVisible] = useState(false);

    const confirm = () => {
        applyEndCoStore(id).then(rsp => {
            setVisible(false)
            onConfirm?.()
        })
    }

    return <View style={style}>
        <RneButton type={'clear'} onPress={() => setVisible(true)} color={'error'}>终止合作</RneButton>
        <Modal
            style={{borderTopLeftRadius: 10, borderTopRightRadius: 10}}
            popup
            visible={visible}
            maskClosable={true}
            animationType="slide-up"
            onClose={() => {
                setVisible(false)
            }}>
            <Stack p={20} spacing={10}>

                <Text variant={'bodySmall'} style={{color: 'red'}}>终止合作需要合作方店铺确认并退还所有剩余的余额,
                    请主动与合作店铺沟通</Text>

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

export default DeleteTrigger