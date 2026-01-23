import {FlatList, Pressable, View} from "react-native";
import React, {useCallback, useEffect, useState} from "react";
import {useInfiniteScroll, useRequest} from "ahooks";
import {fetchMatches, fetchMatchFilters, fetchOrders, pay} from "../../service/api";
import {Empty, Footer} from "../../components/ListComponent";
import {Text, useTheme} from "react-native-paper";
import {useFocusEffect, useNavigation} from "@react-navigation/native";
import {DropdownSelector} from "../../components/DropdownSelector";
import MatchList from "./ZjcList";
import {HStack} from "@react-native-material/core";

const Z14MatchList = ({category}) => {

    const navigation = useNavigation()

    const {data} = useRequest(() => fetchMatchFilters({category}))
    const filters = data?.data?.data

    const [issueOption, setIssueOption] = useState<any>();

    const {colors} = useTheme()

    return <View style={{flex: 1}}>
        <HStack p={10} bg={colors.background}>
            <DropdownSelector items={filters} value={issueOption ?? filters?.[0]} onChange={setIssueOption}/>
        </HStack>
        <MatchList category={category} issue={issueOption?.value}/>
    </View>

}

export default Z14MatchList