import * as dotenv from 'dotenv' // see https://github.com/motdotla/dotenv#how-do-i-use-dotenv-with-import
import { ExpoConfig, ConfigContext } from 'expo/config';

dotenv.config()
export default ({ config }: ConfigContext): ExpoConfig => ({
    ...config,
    slug: 'shopuser',
    name: '福彩',
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
            projectId: "0c3b0d85-391f-4fbd-b2b0-c6d61f3ae755"
        }
    },
    scheme: "lotteryapp",
    version: "1.0.0",
    android: {
        versionCode: 5,
        package: "com.lottery.app",
    },
    ios: {
        supportsTablet: true,
        distribution: 'store',
        bundleIdentifier: "com.lottery.client",
        config: {

        }
    },
    platforms: [
        "android",
        "ios",
        "web"
    ],
    updates: {
        url: "https://u.expo.dev/0c3b0d85-391f-4fbd-b2b0-c6d61f3ae755"
    },
    runtimeVersion: {
        policy: "sdkVersion"
    },
    plugins: [
        "expo-localization"
    ]
});
