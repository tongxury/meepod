import {View} from "react-native";
import Z1SelectorView from "./Z1SelectorView";
import Z3SelectorView from "./Z3SelectorView";
import Z6SelectorView from "./Z6SelectorView";
import {Ticket} from "../../types";
import {Button, SegmentedButtons} from "react-native-paper";
import React, {useState} from "react";
import {Stack} from "@react-native-material/core";


const SelectorView = ({onConfirm}: { onConfirm: (ticket: Ticket) => void }) => {
    const options = [
        {value: 'z1', label: '直选'},
        {value: 'z3', label: '组三'},
        {value: 'z6', label: '组六'},
    ]

    const [option, setOption] = useState<string>(options[0].value)

    return <Stack fill={1} spacing={10}>
        <SegmentedButtons
            // style={{backgroundColor: 'red'}}
            multiSelect={false}
            density={"medium"}
            value={option}
            onValueChange={(value) => setOption(value)}
            buttons={options}
        />
        {option == 'z1' && <Z1SelectorView onConfirm={onConfirm}/>}
        {option == 'z3' && <Z3SelectorView onConfirm={onConfirm}/>}
        {option == 'z6' && <Z6SelectorView onConfirm={onConfirm}/>}
    </Stack>
}

export default SelectorView