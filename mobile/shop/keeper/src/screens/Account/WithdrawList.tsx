import {View} from "react-native";
import React, {useCallback, useEffect} from "react";
import {useTheme, Text, Avatar} from "react-native-paper";
import {HStack, Stack} from "@react-native-material/core";
import {fetchWithdraws} from "../../service/api";
import AcceptWithdrawTrigger from "../../triggers/Withdraw/Accept";
import RejectTrigger from "../../triggers/Withdraw/Reject";
import ListView from "../../components/ListView";
import UserTrigger from "../../triggers/User";
import ImageViewer from "../../components/ImageViewer";

const WithdrawList = ({}) => {
    const {colors} = useTheme()

    return <ListView
        fetch={page => fetchWithdraws({page})}
        renderItem={(x, updateListItem) =>
            <View style={{backgroundColor: colors.background, padding: 15}}>
                <Stack spacing={10}>
                    <HStack items="center" justify="between">
                        <UserTrigger data={x.user} />
                        <HStack items={'center'} spacing={5}>
                            <Text>{x.created_at}</Text>
                            <Text style={{color: x.status?.color}}>{x.status?.name}</Text>
                        </HStack>
                    </HStack>

                    <Text>金额: <Text
                        style={{fontSize: 20, fontWeight: "bold", color: colors.primary}}>{x.amount}</Text> 元</Text>

                    {x.remark && <Text variant={"bodySmall"}>{x.remark}</Text>}
                    {x.image && <ImageViewer size={"small"} images={[{url: x.image}]}/>}

                    <HStack items="center" justify="end">
                        <HStack items="center" spacing={5}>
                            {
                                x.rejectable && <RejectTrigger
                                    id={x.id}
                                    onConfirm={(newValue) => updateListItem(newValue)}>
                                </RejectTrigger>
                            }
                            {
                                x.acceptable && <AcceptWithdrawTrigger
                                    id={x.id}
                                    onConfirm={(newValue) => updateListItem(newValue)}>
                                </AcceptWithdrawTrigger>
                            }
                        </HStack>
                    </HStack>
                </Stack>
            </View>
        }

    />
}

export default WithdrawList