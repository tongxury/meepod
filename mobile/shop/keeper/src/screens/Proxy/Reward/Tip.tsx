import {StyleProp, View, ViewStyle} from "react-native";
import {HStack, Stack} from "@react-native-material/core";
import {Banner, Text} from "react-native-paper";
import {useState} from "react";
import {Modal} from "@ant-design/react-native";
import Markdown from "react-native-markdown-display";

const TipView = ({style}: { style?: StyleProp<ViewStyle> }) => {

    const [open, setOpen] = useState<boolean>(false)


    const tip = `
    佣金结算以自然月为单位，每月2号系统自动生成上一个自然月的结算账单，请店主在月初核对后统一手动操作结算；
    结算操作后佣金将以余额形式先返到推广员在店内的余额账户，推广员可以自行申请提现。
    请您及时为推广员结算佣金，以免因此造成信誉损失。
    
    `

    return <View style={style}>
        <Text onPress={() => setOpen(!open)}>结算说明</Text>
        <Modal
            style={{borderTopLeftRadius: 10, borderTopRightRadius: 10}}
            popup
            visible={open}
            bodyStyle={{padding: 10}}
            maskClosable={true}
            animationType="slide-up"
            onClose={() => {
                setOpen(false)
            }}>
            {/* @ts-ignore*/}
            <Markdown>
                {tip}
            </Markdown>
        </Modal>
    </View>

};

export default TipView
