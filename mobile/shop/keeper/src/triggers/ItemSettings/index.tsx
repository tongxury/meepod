import {ActivityIndicator, FlatList, Pressable, View} from "react-native";
import React, {useContext, useEffect, useState} from "react";
import {Avatar, Button, Checkbox, Dialog, Portal, RadioButton, Text, TextInput, useTheme} from "react-native-paper";
import {ImagePicker, InputItem, Modal, TextareaItem, WhiteSpace, WingBlank} from "@ant-design/react-native";
import {HStack, Stack} from "@react-native-material/core";
import {AppContext} from "../../providers/global";
import {Image} from "@rneui/themed";
import {useRequest} from "ahooks";
import {fetchItems, fetchSettings, updateStoreItems} from "../../service/api";

const ItemSettingsTrigger = ({values, onConfirm, children}: {
    values: string[]
    onConfirm?: () => void,
    children: React.ReactNode
}) => {

    const [visible, setVisible] = useState(false);

    const {data, runAsync, loading} = useRequest(fetchItems, {manual: true})
    const items = data?.data?.data

    const [itemIds, setItemIds] = useState<string[]>(values)

    useEffect(() => {
        if (visible) {
            runAsync().then(rsp => {
                // setItemIds(rsp.data?.data?.selected ?? [])
            })
        }
    }, [visible])

    useEffect(() => {
        setItemIds(values)
    }, [values])


    const onChange = (itemId: string) => {

        let newItemIds = itemIds ?? [];

        if (itemIds?.includes(itemId)) {
            newItemIds = newItemIds.filter(t => t != itemId)
        } else {
            newItemIds = newItemIds.concat(itemId)
        }

        setItemIds(newItemIds)

    }

    const confirm = () => {
        updateStoreItems({itemIds}).then(rsp => {
            setVisible(false)
            onConfirm?.()
        })
    }

    const {colors} = useTheme()

    return <Pressable onPress={() => setVisible(true)}>
        {children}
        <Modal
            style={{borderTopLeftRadius: 10, borderTopRightRadius: 10}}
            popup
            visible={visible}
            bodyStyle={{minHeight: 100}}
            maskClosable={true}
            animationType="slide-up"
            onClose={() => {
                setVisible(false)
            }}>
            <Stack p={20} spacing={30} fill={1}>

                <FlatList style={{flex: 1}} data={items} numColumns={2} renderItem={({item: x}) => {
                    return <HStack fill={1} items={"center"} spacing={5}>
                        <Checkbox onPress={() => onChange(x.id)}
                                  status={itemIds?.includes(x.id) ? "checked" : "unchecked"}/>
                        <Avatar.Image style={{backgroundColor: colors.background}} size={30} source={{uri: x.icon}}/>
                        <Text variant={"titleMedium"}>{x.name}</Text>
                    </HStack>

                }}/>

                <View style={{flexDirection: "row", alignItems: "center"}}>
                    <Button onPress={confirm} mode="contained-tonal"
                            style={{flex: 2}}>确认</Button>
                </View>
            </Stack>
        </Modal>
    </Pressable>
}

export default ItemSettingsTrigger