import {Text, View} from "react-native";
import React, {useCallback, useState} from "react";
import Collapsible from "react-native-collapsible";
import ButtonSelector from "../../components/ButtonSelectorV2";
import {Chip, Searchbar, useTheme} from "react-native-paper";
import {HStack} from "@react-native-material/core";
import IconAntd from "react-native-vector-icons/AntDesign";


declare type FilterOption = {
    value: string
    label: string
}

const FilterView = ({values, items, onChange}: {
    items: { [key: string]: FilterOption[] }
    values: { [key: string]: string },
    onChange: (name: string, value: string) => void
}) => {

    const [collapsed, setCollapsed] = useState<{ [key: string]: boolean }>({})

    const onCollapseChange = (name: string) => {
        collapsed[name] = !collapsed[name]

        const newValues = {...collapsed}
        newValues[name] = !newValues[name]

        setCollapsed(newValues)
    }

    const {colors} = useTheme()

    return <View style={{backgroundColor: colors.background}}>
        <HStack items={"center"} justify={"between"} p={15}>
            {
                Object.keys(items).map((t, i) =>
                    <Chip key={i} onClose={() => undefined}
                          closeIcon={() => <IconAntd
                              color={colors.primary}
                              name={collapsed[t] ? "caretright" : "caretdown"}/>}
                          onPress={() => onCollapseChange(t)}>
                        {items[t]?.filter(k => values[t] === k.value)?.[0]?.label}
                    </Chip>
                )
            }
        </HStack>
        {
            Object.keys(items).map((t, i) =>
                <Collapsible style={{paddingHorizontal: 10}} collapsed={collapsed[t]}>
                    <ButtonSelector options={items[t]?.map(t => ({name: t.label, value: t.value}))} value={values[t]}
                                    onChange={v => {
                                        onChange(t, v)
                                        onCollapseChange(t)
                                    }}/>
                </Collapsible>
            )

        }
    </View>
}

export default FilterView