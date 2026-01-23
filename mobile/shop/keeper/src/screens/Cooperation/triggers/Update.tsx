import editor from "../../Profile/Editor";
import {InputItem, Modal, Toast} from "@ant-design/react-native";
import {HStack, Stack} from "@react-native-material/core";
import {Button as RneButton, Image, Input} from "@rneui/themed";
import {Avatar, Button, Checkbox, Text, useTheme} from "react-native-paper";
import {StyleProp, View, ViewStyle, TextInput, FlatList} from "react-native";
import React, {useEffect, useState} from "react";
import {useForm, Controller} from "react-hook-form";
import UserView from "../../../components/UserView";
import {useRequest} from "ahooks";
import {addProxy, fetchStoreUsers, updateCoStore, updateProxy} from "../../../service/api";
import {CoStore, Proxy, User} from "../../../service/typs";
import StoreView from "../../../components/StoreView";
import ItemSelector from "../../../components/ItemSelector";

const UpdateTrigger = ({data, style, onConfirm}: {
    data: CoStore,
    onConfirm?: () => void,
    style?: StyleProp<ViewStyle>
}) => {

    const [open, setOpen] = useState<boolean>(false)

    const [itemIds, setItemIds] = useState<string[]>(data?.items?.map(t => t.item.id))
    const confirm = () => {


        updateCoStore(data?.co_store?.id, itemIds.reduce((a, v) => ({...a, [v]: 1}), {})).then(rsp => {
            if (rsp.data?.code === 0) {
                setOpen(false)
                onConfirm?.()
            }
        })
    }

    const onChange = (itemId: string) => {

        let newItemIds = itemIds ?? [];

        if (itemIds?.includes(itemId)) {

            if (itemIds.length <= 1) {
                Toast.info("至少要选择一个彩种")
            } else {
                newItemIds = newItemIds.filter(t => t != itemId)
            }

        } else {
            newItemIds = newItemIds.concat(itemId)
        }

        setItemIds(newItemIds)

    }


    const {colors} = useTheme()

    return <View style={style}>
        <RneButton onPress={() => setOpen(true)}>修改彩种</RneButton>
        <Modal
            style={{borderTopLeftRadius: 10, borderTopRightRadius: 10}}
            popup
            visible={open}
            maskClosable={true}
            animationType="slide-up"
            onClose={() => {
                setOpen(false)
            }}>
            <Stack p={20} spacing={30}>
                <Stack spacing={20}>

                    <Stack spacing={8} fill={1}>
                        <Text variant={"titleMedium"}>合作店铺</Text>
                        <StoreView data={data?.co_store}/>
                    </Stack>

                    <Stack spacing={5} fill={1}>
                        <Text variant={"titleMedium"}>合作彩种</Text>

                        <ItemSelector values={itemIds} onChange={values => setItemIds(values)}
                                      items={data?.co_store?.items}/>

                    </Stack>

                </Stack>


                <Button mode="contained" onPress={confirm}>确认</Button>
            </Stack>
        </Modal>
    </View>

}

export default UpdateTrigger