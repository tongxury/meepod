import {createBottomTabNavigator} from '@react-navigation/bottom-tabs';
import {Avatar, BottomNavigation, Text, useTheme} from 'react-native-paper';
import Icon from 'react-native-vector-icons/FontAwesome';
import IconMci from 'react-native-vector-icons/MaterialCommunityIcons';
import IconIon from 'react-native-vector-icons/Ionicons';
import IconFontisto from 'react-native-vector-icons/Fontisto';
import FontAwesome5Icon from 'react-native-vector-icons/FontAwesome5';
import {CommonActions, NavigationContainer, useFocusEffect, useNavigation} from "@react-navigation/native";
import ProfileScreen from "./screens/Profile";
import LoginScreen from "./screens/Login";
import {createStackNavigator} from "@react-navigation/stack";
import * as Linking from "expo-linking";
import {LinkingOptions} from "@react-navigation/native/lib/typescript/src/types";
import {useCallback, useContext, useEffect} from "react";
import {AppContext} from "./providers/global";
import OrderDetailScreen from "./screens/Order/Detail";
import AccountScreen from "./screens/Account";
import ProfileEditorScreen from "./screens/Profile/Editor";
import {bottomBarHeight} from "./utils/dimensions";
import GroupDetailScreen from "./screens/OrderGroup/Detail";
import TicketScreen from "./screens/Ticket";
import ProxyScreen from "./screens/Proxy";
import ProxyDetailScreen from "./screens/Proxy/Detail";
import CooperationScreen from "./screens/Cooperation";
import PaymentScreen from "./screens/Payment";
import {Platform, View} from "react-native";
import nav from "./utils/navigation";
import RegisterScreen from "./screens/Register";
import FeedbackScreen from "./screens/Feedback";
import {StatusBar} from "expo-status-bar";
import CoStorePaymentScreen from "./screens/Cooperation/Payment";


const config = {
    screens: {
        Root: {
            screens: {
                Home: 'home',
                Ticket: 'ticket',
                Account: 'account',
                Proxy: 'proxy',
                // Cooperation: 'cooperation',
                Profile: {
                    path: 'profile',
                    parse: {
                        id: (id) => `user-${id}`,
                    },
                    stringify: {
                        id: (id) => id.replace(/^user-/, ''),
                    },
                },
                Stats: {
                    path: 'stats'
                }
            }
        },
        Login: {
            path: 'login',
        },
        Register: {
            path: 'register',
        },
        OrderDetail: {
            path: 'order-details/:id'
        },
        OrderGroupDetail: {
            path: 'order-group-details/:id'
        },
        ProxyDetail: {
            path: 'proxy-details/:id'
        },
        ProfileEditor: {
            path: 'profile-editor'
        },
        Payment: {
            path: 'payment'
        },
        Feedback: {
            path: 'feedback'
        },
        CoStorePayment: {
            path: 'co-store-payments'
        },
        Cooperation: {
            path: 'cooperation'
        }
    },
};

const linking: LinkingOptions<ReactNavigation.RootParamList> = {
    prefixes: [Linking.createURL("/")],
    config,
};

export default function Main() {

    return <View style={{flex: 1}}>
        <StatusBar style="dark"/>
        <NavigationContainer linking={linking}>
            <StackNavigator/>
        </NavigationContainer>
    </View>
}


