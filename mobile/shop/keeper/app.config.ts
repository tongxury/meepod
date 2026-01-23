import * as dotenv from 'dotenv' // see https://github.com/motdotla/dotenv#how-do-i-use-dotenv-with-import
import { ExpoConfig, ConfigContext } from 'expo/config';

dotenv.config()

export default ({ config }: ConfigContext): ExpoConfig => ({
    ...config,
    slug: 'shopkeeper',
    name: '福彩-店主端',
    icon: './src/assets/icon.png',
    // splash: {
    //     image: './src/assets/splash.png',
    //     resizeMode: "contain",
    //     // backgroundColor: "#ffffff",
    // },
    orientation: "portrait",
    extra: {
        env: process.env.ENV,
        eas: {
            projectId: "c36d8622-6cdc-4454-a84a-6de87a45e337"
        }
    },
    scheme: "lotterykeeper",
    version: "2.0.0",
    android: {
        versionCode: 5,
        package: "com.lottery.keeper",
    },
    // web: {
    //     "bundler": "metro"
    // },
    platforms: [
        "android",
        "web"
    ],
    updates: {
        "url": "https://u.expo.dev/c36d8622-6cdc-4454-a84a-6de87a45e337"
    },
    runtimeVersion: {
        policy: "sdkVersion"
    },
    plugins: [
        "expo-localization",
        [
            "expo-updates",
            {
                "username": "lottery"
            }
        ]
    ]

});
