import { Modal, Toast } from "@ant-design/react-native";
import React, { useState } from "react";
import { Linking, Pressable, StyleProp, View, ViewStyle } from "react-native";
import { Button, Chip, Text, TextInput, useTheme } from "react-native-paper";
import { styles } from "../../utils/styles";
import { HStack, Stack } from "@react-native-material/core";
import { addTopUpOrder, cancelTopup, fetchTopup, pay } from "../../service/api";
import { PaymentOrder, Topup } from "../../service/typs";
import { Badge, Button as RneButton, Dialog, } from '@rneui/themed';
import SvgQRCode from 'react-native-qrcode-svg';
import { PayTrigger } from "../Pay";


export const TopUpBuyingTrigger = ({ orderId, onConfirmed }: {
    orderId: string,
    onConfirmed?: () => void,
}) => {

    const [topup, setTopup] = useState<Topup>()
    const [paying, setPaying] = useState<boolean>(false)
    const onConfirm = () => {

        setPaying(true)

        pay({ orderId }).then(rsp => {
            if (!rsp.data?.data) {
                onConfirmed?.()
            } else {
                if (rsp.data?.data?.payed) {
                    onConfirmed?.()
                    return
                }
                if (rsp.data?.data?.qr_code) {
                    setTopup(rsp.data?.data)
                    // const payUrl = prefix + encodeURI(rsp.data?.data?.qr_code)
                    //
                    // Linking.openURL(payUrl).then(() => {
                    //     onPayingByUrl(rsp.data?.data?.id)
                    // }).finally(() => {
                    //     setOpen(false)
                    // })
                } else {
                    Toast.info("支付错误，请重试")
                }
            }


        }).finally(() => {
            setPaying(false)
        })
    }

    return <View style={{ flex: 1 }}>
        <RneButton loading={paying} color={'warning'} onPress={onConfirm} style={{ flex: 1 }}
            type={"solid"}>付款给店家</RneButton>
        <PayTrigger topup={topup} onClose={() => setTopup(undefined)} onConfirmed={onConfirmed} />
    </View>

}

export const TopUpTriggerV2 = ({ onConfirmed, children }: {
    onConfirmed?: (amount: number) => void,
    children: React.ReactNode,
}) => {

    const [open, setOpen] = useState<boolean>(false)
    const [topup, setTopup] = useState<Topup>()

    const { colors } = useTheme()
    const [amount, setAmount] = useState<number>()

    const onChange = (text: string) => {
        setAmount(isNaN(parseInt(text)) ? 0 : parseInt(text))
    }

    const onConfirm = () => {

        // let prefix = 'alipays://platformapi/startapp?saId=10000007&qrcode='
        // if (method === 'wechat') {
        //     prefix = 'weixin://scanqrcode?url='
        // }

        let req = addTopUpOrder({ amount })
        // if (orderId) {
        //     req = pay({orderId, method})
        // }

        req.then(rsp => {
            if (rsp.data?.data?.payed) {
                onConfirmed?.(amount!!)
                setOpen(false)
                return
            }

            if (rsp.data?.data?.qr_code) {
                setTopup(rsp.data?.data)
                setOpen(false)
                // const payUrl = prefix + encodeURI(rsp.data?.data?.qr_code)
                //
                // Linking.openURL(payUrl).then(() => {
                //     onPayingByUrl(rsp.data?.data?.id)
                // }).finally(() => {
                //     setOpen(false)
                // })
            } else {
                Toast.info("支付错误，请重试")
                setOpen(false)
            }
        })
    }


    const onClose = () => {
        setOpen(false)
        setAmount(0)
    }

    return <Pressable onPress={() => setOpen(true)}>
        {children}
        <PayTrigger topup={topup} onClose={() => setTopup(undefined)} />
        <Modal
            popup
            style={{ ...styles.popup }}
            visible={open}
            maskClosable={true}
            animationType="slide-up"
            onClose={onClose}
        >
            <Stack spacing={30}>
                <Stack spacing={10}>
                    <Text variant="titleMedium" style={{ fontWeight: "bold", textAlign: "center" }}>充值</Text>
                    <TextInput value={amount?.toString() ?? ''} placeholder="请输入金额" dense mode="outlined"
                        outlineStyle={{ borderRadius: 10 }} onChangeText={onChange}></TextInput>
                    <HStack items="center" spacing={8}>
                        {[1000, 500, 200, 100].map(t =>
                            <Chip key={t} onPress={() => onChange(`${t}`)}>{t}</Chip>
                        )}
                    </HStack>
                </Stack>

                <HStack items="center" spacing={10}>
                    <Button mode={'contained'} disabled={!amount} style={{ flex: 1 }}
                        onPress={onConfirm}>确认</Button>
                </HStack>
            </Stack>
        </Modal>
    </Pressable>

}


