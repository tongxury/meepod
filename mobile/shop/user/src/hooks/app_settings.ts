import {useEffect, useState} from "react";
import {fetchSettings} from "../service/api";

function useAppSettings() {
    const defaultSetting = {
        "ssqBlueLimit": 16,
        "ssqRedLimit": 16,
        "ssqDRedLimit": 5,
        "dltBlueLimit": 16,
        "dltRedLimit": 16,
        "dltDRedLimit": 4,
        "rx9DanLimit": 5,
        "minUnionAmount": 30,
        "maxVolume": 150,
        "maxMultiple": 50000,
    }

    const [settings, setSettings] = useState<any>()

    const update = () => {

        fetchSettings().then(rsp => {
            if (rsp?.data?.data) {
                setSettings(rsp?.data?.data)
            } else {
                setSettings(defaultSetting)
            }
        }).catch(() => {
            setSettings(defaultSetting)
        })
    }

    return {
        settings,
        update
    }
}

export default useAppSettings