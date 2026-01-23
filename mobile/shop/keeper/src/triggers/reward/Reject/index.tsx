import {Pressable, StyleProp, View, ViewStyle} from "react-native";
import React, {useContext, useState} from "react";
import {Button, Dialog, Portal, RadioButton, Text, useTheme} from "react-native-paper";
import {AppContext} from "../../../providers/global";
import {Button as RneButton} from "@rneui/themed";

const RejectTrigger = ({onConfirm, style}: {
    onConfirm?: (id) => void,
    style?: StyleProp<ViewStyle> | undefined
}) => {

    const [visible, setVisible] = useState(false);

    const [id, setId] = useState('1');

    const {settingsState: {settings: appSettings}} = useContext<any>(AppContext);

    const {colors} = useTheme()
    return <View style={style}>
        <RneButton size={'sm'} color={colors.error} onPress={() => setVisible(true)}>拒</RneButton>
        <Portal>
            <Dialog style={{borderRadius: 10}} visible={visible} onDismiss={() => setVisible(false)}>
                <Dialog.Title>拒单原因</Dialog.Title>
                <Dialog.Content>
                    <RadioButton.Group onValueChange={newValue => setId(newValue)} value={id}>
                        {
                            appSettings.rejectReasons.map(t => <View key={t.id}
                                                                     style={{
                                                                         flexDirection: "row",
                                                                         alignItems: "center"
                                                                     }}>
                                <RadioButton value={t.id}/>
                                <Text>{t.text}</Text>
                            </View>)
                        }
                    </RadioButton.Group>
                </Dialog.Content>
                <Dialog.Actions>
                    <Button onPress={() => setVisible(false)}>取消</Button>
                    <Button onPress={() => {
                        setVisible(false);
                        onConfirm?.(id)
                    }}>确定</Button>
                </Dialog.Actions>
            </Dialog>
        </Portal>

    </View>
}

export default RejectTrigger