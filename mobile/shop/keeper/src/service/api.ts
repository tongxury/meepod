import axios from "axios";
import Constants from 'expo-constants';
import AsyncStorage from "@react-native-async-storage/async-storage";
import { Toast } from "@ant-design/react-native";
import {
    Account,
    AccountSummary,
    Order,
    OrderFilter,
    PageData,
    PaymentOrder,
    Asset,
    Result,
    Reward,
    Store,
    StoreProfile,
    Counter,
    Withdraw,
    OrderGroup,
    GroupShare,
    Item,
    Proxy,
    User,
    ProxyUser,
    ProxyReward,
    CoStore,
    Feedback, StoreUser, Files
} from "./typs";
import nav from "../utils/navigation";
import { Platform } from "react-native";
import { Config } from "../../config";

const instance = axios.create({});

instance.interceptors.request.use(async function (config) {
    // Do something before request is sent

    const conf = Config()
    if (config.url.startsWith("/api/payment/")) {
        config.baseURL = conf.apiPaymentUrl || conf.apiUrl
    } else {
        config.baseURL = conf.apiUrl
    }

    // config.headers.set('Auth-Token', await AsyncStorage.getItem('user_auth') || '')
    // config.headers.set('Store', await AsyncStorage.getItem('store') || '')
    const token = await AsyncStorage.getItem('user_auth') ?? ''
    config.headers.set('Authorization', token ? 'Bearer ' + token : '')
    config.headers.set('Platform', Platform.OS)
    config.headers.set('Client', 'keeper')
    config.headers.set('Version', Constants.expoConfig.version ?? '')
    config.headers.set('U-Version', conf.version)


    if (__DEV__) {
        console.log(`>>> [API Request] ${config.method?.toUpperCase()} ${config.baseURL + (config.url || '')}`, {
            params: config.params,
            data: config.data,
            headers: config.headers,
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


    if (response.status !== 200) {
        Toast.info({
            content: '服务器繁忙, 请稍后重试',
            mask: false,
        })
    } else {
        if (response.data.code !== 0) {
            console.error(response.data?.error)
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
                nav.navigate('Register')
        }

    }

    return response;
}, function (error) {

    console.error(error)
    // Any status codes that falls outside the range of 2xx cause this function to trigger
    // Do something with response error
    return Promise.reject(error);
});

export const uploadImage = ({ file }: { file: any }) => {

    const fd = new FormData()
    fd.append("files", file)

    return instance.request<Result<string[]>>({
        method: "post",
        url: `/api/v1/images`,
        data: fd
    })
}

// export const updateStoreProfile = ({field, value}: { field: string, value: string }) => {
//     return instance.request<Result<StoreProfile>>({
//         method: "put",
//         url: `/api/v1/stores/me/profiles`,
//         params: {
//             field, value
//         }
//     })
// }

export const updateStoreItems = ({ itemIds }: { itemIds: string[] }) => {
    return instance.request<Result<Store>>({
        method: "put",
        url: `/api/v1/stores/me`,
        params: { category: "items" },
        data: { itemIds }
    })
}

export const updateStoreNotice = ({ notice }: { notice: string }) => {
    return instance.request<Result<Store>>({
        method: "put",
        url: `/api/v1/stores/me`,
        params: { category: "notice" },
        data: { notice }
    })
}


export const updateStore = ({ field, value }: { field: string, value: string }) => {
    return instance.request<Result<Store>>({
        method: "put",
        url: `/api/v1/stores/me`,
        params: {
            field, value
        }
    })
}

// export const fetchStores = (params: { [key: string]: any, page?: number, size?: number }) => {
//     return instance.request<Result<PageData<Store>>>({
//         url: `/api/v1/stores/me/stores`,
//         params
//     })
// }

export const fetchStore = ({ id }) => {
    return instance.request<Result<Store>>({
        url: `/api/v1/stores/` + id,
    })
}

export const fetchMyStore = () => {
    return instance.request<Result<Store>>({
        url: `/api/v1/stores/me/profiles`,
    })
}

// export const fetchStoreProfile = ({storeId}: { storeId: string }) => {
//     return instance.request<Result<StoreProfile>>({
//         url: `/api/v1/stores/${storeId}/profiles`,
//     })
// }


export const payOrder = ({ orderId }: { orderId: string }) => {
    return instance.request<Result<Order>>({
        method: "put",
        url: `/api/v1/stores/me/orders/${orderId}`,
        params: { action: "pay" },
    })
}


export const acceptOrder = ({ orderId }: { orderId: string }) => {
    return instance.request<Result<Order>>({
        method: "put",
        url: `/api/v1/stores/me/orders/${orderId}`,
        params: { action: "accept" },
    })
}
export const rejectOrder = ({ orderId, reasonId }: { orderId: string, reasonId: string }) => {
    return instance.request<Result<Order>>({
        method: "put",
        url: `/api/v1/stores/me/orders/${orderId}`,
        params: { action: "reject", reasonId },
    })
}

export const ticketOrder = ({ orderId, images }: { orderId: string, images: string[] }) => {
    return instance.request<Result<Order>>({
        method: "put",
        url: `/api/v1/stores/me/orders/${orderId}`,
        params: { action: "ticket" },
        data: {
            images
        }
    })
}

export const switchOrder = ({ orderId, toStoreId }: { orderId: string, toStoreId: string }) => {
    return instance.request<Result<Order>>({
        method: "put",
        url: `/api/v1/stores/me/orders/${orderId}`,
        params: { action: "switch", toStoreId },
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


export const fetchOrderFilters = () => {
    return instance.request<Result<OrderFilter>>({
        url: `/api/v1/stores/me/order-filters`,
    })
}

export const fetchOrders = (params: { [key: string]: any; page?, size?}) => {
    return instance.request<Result<PageData<Order>>>({
        url: '/api/v1/stores/me/orders',
        params
    })
}

export const fetchAccounts = (params: { page?, size?}) => {
    return instance.request<Result<PageData<Account>>>({
        url: '/api/payment/v1/stores/me/accounts',
        params
    })
}
export const fetchAccountSummary = (params: { page?, size?}) => {
    return instance.request<Result<AccountSummary>>({
        url: '/api/payment/v1/stores/me/account-summaries',
        params
    })
}

export const topUp = (params: { id: string, amount: number, remark: string }) => {
    return instance.request<Result<any>>({
        url: `/api/payment/v1/stores/me/accounts/${params.id}`,
        method: "put",
        params: { action: 'topUp' },
        data: {
            amount: params.amount,
            remark: params.remark
        }
    })
}

export const decrease = (params: { id: string, amount: number, remark: string }) => {
    return instance.request<Result<any>>({
        url: `/api/payment/v1/stores/me/accounts/${params.id}`,
        method: "put",
        params: { action: 'decrease' },
        data: {
            amount: params.amount,
            remark: params.remark
        }
    })
}


export const fetchRewards = (params: { page?, size?}) => {
    return instance.request<Result<PageData<Reward>>>({
        url: '/api/v1/stores/me/rewards',
        params
    })
}


export const sendReward = ({ rewardId }: { rewardId: string }) => {
    return instance.request<Result<Reward>>({
        method: "put",
        url: `/api/v1/stores/me/rewards/${rewardId}`,
        params: { action: "reward" },
    })
}
export const rejectReward = ({ rewardId, reason }: { rewardId: string, reason: string }) => {
    return instance.request<Result<Reward>>({
        method: "put",
        url: `/api/v1/stores/me/rewards/${rewardId}`,
        params: { action: "reject", reason },
    })
}


export const sendCode = ({ phone, event }: { phone: string, event?: string }) => {
    event = event || 'login'

    return instance.request({
        method: "post",
        url: '/api/v1/auth-codes',
        params: { phone, event }
    })
}
export const login = ({ phone, code }: { phone: string, code: string }) => {
    return instance.request({
        url: '/api/v1/auth-tokens',
        params: {
            phone: phone, code: code
        }
    })
}

export const loginByPassword = ({ phone, password }: { phone: string, password: string }) => {
    return instance.request({
        url: '/api/v1/auth-tokens',
        params: {
            phone: phone, password: password
        }
    })
}
export const fetchAuthStatus = ({ token }: { token: string }) => {
    return instance.request({
        url: '/api/v1/auth-status',
        params: {
            token
        }
    })
}
export const fetchStoreInfo = () => {
    return instance.request({
        url: '/api/v1/stores/me/profiles',
    })
}

export const fetchSettings = () => {
    return instance.request({
        url: '/api/v1/settings'
    })
}

export const fetchOrderGroups = (params: { page?, size?}) => {
    return instance.request<Result<PageData<OrderGroup>>>({
        url: '/api/v1/stores/me/order-groups',
        params
    })
}
export const fetchOrderGroupDetail = ({ id, biz_category }: { id: string, biz_category?: string }) => {
    return instance.request<Result<OrderGroup>>({
        url: `/api/v1/stores/me/order-groups/${id}`,
        params: { biz_category }
    })
}


export const acceptOrderGroup = ({ orderId }: { orderId: string }) => {
    return instance.request<Result<OrderGroup>>({
        method: "put",
        url: `/api/v1/stores/me/order-groups/${orderId}`,
        params: { action: "accept" },
    })
}
export const rejectOrderGroup = ({ orderId, reasonId }: { orderId: string, reasonId: string }) => {
    return instance.request<Result<OrderGroup>>({
        method: "put",
        url: `/api/v1/stores/me/order-groups/${orderId}`,
        params: { action: "reject", reasonId },
    })
}

export const ticketOrderGroup = ({ orderId, images }: { orderId: string, images: string[] }) => {
    return instance.request<Result<OrderGroup>>({
        method: "put",
        url: `/api/v1/stores/me/order-groups/${orderId}`,
        params: { action: "ticket" },
        data: {
            images
        }
    })
}

export const switchOrderGroup = ({ orderId, toStoreId }: { orderId: string, toStoreId: string }) => {
    return instance.request<Result<OrderGroup>>({
        method: "put",
        url: `/api/v1/stores/me/order-groups/${orderId}`,
        params: { action: "switch", toStoreId },
    })
}


export const fetchOrderDetail = ({ id }: { id: string }) => {
    return instance.request<Result<Order>>({
        url: `/api/v1/stores/me/orders/${id}`,
    })
}

export const fetchCurrentIssue = (params: { itemId }) => {
    return instance.request({
        url: '/api/v1/current-issues',
        params
    })
}

export const pay = ({ orderId, method }: { orderId: string, method?: string }) => {
    return instance.request({
        url: '/api/v1/payments',
        method: "post",
        data: { orderId },
        params: { method }
    })
}


export const fetchWithdraws = ({ page, size }: { page?: number, size?: number }) => {
    return instance.request<Result<PageData<Withdraw>>>({
        url: `/api/payment/v1/stores/me/withdraws`,
        params: { page, size }
    })
}


export const acceptWithdraw = ({ id, image }: { id: string, image: string }) => {
    return instance.request<Result<Withdraw>>({
        url: `/api/payment/v1/stores/me/withdraws/${id}`,
        method: "put",
        params: {
            action: 'accept',
            image,
        },
    })
}
export const rejectWithdraw = ({ id, reason }: { id: string, reason: string }) => {
    return instance.request<Result<Withdraw>>({
        url: `/api/payment/v1/stores/me/withdraws/${id}`,
        method: "put",
        params: {
            action: 'reject'
        },
        data: {
            reason
        }
    })
}

export const listCounters = () => {
    return instance.request<Result<Counter>>({
        url: `/api/v1/stores/me/counters`,
    })
}

export const fetchOrderGroupOrders = (params: { groupId, page?, size?}) => {
    return instance.request<Result<PageData<GroupShare>>>({
        url: `/api/v1/stores/me/order-groups/${params.groupId}/shares`,
        params: params
    })
}


export const fetchItems = () => {
    return instance.request<Result<Item[]>>({
        url: `/api/v1/stores/me/items`,
    })
}
// export const updateItemSettings = ({itemIds}: { itemIds: string[] }) => {
//     return instance.request<Result<ItemSettings>>({
//         url: `/api/v1/stores/me/items`,
//         method: 'put',
//         data: {
//             itemIds
//         }
//     })
// }


export const fetchProxies = (params: { page?: number, size?: number }) => {
    return instance.request<Result<PageData<Proxy>>>({
        url: `/api/v1/stores/me/proxies`,
        params: params
    })
}
export const addProxy = (params: { userId: string, rewardRate: string }) => {
    return instance.request<Result<PageData<Proxy>>>({
        url: `/api/v1/stores/me/proxies`,
        method: 'post',
        data: params
    })
}

export const deleteProxy = (proxyId) => {
    return instance.request<Result<Proxy>>({
        url: `/api/v1/stores/me/proxies/` + proxyId,
        method: 'put',
        params: { action: 'delete' }
    })
}
export const recoverProxy = (proxyId) => {
    return instance.request<Result<Proxy>>({
        url: `/api/v1/stores/me/proxies/` + proxyId,
        method: 'put',
        params: { action: 'recover' }
    })
}
export const updateProxy = (proxyId, rewardRate) => {
    return instance.request<Result<Proxy>>({
        url: `/api/v1/stores/me/proxies/` + proxyId,
        method: 'put',
        params: { action: 'updateRewardRate', rewardRate }
    })
}

export const addProxyUser = (proxyId, userId) => {
    return instance.request<Result<Proxy>>({
        url: `/api/v1/stores/me/proxies/${proxyId}/users`,
        method: 'post',
        params: { userId }
    })
}

export const fetchStoreUsers = (params: { phone: string }) => {
    return instance.request<Result<PageData<StoreUser>>>({
        url: `/api/v1/stores/me/users`,
        params: params
    })
}


export const fetchProxyUsers = (params: { proxyId: string, page?: number, size?: number }) => {
    return instance.request<Result<PageData<ProxyUser>>>({
        url: `/api/v1/stores/me/proxies/${params.proxyId}/users`,
        params: params
    })
}

export const fetchProxyRewards = (params: { month: string, cat: string, page?: number, size?: number }) => {
    return instance.request<Result<PageData<ProxyReward>>>({
        url: `/api/v1/stores/me/proxy-rewards`,
        params: params
    })
}

export const payProxyReward = (params: { rewardId: string }) => {
    return instance.request<Result<PageData<ProxyReward>>>({
        url: `/api/v1/stores/me/proxy-rewards/` + params.rewardId,
        method: 'put',
        params: { action: 'pay' }
    })
}

export const fetchCoStores = (params: { [key: string]: any, page?: number, size?: number }) => {
    return instance.request<Result<PageData<CoStore>>>({
        url: `/api/v1/stores/me/co-stores`,
        params: params
    })
}

export const addCoStore = (params: { coStoreId: string, itemIds: string[] }) => {
    return instance.request<Result<PageData<CoStore>>>({
        url: `/api/v1/stores/me/co-stores`,
        method: 'post',
        data: params
    })
}

export const updateCoStore = (coStoreId, items) => {
    return instance.request<Result<CoStore>>({
        url: `/api/v1/stores/me/co-stores/` + coStoreId,
        method: 'put',
        params: { action: 'updateItems' },
        data: { items }
    })
}

export const applyEndCoStore = (coStoreId) => {
    return instance.request<Result<CoStore>>({
        url: `/api/v1/stores/me/co-stores/` + coStoreId,
        method: 'put',
        params: { action: 'endApply' },
    })
}

export const endCoStore = (coStoreId, imageProof) => {
    return instance.request<Result<CoStore>>({
        url: `/api/v1/stores/me/co-stores/` + coStoreId,
        method: 'put',
        params: { action: 'end', imageProof },
    })
}

export const pauseCoStore = (coStoreId) => {
    return instance.request<Result<CoStore>>({
        url: `/api/v1/stores/me/co-stores/` + coStoreId,
        method: 'put',
        params: { action: 'pause' },
    })
}

export const resumeCoStore = (coStoreId) => {
    return instance.request<Result<CoStore>>({
        url: `/api/v1/stores/me/co-stores/` + coStoreId,
        method: 'put',
        params: { action: 'resume' },
    })
}

export const recoverCoStore = (coStoreId) => {
    return instance.request<Result<CoStore>>({
        url: `/api/v1/stores/me/co-stores/` + coStoreId,
        method: 'put',
        params: { action: 'recover' },
    })
}

export const addCoStoreTopUp = ({ storeId, amount }) => {
    return instance.request<Result<any>>({
        url: `/api/payment/v1/stores/me/co-store-payments`,
        method: 'post',
        data: { storeId, amount }
    })
}

export const fetchCoStorePayments = (params: { [key: string]: any, page?: number, size?: number }) => {
    return instance.request<Result<any>>({
        url: `/api/payment/v1/stores/me/co-store-payments`,
        params
    })
}

export const fetchStorePayments = (params: { [key: string]: any, page?: number, size?: number }) => {
    return instance.request<Result<any>>({
        url: `/api/payment/v1/stores/me/store-payments`,
        params
    })
}


export const addFeedback = (params: { text: string }) => {
    return instance.request<Result<Feedback>>({
        url: `/api/v1/stores/me/feedbacks`,
        method: 'post',
        data: params
    })
}

export const fetchFeedbacks = (params: { page?, size?}) => {
    return instance.request<Result<PageData<Feedback>>>({
        url: `/api/v1/stores/me/feedbacks`,
        params
    })
}

export const fetchStoreUser = ({ userId }: { userId: string }) => {
    return instance.request<Result<StoreUser>>({
        url: `/api/v1/stores/me/users/${userId}`,
    })
}

export const updateStoreUser = ({ userId, field, value }: { userId: string, field: string, value: string }) => {
    return instance.request<Result<StoreUser>>({
        url: `/api/v1/stores/me/users/${userId}`,
        method: 'put',
        params: { field, value }
    })
}
