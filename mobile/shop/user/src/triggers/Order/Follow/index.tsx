import {View} from "react-native";
import {Button as RneButton} from "@rneui/themed";
import React from "react";
import {Modal} from "@ant-design/react-native";
import {addOrder, deleteOrder as delete_, followOrder} from "../../../service/api";


const FollowTrigger = ({id, onConfirm}: { id: string, onConfirm?: (id: string) => void }) => {
    const onFollow = () => {
        Modal.alert('', '确认跟随当前订单吗', [
            {text: '取消', onPress: undefined, style: 'cancel'},
            {
                text: '确认', onPress: () => {

                    followOrder({followOrderId: id}).then(rsp => {
                        if (rsp.data?.code == 0) {
                            // mutate(rsp)
                            onConfirm?.(rsp?.data?.data)
                        }
                    })
                }
            },
        ])
    }


    return <View>
        <RneButton size={"sm"} onPress={onFollow}>跟单</RneButton>
    </View>
}


export default FollowTrigger