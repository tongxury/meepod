import {ScrollView, StyleProp, View, ViewStyle} from "react-native";
import {HStack, Stack} from "@react-native-material/core";
import {Banner, Text} from "react-native-paper";
import {useState} from "react";
import Markdown from 'react-native-markdown-display';
import {Modal} from "@ant-design/react-native";
import Ionicons from "react-native-vector-icons/Ionicons";
import {mainBodyHeight} from "../../utils/dimensions";

const TipView = ({style}: { style?: StyleProp<ViewStyle> }) => {

    const [open, setOpen] = useState<boolean>(false)


    const tip = `
**开通店铺合作后需满一个月才可以更换合作店铺，且需要双方账本清空后才可停止合作！**
**关于合作彩种的出票截止时间与最小接单金额，请与合作店铺咨询！**

**普通版只支持添加1家合作店铺，如要更换合作店铺，需先暂停与当前店铺的合作，再添加新的合作店铺**
**如需同时与多家店铺合作，请联系客服开通高级版**

一、什么是合作转单
用户方【A】与服务方【B】基于双方在用户及服务资源上的各自优势，通过本平台提供的合作功能，达成合作意向，共同为用户提供方案代销的服务。

二、合作转单的意义
互通有无、合作共赢。
- 充分发挥双方优势，延长用户方【A】代销时间，丰富其店内代销产品种类；
- 提升服务方【B】销量的同时，使双方达到共同创收、合作共赢的目的。

三、店铺间如何建立合作
用户方【A】通过本平台申请开通店铺合作转单功能；
系统自动为用户方【A】匹配合作店铺服务方【B】
双方合作关系建立成功

四、转单流程
- 用户方【A】根据每日预估销量**提前在服务方【B】做好订单款项预存**，即系统内转单额度预存；
- 用户方【B】需在系统内**提前预存合作转单服务点**，以备后续流程正常操作；
- 订单到达，用户方【A】确认订单收款后，在系统内选择【转单】，将订单转给服务方【B】；
- 服务方【B】在接到用户方【A】提交方案后，进行出票并提供后续服务, 同时向【A】支付相应的转单佣金。
    `

    return <View style={style}>
        {/*<Text onPress={() => setOpen(!open)}>合作说明</Text>*/}
        <Ionicons name="information-circle-outline" onPress={() => setOpen(!open)} size={15}/>
        <Modal
            style={{borderTopLeftRadius: 10, borderTopRightRadius: 10}}
            popup
            visible={open}
            bodyStyle={{padding: 10, height: mainBodyHeight - 200}}
            maskClosable={true}
            animationType="slide-up"
            onClose={() => {
                setOpen(false)
            }}>
            <ScrollView style={{flex: 1}}>
                {/* @ts-ignore*/}
                <Markdown>
                    {tip}
                </Markdown>
            </ScrollView>

        </Modal>
    </View>

};

export default TipView
