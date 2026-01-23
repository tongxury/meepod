import {createBottomTabNavigator} from '@react-navigation/bottom-tabs';
import {Avatar, BottomNavigation, useTheme} from 'react-native-paper';
import Icon from 'react-native-vector-icons/FontAwesome5';
import {CommonActions, NavigationContainer, useFocusEffect, useNavigation,} from "@react-navigation/native";
import HomeScreen from "./screens/Home";
import ProfileScreen from "./screens/Profile";
import LoginScreen from "./screens/Login";
import {CardStyleInterpolators, createStackNavigator, TransitionSpecs} from "@react-navigation/stack";
import * as Linking from "expo-linking";
import {LinkingOptions} from "@react-navigation/native/lib/typescript/src/types";
import {useCallback, useContext, useEffect} from "react";
import StoreScreen from "./screens/Store";
import {AppContext} from "./providers/global";
import HistoryScreen from "./screens/History";
import OrderDetailScreen from "./screens/Order/Detail";
import GroupDetailScreen from "./screens/OrderGroup/Detail";
import PlanScreen from "./screens/Plan";
import ProfileEditorScreen from "./screens/Profile/Editor";
import {bottomBarHeight} from "./utils/dimensions";
import PaymentScreen from "./screens/Payment";
import IconAntd from "react-native-vector-icons/AntDesign";
import MatchScreen from "./screens/Match";
import MoreScreen from "./screens/More";
import More from "./screens/More";
import ProxyScreen from "./screens/Proxy";
import nav from "./utils/navigation";
import FeedbackScreen from "./screens/Feedback";
import {View} from "react-native";
import {StatusBar} from 'expo-status-bar';


const config = {
    screens: {
        Root: {
            screens: {
                Home: 'home',
                Score: 'score',
                History: 'store/history',
                Store: 'store',
                Profile: {
                    path: 'profile',
                    parse: {
                        id: (id) => `user-${id}`,
                    },
                    stringify: {
                        id: (id) => id.replace(/^user-/, ''),
                    },
                }
            }
        },
        Login: {
            path: 'login',
        },
        Plan: {
            path: 'store/plan'
        },
        More: {
            path: 'store/more'
        },
        Proxy: {
            path: 'store/proxy'
        },
        Feedback: {
            path: 'store/feedback'
        },
        OrderDetail: {
            path: 'store/order-details/:id',
        },
        OrderGroupDetail: {
            path: 'store/order-group-details/:id',
        },
        ProfileEditor: {
            path: 'profile-editor',
        },
        Payment: {
            path: 'payment',
        },
    },
};


const linking: LinkingOptions<ReactNavigation.RootParamList> = {
    prefixes: [Linking.createURL("/")],
    config,
};

export default function Main() {

    return (<View style={{flex: 1}}>
            <StatusBar style="dark"/>
            <NavigationContainer linking={linking}>
                <StackNavigator/>
            </NavigationContainer>
        </View>

    );
}

function StackNavigator() {

    const {
        settingsState: {update: updateSettings}
    } = useContext<any>(AppContext);

    useFocusEffect(useCallback(() => {
    }, []))

    const Stack = createStackNavigator();

    const options = {
        // cardStyleInterpolator: CardStyleInterpolators.forHorizontalIOS,
    }

    const navigation = useNavigation()

    nav.setNavigation(navigation)

    return <Stack.Navigator screenOptions={{headerShown: false,}}
                            screenListeners={{
                                state: (e) => {
                                    // @ts-ignore
                                    // checkDeps(e.data.state).then()
                                },

                            }}
    >
        <Stack.Screen name="Root" component={BottomTabNavigator}/>
        <Stack.Screen name="Login" component={LoginScreen} options={options}/>
        <Stack.Screen name="Plan" component={PlanScreen} options={options}/>
        <Stack.Screen name="OrderDetail" component={OrderDetailScreen} options={options}/>
        <Stack.Screen name="OrderGroupDetail" component={GroupDetailScreen} options={options}/>
        <Stack.Screen name="More" component={MoreScreen} options={options}/>
        <Stack.Screen name="Proxy" component={ProxyScreen} options={options}/>
        <Stack.Screen name="Feedback" component={FeedbackScreen} options={options}/>
        <Stack.Screen name="ProfileEditor" component={ProfileEditorScreen} options={options}/>
        <Stack.Screen name="Payment" component={PaymentScreen} options={options}/>
    </Stack.Navigator>
}


