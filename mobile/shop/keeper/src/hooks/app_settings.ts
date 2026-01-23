import { useEffect, useState } from "react";
import { fetchSettings } from "../service/api";

function useAppSettings() {
    const defaultSetting = {
        "showCoStore": false,
        "rejectReasons": [
            { id: '1', text: '联系不上客户' }
        ],
        "rejectWithdrawReasons": [
            "用户异常"
        ],
        "maxUploadTicket": 4,
        "service": {
            "email": "lottery@example.com",
        },
        "filters": {
            "month": [],
            "proxyRewardCats": [],
            "coStorePaymentCats": []
        },
        "proxy": {
            "maxRewardRate": 0.07,
            "rewardRateItems": [
                { label: "1%", value: "0.01" },
                { label: "2%", value: "0.02" },
                { label: "3%", value: "0.03" },
                { label: "4%", value: "0.04" },
                { label: "5%", value: "0.05" },
                { label: "6%", value: "0.06" },
                { label: "7%", value: "0.07" },
            ],
        }
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