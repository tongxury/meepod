import {Pressable} from "react-native";
import React from "react";
import {Modal,} from "@ant-design/react-native";
import {cancelWithdraw} from "../../../service/api";
import {PaymentOrder} from "../../../service/typs";
import {Button} from "@rneui/themed";

const CancelWithdrawTrigger = ({id, onConfirm, children}: {
    id: string,
    onConfirm?: (newValue: PaymentOrder) => void,
    children: React.ReactNode
}) => {

    const onDelete = () => {
        Modal.alert('', '确认撤销当前订单吗', [
            {text: '取消', onPress: undefined, style: 'cancel'},
            {
                text: '确认', onPress: () => {
                    cancelWithdraw({id}).then(rsp => {
                        onConfirm?.(rsp?.data?.data)
                    })
                }
            },
        ])
    }

    return <Button onPress={onDelete} type={'clear'}>撤销</Button>
}

export default CancelWithdrawTrigger