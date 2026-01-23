import axios from "axios";
import AsyncStorage from "@react-native-async-storage/async-storage";
import { Toast } from "@ant-design/react-native";
import {
    Account, Feedback, Files,
    GroupShare, Image,
    Issue, ItemState, Match,
    Order,
    OrderGroup,
    PageData, Payment,
    PaymentOrder,
    PayMethod,
    Plan, Proxy, ProxyUser,
    Result, Store, Topup,
    User,
    UserProfile, Withdraw
} from "./typs";
import { Platform } from "react-native";
import { Config } from "../../config";
import nav from "../utils/navigation";
import Constants from "expo-constants";

const instance = axios.create({});

instance.interceptors.request.use(async function (config) {

    const conf = Config()
    if (config.url.startsWith("/api/payment/")) {
        config.baseURL = conf.apiPaymentUrl || conf.apiUrl
    } else {
        config.baseURL = conf.apiUrl
    }

    const storeId = await AsyncStorage.getItem('store') || ''
    // config.url = config.url.replace("STORE_ID", storeId)

    const token = await AsyncStorage.getItem('user_auth') ?? ''
    config.headers.set('Authorization', token ? 'Bearer ' + token : '')
    config.headers.set('StoreId', storeId)
    config.headers.set('Platform', Platform.OS)
    config.headers.set('Client', 'user')
    config.headers.set('Version', Constants.expoConfig.version ?? '')
    config.headers.set('U-Version', conf.version)

    if (__DEV__) {
        console.log(`>>> [API Request] ${config.method?.toUpperCase()} ${config.baseURL + (config.url || '')}`, {
            params: config.params,
            data: config.data,
        });
    }

    return config;
}, function (error) {
    console.error(error)
    // Do something with request error
    return Promise.reject(error);
});

// Add a response interceptor
instance.interceptors.response.use(function (response) {
    // Any status code that lie within the range of 2xx cause this function to trigger
    // Do something with response data

    if (__DEV__) {
        console.log(`<<< [API Response] ${response.config.method?.toUpperCase()} ${response.config.baseURL + (response.config.url || '')}`, response.data);
    }


    if (response.status !== 200) {
        Toast.info({
            content: '服务器繁忙, 请稍后重试',
            mask: false,
        })
    } else {

        if (response.data.code !== 0) {
            console.error(response.data.error, response)
        }
        if (response.data.message) {
            Toast.info({
                content: response.data.message,
                mask: false,
            })
        }

        switch (response.data.code) {
            case 10401:
                nav.navigate('Login')
                break
            case 10404:
                nav.navigate('Store')
                break
        }

    }

    return response;
}, function (error) {

    console.error(error)
    // Any status codes that falls outside the range of 2xx cause this function to trigger
    // Do something with response error
    return Promise.reject(error);
});

export const sendCode = ({ phone, event }: { phone: string, event?: string }) => {
    event = event || 'login'

    return instance.request({
        method: "POST",
        url: '/api/v1/auth-codes',
        params: { phone, event }
    })
}
export const login = ({ phone, code }: { phone: string, code: string }) => {
    return instance.request({
        // method: "POST",
        url: '/api/v1/auth-tokens',
        params: {
            phone: phone, code: code
        }
    })
}
export const fetchAuthStatus = ({ token }: { token: string }) => {
    return instance.request({
        method: "GET",
        url: '/api/v1/auth-status',
        params: {
            token
        }
    })
}
export const fetchStoreInfo = (storeId) => {
    return instance.request<Result<Store>>({
        method: "GET",
        url: '/api/v1/stores/' + storeId,
    })
}

export const addStoreUser = (storeId, proxyId) => {
    return instance.request({
        url: '/api/v1/stores/' + storeId + '/users',
        method: 'post',
        params: { proxyId }
    })
}
export const fetchItemStates = () => {
    return instance.request<Result<ItemState[]>>({
        method: "GET",
        url: '/api/v1/store-items'
    })
}
export const fetchSettings = () => {
    return instance.request({
        url: '/api/v1/settings'
    })
}
export const addPlan = (plan: {
    itemId: string,
    content: string,
    multiple: number,
}) => {
    return instance.request({
        url: '/api/v1/plans',
        method: "post",
        data: plan
    })
}

