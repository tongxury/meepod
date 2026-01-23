import editor from "../../Profile/Editor";
import {InputItem, Modal} from "@ant-design/react-native";
import {HStack, Stack} from "@react-native-material/core";
import {Button as RneButton, Image, Input} from "@rneui/themed";
import {Button, Text, useTheme} from "react-native-paper";
import {StyleProp, View, ViewStyle, TextInput} from "react-native";
import React, {useEffect, useState} from "react";
import {useForm, Controller} from "react-hook-form";
import UserView from "../../../components/UserView";
import {useRequest} from "ahooks";
import {addCoStore, fetchStore} from "../../../service/api";
import IconText from "../../../components/IconText";
import ItemSelector from "../../../components/ItemSelector";
import {AntDesign as AntDesignIcon} from "@expo/vector-icons";


const CreateTrigger = ({onConfirm, style}: { onConfirm?: () => void, style?: StyleProp<ViewStyle> }) => {

    const [open, setOpen] = useState<boolean>(false)

    const [coStoreId, setCoStoreId] = useState<string>()
    const [itemIds, setItemIds] = useState<string[]>()

    const {data, run} = useRequest(fetchStore, {manual: true})

    const coStore = data?.data?.data


    useEffect(() => {
        if (coStoreId) {
            run({id: coStoreId})
        }
    }, [coStoreId])

    const confirm = () => {
        addCoStore({coStoreId, itemIds}).then(rsp => {
            if (rsp.data?.code === 0) {
                setOpen(false)
                onConfirm?.()
            }
        })
    }


    const {colors} = useTheme()

    return <View style={style}>
        <RneButton icon={<AntDesignIcon name={'plus'} color={colors.onPrimary}/>}
                   onPress={() => setOpen(true)}>添加合作店铺</RneButton>
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
                        <Text variant={"titleMedium"}>店铺邀请码</Text>
                        <TextInput
                            value={coStoreId}
                            onChangeText={text => {
                                setCoStoreId(text)
                            }}
                            style={{
                                minHeight: 40,
                                flex: 1,
                                paddingHorizontal: 10,
                                backgroundColor: colors.primaryContainer,
                                borderRadius: 5
                            }}
                            placeholder={"请输入店铺邀请码"}/>
                        {coStore && <IconText icon={coStore?.icon} title={coStore?.name}/>}
                    </Stack>

                    {coStore &&
                        <Stack spacing={5} fill={1}>
                            <Text variant={"titleMedium"}>合作彩种</Text>

                            <ItemSelector values={itemIds} onChange={values => setItemIds(values)}
                                          items={coStore?.items}/>

                        </Stack>}


                </Stack>


                <Button mode="contained" disabled={!coStore} onPress={confirm}>确认</Button>
            </Stack>
        </Modal>
    </View>

}

export default CreateTrigger