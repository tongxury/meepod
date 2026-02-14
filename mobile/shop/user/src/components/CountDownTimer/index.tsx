import useCountdown from "ahooks/lib/useCountDown";
import { HStack } from "@react-native-material/core";
import { Text } from "react-native-paper";
import { StyleProp, ViewStyle } from "react-native";

const CountDownTimer = ({ timeLeftSecond, color, style }: { timeLeftSecond?: number, color?: string, style?: StyleProp<ViewStyle> }) => {

    const [n, { days, hours, minutes, seconds }] = useCountdown({ leftTime: timeLeftSecond * 1000 })


    return <HStack items="center" spacing={2} style={style}>
        {days >= 2 && <Text style={{ fontWeight: "bold", color: color }} variant={"titleSmall"}>{days}天</Text>}
        <Text style={{ fontWeight: "bold", color: color }} variant={"titleSmall"}>{days >= 2 ? (hours < 10 ? `0${hours}` : hours) : (days * 24 + hours)}</Text>
        <Text style={{ fontWeight: "bold", color: color }} variant={"titleSmall"}>:</Text>
        <Text style={{ fontWeight: "bold", color: color }} variant={"titleSmall"}>{minutes < 10 ? `0${minutes}` : minutes}</Text>
        <Text style={{ fontWeight: "bold", color: color }} variant={"titleSmall"}>:</Text>
        <Text style={{ fontWeight: "bold", color: color }} variant={"titleSmall"}>{seconds < 10 ? `0${seconds}` : seconds}</Text>
    </HStack>

}

export default CountDownTimer