import {HStack, Stack} from "@react-native-material/core";
import {CheckBox} from "@rneui/themed";
import {Text, useTheme} from "react-native-paper";
import {useContext, useState} from "react";
import IconAntd from "react-native-vector-icons/AntDesign";
import {AppContext} from "../../../../providers/global";


const ModeView = ({size, values, onChange, max}: {
    size: number,
    values: string[],
    onChange: (mode) => void,
    max: number
}) => {


    const modeGroups = [
        // [
        //     {value: '1-1', name: '单关', min: 1, pre: true},
        // ],
        [
            {value: '2-1', name: '2串1', min: 2, factors: [2], pre: true},
        ],
        [
            {value: '3-1', name: '3串1', min: 3, factors: [3], pre: true},
            {value: '3-3', name: '3串3', min: 3, factors: [2],},
            {value: '3-4', name: '3串4', min: 3, factors: [2, 3]},
        ],
        [
            {value: '4-1', name: '4串1', min: 4, factors: [4], pre: true},
            {value: '4-4', name: '4串4', min: 4, factors: [3]},
            {value: '4-5', name: '4串5', min: 4, factors: [3, 4]},
            {value: '4-6', name: '4串6', min: 4, factors: [2]},
            {value: '4-11', name: '4串11', min: 4, factors: [2, 3, 4]},
        ],
        [
            {value: '5-1', name: '5串1', min: 5, factors: [5], pre: true},
            {value: '5-5', name: '5串5', min: 5, factors: [4]},
            {value: '5-6', name: '5串6', min: 5, factors: [4, 5]},
            {value: '5-10', name: '5串10', min: 5, factors: [3]},
            {value: '5-16', name: '5串16', min: 5, factors: [3, 4, 5]},
            {value: '5-20', name: '5串20', min: 5, factors: [2, 3]},
            {value: '5-26', name: '5串26', min: 5, factors: [2, 3, 4, 5]},
        ],
        [
            {value: '6-1', name: '6串1', min: 6, factors: [6], pre: true},
            {value: '6-6', name: '6串6', min: 6, factors: [5],},
            {value: '6-7', name: '6串7', min: 6, factors: [5, 6]},
            {value: '6-15', name: '6串15', min: 6, factors: [4]},
            {value: '6-20', name: '6串20', min: 6, factors: [3]},
            {value: '6-22', name: '6串22', min: 6, factors: [4, 5, 6]},
            {value: '6-35', name: '6串35', min: 6, factors: [3, 4]},
            {value: '6-42', name: '6串42', min: 6, factors: [3, 4, 5, 6]},
            {value: '6-50', name: '6串50', min: 6, factors: [2, 3, 4]},
            {value: '6-57', name: '6串57', min: 6, factors: [2, 3, 4, 5, 6]},
        ],
        [
            {value: '7-1', name: '7串1', min: 7, factors: [7], pre: true},
            {value: '7-7', name: '7串7', min: 7, factors: [6]},
            {value: '7-8', name: '7串8', min: 7, factors: [6, 7]},
            {value: '7-21', name: '7串21', min: 7, factors: [5]},
            {value: '7-35', name: '7串35', min: 7, factors: [4]},
            {value: '7-120', name: '7串120', min: 7, factors: [2, 3, 4, 5, 6, 7]},
        ],
        [
            {value: '8-1', name: '8串1', min: 8, factors: [8], pre: true},
            {value: '8-8', name: '8串8', min: 8, factors: [7]},
            {value: '8-9', name: '8串9', min: 8, factors: [7, 8]},
            {value: '8-28', name: '8串28', min: 8, factors: [6]},
            {value: '8-56', name: '8串56', min: 8, factors: [5]},
            {value: '8-70', name: '8串70', min: 8, factors: [4]},
            {value: '8-247', name: '8串247', min: 8, factors: [2, 3, 4, 5, 6, 7, 8]},
        ]
    ]

    const [collapsed, setCollapsed] = useState<boolean>(true)

    const moreCount = values?.length

    const {colors} = useTheme()

    return <Stack>
        <HStack items={"center"} wrap={"wrap"}>
            {modeGroups.map(g => {
                return g.filter(t => size >= t.min && t.pre).map(t =>
                    <CheckBox key={t.value} checked={values.includes(t.value)} size={20}
                              title={t.name}
                              onPress={() => onChange(t)}/>)
            })}

        </HStack>
        {size >= 3 &&
            <HStack items={"center"} ph={10} spacing={5}>
                <Text onPress={() => setCollapsed(!collapsed)}
                      style={{fontSize: 15,}}
                >
                    {collapsed ? '展开更多' : '隐藏'}
                    <Text
                        style={{
                            fontSize: 15,
                            fontWeight: "bold",
                            color: moreCount > 0 ? colors.primary : colors.onSurfaceVariant,
                        }}> {moreCount > 0 && `${moreCount} 项`}</Text>
                </Text>
                <IconAntd
                    color={moreCount > 0 ? colors.primary : colors.onSurfaceVariant}
                    onPress={() => setCollapsed(!collapsed)}
                    name={collapsed ? "caretdown" : "caretup"}/>
            </HStack>
        }
        {
            !collapsed && <HStack items={"center"} spacing={5} wrap={"wrap"}>
                {modeGroups.map(g => {
                    return g.filter(t => size >= t.min && !t.pre).map(t =>
                        <CheckBox key={t.value} checked={values.includes(t.value)} size={20}
                                  title={t.name}
                                  onPress={() => onChange(t)}/>)
                })}
            </HStack>
        }
    </Stack>

}

export default ModeView