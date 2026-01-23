import {
    FlatList, useWindowDimensions, View
} from "react-native";
import {Appbar, Checkbox, Text, Button, useTheme, Divider} from "react-native-paper";
import {useContext, useEffect, useState} from "react";
import {HStack, Stack} from "@react-native-material/core";
import MatchView, {MatchValue} from "./MatchItem";
import ModeView from "./ModeView";
import {createMd5, multiReduce, sumReduce} from "../../../../utils";
import {Ticket} from "../../types";
import {arrange, cnm} from "../../../../utils/math";
import {AppContext} from "../../../../providers/global";
import {Toast} from "@ant-design/react-native";
import {Match} from "../../../../service/typs";

const ZjcSelectorView = ({src, onConfirm, onClear}: {
    src: Match[],
    onConfirm: (plan: Ticket) => void,
    onClear: () => void,
}) => {

    const {colors} = useTheme()

    const minMatches = 2
    const maxModes = 5


    const [matchOptions, setMatchOptions] = useState<Map<string, MatchValue>>(new Map<string, MatchValue>())
    const [modeOptions, setModeOptions] = useState<any[]>([])
    const [amount, setAmount] = useState<number>(0)


    const updateAmount = (modeOptions: any[], matchOptions: Map<string, MatchValue>) => {
        const matchOptionCounts = Array.from(matchOptions.keys()).reduce(function (map: {}, matchId) {
            map[matchId] = sumReduce(Object.keys(matchOptions.get(matchId)?.odds ?? {})
                .map(cat => matchOptions.get(matchId)?.odds?.[cat]?.length))

            return map;
        }, {})
        const matchIds = Array.from(matchOptions.keys())

        const amountValues = modeOptions.map(mode => {
            const parts = mode.value.split('-')
            const m = parts[0]

            if (m > matchIds.length) {
                return 0
            }

            return arrange(matchIds, m).map(matchIds => matchIds.map(matchId => matchOptionCounts[matchId]))
                .map(counts => {
                    return mode.factors.map(f => {
                        const options = arrange(counts, f)
                        return options.map(t => t.reduce((a1, a2) => a1 * a2)).reduce((a1, a2) => a1 + a2)
                    }).reduce((a1, a2) => a1 + a2)
                }).reduce((a1, a2) => a1 + a2)
        })

        setAmount(sumReduce(amountValues) * 2)
    }

    const confirm = () => {
        const ticket = {
            options: Object.fromEntries(matchOptions.entries()),
            modes: Array.from(modeOptions.values()).map(t => t.value),
            matches: src.map(t => ({
                id: t.id,
                home_team: t.home_team,
                r_count: t.r_count
            })).reduce((map, v) => (map[v.id] = v, map), {}),
            amount: amount,
            key: createMd5(JSON.stringify(Array.from(matchOptions.values())) + JSON.stringify(Array.from(modeOptions.values()).map(t => t.value)))
        }

        onConfirm(ticket)
    }

    const onChange = (matchId: string, newValue: MatchValue) => {

        const newValues = new Map(matchOptions.set(matchId, newValue))

        setMatchOptions(newValues)
        updateAmount(modeOptions, newValues)
    }

    const onModeChange = (mode) => {

        let newValues = modeOptions ?? []
        if (newValues?.map(t => t.value)?.includes(mode.value)) {
            newValues = newValues.filter(t => t.value != mode.value)
        } else {

            if (newValues.length < maxModes) {
                newValues = newValues.concat(mode)
            } else {
                Toast.info(`组合方式最多可选${maxModes}个`)
            }
        }

        setModeOptions(newValues)
        updateAmount(newValues, matchOptions)
    }


    const validPlan = matchOptions?.size >= minMatches && modeOptions.length > 0
    // const validPlan = true

    const {height} = useWindowDimensions()

    return <Stack fill={true}>
        <FlatList
            style={{height: height - 200}}
            data={src ?? []}
            ItemSeparatorComponent={() => <Divider bold={true} style={{marginVertical: 15}}/>}
            renderItem={({item: x, index}) => <Stack fill={1}>
                <MatchView hideDan={true} meta={x} onChange={newValue => onChange(x.id, newValue)}/>
            </Stack>
            }
        />

        <Stack spacing={10} pt={10}>
            <View>
                <ModeView size={matchOptions.size} onChange={onModeChange} values={modeOptions?.map(t => t.value) ?? []}
                          max={maxModes}/>
            </View>
            <HStack items={"center"} spacing={10}>
                <View style={{flex: 2, flexDirection: "row", alignItems: "center"}}>
                    {/*<Button mode="text" style={{padding: 3}} onPress={onRandom}>机选</Button>*/}
                    <Button onPress={onClear}>清空重选</Button>
                </View>
                <View style={{flex: 3}}>
                    <Button
                        disabled={!validPlan}
                        onPress={confirm}
                        mode="contained"
                    >
                        {matchOptions?.size >= minMatches ? ` 金额 ${amount} 元` : `已选${matchOptions?.size}场(至少${minMatches}场)`}
                    </Button>
                </View>
            </HStack>
        </Stack>

    </Stack>

}

export default ZjcSelectorView