export const TopUpTrigger = ({ onConfirmed, children }: {
    onConfirmed?: (amount: number) => void,
    children: React.ReactNode,
}) => {

    const [open, setOpen] = useState<boolean>(false)
    const [qrcode, setQrcode] = useState<string>()

    const { colors } = useTheme()
    const [amount, setAmount] = useState<number>()

    const onChange = (text: string) => {
        setAmount(isNaN(parseInt(text)) ? 0 : parseInt(text))
    }

    const onConfirm = (method: string) => {

        let prefix = 'alipays://platformapi/startapp?saId=10000007&qrcode='
        if (method === 'wechat') {
            prefix = 'weixin://scanqrcode?url='
        }

        let req = addTopUpOrder({ amount })
        // if (orderId) {
        //     req = pay({orderId, method})
        // }

        req.then(rsp => {
            if (rsp.data?.data?.payed) {
                onConfirmed?.(amount!!)
                setOpen(false)
                return
            }

            if (rsp.data?.data?.qr_code) {
                setQrcode(rsp.data?.data?.qr_code)

                const payUrl = prefix + encodeURI(rsp.data?.data?.qr_code)

                Linking.openURL(payUrl).then(() => {
                    onPayingByUrl(rsp.data?.data?.id)
                }).finally(() => {
                    setOpen(false)
                })
            } else {
                Toast.info("支付错误，请重试")
                setOpen(false)
            }
        })
    }

    const onPayingByUrl = (id: string) => {


        const job = setInterval(() => {
            fetchTopup({ id }).then(rsp => {
                if (rsp.data?.data?.payed) {
                    setQrcode('')
                    onConfirmed?.(amount)
                }
            })
        }, 3000)

        setTimeout(() => {
            clearInterval(job)
            setQrcode('')
        }, 3000 * 20)

    }

    const onClose = () => {
        setOpen(false)
        setAmount(0)
    }

    return <Pressable onPress={() => setOpen(true)}>
        {children}
        {/*<Dialog isVisible={paying} onBackdropPress={() => {*/}
        {/*}}>*/}
        {/*    <Dialog.Title title={"支付中..."}/>*/}
        {/*    <Dialog.Loading/>*/}
        {/*    <Dialog.Actions>*/}
        {/*        <Text onPress={() => setPaying(false)}>不想等了</Text>*/}
        {/*    </Dialog.Actions>*/}
        {/*</Dialog>*/}
        <Modal style={{ borderTopLeftRadius: 10, borderTopRightRadius: 10 }}
            popup
            visible={!!qrcode}
            maskClosable={false}
            bodyStyle={{ padding: 20 }}
            animationType="slide-up"
            onClose={() => {
                setOpen(false)
            }}>
            <Stack items={'center'} spacing={30}>
                <Text variant={'titleMedium'}>支付中...(支付成功后自动关闭)</Text>
                <View>{qrcode && <SvgQRCode size={200} value={qrcode} />}</View>
                <Text
                    variant={'bodySmall'}>如果无法正常跳转，请自行保存二维码截图后打开微信或者支付宝识别并完成付款</Text>
                <HStack fill={1} justify={'end'}><Text onPress={() => setQrcode('')}
                    style={{ color: colors.primary }}>稍后查看</Text></HStack>
            </Stack>
        </Modal>
        <Modal
            popup
            style={{ ...styles.popup }}
            visible={open}
            maskClosable={true}
            animationType="slide-up"
            onClose={onClose}
        >
            <Stack spacing={30}>
                <Stack spacing={10}>
                    <Text variant="titleMedium" style={{ fontWeight: "bold", textAlign: "center" }}>充值</Text>
                    <TextInput value={amount?.toString() ?? ''} placeholder="请输入金额" dense mode="outlined"
                        outlineStyle={{ borderRadius: 10 }} onChangeText={onChange}></TextInput>
                    <HStack items="center" spacing={8}>
                        {[1000, 500, 200, 100].map(t =>
                            <Chip key={t} onPress={() => onChange(`${t}`)}>{t}</Chip>
                        )}
                    </HStack>
                </Stack>

                <HStack items="center" spacing={10}>
                    <Button textColor={'#31a606'} disabled={!amount} style={{ flex: 1 }}
                        onPress={() => onConfirm('wechat')}>微信支付</Button>
                    <View style={{ flex: 2 }}>
                        <Button mode="contained" buttonColor={'#02a6e7'} disabled={!amount}
                            onPress={() => onConfirm('alipay')}>支付宝支付</Button>
                        <Badge
                            value={"推荐"}
                            containerStyle={{ position: 'absolute', top: -5, right: -5 }}
                        />
                    </View>
                </HStack>
            </Stack>
        </Modal>
    </Pressable>

}


export const CancelTopupTrigger = ({ id, onConfirm, children, style }: {
    id: string,
    onConfirm?: (newValue: PaymentOrder) => void,
    children: React.ReactNode
    style?: StyleProp<ViewStyle>
}) => {

    const onDelete = () => {
        Modal.alert('', '确认撤销当前订单吗', [
            { text: '取消', onPress: undefined, style: 'cancel' },
            {
                text: '确认', onPress: () => {
                    cancelTopup({ id }).then(rsp => {
                        onConfirm?.(rsp?.data?.data)
                    })
                }
            },
        ])
    }


    return <RneButton style={style} onPress={onDelete} color={'error'} size={'sm'}>撤销</RneButton>
}