export const fetchPlans = (params: { category, page?, size?}) => {
    return instance.request<Result<PageData<Plan>>>({
        method: "GET",
        url: '/api/v1/plans',
        params
    })
}

export const deletePlans = ({ id }) => {
    return instance.request({
        url: `/api/v1/plans/${id}`,
        method: "delete"
    })
}

export const deleteOrder = ({ id }) => {
    return instance.request({
        url: `/api/v1/orders/${id}`,
        method: "put",
        params: {
            action: 'cancel'
        }
    })
}

export const fetchOrderGroups = (params: { category, page?, size?}) => {
    return instance.request<Result<PageData<OrderGroup>>>({
        method: "GET",
        url: '/api/v1/order-groups',
        params
    })
}
// export const fetchOrderGroupJoiners = ({groupId, page, size}: { groupId, page?, size? }) => {
//     return instance.request<Result<PageData<GroupJoiner>>>({
//         url: `/api/v1/order-groups/${groupId}/joiners`,
//     })
// }

export const fetchOrderGroupOrders = (params: { groupId, page?, size?}) => {
    return instance.request<Result<PageData<GroupShare[]>>>({
        method: "GET",
        url: `/api/v1/order-groups/${params.groupId}/shares`,
        params: params
    })
}

export const addOrderGroupOrder = ({ groupId, volume }: { groupId: string, volume: number }) => {
    return instance.request<Result<string>>({
        url: `/api/v1/order-groups/${groupId}/shares`,
        method: "post",
        data: { volume }
    })
}
export const fetchOrderGroupDetail = ({ id, biz_category }: { id: string, biz_category?: string }) => {
    return instance.request<Result<OrderGroup>>({
        method: "GET",
        url: `/api/v1/order-groups/${id}`,
        params: { biz_category }
    })
}


export const followOrder = ({ followOrderId }) => {
    return instance.request({
        url: '/api/v1/orders',
        method: "post",
        params: { action: 'follow' },
        data: {
            followOrderId
        }
    })
}

export const addGroupOrder = (params: {
    planId?: string,
    volume?: number,
    totalVolume?: number,
    rewardRate?: number,
    floor?: number,
    remark?: string,
    needUpload?: boolean,
}) => {
    return instance.request({
        url: '/api/v1/order-groups',
        method: "post",
        data: params
    })
}

export const addOrder = (params: {
    planId?: string,
    needUpload?: boolean,
}) => {
    return instance.request({
        url: '/api/v1/orders',
        method: "post",
        params: { action: '' },
        data: params
    })
}

export const fetchOrders = (params: { category, page?, size?}) => {
    return instance.request<Result<PageData<Order>>>({
        method: "GET",
        url: '/api/v1/orders',
        params
    })
}
export const fetchOrderDetail = ({ id }: { id: string }) => {
    return instance.request<Result<Order>>({
        method: "GET",
        url: `/api/v1/orders/${id}`,
    })
}

export const fetchCurrentIssue = (params: { itemId }) => {
    return instance.request<Result<Issue>>({
        method: "GET",
        url: '/api/v1/current-issues',
        params
    })
}

// export const fetchIssueData = (params: { id }) => {
//     return instance.request<Result<any>>({
//         url: `/api/v1/issues/${params.id}/data`
//     })
// }


export const pay = ({ orderId }: { orderId: string }) => {
    return instance.request<Result<Topup | undefined>>({
        url: '/api/payment/v1/payments',
        method: "post",
        data: { orderId },
        // params: {method}
    })
}

export const fetchPayMethods = ({ amount }: { amount: number }) => {
    return instance.request<Result<PayMethod[]>>({
        method: "GET",
        url: '/api/payment/v1/payment-methods',
        params: { amount }
    })
}

export const fetchUserAccount = ({ userId }: { userId: string }) => {
    return instance.request<Result<Account>>({
        method: "GET",
        url: `/api/payment/v1/accounts/${userId}`,
    })
}


export const fetchUserProfile = ({ userId }: { userId: string }) => {
    return instance.request<Result<UserProfile>>({
        method: "GET",
        url: `/api/v1/users/${userId}/profiles`,
    })
}

