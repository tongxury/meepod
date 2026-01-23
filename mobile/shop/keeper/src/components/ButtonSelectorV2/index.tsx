import {View} from "react-native";
import {Option} from "../../service/typs";
import {Flex} from "@ant-design/react-native";
import {Button} from "@rneui/themed";

const ButtonSelector = ({options, value, onChange}: {
    options: Option[],
    value: string,
    onChange: (value: string) => void
}) => {

    return <Flex justify="start" align="center" style={{flex: 1}} wrap="wrap">
        {options.map((t, i) => <View key={t.value} style={{margin: 1}}>
            <Button
                color={value === t.value ? 'primary' : 'gray'}
                onPress={() => onChange?.(t.value)}
                key={i}>{t.name}</Button>
        </View>)}
    </Flex>
}

export default ButtonSelector