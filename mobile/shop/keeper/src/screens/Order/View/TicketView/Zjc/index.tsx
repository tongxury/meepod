import {useTheme, Text} from "react-native-paper";
import {HStack, Stack} from "@react-native-material/core";
import {Button, Chip} from "@rneui/themed";


const ZjcTicketView = ({data}) => {
    const {colors} = useTheme()

    return <Stack fill={1} spacing={10}>

        <Text>共选择
            <Text style={{color: colors.primary, fontWeight: "bold"}}> {Object.keys(data?.options ?? {}).length} </Text>场
            <Text style={{color: colors.primary, fontWeight: "bold"}}>  {data.amount} </Text>元
        </Text>

        <Stack spacing={5}>
            <Text style={{fontWeight: "bold"}}>过关方式</Text>
            <HStack items={"center"} wrap={"wrap"}>
                {data?.modes.sort().map(t => {
                    return <Button
                        key={t}
                        radius={0}
                        buttonStyle={{height: 20, margin: 1}}
                        title={<Text style={{fontSize: 14, color: colors.onPrimary}}>{t}</Text>}
                        size={"sm"}/>
                })}
            </HStack>
        </Stack>

        {Object.keys(data?.options ?? {}).map(matchId => {

            const matchOptions = data?.options[matchId]

            const isDan = matchOptions['dan']?.length > 0

            return <Stack spacing={5} key={matchId}>
                <HStack items={"center"} spacing={10}>
                    {isDan && <Text style={{fontWeight: "bold",}}>胆</Text>}
                    <Text style={{
                        fontWeight: "bold",
                        color: isDan ? colors.primary : colors.onBackground
                    }}>主队: {data?.matches?.[matchId]?.home_team}</Text>
                    <Text>让球: {data?.matches?.[matchId]?.r_count}</Text>
                </HStack>
                {Object.keys(matchOptions).filter(t => t != 'dan').sort().map(cat => {
                    return <HStack items={"center"} wrap={"wrap"} key={cat}>
                        {matchOptions[cat].map(op =>
                            <Button
                                key={op.name}
                                radius={0}
                                buttonStyle={{height: 20, margin: 1}}
                                title={<Text style={{fontSize: 14, color: colors.onPrimary}}>{op.name}</Text>}
                                size={"sm"}/>)
                        }
                    </HStack>
                })}
            </Stack>
        })}
    </Stack>
}

export default ZjcTicketView