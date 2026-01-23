import {Pressable, View} from "react-native";
import React, {useState} from "react";
import {Modal} from "@ant-design/react-native";
import {StyleProp, ViewStyle} from "react-native";
import {Text, useTheme} from "react-native-paper";
import {Stack} from "@react-native-material/core";
import Constants from "expo-constants";
import * as Updates from "expo-updates";
import {Config} from "../../../config";


const DebugTrigger = ({style}: { style?: StyleProp<ViewStyle> }) => {

    const [open, setOpen] = useState<boolean>(false)

    const {colors} = useTheme()


    return <View style={style}>
        <Pressable style={{width: 5, height: 5}} onLongPress={() => setOpen(true)}/>
        <Modal
            style={{borderTopLeftRadius: 10, borderTopRightRadius: 10}}
            popup
            visible={open}
            maskClosable={true}
            animationType="slide-up"
            onClose={() => {
                setOpen(false)
            }}>
            <Stack spacing={8} p={15}>
                <Text>extra: {JSON.stringify(Constants.expoConfig.extra)}</Text>
                <Text>channel: {Updates.channel}</Text>
                <Text>config: {JSON.stringify(Config())}</Text>
            </Stack>
        </Modal>
    </View>
}

export default DebugTrigger