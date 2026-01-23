create table public.p_co_store_payments
(
    id           bigserial
        primary key,
    biz_id       varchar(20) default ''::character varying        not null,
    biz_category varchar(20) default ''::character varying        not null,
    category     varchar(20) default ''::character varying        not null,
    amount       numeric,
    store_id     varchar(50)                                      not null,
    co_store_id  varchar(50)                                      not null,
    created_at   timestamp   default now()                        not null,
    status       varchar(20) default 'pending'::character varying not null,
    extra        jsonb       default '{}'::jsonb
);

alter table public.p_co_store_payments
    owner to postgres;

create unique index p_co_store_payments_category_biz_category_biz_id_store_id_c_idx
    on public.p_co_store_payments (category, biz_category, biz_id, store_id, co_store_id)
    where ((biz_id)::text <> ''::text);

create table public.p_co_store_wallets
(
    id          bigserial
        primary key,
    store_id    varchar(50)                                                   not null,
    co_store_id varchar(50)                                                   not null,
    balance     numeric                  default 0                            not null,
    created_at  timestamp with time zone default now()                        not null,
    status      varchar(20)              default 'pending'::character varying not null
);

alter table public.p_co_store_wallets
    owner to postgres;

create unique index p_co_store_wallets_store_id_co_store_id_idx
    on public.p_co_store_wallets (store_id, co_store_id);

create table public.p_store_payments
(
    id           bigserial
        primary key,
    biz_id       varchar(20) default ''::character varying        not null,
    biz_category varchar(20) default 'plan'::character varying    not null,
    category     varchar(20) default ''::character varying        not null,
    amount       numeric,
    store_id     varchar(50)                                      not null,
    created_at   timestamp   default now()                        not null,
    status       varchar(20) default 'pending'::character varying not null,
    extra        jsonb       default '{}'::jsonb
);

alter table public.p_store_payments
    owner to postgres;

create unique index p_store_payments_biz_category_biz_id_store_id_idx
    on public.p_store_payments (biz_category, biz_id, store_id)
    where ((biz_id)::text <> ''::text);

create table public.p_store_wallets
(
    store_id   varchar(50)                                      not null,
    balance    numeric,
    created_at timestamp   default now()                        not null,
    status     varchar(20) default 'pending'::character varying not null,
    extra      jsonb       default '{}'::jsonb
);

alter table public.p_store_wallets
    owner to postgres;

create unique index p_store_wallets_store_id_idx
    on public.p_store_wallets (store_id);

create table public.p_stores
(
    store_id   varchar                                          not null,
    created_at timestamp   default now()                        not null,
    status     varchar(20) default 'pending'::character varying not null,
    alipay     jsonb       default jsonb_build_object()         not null,
    xinsh      jsonb       default jsonb_build_object()         not null
);

alter table public.p_stores
    owner to postgres;

create unique index p_stores_store_id_idx
    on public.p_stores (store_id);

create table public.p_user_topups
(
    id           bigserial
        primary key,
    store_id     varchar(50)                                                   not null,
    user_id      bigint                                                        not null,
    amount       numeric                                                       not null,
    created_at   timestamp with time zone default now()                        not null,
    status       varchar                  default 'pending'::character varying not null,
    extra        jsonb                    default jsonb_build_object()         not null,
    category     varchar(50)              default ''::character varying        not null,
    biz_id       varchar(50)              default ''::character varying        not null,
    biz_category varchar(20)              default ''::character varying        not null,
    sync_settled integer                  default 1                            not null
);

alter table public.p_user_topups
    owner to postgres;

create unique index p_user_topups_biz_category_biz_id_store_id_idx
    on public.p_user_topups (biz_category, biz_id, store_id)
    where ((biz_id)::text <> ''::text);

create table public.p_user_payments
(
    id           bigserial
        primary key,
    biz_id       varchar(20) default ''::character varying        not null,
    biz_category varchar(20) default ''::character varying        not null,
    amount       numeric,
    user_id      bigint                                           not null,
    created_at   timestamp   default now()                        not null,
    status       varchar(20) default 'pending'::character varying not null,
    store_id     varchar(50)                                      not null,
    extra        jsonb       default '{}'::jsonb,
    category     varchar(20) default ''::character varying        not null
);

alter table public.p_user_payments
    owner to postgres;

create unique index p_payments_biz_category_biz_id_store_id_idx
    on public.p_user_payments (biz_category, biz_id, store_id)
    where ((biz_id)::text <> ''::text);

