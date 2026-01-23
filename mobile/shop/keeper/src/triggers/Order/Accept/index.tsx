import {Pressable, StyleProp, View, ViewStyle} from "react-native";
import React, {useContext, useState} from "react";
import {Button, Dialog, Portal, RadioButton, Text} from "react-native-paper";
import {AppContext} from "../../../providers/global";
import {Button as RneButton} from "@rneui/themed";

const AcceptTrigger = ({onConfirm, style}: { onConfirm?: () => void, style?: StyleProp<ViewStyle> | undefined }) => {

    const [visible, setVisible] = useState(false);

    // const [id, setId] = useState('1');

    const {settingsState: {settings: appSettings}} = useContext<any>(AppContext);

    return <View style={style}>
        <RneButton size={'sm'}  onPress={() => setVisible(true)}>接单</RneButton>
        <Portal>
            <Dialog style={{borderRadius: 10}} visible={visible} onDismiss={() => setVisible(false)}>
                {/*<Dialog.Title></Dialog.Title>*/}
                <Dialog.Content>
                    <Text>确认接单吗</Text>
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

    </View>
}

export default AcceptTrigger