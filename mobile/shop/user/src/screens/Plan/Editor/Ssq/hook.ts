import {Ticket} from "../../types";

const useHook = () => {

    const redOptions = [...Array(33).keys()].map(t => t + 1).map(t => {
        const v = t <= 9 ? `0${t}` : `${t}`
        return {value: v, label: v}
    })
    const blueOptions = [...Array(16).keys()].map(t => t + 1).map(t => {
        const v = t <= 9 ? `0${t}` : `${t}`
        return {value: v, label: v}
    })

    const genKey = (plan: Ticket): string => {
        return `${(plan.red || []).sort().join(",")}-${(plan.redD || []).sort().join(",")}-${(plan.blue || []).sort().join(",")}`
    }


    const getRandom = (): Ticket => {
        const randomRed = [...redOptions].sort(() => Math.random() - 0.5).slice(0, 6)
        const randomBlue = [...blueOptions].sort(() => Math.random() - 0.5).slice(0, 1)

        const red = randomRed.map(t => t.value)
        const redD = []
        const blue = randomBlue.map(t => t.value)

        const plan = {
            key: '',
            red,
            redD,
            blue,
            amount: 2,
        }
        plan.key = genKey(plan)
        return plan
    }

    return {
        redOptions, blueOptions,
        getRandom, genKey
    }
}
export default useHook