import React, {useContext, useState} from "react";
import {AppContext} from "../../../providers/global";
import {Modal} from "@ant-design/react-native";
import {Pressable, View} from "react-native";
import Icon from "react-native-vector-icons/FontAwesome";
import {Button as RneButton} from "@rneui/themed";
import {useTheme} from "react-native-paper";
import {Ticket} from "../types";
import SsqSelectorView from "./Ssq";
import X7cSelectorView from "./X7c";
import F3dSelectorView from "./F3d/SelectorView";
import Z14SelectorView from "./Zc/Z14SelectorView";
import ZjcSelectorView from "./Zc/ZjcSelectorView";
import Pl5SelectorView from "./Pl5/SelectorView";
import DltSelectorView from "./Dlt";

const SelectorTrigger = ({itemId, src, onConfirm, onClear}: {
    itemId: string,
    src?: any,
    onConfirm: (plan: Ticket) => void,
    onClear: () => void,
}) => {

    const [open, setOpen] = useState<boolean>(false)

    const {settingsState: {settings}} = useContext<any>(AppContext);

    const confirm = (ticket: Ticket) => {
        setOpen(false)
        onConfirm(ticket)
    }

    const clear = () => {
        setOpen(false)
        onClear()
    }

    const {colors} = useTheme()
    return <View style={{flex: 1}}>
        {/*<Pressable onPress={() => setOpen(true)}>{children}</Pressable>*/}
        <RneButton
            onPress={() => setOpen(true)}
            icon={<Icon name="plus" color={colors.onPrimary}/>}
            size="sm">添加一注</RneButton>
        <Modal
            popup
            visible={open}
            maskClosable
            animationType="slide-up"
            bodyStyle={{padding: 10, paddingTop: 20}}
            onClose={() => setOpen(false)}>

            {itemId == 'ssq' && <SsqSelectorView onConfirm={confirm}/>}
            {itemId == 'x7c' && <X7cSelectorView onConfirm={confirm}/>}
            {itemId == 'f3d' && <F3dSelectorView onConfirm={confirm}/>}
            {itemId === 'rx9' && <Z14SelectorView min={9} src={src} onConfirm={confirm} onClear={clear}/>}
            {itemId === 'sfc' && <Z14SelectorView min={14} src={src} onConfirm={confirm} onClear={clear}/>}
            {itemId === 'zjc' && <ZjcSelectorView src={src} onConfirm={confirm} onClear={clear}/>}
            {itemId === 'pl3' && <F3dSelectorView onConfirm={confirm}/>}
            {itemId === 'pl5' && <Pl5SelectorView onConfirm={confirm}/>}
            {itemId === 'dlt' && <DltSelectorView onConfirm={confirm}/>}
        </Modal>
    </View>

}

export default SelectorTrigger