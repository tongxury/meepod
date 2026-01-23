import {createContext} from "react";
import useAuthState from "../hooks/auth";
import useAppSettings from "../hooks/app_settings";
import useCounter from "../hooks/counters";
import useAppUpdates from "../hooks/app_update";

export const AppContext = createContext({});
export const AppProvider = ({children}) => {

    const authState = useAuthState()
    const settingsState = useAppSettings()
    const counterState = useCounter()
    const updateState = useAppUpdates()

    return <>
        <AppContext.Provider value={{ authState, settingsState, counterState, updateState}}>
            {children}
        </AppContext.Provider>
    </>
}
