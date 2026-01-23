import {useRequest} from "ahooks";
import {fetchItemStates as query} from "../service/api";
import {useState} from "react";

const useItemState = () => {

    const {loading, runAsync} = useRequest(query, {manual: true})
    const [items, setItems] = useState([])

    const fetch = () => {
        runAsync().then(rsp => {
            if (rsp?.data?.data) {
                setItems(rsp?.data?.data)
            }
        })
    }

    return {
        supported: ['ssq', 'f3d', 'x7c', 'rx9', 'sfc', 'zjc'],
        loading,
        items: items || [],
        fetch
    }
}

export default useItemState