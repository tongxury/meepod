import {Dimensions, ScrollView, TouchableWithoutFeedback, useWindowDimensions, View} from "react-native";
import {Appbar, Avatar, Badge, Button, Text} from "react-native-paper";
import {Flex, Toast, WhiteSpace, WingBlank} from "@ant-design/react-native";
import {useContext, useEffect, useState} from "react";
import {DoubleSelector, Selector} from "../../../../components/Selector";
import useHook from "./hook";
import {HStack, Stack} from "@react-native-material/core";
import {Chip} from "@rneui/themed";
import {Ticket} from "../../types";
import {createMd5} from "../../../../utils";


const emptyX7cTicket = {
    key: '',
    swan: [],
    wan: [],
    qian: [],
    bai: [],
    shi: [],
    gen: [],
    last: [],
    amount: 0,
}

const X7cSelectorView = ({onConfirm}: {
    onConfirm: (ticket: Ticket) => void,
}) => {

    const [selected, setSelected] = useState<Ticket>(emptyX7cTicket)

    const update = (pos, value) => {
        const newValue = {...selected,}
        newValue[pos] = value

        newValue['amount'] = newValue.gen.length * newValue.shi.length * newValue.bai.length * newValue.qian.length *
            newValue.wan.length * newValue.swan.length * newValue.last.length * 2

        setSelected(newValue)
    }

    const {createRandom, options, lastOptions} = useHook()

    const onSelectedGenChange = (pos, item, _selected) => {
        if (_selected) {
            if (selected[pos].includes(item.value)) {
                return;
            } else {
                update(pos, [...selected[pos], item.value])
            }
        } else {
            update(pos, selected[pos].filter(t => t !== item.value))
        }
    }

    const onRandom = () => {
        setSelected(createRandom())
    }
    const onClear = () => {
        setSelected(emptyX7cTicket)
    }

    const confirm = () => {
        if (validPlan) {

            const vl = {...selected}
            vl.key = createMd5(JSON.stringify(vl))
            onConfirm(vl)
        }
    }


    const validPlan = selected.gen.length > 0 && selected.shi.length > 0 && selected.bai.length > 0 &&
        selected.qian.length > 0 && selected.wan.length > 0 && selected.swan.length > 0 && selected.last.length > 0

    const {height} = useWindowDimensions()

    return <Stack h={height - 200}>
        <ScrollView style={{flex: 1}}>
            <Stack spacing={10}>
                {
                    [
                        {pos: 'swan', name: '一位'},
                        {pos: 'wan', name: '二位'},
                        {pos: 'qian', name: '三位'},
                        {pos: 'bai', name: '四位'},
                        {pos: 'shi', name: '五位'},
                        {pos: 'gen', name: '六位'},
                    ].map((t, i) =>
                        <View key={i} style={{flexDirection: "row", alignItems: "center"}}>
                            <Text variant="titleSmall">{t.name}</Text>
                            <WingBlank size="sm"/>
                            <Stack fill={1} spacing={5}>
                                <View>
                                    <Selector
                                        itemStyle={{selected: "#f53a3a", unselected: '#faa199'}}
                                        items={options}
                                        selectedKeys={selected[t.pos]}
                                        onSelectChanged={(item, selected) => onSelectedGenChange(t.pos, item, selected)}
                                    />
                                </View>

                                <HStack items="center" spacing={20}>
                                    <View><Chip
                                        onPress={() => update(t.pos, options.map(x => x.value))}>全</Chip></View>
                                    <View><Chip
                                        onPress={() => update(t.pos, options.filter(x => parseInt(x.value) % 2 !== 0)
                                            .map(x => x.value))}>奇</Chip></View>
                                    <View><Chip onPress={() => update(t.pos, options.filter(x =>
                                        parseInt(x.value) % 2 === 0).map(x => x.value))}>偶</Chip></View>
                                    <View><Chip onPress={() => update(t.pos, [])}>清</Chip></View>
                                </HStack>
                            </Stack>
                        </View>
                    )

                }
                {
                    [{pos: 'last', name: '七位'},].map((t, i) =>
                        <View key={i} style={{flexDirection: "row", alignItems: "center"}}>
                            <Text variant="titleSmall">{t.name}</Text>
                            <WingBlank size="sm"/>
                            <Stack fill={1} spacing={5}>
                                <View>
                                    <Selector
                                        itemStyle={{selected: "#346dfc", unselected: '#99bbfa'}}
                                        items={lastOptions}
                                        selectedKeys={selected[t.pos]}
                                        onSelectChanged={(item, selected) => onSelectedGenChange(t.pos, item, selected)}
                                    />
                                </View>
                                <HStack items="center" spacing={20}>
                                    <Chip onPress={() =>
                                        update(t.pos, lastOptions.map(x => x.value))
                                    }>全</Chip>
                                    <Chip onPress={() =>
                                        update(t.pos, lastOptions.filter(x =>
                                            parseInt(x.value) % 2 !== 0).map(x => x.value))
                                    }>奇</Chip>
                                    <Chip onPress={() => update(t.pos, lastOptions.filter(x =>
                                        parseInt(x.value) % 2 === 0).map(x => x.value))}>偶</Chip>
                                    <Chip onPress={() => update(t.pos, [])}>清</Chip>
                                </HStack>
                            </Stack>
                        </View>
                    )
                }
            </Stack>
        </ScrollView>
        <View style={{flexDirection: "row", alignItems: "center", paddingTop: 10}}>
            <View style={{flex: 1, flexDirection: "row", alignItems: "center"}}>
                <Button mode="text" style={{padding: 3}} onPress={onRandom}>机选</Button>
                <Button onPress={onClear}>清除</Button>
            </View>
            <View style={{flex: 1}}>
                <Button disabled={!validPlan}
                        onPress={confirm}
                        mode="contained"
                >
                    金额 {selected.amount} 元
                </Button>
            </View>
        </View>
    </Stack>

}

export default X7cSelectorView