import {View, StyleSheet, SafeAreaView} from "react-native";
import useLoginState from "../../hooks/auth";
import {useContext, useState} from "react";
import {Button, Card, Text, TextInput, useTheme} from "react-native-paper";
import {HStack, Stack} from "@react-native-material/core";
import {styles} from "../../utils/styles";
import useCountdown from "ahooks/lib/useCountDown";
import {useInterval, useRafInterval} from "ahooks";
import AgreementView from "./Agreement";
import PrivacyPolicy from "./PrivacyPolicy";

function LoginScreen({navigation}) {

    const [phone, setPhone] = useState<string>()
    const [code, setCode] = useState<string>()

    const {doLogin, sendCode} = useLoginState()
    const onLogin = () => {
        if (!(phone && code)) {

        } else {
            doLogin({phone, code}, () => {
                navigation.navigate('Root', {screen: 'Ticket'})
            })
        }
    }


    const [cd, setCd] = useState<number>(0)

    useInterval(
        () => {
            console.log('useInterval')
            if (cd > 0) {
                setCd(cd - 1)
            }
        },
        cd > 0 ? 1000 : undefined, {immediate: false}
    )
    const onSendCode = () => {
        if (!phone) {

        } else {
            sendCode({
                phone, onSuccess: () => {
                    setCd(60)
                }
            })
        }
    }

    const {colors} = useTheme()

    return (
        <View style={{flex: 1}}>
            <Stack fill={true} justify="center" p={20}>
                <Card style={{shadowColor: colors.primary, ...styles.card, paddingVertical: 30, marginBottom: 100,}}>
                    <Stack spacing={50} style={{paddingHorizontal: 20}}>
                        <Text style={{fontWeight: "bold", textAlign: "center"}} variant="titleLarge">欢迎登录</Text>

                        <Stack spacing={10}>
                            <TextInput
                                dense
                                outlineStyle={{borderRadius: 5}}
                                mode="outlined"
                                left={<TextInput.Icon
                                    color={(isTextInputFocused) => isTextInputFocused ? colors.primary : undefined}
                                    icon="phone"/>}
                                label="手机号"
                                underlineStyle={{height: 0}}
                                placeholder="请输入手机号"
                                value={phone}
                                onChangeText={text => setPhone(text)}
                            />

                            <HStack items={'center'} spacing={5}>
                                <TextInput
                                    dense
                                    style={{flex: 1}}
                                    outlineStyle={{borderRadius: 5}}
                                    mode="outlined"
                                    left={
                                        <TextInput.Icon
                                            color={(isTextInputFocused) => isTextInputFocused ? colors.primary : undefined}
                                            icon="account"></TextInput.Icon>}
                                    label="验证码"
                                    underlineStyle={{height: 0}}
                                    placeholder="输入验证码"
                                    value={code}
                                    onChangeText={text => setCode(text)}
                                />
                                <Button style={{marginTop: 6}} mode={"contained"} disabled={!(phone && cd <= 0)}
                                        onPress={onSendCode}>{cd > 0 ? `重发(${cd}秒)` : '获取验证码'}</Button>
                            </HStack>
                            <HStack items={'center'}>
                                <Text style={{fontSize: 10}}>登录即代表您同意 </Text>
                                <AgreementView/>
                                <PrivacyPolicy/>
                            </HStack>
                        </Stack>

                        <Button mode="contained" onPress={onLogin} style={{
                            marginTop: 50
                        }}>
                            登录
                        </Button>
                    </Stack>
                </Card>
            </Stack>
        </View>

    );
}


export default LoginScreen