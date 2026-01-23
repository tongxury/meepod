import {Pressable, View} from "react-native";
import React, {useContext, useState} from "react";
import {Button, Dialog, Portal, RadioButton, Text} from "react-native-paper";
import {AppContext} from "../../../providers/global";

const PayTrigger = ({onConfirm, children}: { onConfirm?: () => void, children: React.ReactNode }) => {

    const [visible, setVisible] = useState(false);

    // const [id, setId] = useState('1');

    const {settingsState: {settings: appSettings}} = useContext<any>(AppContext);

    return <Pressable onPress={() => setVisible(true)}>
        {children}
        <Portal>
            <Dialog style={{borderRadius: 10}} visible={visible} onDismiss={() => setVisible(false)}>
                {/*<Dialog.Title></Dialog.Title>*/}
                <Dialog.Content>
                    <Text>确认结账吗</Text>
                </Dialog.Content>
                <Dialog.Actions>
                    <Button onPress={() => setVisible(false)}>取消</Button>
                    <Button onPress={() => {
                        setVisible(false);
                        onConfirm?.()
                    }}>确定</Button>
                </Dialog.Actions>
            </Dialog>
        </Portal>

    </Pressable>
}

export default PayTrigger