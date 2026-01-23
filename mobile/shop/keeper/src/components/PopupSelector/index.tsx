import {Modal} from "@ant-design/react-native";
import {Pressable, Text, View} from "react-native";
import React, {useState} from "react";
import ButtonSelector from "../ButtonSelector";

export declare type Option = {
    value: string
    label: string
}
const PopupSelector = ({options, value, onChange, mode, children}: {
    options: Option[],
    value: string,
    onChange: (value: string) => void
    mode?: 'slide-down' | 'slide-up',
    children: React.ReactNode
}) => {

    const [open, setOpen] = useState<boolean>()

    const onChange_ = (value: string) => {

        console.log('onChange_', value)

        onChange?.(value)
        setTimeout(() => {
            setOpen(false)
        }, 200)

    }

    return <View>
        {/*<Pressable onPress={() => setOpen(true)}>*/}
        {/*    {children}*/}
        {/*</Pressable>*/}
        {/*<Modal*/}
        {/*    popup*/}
        {/*    style={{borderBottomLeftRadius: 10, borderBottomRightRadius: 10}}*/}
        {/*    maskClosable={true}*/}
        {/*    visible={open}*/}
        {/*    animationType={mode ?? 'slide-down'}*/}
        {/*    onClose={() => setOpen(false)}*/}
        {/*>*/}
        {/*    <View style={{padding: 15}}>*/}
        {/*        <ButtonSelector value={value} options={options} onChange={onChange_}/>*/}
        {/*    </View>*/}
        {/*</Modal>*/}
    </View>

}

export default PopupSelector