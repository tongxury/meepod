import React, {useContext, useState} from "react";
import {Pressable, StyleProp, Text, View, ViewStyle} from "react-native";
import {Controller, useForm} from "react-hook-form";
import {InputItem, Modal, Toast, WhiteSpace} from "@ant-design/react-native";
import {Button, Checkbox, Chip, useTheme} from "react-native-paper";
import {ButtonGroup} from "@rneui/themed";
import {addOrderGroupOrder} from "../../../service/api";
import {Stepper} from "../../../components/Stepper";

const JoinGroupTrigger = ({data, onSubmitted, children, style}: {
    data: { groupId: string, volumeLeft: number, floor: number }
    onSubmitted?: (orderId: string, groupId: string) => void,
    children: React.ReactNode,
    style?: StyleProp<ViewStyle>
}) => {

    const [open, setOpen] = useState(false)
    const onConfirm = ({volume}: { volume: number }) => {

        console.log(volume, data)

        addOrderGroupOrder({groupId: data.groupId, volume}).then(rsp => {
            if (rsp.data?.data) {
                setOpen(false)
                onSubmitted?.(rsp.data?.data, data.groupId)
            }
        })
    }
    const onSubmitOrder = (planId: string, settings) => {
    }

    const {
        control,
        setValue,
        handleSubmit,
        formState: {errors}
    } = useForm({
        defaultValues: {
            volume: 1,

        }
    });

    const {colors} = useTheme()

    return <Pressable style={style} onPress={() => setOpen(true)}>
        {children}
        <Modal
            popup
            visible={open}
            maskClosable
            animationType="slide-up"
            bodyStyle={{padding: 18, paddingVertical: 25}}
            onClose={() => setOpen(false)}>

            <View>
                <Controller
                    control={control}
                    rules={{required: false,}}
                    render={({field: {value}}) => {

                        const options = [
                            // {value: 5, label: "5份"},
                            // {value: 10, label: "10份"},
                            {value: Math.ceil(data.volumeLeft / 2), label: "半包"},
                            {value: data.volumeLeft, label: "全包"},
                        ]

                        const indexTap = new Map(options.map((t, i) => [t.value, i]))

                        return <View>
                            <View style={{flexDirection: "row", alignItems: "center", justifyContent: "space-between"}}>
                                <Text>参与份数</Text>
                                <Stepper
                                    color={colors.primary}
                                    max={data.volumeLeft}
                                    min={1}
                                    value={value}
                                    onChange={value => {
                                        setValue("volume", value)
                                    }}
                                />
                                <Text>剩余{data.volumeLeft - value}份</Text>
                            </View>
                            <WhiteSpace size="xl"/>
                            <ButtonGroup
                                buttons={options.map(t => t.label)}
                                selectedIndex={indexTap.get(value)}
                                onPress={(value) => {
                                    console.log(options[value])
                                    setValue("volume", options[value].value)
                                }}
                                containerStyle={{marginBottom: 20}}
                            />
                        </View>

                    }}
                    name="volume"
                />

                <WhiteSpace/>
                <Button mode="contained" onPress={handleSubmit(onConfirm)}>确认</Button>
            </View>
        </Modal>
    </Pressable>
}


export default JoinGroupTrigger