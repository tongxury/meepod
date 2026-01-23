import useCountdown from "ahooks/lib/useCountDown";
import {HStack} from "@react-native-material/core";
import {Text} from "react-native-paper";

const CountDownTimer = ({timeLeft, color}: { timeLeft?: number, color?: string }) => {

    const [n, {days, hours, minutes, seconds}] = useCountdown({leftTime: timeLeft})


    return <HStack items="center" spacing={2}>
        <Text style={{fontWeight: "bold", color: color}}>{days * 24 + hours}</Text>
        <Text style={{fontWeight: "bold", color: color}}>:</Text>
        <Text style={{fontWeight: "bold", color: color}}>{minutes}</Text>
        <Text style={{fontWeight: "bold", color: color}}>:</Text>
        <Text style={{fontWeight: "bold", color: color}}>{seconds}</Text>
    </HStack>

}

export default CountDownTimer