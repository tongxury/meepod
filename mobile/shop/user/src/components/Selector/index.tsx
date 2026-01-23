import {Pressable, Text, View} from "react-native";
import {Button, Chip} from "react-native-paper";
import {Chip as RneChip} from "@rneui/themed";
import {Badge, Flex} from "@ant-design/react-native";
import {HStack, Stack} from "@react-native-material/core";


export declare type SelectorItem = {
    label: string,
    value: string
}

export const Selector = ({items, selectedKeys, onSelectChanged, itemStyle}: {
    items: SelectorItem[],
    selectedKeys: string[],
    itemStyle?: {
        selected: string,
        unselected: string,
    }
    onSelectChanged: (item: SelectorItem, selected: boolean) => void,
}) => {

    return <View>
        <Flex wrap="wrap" justify="between">
            {
                items.map(t => {
                    const isSelected = (selectedKeys || []).includes(t.value)

                    return <Pressable key={t.value} onPress={() => onSelectChanged(t, !isSelected)}>

                        <Stack center={true} m={2}
                               style={{
                                   width: 26,
                                   borderRadius: 13,
                                   height: 26,
                                   backgroundColor: isSelected ? itemStyle?.selected : itemStyle?.unselected
                               }}
                        >
                            <Text style={{color: "white"}}>{t.label}</Text>
                        </Stack>
                    </Pressable>

                })
            }
        </Flex>
    </View>
}

export const DoubleSelector = (
    {
        items,
        selectedKeys,
        doubleSelectedKeys,
        onSelectChanged,
        onDoubleSelectedChanged,
        itemStyle
    }: {
        items: SelectorItem[],
        selectedKeys: string[],
        doubleSelectedKeys: string[],
        itemStyle?: {
            selected: string,
            doubleSelected: string,
            unselected: string,
        }
        onSelectChanged: (item: SelectorItem, selected: boolean) => void,
        onDoubleSelectedChanged?: (item: SelectorItem, selected: boolean) => void
    }) => {


    const onDoubleSelect = (item: SelectorItem, tagged: boolean) => {
        if (tagged) {
            onSelectChanged(item, true)
        }

        onDoubleSelectedChanged(item, tagged)
    }

    return <Flex wrap="wrap" justify="between">
        {
            items.map(t => {
                const isSelected = (selectedKeys || []).includes(t.value)
                const isDoubleSelected = (doubleSelectedKeys || []).includes(t.value)

                const isSelectedWhatEver = isSelected || isDoubleSelected

                return <Pressable key={t.value} onLongPress={() => onDoubleSelect(t, !isDoubleSelected)}
                                  onPress={() => onSelectChanged(t, !(isSelectedWhatEver))}>

                    <Stack center={true} m={2}

                           style={{
                               width: 30,
                               borderRadius: 15,
                               height: 30,
                               backgroundColor: isDoubleSelected ? itemStyle?.doubleSelected : (isSelected ? itemStyle?.selected : itemStyle?.unselected)
                           }}
                    >
                        <Text style={{color: "white"}}>{t.label}</Text>
                    </Stack>
                </Pressable>

            })
        }
    </Flex>
}


