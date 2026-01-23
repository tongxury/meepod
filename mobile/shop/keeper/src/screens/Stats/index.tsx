import {Dimensions, SafeAreaView, useWindowDimensions, View} from "react-native";
import {
    LineChart,
    BarChart,
    PieChart,
    ProgressChart,
    ContributionGraph,
    StackedBarChart
} from "react-native-chart-kit";
import {Text, useTheme} from 'react-native-paper'
import {useSafeAreaInsets} from "react-native-safe-area-context";
import {Stack} from "@react-native-material/core";

const StatsScreen = () => {

    const insets = useSafeAreaInsets();

    const {colors} = useTheme()

    const {width} = useWindowDimensions()

    const bodyPadding = 10

    return <SafeAreaView style={{flex: 1, paddingTop: insets.top, backgroundColor: colors.background}}>
        <Stack p={bodyPadding}>

            <Text>Bezier Line Chart</Text>
            <LineChart
                data={{
                    labels: ["January", "February", "March", "April", "May", "June"],
                    datasets: [
                        {
                            data: [
                                Math.random() * 100,
                                Math.random() * 100,
                                Math.random() * 100,
                                Math.random() * 100,
                                Math.random() * 100,
                                Math.random() * 100
                            ]
                        }
                    ]
                }}
                width={width - 2 * bodyPadding} // from react-native
                height={220}
                yAxisLabel="$"
                yAxisSuffix="k"
                yAxisInterval={1} // optional, defaults to 1
                chartConfig={{
                    backgroundColor: colors.primary,
                    backgroundGradientFrom: colors.primary,
                    backgroundGradientTo: colors.primaryContainer,
                    decimalPlaces: 2, // optional, defaults to 2dp
                    color: (opacity = 1) => `rgba(255, 255, 255, ${opacity})`,
                    labelColor: (opacity = 1) => `rgba(255, 255, 255, ${opacity})`,
                    style: {
                        // borderRadius: 16
                    },
                    propsForDots: {
                        r: "6",
                        strokeWidth: "2",
                        stroke: colors.primary
                    }
                }}
                bezier
                style={{
                    borderRadius: 16
                }}
            />
        </Stack>
    </SafeAreaView>
}

export default StatsScreen