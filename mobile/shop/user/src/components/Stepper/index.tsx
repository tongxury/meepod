import Icon from "react-native-vector-icons/FontAwesome5";
import {useTheme} from "react-native-paper";
import {View, TextInput, Pressable} from "react-native";
import {HStack, Stack} from "@react-native-material/core";


export const Stepper = ({value, min, max, onChange, color, disabled}: {
    value: number,
    min: number,
    max: number,
    onChange: (value: number) => void,
    disabled?: boolean,
    color?: string
}) => {

    const {colors} = useTheme()

    const change = (newValue) => {

        const reg = new RegExp('^[0-9]*$')

        if (!disabled && (!newValue || (newValue && reg.test(newValue) && parseInt(newValue) >= min && parseInt(newValue) <= max))) {
            onChange(!newValue ? 0 : parseInt(newValue))
        }
    }

    return <View>
        <HStack spacing={5} items="center">
            <Pressable onPress={() => change(value - 1)}>
                <Stack center={true} h={25} w={25} bg={value <= min || disabled ? 'gray' : color} radius={2}>
                    <Icon name="minus" color={colors.onPrimary}/>
                </Stack>
            </Pressable>
            <TextInput
                editable={!disabled}
                // @ts-ignore
                value={value === 0 ? '' : value}
                placeholder={`${min}-${max}`}
                onChangeText={value => change(value)}
                style={{width: 70, height: 25, textAlign: "center"}}
            ></TextInput>
            <Pressable onPress={() => change(value + 1)}>
                <Stack center={true} h={25} w={25} bg={value >= max || disabled ? 'gray' : color} radius={2}>
                    <Icon onPress={() => change(value + 1)} name="plus" color={colors.onPrimary}/>
                </Stack>
            </Pressable>
        </HStack>
    </View>

}


