import {createMd5} from "../../../../utils";
import {Ticket} from "../../types";

const useHook = () => {

    const options = [...Array(10).keys()].map(t => {
        return {value: `${t}`, label: `${t}`}
    })
    const lastOptions = [...Array(15).keys()].map(t => {
        return {value: `${t}`, label: `${t}`}
    })

    const getRandom = (count: number = 1): string[] => {
        return [...options].sort(() => Math.random() - 0.5).slice(0, count).map(t => t.value)
    }
    const getRandomLast = (count: number = 1): string[] => {
        return [...lastOptions].sort(() => Math.random() - 0.5).slice(0, count).map(t => t.value)
    }

    const createRandom = (): Ticket => {

        const gen = getRandom()
        const shi = getRandom()
        const bai = getRandom()
        const qian = getRandom()
        const wan = getRandom()
        const swan = getRandom()
        const last = getRandomLast()

        const vl = {
            key: '',
            swan,
            wan,
            qian,
            bai,
            shi,
            gen,
            last,
            amount: 2
        }

        vl.key = createMd5(JSON.stringify(vl))

        return vl


    }


    return {
        options,
        lastOptions,
        getRandom,
        createRandom,
    }
}
export default useHook