create table public.p_user_wallets
(
    id         bigserial
        primary key,
    user_id    bigint                                                        not null,
    store_id   varchar(50)                                                   not null,
    balance    numeric                  default 0                            not null,
    created_at timestamp with time zone default now()                        not null,
    status     varchar(20)              default 'pending'::character varying not null
);

alter table public.p_user_wallets
    owner to postgres;

create unique index p_accounts_user_id_store_id_idx
    on public.p_user_wallets (user_id, store_id);

create table public.p_user_withdraws
(
    id         bigserial
        primary key,
    store_id   varchar(50)                                                   not null,
    user_id    bigint                                                        not null,
    amount     numeric                                                       not null,
    created_at timestamp with time zone default now()                        not null,
    status     varchar                  default 'pending'::character varying not null,
    extra      jsonb                    default jsonb_build_object()         not null
);

alter table public.p_user_withdraws
    owner to postgres;

create table public.t_co_stores
(
    id          bigserial
        primary key,
    store_id    varchar(50)                                                   not null,
    co_store_id varchar(50)                                                   not null,
    items       jsonb                    default '{}'::jsonb                  not null,
    created_at  timestamp with time zone default now()                        not null,
    status      varchar                  default 'pending'::character varying not null,
    extra       jsonb                    default jsonb_build_object()         not null,
    sync_return integer                  default 1                            not null
);

alter table public.t_co_stores
    owner to postgres;

create unique index t_co_stores_store_id_co_store_id_idx
    on public.t_co_stores (store_id, co_store_id);

create table public.t_feedbacks
(
    id         bigserial
        primary key,
    store_id   varchar(50)                                                   not null,
    user_id    varchar(50)                                                   not null,
    text       text                     default ''::text                     not null,
    created_at timestamp with time zone default now()                        not null,
    status     varchar                  default 'pending'::character varying not null,
    extra      jsonb                    default jsonb_build_object()         not null
);

alter table public.t_feedbacks
    owner to postgres;

create index t_feedbacks_store_id_user_id_idx
    on public.t_feedbacks (store_id, user_id);

create table public.t_issues
(
    id           varchar(50)                                      not null
        constraint issues_pkey
            primary key,
    item_id      varchar(50)                                      not null,
    index        varchar(50)                                      not null,
    result       jsonb       default '[]'::jsonb                  not null,
    started_at   timestamp                                        not null,
    prized_at    timestamp                                        not null,
    status       varchar(20) default 'ongoing'::character varying not null,
    extra        jsonb       default '{}'::jsonb,
    close_at     timestamp                                        not null,
    prize_grades jsonb       default json_build_array()           not null
);

alter table public.t_issues
    owner to postgres;

create unique index issues_item_id_index_idx
    on public.t_issues (item_id, index);

create index issues_status_idx
    on public.t_issues (status);

create table public.t_items
(
    id     varchar(50)                                       not null
        constraint items_pkey
            primary key,
    name   varchar(50)  default ''::character varying        not null,
    icon   varchar(255) default ''::character varying        not null,
    status varchar(10)  default 'pending'::character varying not null,
    sort   integer      default 0
);

alter table public.t_items
    owner to postgres;

create table public.t_matches
(
    id             bigserial
        primary key,
    league         varchar(50)                                                   not null,
    home_team      varchar(50)                                                   not null,
    home_team_tag  varchar(50)              default ''::character varying        not null,
    guest_team     varchar(50)                                                   not null,
    guest_team_tag varchar(50)              default ''::character varying        not null,
    category       varchar(20)                                                   not null,
    issue          varchar(20)              default ''::character varying        not null,
    start_at       timestamp                                                     not null,
    close_at       timestamp                                                     not null,
    result         jsonb                    default jsonb_build_object()         not null,
    status         varchar                  default 'pending'::character varying not null,
    odds           jsonb                    default jsonb_build_object()         not null,
    created_at     timestamp with time zone default now(),
    real_odds      jsonb                    default jsonb_build_object(),
    r_count        integer
);

alter table public.t_matches
    owner to postgres;

create unique index matches_issue_home_team_guest_team_category_idx
    on public.t_matches (issue, home_team, guest_team, category);

create table public.t_order_group_shares
(
    id         bigserial
        primary key,
    group_id   bigint                                                        not null,
    user_id    bigint                                                        not null,
    volume     integer                                                       not null,
    amount     double precision                                              not null,
    created_at timestamp with time zone default now()                        not null,
    status     varchar                  default 'pending'::character varying not null,
    extra      jsonb                    default jsonb_build_object()         not null,
    store_id   varchar(50)                                                   not null
);

