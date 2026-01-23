import {ButtonGroup} from "@rneui/themed";
import {Text, useTheme} from "react-native-paper";
import React from "react";

const ButtonSelector = ({disabled, values, options}: { values: number[], options: string[], disabled?: boolean }) => {

    const {colors} = useTheme()

    return <ButtonGroup
        disabled={disabled}
        containerStyle={{borderRadius: 5, backgroundColor: colors.primaryContainer}}
        selectMultiple
        selectedButtonStyle={{backgroundColor: colors.primary}}
        disabledSelectedStyle={{backgroundColor: colors.primary}}
        selectedIndexes={values}
        buttons={
            options.map((t, i) => {
                return {
                    element: () => <Text
                        style={{color: values?.includes(i) ? colors.onPrimary : colors.onPrimaryContainer}}>{t}</Text>
                }
            })
        }
    />
}

export default ButtonSelector