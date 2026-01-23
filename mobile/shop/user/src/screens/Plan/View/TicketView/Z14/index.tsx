import {useTheme, Text} from "react-native-paper";
import {HStack, Stack} from "@react-native-material/core";
import {Button, Chip} from "@rneui/themed";


const Z14TicketView = ({data}) => {
    const {colors} = useTheme()

    return <Stack fill={1} spacing={10}>
        <Text>共选择
            <Text style={{color: colors.primary, fontWeight: "bold"}}> {Object.keys(data?.options ?? {}).length} </Text>场
            <Text style={{color: colors.primary, fontWeight: "bold"}}>  {data.amount} </Text>元
        </Text>

        {Object.keys(data?.options ?? {}).map(matchId => {

            const matchOptions = data?.options?.[matchId]?.odds
            const isDan = data?.options[matchId]?.dan

            return <HStack spacing={5} key={matchId}>
                <HStack items={"center"} spacing={10}>
                    <Text style={{
                        fontWeight: "bold",
                        color: isDan ? colors.primary : colors.onBackground
                    }}>主队: {data?.matches?.[matchId]?.home_team}</Text>
                </HStack>
                {Object.keys(matchOptions ?? {}).sort().map(cat => {
                    return <HStack items={"center"} wrap={"wrap"} key={cat}>
                        {matchOptions?.[cat].map(op =>
                            <Button
                                key={op.name}
                                radius={0}
                                // titleStyle={{fontSize: 14}}
                                buttonStyle={{height: 20, padding: 1, margin: 1}}
                                title={<Text style={{color: colors.onPrimary}}>{op.name}</Text>}
                                size={"sm"}/>)

                        }
                    </HStack>
                })}
                {isDan && <Text style={{fontWeight: "bold",}}>胆</Text>}
            </HStack>
        })}
    </Stack>
}

export default Z14TicketView