export const fetchUser = ({ userId }: { userId: string }) => {
    return instance.request<Result<User>>({
        method: "GET",
        url: `/api/v1/users/${userId}`,
    })
}

export const uploadImageV2 = ({ files }: { files: { name: string, src: string }[] }) => {

    return instance.request<Result<Files>>({
        method: "post",
        url: `/api/v2/images`,
        data: {
            base64Resources: files.map(t => `${t.name},${t.src}`),
        }
    })
}

export const updateUser = ({ field, value }: { field: string, value: string }) => {
    return instance.request<Result<User>>({
        method: "put",
        url: `/api/v1/users/me`,
        params: {
            field, value
        }
    })
}

export const addWithdrawOrder = ({ amount }: { amount: number }) => {
    return instance.request<Result<Withdraw>>({
        method: "post",
        url: `/api/payment/v1/withdraws`,
        data: {
            amount
        }
    })
}

export const fetchWithdraws = ({ page, size }: { page?: number, size?: number }) => {
    return instance.request<Result<PageData<Withdraw>>>({
        method: "GET",
        url: `/api/payment/v1/withdraws`,
        params: { page, size }
    })
}

export const fetchPayments = ({ page, size }: { page?: number, size?: number }) => {
    return instance.request<Result<PageData<Payment>>>({
        method: "GET",
        url: `/api/payment/v1/payments`,
        params: { page, size }
    })
}

export const cancelWithdraw = ({ id }: { id: string }) => {
    return instance.request<Result<Withdraw>>({
        url: `/api/payment/v1/withdraws/${id}`,
        method: "put",
        params: {
            action: 'cancel'
        }
    })
}


export const addTopUpOrder = ({ amount }: { amount: number }) => {
    return instance.request<Result<Topup>>({
        method: "post",
        url: `/api/payment/v1/topups`,
        data: {
            // method,
            amount
        }
    })
}

export const fetchPaymentOrders = ({ category, page, size }: { category: string, page?: number, size?: number }) => {
    return instance.request<Result<PageData<PaymentOrder>>>({
        method: "GET",
        url: `/api/payment/v1/orders`,
        params: { category, page, size }
    })
}

export const fetchTopups = ({ page, size }: { page?: number, size?: number }) => {
    return instance.request<Result<PageData<Topup>>>({
        method: "GET",
        url: `/api/payment/v1/topups`,
        params: { page, size }
    })
}

export const cancelTopup = ({ id }: { id: string }) => {
    return instance.request<Result<Topup>>({
        url: `/api/payment/v1/topups/${id}`,
        method: "put",
        params: {
            action: 'cancel'
        }
    })
}

export const fetchTopup = ({ id }: { id: string }) => {
    return instance.request<Result<Topup>>({
        method: "GET",
        url: `/api/payment/v1/topups/${id}`,
    })
}


export const fetchMatchFilters = (params: {}) => {
    return instance.request<Result<any>>({
        method: "GET",
        url: '/api/v1/match-filters',
        params
    })
}

export const fetchMatches = (params: { category, issue?, page?, size?}) => {
    return instance.request<Result<PageData<Match>>>({
        method: "GET",
        url: '/api/v1/matches',
        params
    })
}

export const fetchProxy = () => {
    return instance.request<Result<Proxy>>({
        method: "GET",
        url: `/api/v1/proxies/me`,
    })
}
export const fetchProxyUsers = (params: { page?, size?}) => {
    return instance.request<Result<PageData<ProxyUser>>>({
        method: "GET",
        url: `/api/v1/proxies/me/users`,
        params
    })
}


export const addFeedback = (params: { text: string }) => {
    return instance.request<Result<Feedback>>({
        url: `/api/v1/feedbacks`,
        method: 'post',
        data: params
    })
}

export const fetchFeedbacks = (params: { page?, size?}) => {
    return instance.request<Result<PageData<Feedback>>>({
        method: "GET",
        url: `/api/v1/feedbacks`,
        params
    })
}

export const resolve = (params: { id: string }) => {
    return instance.request<Result<Feedback>>({
        url: `/api/v1/feedbacks/` + params.id,
        method: 'put',
        params: { action: 'resolved' }
    })
}
