import {ActivityIndicator, FlatList, Pressable, View} from "react-native";
import React, {useCallback, useContext, useEffect, useState} from "react";
import {ImagePicker, InputItem, Modal, TextareaItem, WhiteSpace, WingBlank} from "@ant-design/react-native";
import {Text} from "react-native-paper";
import {deleteOrder as delete_} from "../../../service/api";
import {Order} from "../../../service/typs";


const CancelTrigger = ({id, onConfirm, children}: {
    id: string
    onConfirm?: (newValue) => void,
    children?: React.ReactNode
}) => {

    const onDelete = () => {
        Modal.alert('', '确认撤销当前订单吗', [
            {text: '取消', onPress: undefined, style: 'cancel'},
            {
                text: '确认', onPress: () => {
                    delete_({id}).then(rsp => {
                        if (rsp.data?.code == 0) {
                            // mutate(rsp)
                            onConfirm?.(rsp)
                        }
                    })
                }
            },
        ])
    }


    return <Pressable onPress={onDelete}>
        <Text>撤单</Text>
    </Pressable>

    // return <Pressable onPress={onDelete}>
    //     {children}
    // </Pressable>
}

export default CancelTrigger