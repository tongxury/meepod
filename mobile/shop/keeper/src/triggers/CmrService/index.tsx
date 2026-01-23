import {ActivityIndicator, FlatList, Pressable, StyleProp, View, ViewStyle} from "react-native";
import React, {useContext, useState} from "react";
import {Avatar, Button, Dialog, Portal, RadioButton, Text, TextInput} from "react-native-paper";
import {ImagePicker, InputItem, Modal, TextareaItem, WhiteSpace, WingBlank} from "@ant-design/react-native";
import {Stack} from "@react-native-material/core";
import {AppContext} from "../../providers/global";
import {Image} from "@rneui/themed";


const CmrServiceTrigger = ({children, style}: {
    children: React.ReactNode, style?: StyleProp<ViewStyle>
}) => {

    const [visible, setVisible] = useState(false);

    const {settingsState: {settings}} = useContext<any>(AppContext);

    return <Pressable style={style} onPress={() => setVisible(true)}>
        {children}
        <Modal
            style={{borderTopLeftRadius: 10, borderTopRightRadius: 10}}
            popup
            visible={visible}
            maskClosable={true}
            animationType="slide-up"
            onClose={() => {
                setVisible(false)
            }}>
            <Stack p={20} spacing={30}>
                <Stack spacing={20}>
                    <Image
                        source={{uri: settings?.service?.wechat}}
                        containerStyle={{
                            aspectRatio: 1,
                            width: '100%',
                            flex: 1,
                        }}
                        // PlaceholderContent={<ActivityIndicator style={{flex: 1}}/>}
                    />
                    <Text variant={"titleMedium"}>{settings?.service?.email}</Text>
                </Stack>

                <View style={{flexDirection: "row", alignItems: "center"}}>
                    <Button onPress={() => setVisible(false)} mode="contained-tonal"
                            style={{flex: 2}}>确认</Button>
                </View>
            </Stack>
        </Modal>

    </Pressable>
}

export default CmrServiceTrigger