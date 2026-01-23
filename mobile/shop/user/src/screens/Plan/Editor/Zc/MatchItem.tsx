import {HStack, Stack} from "@react-native-material/core";
import {View} from "react-native";
import {Button, Checkbox, Text, useTheme} from "react-native-paper";
import IconAntd from "react-native-vector-icons/AntDesign";
import {useEffect, useState} from "react";
import {Chip, ListItem, Button as RneButton, CheckBox} from "@rneui/themed";
import {Match, Odds, OddsItem} from "../../../../service/typs";
import {sumReduce} from "../../../../utils";


export declare type MatchValue = { odds?: Odds, dan?: boolean } | undefined

const MatchItem = ({meta: x, onChange, hideMore, hideDan, disableDan}: {
    meta: Match,
    onChange: (newValue: MatchValue) => void,
    hideMore?: boolean,
    hideDan?: boolean,
    disableDan?: boolean
}) => {

    // state 在组件内部管理，因为放在外面重新渲染的成本过高
    const [value, setValue] = useState<MatchValue>()
    const onOptionChange = (cat: string, option: OddsItem, selected: boolean) => {

        let tmpValues
        if (selected) {
            tmpValues = value ?? {odds: {}}
            tmpValues.odds[cat] = (tmpValues.odds[cat] ?? []).concat(option)
        } else {
            tmpValues = value
            tmpValues.odds[cat] = tmpValues.odds[cat].filter(t => t.result != option.result)

        }

        const optionCount = sumReduce(Object.keys(tmpValues.odds).map(k => tmpValues.odds[k]?.length || 0))

        if (optionCount === 0 && !tmpValues?.dan) {
            tmpValues = undefined
        }

        setValue(tmpValues)
        onChange(tmpValues)
    }


    const onDanChange = (selected: boolean) => {

        let tmpValues = {...(value || {}), dan: selected}
        const optionCount = sumReduce(Object.keys(tmpValues.odds ?? {}).map(k => tmpValues.odds[k]?.length || 0))

        if (optionCount === 0 && !tmpValues?.dan) {
            tmpValues = undefined
        }

        setValue(tmpValues)
        onChange(tmpValues)
    }

    useEffect(() => {
        if (hideDan && value?.dan) {
            onDanChange(false)
        }
    }, [hideDan])


    const counts = Object.keys(value?.odds ?? {}).filter(t => t != 'spf' && t != 'rspf').map(t => value?.odds?.[t]?.length ?? 0)
    const moreCount = counts.length == 0 ? 0 : counts.reduce((a, c) => a + c)

    const [collapsed, setCollapsed] = useState<boolean>(false)

    const oddsOptions = [
        {title: '', cat: 'items', options: x.odds?.items ?? []},
        x.odds?.r_items && {title: '', cat: 'r_items', options: x.odds?.r_items ?? []},
    ]

    const moreOddsOptions = [
        {title: '进球数', cat: 'goals_items', options: x.odds?.goals_items ?? []},
        {title: '半全场', cat: 'half_full_items', options: x.odds?.half_full_items ?? []},
        {title: '总比分(胜)', cat: 'score_victory_items', options: x.odds?.score_victory_items ?? []},
        {title: '总比分(平)', cat: 'score_dogfall_items', options: x.odds?.score_dogfall_items ?? []},
        {title: '总比分(负)', cat: 'score_defeat_items', options: x.odds?.score_defeat_items ?? []},
    ]
    const {colors} = useTheme()


    return <Stack style={{flex: 1}} spacing={8}>
        <HStack items="center" spacing={5}>
            <View><Chip buttonStyle={{backgroundColor: colors.primaryContainer}}
                        titleStyle={{color: colors.onPrimaryContainer}}
                        radius={0}>{x.league}</Chip></View>
            <HStack items="center" justify={"between"} fill={true} spacing={10}>
                <Text variant="titleMedium" numberOfLines={1}
                      ellipsizeMode="tail">{x.home_team_tag}{x.home_team}</Text>
                <Text variant="titleMedium">VS</Text>
                <Text variant="titleMedium"
                      numberOfLines={1}>{x.guest_team_tag}{x.guest_team}</Text>
            </HStack>
        </HStack>
        <Stack fill={1} spacing={5}>
            {oddsOptions.filter(t => !!t).map(t =>
                <HStack key={t.cat} items={"center"} fill={1}>
                    {
                        t.options.map(op => {
                            const selected = value?.odds?.[t.cat]?.map(t => t.result)?.includes(op.result)

                            return <RneButton
                                key={op.result}
                                title={<Text style={{
                                    fontSize: 14,
                                    color: selected ? colors.onPrimary : colors.onSecondaryContainer
                                }}>{op.name} {op.value}</Text>}
                                containerStyle={{minWidth: 100, height: 30, flex: 3, borderRadius: 0}}
                                size="sm"
                                radius={0}
                                color={selected ? colors.primary : colors.secondaryContainer}
                                onPress={() => onOptionChange(t.cat, op, !selected)}
                            />
                        })
                    }
                </HStack>
            )}
            <HStack items={"center"} justify={"between"} spacing={5}>
                <Text variant={"bodySmall"}>{x.start_at}</Text>

                <HStack items={"center"} spacing={2}>
                    {!hideDan &&
                        <CheckBox disabled={disableDan} checked={value?.dan} size={20} title="胆"
                                  onPress={() => onDanChange(!value?.dan)}/>
                    }
                    {
                        !hideMore && <HStack spacing={2} items={"center"}>
                            <Text
                                onPress={() => setCollapsed(!collapsed)}
                                style={{
                                    fontSize: 15,
                                    color: moreCount > 0 ? colors.primary : colors.onSurfaceVariant,
                                    fontWeight: moreCount > 0 ? 'bold' : 'normal'
                                }}
                            >
                                {!collapsed ? '展开' : '隐藏'} {moreCount} 项
                            </Text>
                            <IconAntd
                                color={moreCount > 0 ? colors.primary : colors.onSurfaceVariant}
                                onPress={() => setCollapsed(!collapsed)}
                                name={!collapsed ? "caretdown" : "caretup"}/>
                        </HStack>
                    }
                </HStack>
            </HStack>

            {
                collapsed &&
                <Stack spacing={10}>
                    {
                        moreOddsOptions.map((t, i) => <Stack key={i} spacing={5}>
                            <Text variant="titleSmall">{t.title}</Text>
                            <HStack items="center" fill={1} wrap={true}>
                                {
                                    t.options?.map((op, i) => {

                                        // const selected = options.get(t.cat)?.includes(op.result)
                                        const selected = value?.odds?.[t.cat]?.map(t => t.result)?.includes(op.result)

                                        return <RneButton
                                            key={i}
                                            title={`${op.name} ${op.value}`}
                                            containerStyle={{minWidth: 100, flex: 3, borderRadius: 0}}
                                            buttonStyle={{height: 30}}
                                            size="sm"
                                            radius={0}
                                            color={selected ? colors.primary : colors.secondaryContainer}
                                            titleStyle={{
                                                fontSize: 14,
                                                color: selected ? colors.onPrimary : colors.onSecondaryContainer
                                            }}
                                            onPress={() => onOptionChange(t.cat, op, !selected)}
                                        />

                                    })
                                }
                            </HStack>
                        </Stack>)
                    }
                </Stack>
            }

            {/*</Collapsible>*/}
        </Stack>
    </Stack>
}

export default MatchItem