alter table public.t_order_group_shares
    owner to postgres;

create table public.t_order_groups
(
    id             bigserial
        primary key,
    plan_id        bigint                                                        not null,
    store_id       varchar(50)                                                   not null,
    volume         integer                                                       not null,
    user_id        bigint                                                        not null,
    status         varchar(10)              default 'pending'::character varying not null,
    created_at     timestamp with time zone default now()                        not null,
    volume_ordered integer                  default 0,
    floor          integer,
    remark         text,
    reward_rate    double precision,
    extra          jsonb                    default json_build_object(),
    shares         jsonb                    default json_build_object(),
    clean          integer                  default 0                            not null,
    to_store_id    varchar(50),
    issue_index    varchar(50),
    item_id        varchar(20)                                                   not null,
    sync_rollback  integer                  default 1                            not null
);

alter table public.t_order_groups
    owner to postgres;

create unique index order_groups_plan_id_idx
    on public.t_order_groups (plan_id);

create table public.t_orders
(
    id              bigserial
        primary key,
    plan_id         bigint                                                        not null,
    item_id         varchar(50)              default ''::character varying        not null,
    store_id        varchar(50)              default ''::character varying        not null,
    user_id         bigint                                                        not null,
    amount          numeric                                                       not null,
    created_at      timestamp with time zone default now()                        not null,
    status          varchar                  default 'pending'::character varying not null,
    extra           jsonb                    default jsonb_build_object()         not null,
    follow_order_id bigint,
    clean           integer                  default 0                            not null,
    to_store_id     varchar(50)              default ''::character varying,
    issue_index     varchar(50),
    sync_switch     integer                  default 1                            not null,
    sync_rollback   integer                  default 1                            not null,
    sync_proxy      integer                  default 1                            not null
);

alter table public.t_orders
    owner to postgres;

create index orders_item_id_status_idx
    on public.t_orders (item_id, status);

create index orders_user_id_follow_order_id_idx
    on public.t_orders (user_id, follow_order_id);

create table public.t_plans
(
    id         bigserial
        primary key,
    type       varchar(20)              default 'normal'::character varying  not null,
    item_id    varchar(50)              default ''::character varying        not null,
    issue      varchar(50)              default ''::character varying        not null,
    content    jsonb                    default jsonb_build_object()         not null,
    multiple   integer                  default 1                            not null,
    amount     numeric                                                       not null,
    user_id    bigint                                                        not null,
    store_id   varchar(50)                                                   not null,
    status     varchar(10)              default 'pending'::character varying not null,
    created_at timestamp with time zone default now()                        not null
);

alter table public.t_plans
    owner to postgres;

create table public.t_proxies
(
    id          varchar(20)                                                   not null
        primary key,
    store_id    varchar(50)                                                   not null,
    user_id     bigint,
    reward_rate double precision                                              not null,
    created_at  timestamp with time zone default now()                        not null,
    status      varchar                  default 'pending'::character varying not null,
    extra       jsonb                    default jsonb_build_object()         not null
);

alter table public.t_proxies
    owner to postgres;

create unique index t_proxies_store_id_user_id_idx
    on public.t_proxies (store_id, user_id);

create table public.t_proxy_rewards
(
    id            bigserial
        primary key,
    month         varchar(20)                                                   not null,
    proxy_id      varchar(20)                                                   not null,
    user_count    integer,
    order_count   integer,
    order_amount  integer,
    reward_rate   double precision,
    reward_amount double precision,
    created_at    timestamp with time zone default now()                        not null,
    status        varchar                  default 'pending'::character varying not null,
    pay_at        timestamp with time zone,
    extra         jsonb                    default jsonb_build_object()         not null,
    proxy_user_id bigint                                                        not null,
    store_id      varchar(50)                                                   not null,
    clean         integer                  default 1                            not null,
    sync_payed    integer                  default 1                            not null
);

alter table public.t_proxy_rewards
    owner to postgres;

create unique index t_proxy_rewards_month_proxy_id_idx
    on public.t_proxy_rewards (month, proxy_id);

create table public.t_proxy_users
(
    id            bigserial
        primary key,
    user_id       bigint                                                not null,
    proxy_id      varchar(20)                                           not null,
    created_at    timestamp with time zone default now()                not null,
    extra         jsonb                    default jsonb_build_object() not null,
    proxy_user_id bigint                                                not null,
    store_id      varchar(50)                                           not null
);

