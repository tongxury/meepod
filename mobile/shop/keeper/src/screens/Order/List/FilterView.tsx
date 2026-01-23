import {Text, View} from "react-native";
import {Controller, useForm} from "react-hook-form";
import React, {useCallback, useState} from "react";
import Collapsible from "react-native-collapsible";
import ButtonSelector from "../../../components/ButtonSelectorV2";
import {Chip, Searchbar, useTheme} from "react-native-paper";
import {useRequest} from "ahooks";
import {fetchOrderFilters} from "../../../service/api";
import {HStack} from "@react-native-material/core";
import IconAntd from "react-native-vector-icons/AntDesign";
import Icon from "react-native-vector-icons/FontAwesome";
import {color} from "@rneui/base";
import {useFocusEffect} from "@react-navigation/native";
import {Option} from "../../../service/typs";

const FilterView = ({onValueChange}: { onValueChange: (values) => void }) => {


    const {data, run} = useRequest(fetchOrderFilters, {manual: true})
    const filters = data?.data?.data ?? {
        order: {
            dateRange: [{name: "全部", value: ""}],
            category: [{name: "全部状态", value: ""}],
            item: [{name: "全部彩种", value: ""}],
        }
    }

    useFocusEffect(useCallback(() => {
        run()
    }, []))

    const {control, getFieldState, getValues, setValue, handleSubmit, formState: {errors}} = useForm({
        defaultValues: {
            dateRange: '',
            status: '',
            item: ''
        }
    });


    const [collapsedDateRange, setCollapsedDateRange] = useState<boolean>(true)
    const [collapsedStatus, setCollapsedStatus] = useState<boolean>(true)
    const [collapsedItem, setCollapsedItem] = useState<boolean>(true)

    const [dateRange, setDateRange] = useState<string>('')
    const [status, setStatus] = useState<string>('')
    const [item, setItem] = useState<string>('')

    const {colors} = useTheme()

    return <View style={{backgroundColor: colors.background}}>
        <HStack items={"center"} justify={"between"} p={15}>
            <Chip onClose={() => undefined}
                  closeIcon={() => <IconAntd
                      color={colors.primary}
                      name={collapsedDateRange ? "caretright" : "caretdown"}/>}
                  onPress={() => setCollapsedDateRange(!collapsedDateRange)}>
                {new Map(filters?.order.dateRange.map(t => [t.value, t])).get(dateRange)?.name}
            </Chip>
            <Chip
                onClose={() => undefined}
                closeIcon={() => <IconAntd
                    color={colors.primary}
                    name={collapsedStatus ? "caretright" : "caretdown"}/>}
                onPress={() => setCollapsedStatus(!collapsedStatus)}>
                {new Map(filters?.order.category.map(t => [t.value, t])).get(status)?.name}
            </Chip>
            <Chip
                onClose={() => undefined}
                closeIcon={() => <IconAntd
                    color={colors.primary}
                    name={collapsedItem ? "caretright" : "caretdown"}/>}
                onPress={() => setCollapsedItem(!collapsedItem)}>
                {new Map(filters?.order.item.map(t => [t.value, t])).get(item)?.name}
            </Chip>
        </HStack>
        <Collapsible style={{paddingHorizontal: 10}} collapsed={collapsedDateRange}>
            <Controller
                control={control}
                rules={{required: false,}}
                render={({field: {value}}) => (
                    <View style={{flex: 1}}>
                        <ButtonSelector options={filters?.order?.dateRange || []} value={value}
                                        onChange={v => {
                                            setValue('dateRange', v);
                                            setDateRange(v)
                                            setCollapsedDateRange(true)
                                            handleSubmit(onValueChange)()
                                        }}/>
                    </View>
                )}
                name="dateRange"
            />
        </Collapsible>
        <Collapsible style={{padding: 10}} collapsed={collapsedStatus}>
            <Controller
                control={control}
                rules={{required: false,}}
                render={({field: {value}}) => (
                    <View style={{flex: 1}}>
                        <ButtonSelector options={filters?.order?.category || []} value={value}
                                        onChange={v => {
                                            setValue('status', v)
                                            setStatus(v)
                                            setCollapsedStatus(true)
                                            handleSubmit(onValueChange)()
                                        }}/>
                    </View>
                )}
                name="status"
            />
        </Collapsible>
        <Collapsible style={{padding: 10}} collapsed={collapsedItem}>
            <Controller
                control={control}
                rules={{required: false,}}
                render={({field: {value}}) => (
                    <View style={{flex: 1}}>
                        <ButtonSelector options={filters?.order?.item || []} value={value}
                                        onChange={v => {
                                            setValue('item', v)
                                            setItem(v)
                                            setCollapsedItem(true)
                                            handleSubmit(onValueChange)()
                                        }}/>
                    </View>
                )}
                name="item"
            />
        </Collapsible>
    </View>
}

export default FilterView