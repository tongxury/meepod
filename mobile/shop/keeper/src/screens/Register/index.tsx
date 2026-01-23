import {StyleProp, View, ViewStyle} from "react-native";
import {Stack} from "@react-native-material/core";
import {Card, Text, useTheme} from "react-native-paper";
import {Button} from "@rneui/themed";
import React, {useContext} from "react";
import {AppContext} from "../../providers/global";
import {Image} from 'expo-image';

const RegisterScreen = ({navigation}) => {

    const {colors} = useTheme()

    const {settingsState: {settings}} = useContext<any>(AppContext);


    return <Stack fill={1} p={30} spacing={20}>
            <Text style={{fontWeight: "bold", textAlign: "center"}}
                  variant="titleLarge">联系客服开通店铺</Text>

            <Stack spacing={20} fill={1}>
                <Image
                    style={{
                        flex: 1,
                        width: '100%',
                    }}
                    source={settings?.service?.wechat}
                    contentFit="cover"
                    // PlaceholderContent={<ActivityIndicator style={{flex: 1}}/>}
                />
                <Text variant={"titleMedium"}>{settings?.service?.email}</Text>
                <Button onPress={() => navigation.navigate('Root', {screen: 'Profile'})}>客服已为我开通</Button>
            </Stack>

        </Stack>
};

export default RegisterScreen
