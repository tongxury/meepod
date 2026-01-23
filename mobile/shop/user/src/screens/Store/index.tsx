import React, {useCallback, useContext, useEffect, useState} from "react";
import {Avatar, Button, Card, IconButton, useTheme, Text} from "react-native-paper";
import StoreSwitcher from "./StoreSwitcher";
import {useFocusEffect} from "@react-navigation/native";
import {AppContext} from "../../providers/global";
import {HStack, Stack} from "@react-native-material/core";
import {useSafeAreaInsets} from "react-native-safe-area-context";
import {Chip, Image} from "@rneui/themed";
import {addStoreUser, fetchStoreInfo} from "../../service/api";

export const StoreScreen = ({navigation}) => {

    const [open, setOpen] = useState<boolean>(false)
    const {storeState: {store, loading, refresh, fetch, update}} = useContext<any>(AppContext);


    useFocusEffect(useCallback(() => {
        refresh()
    }, []))


    const onSwitch = (storeId, proxyId) => {

        addStoreUser(storeId, proxyId).then(rsp => {
            if (rsp.data?.code == 0) {
                update(storeId)
                setOpen(false)
            }
        })

    }

    const {colors} = useTheme()
    const {top} = useSafeAreaInsets();

    return (
        <Stack fill={true} ph={15} pt={top + 15} pb={15} justify="end">
            <Card style={{
                shadowOffset: {width: 0, height: 15},
                shadowColor: colors.primary,
                shadowOpacity: 0.2,
                shadowRadius: 10,
                borderRadius: 10,
                marginBottom: 100,
            }}>
                {
                    store ? <Stack p={20} spacing={20}>
                            <HStack items={"center"} justify={"between"}>
                                <HStack items={"center"} spacing={10}>
                                    <Avatar.Image style={{backgroundColor: colors.background}} size={60}
                                                  source={{uri: store?.icon}}/>

                                    <Stack spacing={10}>
                                        <Text style={{fontSize: 20}} variant={"titleMedium"}>{store?.name}</Text>
                                        <Text variant={"bodyMedium"}>{store?.id}</Text>
                                    </Stack>
                                </HStack>

                                <IconButton icon="dots-vertical" onPress={() => setOpen(true)}/>
                            </HStack>

                            <Button mode={"contained"} color={store?.status?.color}>{store?.status?.name}</Button>

                            {store?.notice && <Stack spacing={5}>
                                <Text variant={"titleMedium"}>店铺公告</Text>
                                <Text variant={"bodyMedium"}>{store?.notice}</Text>
                            </Stack>}


                            {/*<Stack center={true}>*/}
                            {/*    <Image*/}
                            {/*        source={{uri: store?.wechat_image}}*/}
                            {/*        containerStyle={{*/}
                            {/*            aspectRatio: 1,*/}
                            {/*            width: '100%',*/}
                            {/*            flex: 1,*/}
                            {/*            borderRadius: 10*/}
                            {/*        }}*/}
                            {/*        // PlaceholderContent={<ActivityIndicator style={{flex: 1}}/>}*/}
                            {/*    />*/}
                            {/*</Stack>*/}


                        </Stack> :
                        <Stack spacing={100} center={true} style={{paddingVertical: 50}}>
                            <Avatar.Image size={150} source={{uri: store?.icon}}/>
                            <Text onPress={() => setOpen(true)} style={{fontSize: 28}}>点击选择</Text>
                        </Stack>

                }

            </Card>
            <StoreSwitcher onSwitch={onSwitch} open={open} onClose={() => setOpen(false)}/>
        </Stack>

    )
        ;
}

export default StoreScreen

