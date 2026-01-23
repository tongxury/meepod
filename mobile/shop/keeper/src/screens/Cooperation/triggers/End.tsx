import {ActivityIndicator, FlatList, Pressable, StyleProp, View, ViewStyle} from "react-native";
import React, {useContext, useState} from "react";
import {Avatar, Button, Dialog, Portal, RadioButton, Text} from "react-native-paper";
import {AppContext} from "../../../providers/global";
import {ImagePicker, InputItem, Modal, TextareaItem, WhiteSpace, WingBlank} from "@ant-design/react-native";
import {Stack} from "@react-native-material/core";
import {Asset, Withdraw} from "../../../service/typs";
import {acceptWithdraw, applyEndCoStore, endCoStore} from "../../../service/api";
import {Button as RneButton, Image} from "@rneui/themed"
import UploaderTrigger from "../../../triggers/Uploader";


const EndTrigger = ({coStoreId, onConfirm, style}: {
    coStoreId: string,
    onConfirm?: () => void,
    style?: StyleProp<ViewStyle>
}) => {

    const [visible, setVisible] = useState(false);
    const [images, setImages] = useState<Asset[]>([])

    const confirm = () => {
        endCoStore(coStoreId, images?.[0].url).then(rsp => {
            setVisible(false)
            onConfirm?.()
        })
    }


    const {settingsState: {settings: appSettings}} = useContext<any>(AppContext);

    return <View style={style}>
        <RneButton size={'sm'} onPress={() => setVisible(true)}>确认终止</RneButton>
        <Modal
            style={{borderTopLeftRadius: 10, borderTopRightRadius: 10}}
            popup
            visible={visible}
            maskClosable={true}
            animationType="slide-up"
            onClose={() => {
                setVisible(false)
            }}>
            <Stack p={20} spacing={10}>

                <Stack center={true} spacing={8}>
                    <Text variant={"titleMedium"}>上传全额退款凭证</Text>
                    <Text variant={"bodySmall"} style={{color: 'red'}}>请正确上传支付凭证截图</Text>
                </Stack>

                <FlatList
                    // style={{flex: 1}}
                    // scrollEnabled={true}
                    style={{margin: 10}}
                    // ListEmptyComponent={<Empty/>}
                    data={images}
                    numColumns={2}
                    keyExtractor={(e, i) => i.toString()}
                    // ItemSeparatorComponent={() => <View style={{height: 10, width: '100%'}}/>}
                    renderItem={({item: x}) => {
                        return <Image
                            style={{margin: 2}}
                            source={{uri: x.url}}
                            containerStyle={{
                                aspectRatio: 1,
                                width: '100%',
                                flex: 1,
                            }}
                            PlaceholderContent={<ActivityIndicator/>}
                        />
                    }}
                />

                <UploaderTrigger onUploaded={assets => setImages(assets)}>
                    <Button>点击上传</Button>
                </UploaderTrigger>

                <View style={{flexDirection: "row", alignItems: "center"}}>
                    <Button mode="contained" style={{flex: 1}} onPress={() => setVisible(false)}>取消</Button>
                    <WingBlank size="sm"/>
                    <Button disabled={images?.length == 0} onPress={confirm} mode="contained-tonal"
                            style={{flex: 2}}>确认</Button>
                </View>
            </Stack>

        </Modal>

    </View>

}

export default EndTrigger

