import {Dimensions, ScrollView, TouchableWithoutFeedback, View} from "react-native";
import {Appbar, Avatar, Badge, Button, Text, IconButton} from "react-native-paper";
import {Flex, Toast, WhiteSpace, WingBlank} from "@ant-design/react-native";
import {useContext, useEffect, useState} from "react";
import {DoubleSelector, Selector} from "../../../../components/Selector";
import useHook from "./hook";
import {Ticket} from "../../types";
import {createMd5} from "../../../../utils";


const Z1SelectorView = ({onConfirm}: {
    onConfirm: (plan: Ticket) => void,
}) => {

    const {options, getRandom} = useHook()
    const [amount, setAmount] = useState<number>()

    const [selectedGen, setSelectedGen] = useState<string[]>([])
    const [selectedShi, setSelectedShi] = useState<string[]>([])
    const [selectedBai, setSelectedBai] = useState<string[]>([])

    const updateAmount = () => {
        setAmount(selectedGen.length * selectedShi.length * selectedBai.length * 2)
    }

    useEffect(() => {
        updateAmount()
    }, [selectedGen, selectedShi, selectedBai])

    const onSelectedGenChange = (item, selected) => {
        if (selected) {
            if (selectedGen.includes(item.value)) {
                return;
            } else {
                setSelectedGen([...selectedGen, item.value])
            }
        } else {
            setSelectedGen(selectedGen.filter(t => t !== item.value))
        }
    }
    const onSelectedShiChange = (item, selected) => {
        if (selected) {
            if (selectedShi.includes(item.value)) {
                return;
            } else {
                setSelectedShi([...selectedShi, item.value])
            }
        } else {
            setSelectedShi(selectedShi.filter(t => t !== item.value))
        }
    }
    const onSelectedBaiChange = (item, selected) => {
        if (selected) {
            if (selectedBai.includes(item.value)) {
                return;
            } else {
                setSelectedBai([...selectedBai, item.value])
            }
        } else {
            setSelectedBai(selectedBai.filter(t => t !== item.value))
        }
    }

    const onRandom = () => {
        setSelectedGen(getRandom())
        setSelectedShi(getRandom())
        setSelectedBai(getRandom())
    }
    const onClear = () => {
        setSelectedGen([])
        setSelectedShi([])
        setSelectedBai([])
    }

    const confirm = () => {
        if (validPlan) {
            onConfirm({
                cat: 'z1',
                key: `z1${selectedBai.join("")}${selectedShi.join("")}${selectedGen.join("")}`,
                bai: selectedBai,
                shi: selectedShi,
                gen: selectedGen,
                amount: amount
            })
        }
    }


    const validPlan = selectedGen.length > 0 && selectedShi.length > 0 && selectedBai.length > 0


    return <View>
        <View style={{flexDirection: "row", alignItems: "center"}}>
            <Text variant="titleSmall">百位</Text>
            <WingBlank size="sm"/>
            <View style={{flex: 1}}>
                <Selector
                    itemStyle={{selected: "#f53a3a", unselected: '#faa199'}}
                    items={options}
                    selectedKeys={selectedBai}
                    onSelectChanged={onSelectedBaiChange}
                />
                <View style={{flexDirection: "row", alignItems: "center"}}>
                    <Button onPress={() => setSelectedBai(options.map(t => t.value))}>全</Button>
                    <Button onPress={() => setSelectedBai(options.filter(t =>
                        parseInt(t.value) % 2 !== 0).map(t => t.value))}>奇</Button>
                    <Button onPress={() => setSelectedBai(options.filter(t =>
                        parseInt(t.value) % 2 == 0).map(t => t.value))}>偶</Button>
                    <Button onPress={() => setSelectedBai([])}>清</Button>
                </View>
            </View>
        </View>
        <WhiteSpace size="lg"/>
        <View style={{flexDirection: "row", alignItems: "center"}}>
            <Text variant="titleSmall">十位</Text>
            <WingBlank size="sm"/>
            <View style={{flex: 1}}>
                <Selector
                    itemStyle={{selected: "#f53a3a", unselected: '#faa199'}}
                    items={options}
                    selectedKeys={selectedShi}
                    onSelectChanged={onSelectedShiChange}
                />
                <View style={{flexDirection: "row", alignItems: "center"}}>
                    <Button onPress={() => setSelectedShi(options.map(t => t.value))}>全</Button>
                    <Button onPress={() => setSelectedShi(options.filter(t =>
                        parseInt(t.value) % 2 !== 0).map(t => t.value))}>奇</Button>
                    <Button onPress={() => setSelectedShi(options.filter(t =>
                        parseInt(t.value) % 2 == 0).map(t => t.value))}>偶</Button>
                    <Button onPress={() => setSelectedShi([])}>清</Button>
                </View>
            </View>
        </View>
        <WhiteSpace size="lg"/>
        <View style={{flexDirection: "row", alignItems: "center"}}>
            <Text variant="titleSmall">个位</Text>
            <WingBlank size="sm"/>
            <View style={{flex: 1}}>
                <Selector
                    itemStyle={{selected: "#f53a3a", unselected: '#faa199'}}
                    items={options}
                    selectedKeys={selectedGen}
                    onSelectChanged={onSelectedGenChange}
                />
                <View style={{flexDirection: "row", alignItems: "center"}}>
                    <Button onPress={() => setSelectedGen(options.map(t => t.value))}>全</Button>
                    <Button onPress={() => setSelectedGen(options.filter(t =>
                        parseInt(t.value) % 2 !== 0).map(t => t.value))}>奇</Button>
                    <Button onPress={() => setSelectedGen(options.filter(t =>
                        parseInt(t.value) % 2 == 0).map(t => t.value))}>偶</Button>
                    <Button onPress={() => setSelectedGen([])}>清</Button>
                </View>
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

export default Z1SelectorView