import {Dimensions, ScrollView, TouchableWithoutFeedback, View} from "react-native";
import {Appbar, Avatar, Badge, Button, Text, IconButton} from "react-native-paper";
import {Flex, Toast, WhiteSpace, WingBlank} from "@ant-design/react-native";
import {useContext, useEffect, useState} from "react";
import {DoubleSelector, Selector} from "../../../../components/Selector";
import useHook from "./hook";
import {cnm} from "../../../../utils/math";
import {Ticket} from "../../types";
import {createMd5} from "../../../../utils";


const Z3SelectorView = ({onConfirm}: {
    onConfirm: (plan: Ticket) => void,
}) => {

    const {options, getRandom} = useHook()
    const [amount, setAmount] = useState<number>()

    const [selectedTon, setSelectedTon] = useState<string[]>([])

    const updateAmount = () => {
        setAmount((cnm(selectedTon.length, 2) * 2) * 2)
    }

    useEffect(() => {
        updateAmount()
    }, [selectedTon])

    const onSelectedTonChange = (item, selected) => {
        if (selected) {
            if (selectedTon.includes(item.value)) {
                return;
            } else {
                setSelectedTon([...selectedTon, item.value])
            }
        } else {
            setSelectedTon(selectedTon.filter(t => t !== item.value))
        }
    }

    const onRandom = () => {
        setSelectedTon(getRandom(2))
    }
    const onClear = () => {
        setSelectedTon([])
    }

    const confirm = () => {
        if (validPlan) {
            onConfirm({
                cat: 'z3',
                cat_name: '组三',
                key: `z3${selectedTon.join("")}`,
                ton: selectedTon,
                amount: amount
            })
        }
    }


    const validPlan = selectedTon.length >= 2

    return <View>
        <View style={{flex: 1}}>
            <Selector
                itemStyle={{selected: "#f53a3a", unselected: '#faa199'}}
                items={options}
                selectedKeys={selectedTon}
                onSelectChanged={onSelectedTonChange}
            />
            <View style={{flexDirection: "row", alignItems: "center"}}>
                <Button onPress={() => setSelectedTon(options.map(t => t.value))}>全</Button>
                <Button onPress={() => setSelectedTon(options.filter(t =>
                    parseInt(t.value) % 2 !== 0).map(t => t.value))}>奇</Button>
                <Button onPress={() => setSelectedTon(options.filter(t =>
                    parseInt(t.value) % 2 == 0).map(t => t.value))}>偶</Button>
                <Button onPress={() => setSelectedTon([])}>清</Button>
            </View>
        </View>
        <WhiteSpace/>
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
                    金额 {amount} 元
                </Button>
            </View>
        </View>
    </View>

}

export default Z3SelectorView