import React, {useEffect, useState} from "react";
import {fetchAuthStatus, fetchStoreInfo, login, sendCode} from "../service/api";
import AsyncStorage from "@react-native-async-storage/async-storage";
import {useRequest} from "ahooks";


const useStoreState = () => {

    const {data, loading, run, runAsync} = useRequest(fetchStoreInfo, {manual: true})

    const [storeValid, setStoreValid] = useState<boolean>(false)

    useEffect(() => {
        refresh()
    }, [])

    const checkStatus = async () => {

        const store = await AsyncStorage.getItem('store')
        if (!store) {
            setStoreValid(false)
        } else {
            runAsync(store).then(rsp => {
                setStoreValid(rsp.data?.code === 0)
            })
        }
    }

    const refresh = () => {
        AsyncStorage.getItem('store').then(storeId => {
            if (storeId) {
                runAsync(storeId).then(rsp => {
                })
            }
        })
    }

    const fetch = () => {
        AsyncStorage.getItem('store').then(storeId => {
            if (storeId) {
                runAsync(storeId).then(async rsp => {
                })
            }
        }).catch(e => {
            console.error(e)
        })
    }


    const update = (storeId) => {
        AsyncStorage.setItem('store', storeId).then(() => {
            runAsync(storeId).then(rsp => {
            })
        })
    }

    return {
        store: data?.data?.data, loading, fetch, refresh, update, storeValid, checkStatus
    }
}

export default useStoreState