const Tab = createBottomTabNavigator();

function BottomTabNavigator() {

    const {storeState: {store}, updateState: {update, checkForUpdates}} = useContext<any>(AppContext);

    const {colors} = useTheme()

    function PaperBar({navigation, state, descriptors, insets}) {

        return <BottomNavigation.Bar
            navigationState={state}
            safeAreaInsets={insets}
            onTabPress={({route, preventDefault}) => {
                checkForUpdates()

                const event = navigation.emit({
                    type: 'tabPress',
                    target: route.key,
                    canPreventDefault: true,
                });

                if (event.defaultPrevented) {
                    preventDefault();
                } else {
                    navigation.dispatch({
                        // @ts-ignore
                        ...CommonActions.navigate(route.name, route.params),
                        target: state.key,
                    });
                }
            }}
            activeColor={colors.primary}
            getBadge={props => {
                // @ts-ignore
                if (props.route.name === 'Profile') {
                    return update ? 1 : undefined
                } else {
                    return undefined
                }
            }}
            style={{height: bottomBarHeight}}
            renderIcon={({route, focused, color}) => {
                const {options} = descriptors[route.key];
                return options.tabBarIcon?.({focused: focused, color, size: 24});
            }}

            getLabelText={({route}) => {
                const {options} = descriptors[route.key];
                return options.tabBarLabel ?? '';
            }}
        />
    }

    return (
        <Tab.Navigator
            screenOptions={{
                headerShown: false,
            }}
            screenListeners={props => undefined}
            initialRouteName="Store"
            tabBar={props => <PaperBar {...props} />}
            // tabBar={props => <MyTabBar {...props} />}

        >
            <Tab.Screen
                name="Home"
                // component={HomeNavigator}
                component={HomeScreen}
                options={{
                    tabBarLabel: '首页',
                    tabBarIcon: ({color, size}) => {
                        return <Icon name="home" size={size} color={color}/>;
                    },
                }}
            />
            <Tab.Screen
                name="Score"
                component={MatchScreen}
                options={{
                    tabBarLabel: '比分',
                    tabBarIcon: ({color, size}) => {
                        return <Icon name="vimeo-square" size={size} color={color}/>;
                    },
                }}
            />
            <Tab.Screen
                name="Store"
                component={StoreScreen}
                options={{
                    tabBarLabel: store?.name ? '' : '请选择店铺',
                    tabBarIcon: ({color, size}) => {
                        return store?.icon ? <Avatar.Image style={{backgroundColor: colors.background}} size={45}
                                                           source={{uri: store?.icon}}/> :
                            <Icon name="store" size={size} color={color}/>;
                    },
                }}
            />
            <Tab.Screen
                name="History"
                component={HistoryScreen}
                options={{
                    tabBarLabel: '记录',
                    tabBarIcon: ({color, size}) => {
                        return <IconAntd name="clockcircle" size={size} color={color}/>;
                    },
                }}
            />
            <Tab.Screen
                name="Profile"
                component={ProfileScreen}
                options={{
                    tabBarLabel: '我的',
                    tabBarIcon: ({color, size}) => {
                        return <Icon name="user-secret" size={size} color={color}/>;
                    },
                }}
            />
        </Tab.Navigator>
    );
}

// Each tab has its own navigation stack, you can read more about this pattern here:
// https://reactnavigation.org/docs/tab-based-navigation#a-stack-navigator-for-each-tab
const HomeStack = createStackNavigator();

function HomeNavigator() {
    return (
        <HomeStack.Navigator>
            <HomeStack.Screen
                name="HomeScreen"
                component={HomeScreen}
                options={{headerShown: false}}

            />
        </HomeStack.Navigator>
    );
}


