import AsyncStorage, {useAsyncStorage} from '@react-native-async-storage/async-storage';
import {useEffect, useState} from "react";
import {useLocalStorageState} from "ahooks";

export function useStorageState<S>(key: string, initialValue?: S) {


    const [state, setState] = useState<S>(initialValue)

    useEffect(() => {
        AsyncStorage.getItem(key).then(rsp => {
            if (rsp) {
                setState(JSON.parse(rsp))
            }
        })
    })

    const setValue = (value: S) => {
        AsyncStorage.setItem(key, JSON.stringify(value)).then(() => {
            setState(value)
        })
    }

    return [state, setValue]
}

export function useListStorageStateV2<S>(key: string) {
    const {getItem, setItem, removeItem, mergeItem} = useAsyncStorage(key)

    const [value, setValue] = useState<S[]>()

    useEffect(() => {
        getItem().then(rsp => {
            if (rsp) {
                setValue(JSON.parse(rsp))
            }
        })
    }, [])

    const replaceAll = (items: S[]) => {

        setItem(JSON.stringify(items)).then(rsp => {
            setValue(items)
        })
    }


    return {
        value, replaceAll
    }
}

export function useListStorageState<S>(key: string, initialValue?: S[]) {
    const [value, setValue] = useLocalStorageState<S[]>(key, {
        defaultValue: initialValue || [],
    });

    const append = (items: S[]) => {
        const newValues = (value ?? []).concat(items)
        setValue(newValues)
    }

    const replaceAll = (items: S[]) => {
        setValue(items)
    }


    const clear = () => {
        setValue([])
    }

    return {
        value,
        append, clear, replaceAll
    }

}

