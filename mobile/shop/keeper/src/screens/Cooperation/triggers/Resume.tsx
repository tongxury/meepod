import {Pressable, StyleProp, View, ViewStyle} from "react-native";
import React, {useState} from "react";
import {Modal} from "@ant-design/react-native";
import {HStack, Stack} from "@react-native-material/core";
import {Button as RneButton} from "@rneui/themed";
import {Button} from "react-native-paper";
import {deleteProxy, pauseCoStore, recoverCoStore, recoverProxy, resumeCoStore} from "../../../service/api";

const ResumeTrigger = ({coStoreId, onConfirm, style}: {
    coStoreId: string,
    onConfirm?: () => void,
    style?: StyleProp<ViewStyle>
}) => {

    const [visible, setVisible] = useState(false);

    const confirm = () => {
        resumeCoStore(coStoreId).then(rsp => {
            setVisible(false)
            onConfirm?.()
        })
    }

    return <View style={style}>
        <RneButton onPress={confirm} >恢复</RneButton>
    </View>
}

export default ResumeTrigger