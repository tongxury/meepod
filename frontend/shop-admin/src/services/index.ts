import {request} from "@umijs/max";

export async function fetchStores(
    params?: { [key: string]: any },
    options?: { [key: string]: any },
) {
    return request('/api/v1/stores', {
        method: 'GET',
        params: {
            ...params,
        },
        ...(options || {}),
    });
}

export async function fetchUsers(
    params?: { [key: string]: any },
    options?: { [key: string]: any },
) {
    return request('/api/v1/users', {
        method: 'GET',
        params: {
            ...params,
        },
        ...(options || {}),
    });
}

export async function fetchOrders(
    params?: { [key: string]: any },
    options?: { [key: string]: any },
) {
    return request('/api/v1/orders', {
        method: 'GET',
        params: {
            ...params,
        },
        ...(options || {}),
    });
}


export async function addStore(
    data?: { [key: string]: any },
    options?: { [key: string]: any },
) {
    return request('/api/v1/stores', {
        method: 'post',
        data: {
            ...data,
        },
        ...(options || {}),
    });
}

export async function updateStoreMember(
    ids: string[],
    data: { level: number, until: number },
    options?: { [key: string]: any },
) {
    return request('/api/v1/stores', {
        method: 'put',
        params: {
            ids: ids.join(','),
            action: 'updateMember'
        },
        data: {
            ...data,
        },
        ...(options || {}),
    });
}

export async function updateStore(
    ids: string[],
    data?: { [key: string]: any },
    options?: { [key: string]: any },
) {
    return request('/api/v1/stores', {
        method: 'put',
        params: {
            ids: ids.join(','),
            action: 'update'
        },
        data: {
            ...data,
        },
        ...(options || {}),
    });
}

export async function deleteStore(
    ids: string[],
    options?: { [key: string]: any },
) {
    return request('/api/v1/stores', {
        method: 'put',
        params: {
            ids: ids.join(','),
            action: 'delete'
        },
        ...(options || {}),
    });
}

export async function confirmStore(
    ids: string[],
    options?: { [key: string]: any },
) {
    return request('/api/v1/stores', {
        method: 'put',
        params: {
            ids: ids.join(','),
            action: 'confirm'
        },
        ...(options || {}),
    });
}


export async function fetchLocs(
    options?: { [key: string]: any },
) {
    return request('/api/v1/locs', {
        method: 'GET',
        ...(options || {}),
    });
}

export async function fetchStats(
    options?: { [key: string]: any },
) {
    return request('/api/v1/stats', {
        method: 'GET',
        ...(options || {}),
    });
}


export async function fetchPaymentStores(
    params?: { [key: string]: any },
    options?: { [key: string]: any },
) {
    return request('/api/v1/payment-stores', {
        method: 'GET',
        params: {
            ...params,
        },
        ...(options || {}),
    });
}

export async function fetchBanks(
    params?: { branchName?: string, branchNo?: string },
    options?: { [key: string]: any },
) {
    return request('/api/v1/banks', {
        method: 'GET',
        params: {
            ...params,
        },
        ...(options || {}),
    });
}

export async function addPaymentStore(
    data: { store_id: string, app_id: string },
    options?: { [key: string]: any },
) {
    return request('/api/v1/payment-stores', {
        method: 'post',
        data: {
            ...data,
        },
        ...(options || {}),
    });
}


export async function addStorePayment(
    storeId: string,
    data: { amount: number },
    options?: { [key: string]: any },
) {
    return request(`/api/v1/stores/${storeId}/payments`, {
        method: 'post',
        data: {
            ...data,
        },
        ...(options || {}),
    });
}

export async function updatePaymentStoreApplyXinsh(
    data: { store_id: string},
    options?: { [key: string]: any },
) {
    return request(`/api/v1/payment-stores/${data.store_id}`, {
        method: 'patch',
        params: {
            action: 'applyXinsh',
        },
        ...(options || {}),
    });
}
