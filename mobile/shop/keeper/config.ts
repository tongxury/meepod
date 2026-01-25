import * as Updates from 'expo-updates';

import Constants from "expo-constants";

export const Config = () => {

    const version = "1.1.0"

    const devConfig = {
        env: 'development',
        version: version,
        // apiUrl: 'http://192.168.1.102:8080',
        // apiPaymentUrl: 'http://192.168.1.102:6066',
        apiUrl: 'https://x.lansess.com',
        apiPaymentUrl: 'https://x.lansess.com',
    }

    const testConfig = {
        env: 'test',
        version: version,
        apiUrl: 'https://x.lansess.com',
    }

    const productionConfig = {
        env: 'production',
        version: version,
        // apiUrl: 'https://api.lottery.local',
        apiUrl: 'https://x.lansess.com',
        apiPaymentUrl: 'https://x.lansess.com',

    }

    const configs = {
        'development': devConfig,
        'test': testConfig,
        'production': productionConfig
    }

    if (Updates.channel) {
        return configs[Updates.channel] ?? configs['production']
    } else {
        return configs[Constants.expoConfig.extra.env] ?? configs['production']
    }
}
