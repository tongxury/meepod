import {FlatList, StyleProp, ViewStyle} from "react-native";
import {HStack, Stack} from "@react-native-material/core";
import {Avatar, Checkbox, Text, useTheme} from "react-native-paper";
import React, {useState} from "react";
import {Item} from "../../service/typs";
import {updateCoStore} from "../../service/api";
import {Toast} from "@ant-design/react-native";

const ItemSelector = ({values, onChange, items, style}: {
    values: string[],
    onChange: (values: string[]) => void,
    items: Item[],
    style?: StyleProp<ViewStyle>
}) => {


    const change = (itemId: string) => {

        let newItemIds = values ?? [];

        if (values?.includes(itemId)) {

            if (values.length <= 1) {
                Toast.info("至少要选择一个彩种")
            } else {
                newItemIds = newItemIds.filter(t => t != itemId)
            }

        } else {
            newItemIds = newItemIds.concat(itemId)
        }

        onChange(newItemIds)

    }

    const {colors} = useTheme()


    return <Stack style={style}>
        <FlatList style={{flex: 1}} data={items} numColumns={2}
                  renderItem={({item: x}) => {
                      return <HStack fill={1} items={"center"} spacing={5}>
                          <Checkbox onPress={() => change(x.id)}
                                    status={values?.includes(x.id) ? "checked" : "unchecked"}/>
                          <Avatar.Image style={{backgroundColor: colors.background}} size={30}
                                        source={{uri: x.icon}}/>
                          <Text variant={"titleMedium"}>{x.name}</Text>
                      </HStack>

                  }}/>


    </Stack>
};

export default  ItemSelector