alter table public.t_proxy_users
    owner to postgres;

create unique index t_proxy_users_proxy_id_user_id_idx
    on public.t_proxy_users (proxy_id, user_id);

create index t_proxy_users_store_id_proxy_id_idx
    on public.t_proxy_users (store_id, proxy_id);

create table public.t_user_rewards
(
    id           bigserial
        primary key,
    biz_id       varchar(20)              default ''::character varying        not null,
    user_id      bigint                                                        not null,
    amount       numeric                  default 0                            not null,
    created_at   timestamp with time zone default now()                        not null,
    status       varchar                  default 'pending'::character varying not null,
    extra        jsonb                    default jsonb_build_object()         not null,
    store_id     varchar(50)                                                   not null,
    biz_category varchar(20)              default ''::character varying        not null,
    sync_pay     integer                  default 1                            not null
);

alter table public.t_user_rewards
    owner to postgres;

create unique index rewards_order_id_idx
    on public.t_user_rewards (biz_id);

create unique index rewards_order_category_order_id_idx
    on public.t_user_rewards (biz_category, biz_id);

create table public.t_store_users
(
    id         bigserial
        primary key,
    store_id   varchar(50)                                                   not null,
    user_id    bigint                                                        not null,
    created_at timestamp with time zone default now()                        not null,
    status     varchar                  default 'pending'::character varying not null,
    extra      jsonb                    default jsonb_build_object()         not null
);

alter table public.t_store_users
    owner to postgres;

create unique index t_store_users_store_id_user_id_idx
    on public.t_store_users (store_id, user_id);

create table public.t_stores
(
    id         varchar(50)                                                   not null
        constraint stores_pkey
            primary key,
    name       varchar(200)             default ''::character varying        not null,
    owner_id   bigint                                                        not null,
    created_at timestamp with time zone default now()                        not null,
    status     varchar                  default 'pending'::character varying not null,
    extra      jsonb                    default jsonb_build_object()         not null,
    icon       varchar(255)             default ''::character varying,
    member     jsonb                    default json_build_object(),
    settings   jsonb                    default json_build_object()
);

alter table public.t_stores
    owner to postgres;

create index stores_created_at_idx
    on public.t_stores (created_at);

create unique index stores_owner_id_idx
    on public.t_stores (owner_id);

create table public.t_banks
(
    bank_id       varchar(20),
    bank_code     varchar(20),
    bank_name     varchar(20),
    branch_name   varchar(50),
    branch_no     varchar(20),
    province_name varchar(20),
    city_name     varchar(20)
);

alter table public.t_banks
    owner to postgres;

create index t_banks_branch_name_idx
    on public.t_banks (branch_name);

create table public.t_users
(
    id         bigint generated by default as identity
        primary key,
    phone      varchar   not null,
    status     varchar   not null,
    created_at timestamp not null,
    extra      jsonb     not null
);

alter table public.t_users
    owner to postgres;

create unique index user_phone
    on public.t_users (phone);


insert into public.t_issues (id, item_id, index, started_at, prized_at, close_at)
VALUES ('x7c-2023062', 'x7c', '2023062', '2023-05-30 12:30:00', '2023-06-02 12:30:00', '2023-06-02 10:30:00'),
       ('ssq-2023062', 'ssq', '2023062', '2023-05-30 13:25:00', '2023-06-02 13:25:00', '2023-06-02 11:25:00'),
       ('f3d-2023142', 'f3d', '2023142', '2023-05-31 13:15:00', '2023-06-01 13:15:00', '2023-06-01 10:15:00'),
       ('pl3-2023142', 'pl3', '2023142', '2023-05-31 13:15:00', '2023-06-01 13:15:00', '2023-06-01 10:15:00'),
       ('pl5-2023142', 'pl5', '2023142', '2023-05-31 13:15:00', '2023-06-01 13:15:00', '2023-06-01 10:15:00'),
       ('zjc-1', 'zjc', '1', '2021-05-31 13:15:00', '2073-06-01 13:15:00', '2073-06-01 10:15:00')
    on conflict do nothing;

insert into public.t_issues (id, item_id, index, started_at, prized_at, close_at)
VALUES ('dlt-2023092', 'dlt', '2023092', '2023-08-09 12:26', '2023-08-12 12:25', '2023-08-12 10:25');


select setval('t_orders_id_seq', 100000000);
select setval('t_users_id_seq', 10000);