import {StyleProp, Text, View, ViewStyle} from "react-native";
import {Button} from "react-native-paper";
import {Badge, Flex} from "@ant-design/react-native";
import {Chip} from "@rneui/themed";
import {HStack} from "@react-native-material/core";


export declare type SelectorItem = {
    label: string,
    value: string
}

export const Selector = ({items, selectedKeys, onSelectChanged, itemStyle, style}: {
    items: SelectorItem[],
    selectedKeys: string[],
    itemStyle?: {
        selected: string,
        unselected: string,
    }
    onSelectChanged: (item: SelectorItem, selected: boolean) => void,
    style?: StyleProp<ViewStyle>
}) => {

    return <HStack style={style} wrap="wrap" justify="start" spacing={2}>
        {
            items.map(t => {
                const isSelected = (selectedKeys || []).includes(t.value)

                return <Chip
                    style={{margin: 2}}
                    onPress={() => onSelectChanged(t, !isSelected)}
                    key={t.value}
                    color={isSelected ? itemStyle?.selected : itemStyle?.unselected}
                >
                    {t.label}
                </Chip>

            })
        }
    </HStack>
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

                return <Button
                    onLongPress={() => onDoubleSelect(t, !isDoubleSelected)}
                    onPress={() => onSelectChanged(t, !(isSelectedWhatEver))}
                    key={t.value}
                    style={{
                        margin: 3,
                        backgroundColor: isDoubleSelected ? itemStyle?.doubleSelected : (isSelected ? itemStyle?.selected : itemStyle?.unselected)
                    }}
                    mode="elevated">
                    <Text style={{color: "white"}}>{t.label}</Text>
                </Button>

            })
        }
    </Flex>
}


