import {Ticket} from "../../types";

const useHook = () => {

    const options = [...Array(10).keys()].map(t => {
        return {value: `${t}`, label: `${t}`}
    })

    const getRandom = (count: number = 1): string[] => {
        return [...options].sort(() => Math.random() - 0.5).slice(0, count).map(t => t.value)
    }

    const createRandom = (cat: 'z1'| 'z3'| 'z6'): Ticket => {

        if (cat === 'z1') {
            const gen = getRandom()
            const shi = getRandom()
            const bai = getRandom()

            return {
                cat: 'z1',
                key: bai.join('') + shi.join('') + gen.join(''),
                bai,
                shi,
                gen,
                amount: 2
            }
        } else if (cat === 'z3') {
            const ton = getRandom(2)

            return {
                cat: 'z3',
                key: ton.join(''),
                ton,
                amount: 4
            }
        } else if (cat === 'z6') {
            const ton = getRandom(3)

            return {
                cat: 'z6',
                key: ton.join(''),
                ton,
                amount: 2
            }
        }


    }


    return {
        options,
        getRandom,
        createRandom,
    }
}
export default useHook