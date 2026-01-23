import {Appbar, Button, Text, useTheme} from "react-native-paper";
import {Toast, WhiteSpace} from "@ant-design/react-native";
import React from "react";
import {View} from "react-native";
import {HStack, Stack} from "@react-native-material/core";
import {useRequest} from "ahooks";
import {fetchProxy} from "../../service/api";
import UserList from "./UserList";
import Tabs from "../../components/Tabs";
import StatsView from "../../components/StatsView";
import AmountView from "../../components/AmountView";


const ProxyScreen = ({navigation}) => {

    const {data, loading, run} = useRequest(fetchProxy)
    const proxy = data?.data?.data

    const {colors} = useTheme()
    const Body = () => {
        if (loading) {
            return <View></View>
        }

        if (!proxy) {
            return <Stack fill={1} center={true} bg={colors.background} m={8} p={20} spacing={20}>

                <Text variant={"bodySmall"} style={{fontSize: 14}}>
                    成为店铺推广员后，帮助推广可获取佣金。
                    佣金结算以自然月为单位，每月2号系统自动生成上一个自然月的结算账单，由店主在月初核对后统一手动操作结算；
                    结算操作后佣金将以余额形式先返到推广员在店内的余额账户，推广员可以自行申请提现。
                </Text>

                <Button mode={"contained"} onPress={() => Toast.info('请与店主联系')}>我要成为推广员</Button>
            </Stack>
        }

        return <Stack fill={1} spacing={10}>

            <Stack p={10} bg={colors.background} spacing={10}>
                <HStack items={"center"} spacing={5}>
                    <Text>邀请码</Text>
                    <Text variant={"titleMedium"} style={{fontWeight: 'bold'}}>{proxy?.id}</Text>
                </HStack>
                <HStack items={"center"} spacing={20}>

                    <AmountView amount={proxy?.reward_amount} size={"large"}/>
                    <Text><Text style={{
                        fontSize: 20,
                        fontWeight: 'bold',
                        color: colors.primary
                    }}>{(proxy?.reward_rate * 100).toFixed(1)}</Text> %</Text>
                </HStack>
                <HStack items={"center"} justify={"between"}>
                    <StatsView style={{flex: 1}} title="用户数(下过单)" value={proxy?.user_count} unit={"人"}/>
                    <StatsView style={{flex: 1}} title="订单数" value={proxy?.order_count} unit={"单"}/>
                    <StatsView style={{flex: 1}} title="订单总额" value={proxy?.order_amount} unit={"元"}/>
                    {/*<StatsView title="佣金比例" value={3000} unit={"%"}/>*/}
                </HStack>
                <Text variant={'bodySmall'}>数据更新会有一定的延迟</Text>
            </Stack>

            <Tabs style={{flex: 1}} tabs={[
                {key: '1', title: '用户列表', component: () => <UserList/>},
            ]}/>
        </Stack>


    }


    return <View style={{flex: 1}}>
        <Appbar.Header>
            <Appbar.BackAction onPress={() => {
                navigation.goBack()
            }}/>
            <Appbar.Content title={<Text variant="titleMedium">合作推广</Text>}/>
        </Appbar.Header>
        <Body/>
    </View>
}

export default ProxyScreen