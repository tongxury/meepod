import {SplashLoader} from "./src/providers/splash";
import ThemeProvider from "./src/providers/theme";
import Main from "./src/main";
import {SafeAreaProvider} from "react-native-safe-area-context";
import {AppProvider} from "./src/providers/global";
import * as Updates from 'expo-updates';

export default function App() {
    return (
        <SafeAreaProvider>
            <AppProvider>
                <SplashLoader>
                    <ThemeProvider>
                        <Main/>
                    </ThemeProvider>
                </SplashLoader>
            </AppProvider>
        </SafeAreaProvider>
    );
}
