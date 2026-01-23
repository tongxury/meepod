export declare type Result<T> = {
    code: number
    message: string
    data: T
}

export declare type PageData<T> = {
    current: number,
    list: T[],
    size: number,
    total: number
    no_more: boolean
}

export declare type Account = {
    id: string
    user: User
    store: Store
    balance: number
    created_at: number
    status: Enum
}

declare declare type PayMethod = {
    id: string
    name: string
    color: string
}

export declare type Order = {
    id: string
    plan: Plan
    store: Store
    user: User
    volume: number
    amount: number
    group: OrderGroup
    follow_order_id: string
    created_at: number
    status: Enum
    ticket_images: string[]
    prized: boolean
    cancelable: boolean
    rejectable: boolean
    acceptable: boolean
    ticketable: boolean
    followable: boolean
    payable: boolean
    need_upload: boolean
    tags: Tag[]
}


export declare type Item = {
    id: string
    name: string
    icon: string
}
export declare type ItemState = {
    id: string
    name: string
    icon: string
    latest_issue: Issue
    disabled: boolean
    status: Enum
    extra: Extra
}

export declare type Extra = {
    type: 'countdown' | 'text'
    value: any
}

export declare type Enum = {
    name: string
    value: string
    color: string
    desc: string
}

export declare type User = {
    id: string
    phone: string
    nickname: string
    wechat_pay_qrcode: string
    ali_pay_qrcode: string
    desc: string
    icon: string
    created_at: number,
    status: Enum
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
    multiple: number
    amount: number
    user: User
    result: string
    created_at: number
    status: Enum
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
    status: Enum,
    extra: any
}

export declare type Store = {
    id: string
    name: string
    owner: User
    created_at: number
    icon: string
    status: Enum
    notice: string
}

export declare type OrderGroup = {
    id: string
    plan: Plan
    user: User
    volume: number,
    volume_ordered: number
    floor: number,
    reward_rate: number,
    remark: string,
    created_at: number
    status: Enum
    joiner_count: number,
    tags: Tag[],
    joinable: boolean
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
    status: Enum
    tags: Tag[]
}

export declare type Reward = {
    id: string
    user: User
    order_id: string
    amount: number
    store: Store
    created_at: number
    status: Enum
    rewardable: boolean
}

export declare type UserProfile = {
    user: User
    // account: Account
}


export declare type PaymentOrder = {
    id: string
    store: Store
    user: User
    amount: number
    created_at: number
    status: Enum,
    cancelable: boolean
}

export declare type Topup = {
    id: string
    store: Store
    user: User
    amount: number
    created_at: number
    status: Enum,
    category: Enum,
    payable: boolean
    payed: boolean
    // pay_url: string,
    qr_code: string,
    pay_method: string,
    cancelable: boolean
    time_left: number
}


export declare type Withdraw = {
    id: string
    store: Store
    user: User
    amount: number
    created_at: number
    status: Enum,
    payable: boolean
    cancelable: boolean
    remark: string
    image: string
}


export declare type Payment = {
    id: string
    store: Store
    user: User
    amount: number
    created_at: number
    status: Enum,
    biz_id: string
    biz_category: Enum
    category: Enum
    remark: string
}
export declare type Proxy = {
    id: string
    user: User
    created_at: string
    status: Enum
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

export declare type MatchCategory = {
    name: string
    value: string
}

export declare type OddsItem = {
    name: string
    result: string
    value: number
}

export declare type Odds = {
    items?: OddsItem[]
    r_items?: OddsItem[]
    // r_count: number,
    goals_items?: OddsItem[]
    half_full_items?: OddsItem[]
    score_victory_items?: OddsItem[]
    score_dogfall_items?: OddsItem[]
    score_defeat_items?: OddsItem[]
}

export declare type Match = {
    id: string,
    league: string,
    home_team: string,
    home_team_tag: string,
    guest_team: string,
    guest_team_tag: string,
    category: MatchCategory,
    issue: null,
    start_at: string,
    start_at_ts: number,
    close_at_ts: number,
    result: {
        goals: string,
        half_goals: string,
        value: string
        half_value: string
    },
    status: Enum,
    odds: Odds
    r_count: number
}

export declare type Files = {
    files: Image[]
}

export declare type Image = {
    key: string
    url: string
}


export declare type Feedback = {
    id: string
    user: User
    text: string
    created_at: number
    status: Enum,
    resolvable: boolean
}

export declare type ProxyUser = {
    user: User
    created_at: string
    created_at_ts: number
    order_count: number
    order_amount: number
}
