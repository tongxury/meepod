export declare type Result<T> = {
    code: number
    message: string
    data: T
}

export declare type PageData<T> = {
    current: number,
    list: T[],
    size: number,
    total: number,
    no_more: boolean
}

export declare type Account = {
    id: string
    user: User
    store: Store
    balance: number
    created_at: string
    status: Status,
    decrable: boolean
}

export declare type AccountSummary = {
    user_count: number
    total_balance: number
}

export declare type StoreProfile = {
    store: Store
}


export declare type Order = {
    id: string
    plan: Plan
    store: Store
    to_store: Store
    user: User
    volume: number
    amount: number
    group: OrderGroup
    follow_order_id: string
    created_at: string
    status: Status
    ticket_images: string[]
    prized: boolean
    cancelable: boolean
    rejectable: boolean
    acceptable: boolean
    ticketable: boolean
    switchable: boolean
    payable: boolean
    keeper_payable: boolean
    need_upload: boolean
    tags: Tag[]
}

export declare type Option = {
    name: string
    value: string
}

export declare type OrderFilter = {
    order: {
        dateRange: Option[]
        category: Option[]
        item: Option[]
    }
}

export declare type Item = {
    id: string
    name: string
    icon: string
}

export declare type Status = {
    name: string
    value: string
    color: string
}

export declare type User = {
    id: string
    phone: string
    icon: string
    wechat_pay_qrcode: string
    ali_pay_qrcode: string
    nickname: string
    remark: string
    desc: string
    created_at: string,
    status: Status
    tags: Tag[]
}

export declare type Tag = {
    title: string
    color: string
}

export declare type Plan = {
    id: string
    item: Item
    issue: Issue
    tickets: any[]
    split_tickets: any[]

    multiple: number
    amount: number
    user: User
    result: string
    created_at: string
    status: Status
}


export declare type Store = {
    id: string
    name: string
    owner: User
    created_at: string
    icon: string
    wechat_image: string
    status: Status
    selected_item_ids: string[]
    items: Item[]
    tags: Tag[]
    balance: number
    member_level: Status
    notice: string
}

export declare type StoreUser = {
    user: User
    remark: string
}

export declare type OrderGroup = {
    id: string
    plan: Plan
    user: User
    store: Store,
    to_store: Store,
    amount: number,
    volume: number,
    volume_ordered: number
    created_at: string
    status: Status
    joiner_count: number,
    tags: Tag[],
    remark: string,
    floor: number,
    reward_rate: number,
    acceptable: boolean,
    rejectable: boolean,
    ticketable: boolean,
    switchable: boolean,
    need_upload: boolean
    ticket_images: string[]
    prized: boolean
}


export declare type GroupShare = {
    id: string,
    user: User
    is_creator: boolean
    volume: number
    amount: number
    group: OrderGroup
    created_at_ts: number,
    create_at: string,
    status: Status
    tags: Tag[]
}


export declare type GroupJoiner = {
    user: User
    is_creator: boolean
    volume: number
}


export declare type Issue = {
    item: Item
    index: string
    close_time_left: number
    prize_time_left: number
    result: any
    prized_at: number
    started_at: number
    close_at: number
    extra: any
}


export declare type Reward = {
    id: string
    user: User
    biz_id: string
    biz_category: Status
    amount: number
    store: Store
    created_at: string
    status: Status
    rewardable: boolean
    rejectable: boolean
}


export declare type PaymentOrder = {
    id: string
    store: Store
    user: User
    amount: number
    created_at: string
    status: Status,
    cancelable: boolean,
    acceptable: boolean,
    rejectable: boolean

}


export declare type Files = {
    files: Asset[]
}

export declare type Asset = {
    key: string
    url: string
}

export declare type Counter = {
    [key: string]: number
}


export declare type Withdraw = {
    id: string
    store: Store
    user: User
    amount: number
    created_at: number
    status: Status,
    payable: boolean
    cancelable: boolean
}

export declare type Proxy = {
    id: string
    user: User
    created_at: string
    status: Status
    reward_rate: number
    reward_amount: number
    user_count: number
    order_count: number
    order_amount: number
    deletable: boolean
    updatable: boolean
    addable: boolean
    recoverable: boolean

}

export declare type ProxyUser = {
    user: User
    created_at: string
    created_at_ts: number
    order_count: number
    order_amount: number
}


export declare type ProxyReward = {
    id: string
    user: User
    created_at: string
    status: Status
    reward_rate: number
    reward_amount: number
    user_count: number
    order_count: number
    order_amount: number
    month: string,
    pay_at: string,
    payable: boolean
}

export declare type CoStore = {
    id: string
    store: Store
    co_store: Store
    items: CoItem[]
    created_at: string
    status: Status
    balance: number
    topupable: boolean
    updatable: boolean
    resumable: boolean
    pausable: boolean
    endapplyable: boolean
    endable: boolean
    recoverable: boolean
}

export declare type CoItem = {
    item: Item
    rate: number
}

export declare type Feedback = {
    id: User
    user: User
    text: string
    created_at: number
    status: Enum,
    resolvable: boolean
}
