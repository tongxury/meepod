import {Platform, Pressable, StyleProp, View, ViewStyle} from "react-native";
import * as ImagePicker from "expo-image-picker";
import {MediaTypeOptions} from "expo-image-picker";
import * as FileSystem from "expo-file-system";
import {uploadImageV2} from "../../service/api";
import {v4 as uuidv4} from 'uuid';
import React from "react";
import {Asset} from "../../service/typs";

const UploaderTrigger = ({children, onUploaded, style}: {
    children: React.ReactNode,
    onUploaded?: (assets: Asset[]) => void,
    style?: StyleProp<ViewStyle>
}) => {

    const onUpload = () => {

        ImagePicker.launchImageLibraryAsync({mediaTypes: MediaTypeOptions.Images,}).then(response => {

            if (Platform.OS == 'web') {

                const b64 = response.assets[0].uri

                uploadImageV2({files: [{name: uuidv4(), src: b64.split(',')[1]}]}).then(rsp => {
                    if (rsp.data?.data?.files?.length > 0) {
                        onUploaded?.(rsp.data?.data?.files)
                    }
                })
            } else {
                FileSystem.readAsStringAsync(response.assets[0].uri, {encoding: "base64"}).then(file => {
                    uploadImageV2({files: [{name: uuidv4(), src: file}]}).then(rsp => {
                        if (rsp.data?.data?.files?.length > 0) {
                            onUploaded?.(rsp.data?.data?.files)
                        }
                    })
                })
            }

        })

    }

    return <View style={style}>
        <Pressable onPress={() => {
            console.log('onUpload');
            onUpload()
        }}>{children}</Pressable>
    </View>
}

export default UploaderTrigger