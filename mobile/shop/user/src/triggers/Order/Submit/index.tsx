import {InputItem, Modal, TextareaItem, WhiteSpace} from "@ant-design/react-native";
import React, {useContext, useState} from "react";
import {AppContext} from "../../../providers/global";
import {useForm, Controller} from "react-hook-form";
import {Button, Checkbox, Text, TextInput} from "react-native-paper";
import {Pressable, StyleProp, View, ViewStyle} from "react-native";
import {addGroupOrder, addOrder, addPlan, pay} from "../../../service/api";
import {HStack, Stack} from "@react-native-material/core";
import {Input} from "@rneui/themed";
import {useNavigation} from "@react-navigation/native";

type OrderSettings = {
    groupBuy: boolean,
    totalVolume: number,
    // volume: number,
    // rewardRate: number,
    // floor: number,
    remark: string,
    needUpload: boolean,
    // confirm: boolean,
}


export const OrderSubmitterTrigger = ({data, onSubmitted, disabled, children, style}: {
    data: { planId?: string, itemId?: string, plans?: any, multiple?: number, amount?: number }
    onSubmitted?: (orderId: string, groupBuy: boolean) => void,
    disabled?: boolean,
    children: React.ReactNode,
    style?: StyleProp<ViewStyle>
}) => {

    const navigation = useNavigation()

    const {settingsState: {settings: appSettings}} = useContext<any>(AppContext);

    // form组件里拿不到某个字段的实时值 新开个state
    const [open, setOpen] = useState(false)
    const [groupBuy, setGroupBuy] = useState<boolean>()
    const [needUpload, setNeedUpload] = useState<boolean>(false)
    const [confirm, setConfirm] = useState<boolean>(true)
    const [totalVolume, setTotalVolume] = useState<number>(100)
    const [remark, setRemark] = useState<string>('众人拾柴火焰高，一起中奖一起分！')


    const onConfirm = () => {


        if (data.planId) {
            onSubmitOrder(data.planId, {needUpload, remark, groupBuy, totalVolume})
        } else {
            addPlan({itemId: data.itemId, content: JSON.stringify(data.plans), multiple: data.multiple}).then(rsp => {
                if (rsp?.data?.data) {
                    onSubmitOrder(rsp?.data?.data, {needUpload, remark, groupBuy, totalVolume})
                }
            })
        }
    }
    const onSubmitOrder = (planId: string, settings: OrderSettings) => {


        if (groupBuy) {
            addGroupOrder({
                planId: planId,
                needUpload: settings.needUpload,
                totalVolume: settings.totalVolume,
                // volume: settings.volume,
                // rewardRate: settings.rewardRate,
                // floor: settings.floor,
                remark: settings.remark,
            }).then(rsp => {

                const orderId = rsp?.data?.data

                if (orderId) {
                    // @ts-ignore
                    navigation.navigate('OrderGroupDetail', {id: orderId})
                    onSubmitted?.(orderId, true)
                    setOpen(false)
                }
            })
        } else {
            addOrder({
                planId: planId,
                needUpload: settings.needUpload,
            }).then(rsp => {

                const orderId = rsp?.data?.data

                if (orderId) {

                    pay({orderId: orderId}).then(rsp => {
                        // if (rsp?.data?.code == 0) {
                        //
                        // }
                        setOpen(false)
                        // @ts-ignore
                        navigation.navigate('OrderDetail', {id: orderId})
                        onSubmitted?.(orderId, false)
                    })
                }
            })
        }

    }

    return <Pressable disabled={disabled} style={style} onPress={() => setOpen(true)}>
        {children}
        <Modal
            popup
            visible={open}
            maskClosable
            animationType="slide-up"
            style={{borderTopLeftRadius: 10, borderTopRightRadius: 10}}
            bodyStyle={{padding: 18}}
            onClose={() => setOpen(false)}>
            <Stack>

                <View style={{flexDirection: "row", alignItems: "center"}}>
                    <Checkbox
                        disabled={data.amount * data.multiple < appSettings.minUnionAmount}
                        status={groupBuy ? 'checked' : 'unchecked'}
                        onPress={(e) => {
                            setGroupBuy(!groupBuy)
                        }}
                    />
                    <Text>是否合买<Text>{`(总金额需要大于${appSettings.minUnionAmount})`}</Text></Text>
                </View>
                {
                    groupBuy && (

                        <Stack spacing={5} pl={8}>
                            <Text>合买留言</Text>
                            <TextInput
                                style={{flex: 1}}
                                underlineStyle={{width: 0,}}
                                contentStyle={{padding: 8}}
                                value={remark}
                                onChangeText={text => setRemark(text)}
                                multiline
                                numberOfLines={4} maxLength={50}
                                placeholder="合买留言"/>
                        </Stack>
                    )
                }

                <View style={{flexDirection: "row", alignItems: "center"}}>
                    <Checkbox
                        disabled={false}
                        status={needUpload ? 'checked' : 'unchecked'}
                        onPress={(e) => setNeedUpload(!needUpload)}
                    />
                    <Text>是否需要上传票样</Text>
                </View>

                <View style={{flexDirection: "row", alignItems: "center"}}>
                    <Checkbox
                        disabled={false}
                        status={confirm ? 'checked' : 'unchecked'}
                        onPress={(e) => setConfirm(!confirm)}
                    />
                    <Text>同意隐私政策</Text>
                </View>
            </Stack>

            <WhiteSpace size="lg"/>
            <Button mode="contained" disabled={!confirm} onPress={onConfirm}>确认</Button>
        </Modal>
    </Pressable>
}

