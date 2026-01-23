import {FlatList, Platform, View, ViewStyle} from "react-native";
import {Modal, Toast, WhiteSpace, WingBlank} from "@ant-design/react-native";
import {Button, Text, TextInput, useTheme} from "react-native-paper";
import React, {useContext, useEffect, useState} from "react";
import {useNavigation} from "@react-navigation/native";
import {useListStorageStateV2} from "../../../hooks/storage";
import {AppContext} from "../../../providers/global";
import {addPlan, fetchCurrentIssue} from "../../../service/api";
import {useRequest} from "ahooks";
import CountDownTimer from "../../../components/CountDownTimer";
import {HStack, Stack} from "@react-native-material/core";
import {Stepper} from "../../../components/Stepper";
import {Ticket} from "../types";
import {mainBodyHeight} from "../../../utils/dimensions";
import {Chip} from "@rneui/themed";
import IconMci from "react-native-vector-icons/MaterialCommunityIcons";
import {Empty} from "../../../components/ListComponent";
import {TicketView} from "../View/TicketView";
import SelectorTrigger from "./SelectorTrigger";
import {OrderSubmitterTrigger} from "../../../triggers/Order/Submit";

const PlanEditor = ({itemId}: { itemId: string, }) => {

    const navigation = useNavigation()


    const {data: issueResult, loading, runAsync} = useRequest(fetchCurrentIssue, {manual: true})
    const issue = issueResult?.data?.data

    useEffect(() => {
        runAsync({itemId}).then(rsp => {
        })
    }, [])

    const {value: plans, replaceAll} = useListStorageStateV2<Ticket>(`${itemId}_plans`,)
    const onAdd = (newPlan: Ticket) => {
        if (isDuplicate(newPlan)) {
            Toast.info('方案重复, 请检查并重新选择')
        } else {
            const newPlans = plans?.concat(newPlan) ?? [newPlan]
            replaceAll(newPlans)
        }
    }

    const onReset = () => {
        replaceAll([])
    }

    const onRemove = async (plan: Ticket) => {
        const newPlans = plans.filter(t => t.key !== plan.key) ?? []
        replaceAll(newPlans)
    }

    const isDuplicate = (plan: Ticket) => {
        for (let i = 0; i < plans?.length; i++) {
            if (plans[i].key === plan.key) {
                return true
            }
        }
        return false
    }


    const totalAmount = plans?.length > 0 ? plans?.map(t => t.amount ?? 0).reduce((a, c) => a + c) : 0

    const {settingsState: {settings}} = useContext<any>(AppContext);

    const [multiple, setMultiple] = useState<number>(1)

    const disabled = !(plans && plans.length)

    const onSave = () => {
        Modal.alert('', '确认保存当前方案吗', [
            {text: '取消', onPress: undefined, style: 'cancel'},
            {
                text: '确认', onPress: () => {
                    addPlan({itemId, content: JSON.stringify(plans), multiple}).then(rsp => {
                        if (rsp?.data?.data) {
                            replaceAll([])
                            setMultiple(multiple)
                        }
                    })
                }
            },
        ])
    }

    const {colors} = useTheme()


    if (loading) {
        return <Stack></Stack>
    }

    if (issue?.index && issue?.close_time_left == 0) {
        return <Stack fill={true} center={true} bg={colors.background}>
            <Text variant="titleLarge" style={{color: colors.primary}}>{issue?.index}</Text>
            <WhiteSpace size="lg"/>
            <Text variant="titleMedium">本期已截止</Text>
            <WhiteSpace/>
            <View style={{alignItems: "center"}}>
                <View style={{flexDirection: "row", alignItems: "center"}}>
                    <Text>距离开奖</Text>
                    <WingBlank size="sm"/>
                    <CountDownTimer color={colors.primary} timeLeftSecond={issue?.prize_time_left}/>
                </View>
            </View>
        </Stack>
    }

    return <Stack {...Platform.OS == 'web' ? {h: mainBodyHeight} : {fill: 1}} bg={colors.background} p={10}
                  spacing={10}>
        {
            issue?.index && <HStack items="center" justify="between" p={10}>
                <Text><Text style={{fontWeight: "bold", color: colors.primary}}>{issue?.index}</Text> 期</Text>
                {
                    issue?.close_time_left > 0 &&
                    <HStack spacing={5}>
                        <Text>距离截止</Text>
                        <CountDownTimer color={colors.primary} timeLeftSecond={issue?.close_time_left}/>
                    </HStack>
                }
            </HStack>
        }

        <Stack style={{flex: 1}}>
            <HStack items="center" justify={"between"} spacing={10}>
                <View style={{flex: 1}}>
                    <SelectorTrigger itemId={itemId} src={issue?.extra} onConfirm={onAdd} onClear={onReset}/>
                </View>
                <Chip buttonStyle={{padding: 5}} color={colors.primary}
                      icon={<IconMci name="link" color={colors.onPrimary}/>}
                    // @ts-ignore
                      onPress={() => navigation.navigate('Root', {screen: 'History'})}>选号记录</Chip>
            </HStack>
            <WhiteSpace/>
            <FlatList
                style={{flex: 1}}
                data={plans}
                numColumns={1}
                // onEndReachedThreshold={400}
                ItemSeparatorComponent={() => <View style={{height: 10}}/>}
                ListEmptyComponent={<Empty/>}
                renderItem={({item: x, index: i}) =>
                    <HStack items="center" justify="between" fill={1}>
                        <TicketView itemId={itemId} data={x} style={{flex: 1}}/>
                        <IconMci name="delete" onPress={() => onRemove(x)} style={{marginHorizontal: 0}} size={20}/>
                    </HStack>
                }
                keyExtractor={(x, i) => i.toString()}
            />
        </Stack>
        <Stack spacing={5}>
            <HStack justify="between" items="center">
                <Text>共<Text>{totalAmount * multiple}</Text>元</Text>
                <Stepper disabled={disabled} color={colors.primary}
                         value={multiple} min={1}
                         max={settings.maxMultiple}
                         onChange={setMultiple}/>
            </HStack>
            <View style={{flexDirection: "row"}}>
                <Button mode="contained-tonal" onPress={onSave} disabled={disabled}>
                    保存方案
                </Button>
                <WingBlank size="sm"/>
                <OrderSubmitterTrigger
                    disabled={disabled} style={{flex: 1}}
                    data={{itemId, plans, multiple, amount: totalAmount}}
                    onSubmitted={orderId => {
                        replaceAll([])
                    }}>
                    <Button mode="contained" style={{flex: 1}} disabled={disabled}>
                        提交到店
                    </Button>
                </OrderSubmitterTrigger>
            </View>
        </Stack>
    </Stack>
}

export default PlanEditor