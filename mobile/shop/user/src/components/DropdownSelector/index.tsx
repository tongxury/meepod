import SelectDropdown from "react-native-select-dropdown";
import {Text, useTheme} from "react-native-paper";
import Icon from "react-native-vector-icons/Entypo";
import React from "react";
import {color} from "@rneui/base";


export declare type Item = {
    name: string
    value: string
}
export const DropdownSelector = ({items, value, onChange}: {
    items: Item[],
    value: Item,
    onChange: (item: Item) => void
}) => {

    const {colors} = useTheme()

    return <SelectDropdown
        defaultValue={value}
        data={items}
        onSelect={(selectedItem, index) => {
            onChange(selectedItem)
        }}
        buttonTextAfterSelection={(selectedItem, index) => {
            return selectedItem.name
        }}
        rowTextForSelection={(item, index) => {
            return item.name
        }}
        searchPlaceHolder={"选择期数"}
        dropdownStyle={{maxHeight: 200, backgroundColor: colors.background, overflow: 'scroll'}}
        rowStyle={{padding: 5, height: 30, backgroundColor: colors.background}}
        rowTextStyle={{fontSize: 15, color: colors.onBackground}}
        selectedRowStyle={{padding: 5, backgroundColor: colors.primary}}
        selectedRowTextStyle={{fontSize: 15, color: colors.onPrimary}}
        // renderDropdownIcon={() => <Icon name={"chevron-down"} />}
        buttonStyle={{padding: 2, width: 120, height: 30, borderRadius: 5}}
        buttonTextStyle={{color: colors.primary, fontSize: 14}}
        renderDropdownIcon={() => <Icon name={"chevron-down"} color={colors.primary}/>}

        // renderCustomizedButtonChild={() => <Text
        //     // contentStyle={{borderRadius: 5}}
        //     // labelStyle={{margin: 2}}
        // >{value?.name}
        //     <Icon name={"chevron-down"}/></Text>}
    />
}