function StackNavigator() {

    const {
        authState: {authed},
        settingsState: {settings, update: updateSettings},
        counterState: {counter, update: updateCounter},
    } = useContext<any>(AppContext);

    useFocusEffect(useCallback(() => {

        // updateCounter()
    }, []))

    const Stack = createStackNavigator();

    const options = {}
    const navigation = useNavigation()


    nav.setNavigation(navigation)

    return <View style={{flex: 1}}>
        <Stack.Navigator screenOptions={{headerShown: false,}}
                         screenListeners={{
                             state: (e) => {
                                 // @ts-ignore
                                 // checkDeps(e.data.state).then()
                                 // updateCounter()
                                 // console.log('screenListeners', e.data)

                                 // const au = await authed()
                                 //
                                 // console.log('au', au, e.data?.state?.index)
                                 //
                                 // // @ts-ignore
                                 // if (e.data?.state?.index === 0 && au) {
                                 //     updateCounter()
                                 // }

                                 // console.log('e.data.state', e.data.state)
                             },
                         }}
        >
            <Stack.Screen name="Root" component={BottomTabNavigator}/>
            <Stack.Screen name="Login" component={LoginScreen} options={options}/>
            <Stack.Screen name="Register" component={RegisterScreen} options={options}/>
            <Stack.Screen name="OrderDetail" component={OrderDetailScreen} options={options}/>
            <Stack.Screen name="OrderGroupDetail" component={GroupDetailScreen} options={options}/>
            <Stack.Screen name="ProxyDetail" component={ProxyDetailScreen} options={options}/>
            <Stack.Screen name="ProfileEditor" component={ProfileEditorScreen} options={options}/>
            <Stack.Screen name="Payment" component={PaymentScreen} options={options}/>
            <Stack.Screen name="Feedback" component={FeedbackScreen} options={options}/>
            <Stack.Screen name="CoStorePayment" component={CoStorePaymentScreen} options={options}/>
            <Stack.Screen name="Cooperation" component={CooperationScreen} options={options}/>
        </Stack.Navigator>
    </View>

}


const Tab = createBottomTabNavigator();

function BottomTabNavigator() {

    const {
        authState: {authed},
        counterState: {counter, update: updateCounter},
        updateState: {update, checkForUpdates},
    } = useContext<any>(AppContext);


    const {colors} = useTheme()

    function PaperBar({navigation, state, descriptors, insets}) {
        return <BottomNavigation.Bar
            navigationState={state}
            safeAreaInsets={insets}
            onTabPress={({route, preventDefault}) => {

                updateCounter()
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
                if (props.route.name == 'Profile') {
                    if (update) {
                        // @ts-ignore
                       return 1 + (counter?.[props.route.name.toLowerCase()] ?? 0)
                    }
                }
                return counter?.[props.route.name.toLowerCase()]
            }}
            style={{height: bottomBarHeight, margin: 0}}
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
            initialRouteName="Ticket"
            tabBar={props => <PaperBar {...props} />}
            // tabBar={props => <MyTabBar {...props} />}

        >

            <Tab.Screen
                name="Ticket"
                component={TicketScreen}
                options={{
                    tabBarLabel: '票务',
                    tabBarIcon: ({color, size}) => {
                        return <IconFontisto name="ticket" size={size} color={color}/>;
                    },
                }}
            />
            <Tab.Screen
                name="Account"
                component={AccountScreen}
                options={{
                    tabBarLabel: '账本',
                    tabBarIcon: ({color, size}) => {
                        return <IconMci name="book-account" size={size} color={color}/>;
                    },
                }}
            />
            <Tab.Screen
                name="Proxy"
                // component={HomeNavigator}
                component={ProxyScreen}
                options={{
                    tabBarLabel: '推广',
                    tabBarIcon: ({color, size}) => {
                        return <FontAwesome5Icon name="users" size={size} color={color}/>;
                    },
                }}
            />
            {/*<Tab.Screen*/}
            {/*    name="Cooperation"*/}
            {/*    component={CooperationScreen}*/}
            {/*    options={{*/}
            {/*        tabBarLabel: '合作',*/}
            {/*        tabBarIcon: ({color, size}) => {*/}
            {/*            return <FontAwesome5Icon name="hands-helping" size={size} color={color}/>*/}
            {/*        },*/}
            {/*    }}*/}
            {/*/>*/}
            <Tab.Screen
                name="Profile"
                component={ProfileScreen}
                options={{
                    tabBarLabel: '我的',
                    tabBarIcon: ({color, size}) => {
                        return <Icon name="user" size={size} color={color}/>;
                    },
                }}
            />
        </Tab.Navigator>
    );
}
