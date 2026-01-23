import {createContext, useState} from "react";
import useStoreState from "../hooks/store";
import useAuthState from "../hooks/auth";
// import useItemState from "../hooks/item";
import useAppSettings from "../hooks/app_settings";
import useAppUpdates from "../hooks/app_update";

export const AppContext = createContext({});
export const AppProvider = ({children}) => {

    const storeState = useStoreState()
    const authState = useAuthState()
    // const itemState = useItemState()
    const settingsState = useAppSettings()
    const updateState = useAppUpdates()


    return <AppContext.Provider value={{ storeState,authState, settingsState, updateState}}>
        {children}
    </AppContext.Provider>
}
