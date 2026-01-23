import {Button as RneButton, ButtonGroup} from "@rneui/themed";
import {Text, useTheme} from "react-native-paper";
import React from "react";
import {View} from "react-native";
import {HStack} from "../Layout/StackLayout";


type Option = {
    name: string, value: string
}
export const ButtonSelector = ({values, options}: {
    values?: string[],
    options: Option[]
}) => {

    const {colors} = useTheme()

    return <HStack fill={1}>
        {options?.map(t => <RneButton
            style={{flex: 1}}
            key={t.value}
            color={values?.includes(t.value) ? "primary" : "secondary"}
            title={<Text style={{
                fontSize: 14,
                color: values?.includes(t.value) ? colors.onPrimary : colors.primary
            }}>{t.name}</Text>}
            containerStyle={{minWidth: 100, height: 30, flex: 3, borderRadius: 0}}
            size="sm"
            radius={0}
        />)}
    </HStack>
    {/*    disabled={disabled}*/
    }
    {/*    containerStyle={{borderRadius: 5, backgroundColor: colors.primaryContainer}}*/
    }
    {/*    selectMultiple*/
    }
    {/*    selectedButtonStyle={{backgroundColor: colors.primary}}*/
    }
    {/*    disabledSelectedStyle={{backgroundColor: colors.primary}}*/
    }
    {/*    selectedIndexes={values}*/
    }
    {/*    buttons={*/
    }
    {/*        options.map((t, i) => {*/
    }
    {/*            return {*/
    }
    {/*                element: () => <Text*/
    }
    {/*                    style={{color: values?.includes(i) ? colors.onPrimary : colors.onPrimaryContainer}}>{t}</Text>*/
    }
    {/*            }*/
    }
    {/*        })*/
    }
    {/*    }*/
    }
    {/*/>*/
    }
}
