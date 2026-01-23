import {Modal} from "@ant-design/react-native";
import {styles} from "../../../utils/styles";
import {HStack, Stack} from "@react-native-material/core";
import {Button, Chip, Text, TextInput} from "react-native-paper";
import React, {useState} from "react";
import {View} from "react-native";
import Tabs from "../../../components/Tabs";
import ShareList from "./ShareList";
import {mainBodyHeight} from "../../../utils/dimensions";


const DetailTrigger = ({orderId}: { orderId: string }) => {

    const [open, setOpen] = useState<boolean>(false)

    const onClose = () => {
        setOpen(false)
    }

    return <View>
        <Button onPress={() => setOpen(true)}>明细</Button>
        <Modal
            popup
            style={{...styles.popup}}
            visible={open}
            maskClosable={true}
            animationType="slide-up"
            bodyStyle={{maxHeight: mainBodyHeight - 200,}}
            onClose={onClose}
        >
            <Tabs style={{flex: 1}} tabs={[
                // {
                //     key: 'group',
                //     title: '选单详情',
                //     component: () => <Stack p={10} bg={colors.background}>
                //         <TicketListView itemId={group?.plan?.item?.id} data={group?.plan?.tickets}/>
                //     </Stack>
                // },
                {key: 'shares', title: '合买订单', component: () => <ShareList groupId={orderId}/>},
            ]}/>
        </Modal>
    </View>

}

export default DetailTrigger