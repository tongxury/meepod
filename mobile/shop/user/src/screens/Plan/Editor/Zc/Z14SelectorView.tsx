import {
    Dimensions,
    FlatList,
    ScrollView,
    TouchableWithoutFeedback, useWindowDimensions,
    View
} from "react-native";
import {Appbar, Checkbox, Chip, Button, Text, useTheme} from "react-native-paper";
import {useContext, useEffect, useState} from "react";
import {HStack, Stack} from "@react-native-material/core";
import {arrange, cnm} from "../../../../utils/math";
import {Ticket} from "../../types";
import MatchView, {MatchValue} from "./MatchItem";
import {createMd5, multiReduce, sumReduce} from "../../../../utils";
import {AppContext} from "../../../../providers/global";

const SelectorView = ({src, min, onConfirm, onClear}: {
    src: any,
    min: number,
    onConfirm: (ticket: Ticket) => void,
    onClear: () => void
}) => {

    const {colors} = useTheme()

    const [amount, setAmount] = useState<number>(0)
    const [matchOptions, setMatchOptions] = useState<Map<string, MatchValue>>(new Map<string, MatchValue>())

    const {settingsState: {settings}} = useContext<any>(AppContext);

    const updateAmount = (matchOptions: Map<string, MatchValue>) => {
        const matchOptionCounts = Array.from(matchOptions.keys()).reduce(function (map: {}, matchId) {

            map[matchId] = sumReduce(Object.keys(matchOptions.get(matchId)?.odds ?? {})
                .map(cat => matchOptions.get(matchId)?.odds?.[cat]?.length))

            return map;
        }, {})

        const count = () => {
            const matchIds = Array.from(matchOptions.keys())

            if (matchIds.length < min) {
                return 0
            }

            const notDanMatchIds = matchIds.filter(matchId =>
                !(matchOptions.get(matchId)?.dan))

            const danMatchIds = matchIds.filter(matchId =>
                (matchOptions.get(matchId)?.dan))

            const danCount = danMatchIds.length > 0 ? multiReduce(danMatchIds.map(matchId => matchOptionCounts[matchId])) : 1

            return arrange(notDanMatchIds, min - (matchIds.length - notDanMatchIds.length))
                .map(matchIds => matchIds.map(matchId => matchOptionCounts[matchId]))
                .map(counts => counts.reduce((a1, a2) => a1 * a2))
                .reduce((a1, a2) => a1 + a2) * danCount
        }

        setAmount(count() * 2)
    }


    const onChange = (matchId: string, newValue: MatchValue) => {

        let newValues = new Map(matchOptions.set(matchId, newValue))

        if (newValue === undefined) {
            matchOptions.delete(matchId)
            newValues = new Map(matchOptions)
        }

        console.log('newValues', newValues)

        setMatchOptions(newValues)
        updateAmount(newValues)
    }

    const confirm = () => {
        const ticket = {
            options: Object.fromEntries(matchOptions.entries()),
            matches: src.map(t => ({
                id: t.id,
                home_team: t.home_team,
                r_count: t.odds?.r_count
            })).reduce((map, v) => (map[v.id] = v, map), {}),
            amount: amount,
            key: createMd5(JSON.stringify(Array.from(matchOptions.values())))
        }

        onConfirm(ticket)
    }

    const validPlan = matchOptions?.size >= min
    const danMatchIds = Array.from(matchOptions.keys())
        .filter(matchId => matchOptions.get(matchId)?.dan)

    const {height} = useWindowDimensions()


    return <Stack fill={true}>
        <FlatList
            style={{height: height - 300}}
            showsVerticalScrollIndicator={false}
            data={src ?? []}
            ItemSeparatorComponent={() => <View style={{marginVertical: 8}}/>}
            renderItem={({item: x, index}) =>
                <Stack fill={1}>
                    <MatchView
                        meta={x}
                        disableDan={matchOptions.size === 0 || danMatchIds.length >= settings.rx9DanLimit}
                        hideDan={matchOptions.size <= min}
                        onChange={(newValue) => onChange(x.id, newValue)}
                        hideMore/>
                </Stack>
            }/>
        <HStack pt={10} items={"center"} spacing={10}>
            <Button onPress={onClear}>清空重选</Button>
            <View style={{flex: 3}}>
                <Button
                    disabled={!validPlan}
                    onPress={confirm}
                    mode="contained"
                >
                    {matchOptions?.size >= min ? ` 金额 ${amount} 元` : `已选${matchOptions?.size}场(至少${min}场)`}
                </Button>
            </View>
        </HStack>
    </Stack>

}

export default SelectorView