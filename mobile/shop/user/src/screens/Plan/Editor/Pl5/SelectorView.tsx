import {Dimensions, ScrollView, TouchableWithoutFeedback, View} from "react-native";
import {Appbar, Avatar, Badge, Button, Text, IconButton} from "react-native-paper";
import {Flex, Toast, WhiteSpace, WingBlank} from "@ant-design/react-native";
import {useContext, useEffect, useState} from "react";
import {DoubleSelector, Selector} from "../../../../components/Selector";
import useHook from "./hook";
import {Ticket} from "../../types";
import {createMd5} from "../../../../utils";
import {HStack, Stack} from "@react-native-material/core";


const SelectorView = ({onConfirm}: {
    onConfirm: (ticket: Ticket) => void,
}) => {

    const {options, getRandom} = useHook()

    const [ticket, setTicket] = useState<Ticket>({amount: 0, key: ''})

    const onSelectedChange = (pos, values) => {

        console.log(pos, values)

        const newValues = {...ticket}
        newValues[pos] = values

        newValues.key = newValues.gen?.join('') + newValues.shi?.join('') + newValues.bai?.join('') +
            newValues.qian?.join('') + newValues.wan?.join('')

        newValues.amount = newValues.gen?.length *
            newValues.shi?.length *
            newValues.bai?.length *
            newValues.qian?.length *
            newValues.wan?.length * 2

        setTicket(newValues)
    }


    const onRandom = () => {

        const wan = getRandom()
        const qian = getRandom()
        const bai = getRandom()
        const shi = getRandom()
        const gen = getRandom()

        setTicket({
            wan,
            qian,
            bai,
            shi,
            gen,
            amount: 2,
            key: wan?.join('') + qian?.join('') + bai?.join('') + shi?.join('') + gen?.join(''),
        })
    }
    const onClear = () => {
        setTicket({amount: 0, key: ''})
    }

    const confirm = () => {
        if (validPlan) {
            onConfirm(ticket)
        }
    }


    const validPlan =
        ticket.wan?.length > 0 &&
        ticket.qian?.length > 0 &&
        ticket.bai?.length > 0 &&
        ticket.qian?.length > 0 &&
        ticket.wan?.length > 0


    const PosView = ({title, values, onChange}) => {
        const change = (item, selected) => {

            console.log('xxxxxxx', item, selected, values, values?.filter(t => t !== item.value))

            if (selected) {
                onChange((values ?? []).concat(item.value))
            } else {
                onChange(values?.filter(t => t !== item.value))
            }
        }


        return <HStack items={"center"}>
            <Text variant="titleSmall">{title}</Text>
            <WingBlank size="sm"/>
            <View style={{flex: 1}}>
                <Selector
                    itemStyle={{selected: "#f53a3a", unselected: '#faa199'}}
                    items={options}
                    selectedKeys={values}
                    onSelectChanged={change}
                />
                <View style={{flexDirection: "row", alignItems: "center"}}>
                    <Button onPress={() => onChange(options.map(t => t.value))}>全</Button>
                    <Button onPress={() => onChange(options.map(t => t.value)
                        .filter(t => parseInt(t) % 2 !== 0))}>奇</Button>
                    <Button onPress={() => onChange(options.map(t => t.value)
                        .filter(t => parseInt(t) % 2 === 0))}>偶</Button>
                    <Button onPress={() => onChange([])}>清</Button>
                </View>
            </View>
        </HStack>
    }


    return <View>

        <Stack spacing={10}>
            <PosView title={'第一位'} values={ticket['wan'] ?? []}
                     onChange={(values) => onSelectedChange('wan', values)}/>
            <PosView title={'第二位'} values={ticket['qian'] ?? []}
                     onChange={(values) => onSelectedChange('qian', values)}/>
            <PosView title={'第三位'} values={ticket['bai'] ?? []}
                     onChange={(values) => onSelectedChange('bai', values)}/>
            <PosView title={'第四位'} values={ticket['shi'] ?? []}
                     onChange={(values) => onSelectedChange('shi', values)}/>
            <PosView title={'第五位'} values={ticket['gen'] ?? []}
                     onChange={(values) => onSelectedChange('gen', values)}/>

        </Stack>

        <View style={{
            flexDirection: "row", alignItems: "center",
            flex: 1
        }}>
            <View style={{flex: 1, flexDirection: "row", alignItems: "center", justifyContent: "flex-start"}}>
                <Button mode="text" style={{padding: 3}} onPress={onRandom}>机选</Button>
                <Button onPress={onClear}>清除</Button>
            </View>
            <View style={{flex: 1}}>
                <Button disabled={!validPlan}
                        onPress={confirm}
                        mode="contained"
                >
                    金额 {ticket?.amount} 元
                </Button>
            </View>
        </View>
    </View>

}

export default SelectorView