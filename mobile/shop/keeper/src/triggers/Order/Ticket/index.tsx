import {ActivityIndicator, FlatList, Pressable, StyleProp, View, ViewStyle} from "react-native";
import React, {useContext, useState} from "react";
import {Avatar, Button, Dialog, Portal, RadioButton, Text} from "react-native-paper";
import {AppContext} from "../../../providers/global";
import {ImagePicker, Modal, WhiteSpace, WingBlank} from "@ant-design/react-native";
import {Image} from "@rneui/themed";
import {base64ToFile} from "../../../utils";
import {v4 as uuidv4} from 'uuid';
import {Order, Asset} from "../../../service/typs";
import UploaderTrigger from "../../Uploader";
import {HStack} from "@react-native-material/core";

const TicketTrigger = ({needUpload, onConfirm, style}: {
    needUpload: boolean,
    onConfirm?: (images: string[]) => void,
    style?: StyleProp<ViewStyle> | undefined
}) => {

    const [visible, setVisible] = useState(false);

    const {settingsState: {settings: appSettings}} = useContext<any>(AppContext);

    const [tickets, setTickets] = useState<Asset[]>([])

    return <View style={style}>
        <Button onPress={() => setVisible(true)} mode="contained">出票</Button>
        <Modal
            style={{borderTopLeftRadius: 10, borderTopRightRadius: 10}}
            popup
            visible={visible}
            maskClosable={true}
            animationType="slide-up"
            onClose={() => {
                setVisible(false)
            }}>
            <View style={{alignItems: "center", margin: 20}}>
                <Text variant="titleMedium">上传票样</Text>
                <WhiteSpace size="sm"/>
                <Text variant="bodySmall">
                    上传票样请注意遮挡票号等重要信息，最多可上传{appSettings.maxUploadTicket}张
                </Text>
                {
                    needUpload && <View>
                        <WhiteSpace size="sm"/>
                        <Text style={{color: 'red'}} variant="bodySmall">
                            用户要求上传票样
                        </Text>
                    </View>
                }

            </View>

            <FlatList
                // style={{flex: 1}}
                // scrollEnabled={true}
                style={{margin: 10}}
                // ListEmptyComponent={<Empty/>}
                data={tickets}
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

            {
                tickets.length < appSettings.maxUploadTicket &&

                <UploaderTrigger onUploaded={assets => setTickets(tickets.concat(assets))}>
                    <Button>点击上传票样</Button>
                </UploaderTrigger>
            }

            <View style={{flexDirection: "row", alignItems: "center", padding: 15}}>
                <Button mode="contained-tonal" style={{flex: 1}} onPress={() => {
                    setVisible(false);
                }}>取消</Button>
                <WingBlank size="sm"/>
                <Button disabled={needUpload && tickets.length == 0} onPress={() => {
                    onConfirm(tickets.map(t => t.key))
                    setVisible(false)
                }} mode="contained"
                        style={{flex: 2}}>确认</Button>
            </View>
        </Modal>
    </View>
}

export default TicketTrigger