import {useEffect, useState} from "react";
import {fetchSettings, listCounters} from "../service/api";
import {Counter} from "../service/typs";

function useCounter() {

    const [counter, setCounter] = useState<Counter>()

    const update = () => {

        listCounters().then(rsp => {
            if (rsp?.data?.data) {
                setCounter(rsp?.data?.data)
            }
        })
    }

    return {
        counter,
        update
    }
}

export default useCounter