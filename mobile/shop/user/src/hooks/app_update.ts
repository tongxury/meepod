import {useEffect, useState} from "react";
import {fetchSettings} from "../service/api";
import {checkForUpdateAsync, fetchUpdateAsync, reloadAsync} from "expo-updates";
import {Toast} from "@ant-design/react-native";
import {Platform} from "react-native";

function useAppUpdates() {
    const [update, setUpdate] = useState<boolean>(false)


    const checkToUpdate = () => {
        checkForUpdateAsync().then(rsp => {
            if (rsp.isAvailable) {
                fetchUpdateAsync().then(rsp => {
                    reloadAsync().then(() => {
                        Toast.info('已更新为最新版本')
                    })
                })
            } else {
                Toast.info('已是最新版本')
            }
        });
    }
    const checkForUpdates = () => {
        if (Platform.OS !== "web") {
            checkForUpdateAsync().then(rsp => {
                setUpdate(rsp.isAvailable)
                // if (rsp.isAvailable) {
                //     fetchUpdateAsync().then(rsp => {
                //         reloadAsync().then(() => {
                //             Toast.info('已更新为最新版本')
                //         })
                //     })
                // } else {
                //     Toast.info('已是最新版本')
                // }
            });
        }
    }

    return {
        update,
        checkForUpdates,
        checkToUpdate,
    }
}

export default useAppUpdates