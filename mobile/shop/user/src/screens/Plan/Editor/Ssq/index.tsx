import {View} from "react-native";
import {Button, Text} from "react-native-paper";
import {Flex, Toast, WhiteSpace} from "@ant-design/react-native";
import {useContext, useEffect, useState} from "react";
import {factorial} from "../../../../utils/math";
import {DoubleSelector, Selector} from "../../../../components/Selector";
import {HStack, Stack} from "@react-native-material/core";
import {AppContext} from "../../../../providers/global";
import {Ticket} from "../../types";
import useHook from "./hook";

const SsqSelectorView = ({onConfirm}: {
    onConfirm: (plan: Ticket) => void,
    // maxBlue: number,
    // maxRed: number,
    // maxDRed: number,
}) => {

    const {settingsState: {settings}} = useContext<any>(AppContext);
    const maxBlue = settings.ssqBlueLimit
    const maxRed = settings.ssqRedLimit
    const maxDRed = settings.ssqDRedLimit


    const {redOptions, blueOptions, getRandom, genKey} = useHook()

    const [selectedRed, setSelectedRed] = useState<string[]>([])
    const [doubleSelectedRed, setDoubleSelectedRed] = useState<string[]>([])
    const [selectedBlue, setSelectedBlue] = useState<string[]>([])

    const [amount, setAmount] = useState<number>()

    useEffect(() => {
        updateAmount()
    }, [selectedRed, selectedBlue, doubleSelectedRed])

    const confirm = () => {
        if (validPlan()) {
            const key = `${JSON.stringify(selectedRed.sort())}-${JSON.stringify(doubleSelectedRed.sort())}-${JSON.stringify(selectedBlue.sort())}`

            const meta = {red: selectedRed, redD: doubleSelectedRed, blue: selectedBlue, amount, key: ''}
            meta.key = genKey(meta)
            onConfirm(meta)
        }
    }

    const onSelectedRedChange = (item, selected) => {
        if (selected) {
            if (selectedRed.length >= maxRed) {
                Toast.info({
                    content: `红球最多只能选${maxRed}个`,
                    mask: false,
                })
            } else {
                if (selectedRed.includes(item.value)) {
                    return;
                } else {
                    setSelectedRed([...selectedRed, item.value])
                }

            }

        } else {
            setSelectedRed(selectedRed.filter(t => t !== item.value))
            setDoubleSelectedRed(doubleSelectedRed.filter(t => t !== item.value))
        }
    }
    const onDoubleSelectedChange = (item, selected) => {
        if (selected) {
            if (doubleSelectedRed.length >= maxDRed) {
                Toast.info({
                    content: `胆拖最多只能选${maxDRed}个`,
                    mask: false,
                })

            } else {
                setDoubleSelectedRed([...doubleSelectedRed, item.value])
                setSelectedRed(selectedRed.filter(t => t !== item.value))
            }
        } else {
            setDoubleSelectedRed(doubleSelectedRed.filter(t => t !== item.value))
            setSelectedRed([...selectedRed, item.value])
        }
    }
    const onSelectedBlueChange = (item, selected) => {
        if (selected) {
            if (selectedBlue.length >= maxBlue) {
                Toast.info({
                    content: `蓝球最多只能选${maxBlue}个`,
                    mask: false,
                })

            } else {
                if (selectedBlue.includes(item.value)) {
                    return;
                } else {
                    setSelectedBlue([...selectedBlue, item.value])
                }
            }
        } else {
            setSelectedBlue(selectedBlue.filter(t => t !== item.value))
        }
    }

    const onRandom = () => {

        const randomOption = getRandom()

        setSelectedRed(randomOption.red)
        setSelectedBlue(randomOption.blue)
    }

    const onClear = () => {
        setSelectedRed([])
        setSelectedBlue([])
        setDoubleSelectedRed([])
    }

    const updateAmount = () => {

        const redCount = selectedRed.length
        const doubleRedCount = doubleSelectedRed.length
        const blueCount = selectedBlue.length

        if (redCount + doubleRedCount < 6) {
            setAmount(0)
            return
        }

        const num = blueCount * (factorial(redCount) / (factorial(6 - doubleRedCount) * factorial(redCount - (6 - doubleRedCount))))

        setAmount(num * 2)

    }

    const validPlan = () => {
        return selectedRed.length + doubleSelectedRed.length >= 6 && selectedBlue.length >= 1
    }

    return <Stack spacing={10}>
        <View>
            <DoubleSelector
                doubleSelectedKeys={doubleSelectedRed}
                itemStyle={{doubleSelected: "#f53aba", selected: "#f53a3a", unselected: '#faa199'}}
                items={redOptions}
                selectedKeys={selectedRed}
                onSelectChanged={onSelectedRedChange}
                onDoubleSelectedChanged={onDoubleSelectedChange}
            />
        </View>
        <View>
            <Selector
                itemStyle={{selected: "#346dfc", unselected: '#99bbfa'}}
                items={blueOptions}
                selectedKeys={selectedBlue}
                onSelectChanged={onSelectedBlueChange}/>
        </View>
        <Text variant={"bodySmall"}>红球长按可设置胆拖</Text>
        <HStack items="center" spacing={10}>
            <HStack items="center" spacing={5}>
                <Button mode="contained-tonal" onPress={onRandom}>机选</Button>
                <Button onPress={onClear}>清除</Button>
            </HStack>
            <View style={{flex: 1}}>
                <Button disabled={!validPlan()}
                        onPress={confirm}
                        mode="contained"
                >
                    金额 {amount} 元
                </Button>
            </View>
        </HStack>
    </Stack>

}

export default SsqSelectorView