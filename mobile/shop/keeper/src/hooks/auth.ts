import React, {useEffect, useState} from "react";
import {fetchAuthStatus, login, loginByPassword, sendCode} from "../service/api";
import AsyncStorage from "@react-native-async-storage/async-storage";

const useAuthState = () => {

    const [authToken, setAuthToken] = useState<string>()

    // const {runAsync: fetch} = useRequest(fetchAuthStatus, {manual: true,  throttleWait: 500})

    useEffect(() => {
        const fetchItem = async () => {
            return await AsyncStorage.getItem('user_auth')
        }

        fetchItem().then(setAuthToken)

    }, [])


    const loginByPassword_ = ({phone, password}: { phone: string, password: string }, onLoginSuccess?: () => void) => {
        loginByPassword({phone, password}).then(rsp => {
            if (rsp.data?.data) {
                AsyncStorage.setItem('user_auth', rsp.data.data).then(() => {
                    setAuthToken(rsp.data.data)
                    onLoginSuccess?.()
                })
            }
        })
    }


    const login_ = ({phone, code}: { phone: string, code: string }, onLoginSuccess?: (token: string) => void) => {
        login({phone, code}).then(rsp => {
            if (rsp.data?.data) {
                AsyncStorage.setItem('user_auth', rsp.data.data).then(() => {
                    setAuthToken(rsp.data.data)
                    onLoginSuccess?.(rsp.data.data)
                })
            }
        })
    }

    const logout = (onLogoutSuccess?: () => void) => {

        AsyncStorage.removeItem('user_auth').then(() => {
            onLogoutSuccess?.()
        })
    }

    const sendCode_ = ({phone, event, onSuccess}: { phone: string, event?: string, onSuccess?: () => void }) => {
        sendCode({phone, event}).then(rsp => {
            if (rsp.data.code === 0) {
                onSuccess?.()
            }
        })
    }

    // const isTokenValid = async (): Promise<boolean> => {
    //     const token = await AsyncStorage.getItem('user_auth')
    //     if (!token) {
    //         return new Promise(resolve => resolve(false))
    //     } else {
    //         const rsp = await fetchAuthStatus({token: token})
    //         return new Promise(resolve => resolve(rsp?.data?.code === 0))
    //     }
    // }


    // const authed = async (): Promise<boolean> => {
    //     const token = await AsyncStorage.getItem('user_auth')
    //     return new Promise(resolve => resolve(!!token))
    // }


    return {
        authToken, logout,
        doLogin: login_, sendCode: sendCode_, loginByPassword: loginByPassword_
    }
}

export default useAuthState