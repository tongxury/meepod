import {Appbar, Button, Text, useTheme} from "react-native-paper";
import {Toast, WhiteSpace} from "@ant-design/react-native";
import React from "react";
import {View} from "react-native";
import {Stack} from "@react-native-material/core";
import {useRequest} from "ahooks";
import {fetchFeedbacks, fetchProxy} from "../../service/api";
import FeedList from "./List";


const FeedbackScreen = ({navigation}) => {


    const {colors} = useTheme()


    return <View style={{flex: 1}}>
        <Appbar.Header>
            <Appbar.BackAction onPress={() => {
                navigation.goBack()
            }}/>
            <Appbar.Content title={<Text variant="titleMedium">投诉建议</Text>}/>
        </Appbar.Header>
        <WhiteSpace />
        <FeedList />
    </View>
}

export default FeedbackScreen