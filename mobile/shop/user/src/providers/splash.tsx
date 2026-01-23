import {Image, View} from "react-native";
import {useContext, useEffect, useState} from "react";
import {AppContext} from "./global";
import {Stack} from "@react-native-material/core";

export function SplashLoader({children}) {

    const {settingsState: {settings, update: updateSettings}} = useContext<any>(AppContext);
    const [ready, setReady] = useState<boolean>(false)

    useEffect(() => {
        updateSettings()
    }, [])

    useEffect(() => {

        if (settings) {
            setTimeout(() => {
                setReady(true)
            }, 1500)
        }

    }, [settings])


    const SplashView = () => {
        return <Stack fill={true} center={true}>
            <Image style={{width: 200, height: 135}} source={require("../assets/splash.png")}/>
        </Stack>
    }


    return (
        <View style={{flex: 1}}>
            {ready ? children : <SplashView/>}
        </View>
    );
}
