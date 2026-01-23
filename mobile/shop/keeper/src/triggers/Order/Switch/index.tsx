import {ActivityIndicator, FlatList, Pressable, StyleProp, View, ViewStyle} from "react-native";
import React, {useContext, useEffect, useState} from "react";
import {Avatar, Button, Checkbox, Dialog, Portal, RadioButton, Searchbar, Text} from "react-native-paper";
import {AppContext} from "../../../providers/global";
import {ImagePicker, Modal, WhiteSpace, WingBlank} from "@ant-design/react-native";
import {HStack, Stack} from "@react-native-material/core";
import {useRequest} from "ahooks";
import {fetchCoStores} from "../../../service/api";
import {mainBodyHeight} from "../../../utils/dimensions";
import AmountView from "../../../components/AmountView";
import {Button as RneButton} from "@rneui/themed";

const SwitchTrigger = ({amount, itemId, onConfirm, style}: {
    amount: number,
    itemId: string,
    onConfirm?: (storeId: string) => void,
    style?: StyleProp<ViewStyle> | undefined
}) => {

    const [visible, setVisible] = useState(false);

    const {settingsState: {settings: appSettings}} = useContext<any>(AppContext);

    const [keyword, setKeyword] = useState<string>()

    const [storeId, setStoreId] = useState<string>()

    const {data, run} = useRequest(() => fetchCoStores({keyword, itemId, enableOnly: "1", page: 1, size: 10}), {manual: true})
    const stores = data?.data?.data

    useEffect(() => {
        if (visible) {

            run()
        } else {
            setStoreId('')
        }
    }, [visible, keyword])

    return <HStack style={style}>
        <RneButton onPress={() => setVisible(true)}>转单</RneButton>
        <Modal
            style={{borderTopLeftRadius: 10, borderTopRightRadius: 10}}
            popup
            visible={visible}
            bodyStyle={{maxHeight: mainBodyHeight - 200, padding: 15}}
            maskClosable={true}
            animationType="slide-up"
            onClose={() => {
                setVisible(false)
            }}>

            <Stack spacing={10}>
                <Searchbar
                    placeholder="店铺名称或邀请码搜索"
                    onChangeText={text => setKeyword(text)}
                    value={keyword}
                />
                <FlatList
                    data={stores?.list ?? []}
                    numColumns={1}
                    ItemSeparatorComponent={() => <View style={{height: 8}}/>}
                    renderItem={({item: x, index}) => {
                        return <HStack items={"center"} justify={"between"}>
                            <HStack items={"center"} spacing={10}>
                                <Avatar.Image size={40} source={{uri: x.co_store?.icon}}/>
                                <Stack>
                                    <Text variant={"titleMedium"}>{x.co_store?.name}</Text>
                                    <Text variant={"bodySmall"}>{x.co_store?.id}</Text>
                                </Stack>
                            </HStack>
                            <HStack spacing={8} items={"center"}>
                                <AmountView amount={x.balance}/>
                                <Checkbox
                                    disabled={!(x.balance > amount && x.items?.map(t => t.item?.id).includes(itemId))}
                                    onPress={() => setStoreId(x.co_store?.id)}
                                    status={x.co_store?.id === storeId ? "checked" : "unchecked"}/>
                            </HStack>

                        </HStack>
                    }}
                />

                <HStack items={'center'} spacing={10}>
                    <Button mode="contained-tonal" style={{flex: 1}} onPress={() => {
                        setVisible(false);
                    }}>取消</Button>
                    <Button
                        disabled={!storeId}
                        onPress={() => {
                            onConfirm?.(storeId)
                            setVisible(false)
                        }}
                        mode="contained"
                        style={{flex: 2}}>确认</Button>
                </HStack>

            </Stack>
        </Modal>
    </HStack>
}

export default SwitchTrigger