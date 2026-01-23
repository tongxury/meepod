import {Linking, StyleProp, View, ViewStyle} from "react-native";
import React, {useState} from "react";
import {Badge, Button, Dialog} from "@rneui/themed";
import {fetchTopup} from "../../service/api";
import {Text, useTheme} from "react-native-paper";
import {Modal, Toast} from "@ant-design/react-native";
import {HStack, Stack} from "@react-native-material/core";
import SvgQRCode from "react-native-qrcode-svg";
import {Topup} from "../../service/typs";


export const PayTrigger = ({topup, onConfirmed, onClose, style}: {
    topup: Topup,
    onConfirmed?: () => void,
    onClose: () => void,
    style?: StyleProp<ViewStyle>
}) => {

    const [paying, setPaying] = useState<boolean>(false)
    const onPay = (method: string) => {

        let prefix = 'alipays://platformapi/startapp?saId=10000007&qrcode='
        if (method === 'wechat') {
            prefix = 'weixin://scanqrcode?url='
            // prefix = 'weixin://dl/scan'
            // prefix = 'weixin://wxpay/bizpayurl?pr='
        }

        Toast.info(prefix+encodeURIComponent(topup.qr_code))



        Linking.openURL(prefix+encodeURIComponent(topup.qr_code)).then(rsp => {
            onPaying()

        }).finally(() => {

        })
    }

    const onPaying = () => {

        setPaying(true)

        const job = setInterval(() => {
            fetchTopup({id: topup.id}).then(rsp => {
                if (rsp.data?.data?.payed) {
                    onClose()
                    setPaying(false)
                    onConfirmed?.()
                }
            })
        }, 3000)

        setTimeout(() => {
            clearInterval(job)
            onClose()
            setPaying(false)
        }, 3000 * 10)

    }

    const onStop = () => {
        onClose()
        setPaying(false)
    }

    const {colors} = useTheme()

    return <View style={style}>
        <Modal style={{borderTopLeftRadius: 10, borderTopRightRadius: 10}}
               popup
               visible={!!topup?.qr_code}
               maskClosable={false}
               bodyStyle={{padding: 20}}
               animationType="slide-up"
               onClose={onClose}>
            <Stack spacing={30}>
                <Stack items={'center'} spacing={30}>
                    <Text variant={'titleMedium'}>支付中...(支付成功后自动关闭)</Text>
                    <View>{topup?.qr_code && <SvgQRCode size={250} value={topup?.qr_code}/>}</View>
                    <Text
                        variant={'bodySmall'}>如果无法正常跳转，请自行保存二维码截图后打开微信或者支付宝识别并完成付款</Text>
                    {/*<HStack fill={1} justify={'end'}><Text onPress={onClose}*/}
                    {/*                                       style={{color: colors.primary}}>稍后查看</Text></HStack>*/}
                </Stack>
                <Stack spacing={10}>
                    <View style={{flex: 2}}>
                        <Button loading={paying} onPress={() => onPay('alipay')} color={'#02a6e7'}>
                            {paying ? '支付中...': '支付宝支付'}
                        </Button>
                        <Badge
                            value={"推荐"}
                            containerStyle={{position: 'absolute', top: -5, right: -5}}
                        />
                    </View>
                    <Button loading={paying}  onPress={() => onPay('wechat')} color={'#31a606'}>
                        {paying ? '支付中...': '微信支付'}
                    </Button>
                    <HStack justify={'end'}>
                        <Text onPress={onStop}>放弃</Text>
                    </HStack>
                </Stack>
            </Stack>

        </Modal>
    </View>
}
