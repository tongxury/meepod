import useCountdown from "ahooks/lib/useCountDown";
import { HStack } from "@react-native-material/core";
import { Text } from "react-native-paper";

const CountDownTimer = ({ timeLeft, color }: { timeLeft?: number, color?: string }) => {

    const [n, { days, hours, minutes, seconds }] = useCountdown({ leftTime: timeLeft })


    return <HStack items="center" spacing={2}>
        {days >= 2 && <Text style={{ fontWeight: "bold", color: color }}>{days}天</Text>}
        <Text style={{ fontWeight: "bold", color: color }}>{days >= 2 ? (hours < 10 ? `0${hours}` : hours) : (days * 24 + hours)}</Text>
        <Text style={{ fontWeight: "bold", color: color }}>:</Text>
        <Text style={{ fontWeight: "bold", color: color }}>{minutes < 10 ? `0${minutes}` : minutes}</Text>
        <Text style={{ fontWeight: "bold", color: color }}>:</Text>
        <Text style={{ fontWeight: "bold", color: color }}>{seconds < 10 ? `0${seconds}` : seconds}</Text>
    </HStack>

}

export default CountDownTimer