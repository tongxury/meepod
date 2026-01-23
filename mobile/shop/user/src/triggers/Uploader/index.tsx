import {Platform, Pressable, View} from "react-native";
import * as ImagePicker from "expo-image-picker";
import {MediaTypeOptions} from "expo-image-picker";
import * as FileSystem from "expo-file-system";
import {updateUser, uploadImageV2} from "../../service/api";
import {v4 as uuidv4} from 'uuid';
import React from "react";


const UploaderTrigger = ({children, onUploaded}: {
    children: React.ReactNode,
    onUploaded?: (keys: string[]) => void
}) => {

    const onUpload = () => {

        ImagePicker.launchImageLibraryAsync({mediaTypes: MediaTypeOptions.Images,}).then(response => {

            if (Platform.OS == 'web') {

                const b64 = response.assets[0].uri

                uploadImageV2({files: [{name: uuidv4(), src: b64.split(',')[1]}]}).then(rsp => {
                    if (rsp.data?.data?.files?.length > 0) {
                        onUploaded?.(rsp.data?.data?.files?.map(t => t.key))
                    }
                })
            } else {
                FileSystem.readAsStringAsync(response.assets[0].uri, {encoding: "base64"}).then(file => {
                    uploadImageV2({files: [{name: uuidv4(), src: file}]}).then(rsp => {
                        if (rsp.data?.data?.files?.length > 0) {
                            onUploaded?.(rsp.data?.data?.files?.map(t => t.key))
                        }
                    })
                })
            }

        })

    }

    return <View>
        <Pressable onPress={() => {
            console.log('onUpload');
            onUpload()
        }}>{children}</Pressable>
    </View>
}

export default UploaderTrigger