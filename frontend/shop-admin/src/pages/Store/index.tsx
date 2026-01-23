import services from '@/services/demo';
import {
    ActionType,
    FooterToolbar,
    PageContainer,
    ProDescriptions,
    ProDescriptionsItemProps,
    ProTable,
} from "@ant-design/pro-components";
import { Avatar, Button, Divider, Drawer, Input, message, Tag } from 'antd';
import React, { useRef, useState } from 'react';
import CreateForm from './components/CreateForm';
import MemberForm from './components/MemberForm';
import {

    updateStoreMember,
    confirmStore,
    deleteStore,
    fetchStores,
    addStorePayment
} from "@/services";
import { PlusOutlined } from "@ant-design/icons";
import TopUpForm from "@/pages/Store/components/TopUpForm";

const TableList: React.FC<unknown> = () => {
    const actionRef = useRef<ActionType>();
    const [selectedRowsState, setSelectedRows] = useState<any[]>([]);
    const columns: ProDescriptionsItemProps<any>[] = [
        {
            title: '邀请码',
            dataIndex: 'id',
            hideInForm: true,
        },
        {
            title: '头像',
            dataIndex: 'icon',
            render: (dom, entity, index, action, schema) => {
                return <Avatar src={entity?.icon} />
            },
            hideInForm: true,
            hideInSearch: true,
        },
        {
            title: '名称',
            dataIndex: 'name',
            // tip: '名称是唯一的 key',
            formItemProps: {
                rules: [
                    {
                        required: true,
                        message: '名称为必填项',
                    },
                ],
            },

        },
        {
            title: '店主手机',
            dataIndex: 'owner',
            render: (dom, entity, index, action, schema) => {
                return <div>{entity.owner?.phone}</div>
            },
            // renderFormItem: (schema, config, form, action) => {
            //     return <Input status={}/>
            // },
            formItemProps: {
                rules: [
                    {
                        required: true,
                        message: '店主手机为必填项',
                    },
                ],
            },
        },
        {
            title: '创建时间',
            dataIndex: 'created_at',
            valueType: 'text',
            hideInForm: true,
            hideInSearch: true,
        },
        {
            title: '状态',
            dataIndex: 'status',
            render: (dom, entity, index, action, schema) => {
                return <Tag color={entity.status?.color}>{entity.status?.name}</Tag>
            },
            hideInForm: true,
            hideInSearch: true,

        },
        {
            title: '余额',
            dataIndex: 'balance',
        },
        {
            title: '会员等级',
            dataIndex: 'member_level',
            render: (_, record) => (
                <>{record.member_level?.name}(截止到{record.member_until})</>
            )
        },
        {
            title: '操作',
            dataIndex: 'option',
            valueType: 'option',
            render: (_, record) => (
                <>
                    <CreateForm
                        initialValues={record}
                        onSubmit={() => actionRef.current?.reloadAndRest?.()}
                    >
                        <a>编辑</a>
                    </CreateForm>
                    <Divider type="vertical" />
                    <MemberForm
                        onSubmit={async (values: any) => {
                            const params = {
                                ...values,
                            }

                            await updateStoreMember([record.id], params)

                            actionRef.current?.reload?.()
                            return true
                        }}
                        initialValues={{ level: record.member_level?.value, until: record.member_until }}>
                        <a>会员</a>
                    </MemberForm>
                    <Divider type="vertical" />
                    <TopUpForm
                        onSubmit={async (values: any) => {

                            const params = {
                                ...values,
                            }

                            await addStorePayment(record.id, params)

                            actionRef.current?.reload?.()
                            return true
                        }}>
                        <a>充值</a>
                    </TopUpForm>

                </>
            ),
        },
    ];

    return (
        <PageContainer
            header={{ title: '店铺', }}
        >
            <ProTable<any>
                headerTitle="店铺列表"
                actionRef={actionRef}
                rowKey="id"
                search={{
                    span: 6
                    // labelWidth: 120,
                }}
                toolBarRender={() => [
                    <CreateForm onSubmit={() => actionRef.current?.reloadAndRest?.()}>
                        <Button type="primary">
                            <PlusOutlined />
                            创建店铺
                        </Button>
                    </CreateForm>
                ]}
                request={async (params, sorter, filter) => {
                    const { data, success } = await fetchStores({
                        ...params,
                        sorter,
                        filter,
                    });
                    return {
                        data: data?.list || [],
                        total: data?.total,
                        success,
                    };
                }}
                columns={columns as any}
                rowSelection={{
                    onChange: (_, selectedRows) => setSelectedRows(selectedRows),
                }}
            />
            {selectedRowsState?.length > 0 && (
                <FooterToolbar
                    extra={
                        <div>
                            已选择{' '}
                            <a style={{ fontWeight: 600 }}>{selectedRowsState.length}</a>{' '}
                            项&nbsp;&nbsp;
                        </div>
                    }
                >
                    <Button
                        onClick={async () => {

                            await deleteStore(selectedRowsState.map(t => t.id))

                            setSelectedRows([]);
                            actionRef.current?.reloadAndRest?.();
                        }}
                    >
                        批量删除
                    </Button>
                    <Button onClick={async () => {
                        await confirmStore(selectedRowsState.map(t => t.id))

                        setSelectedRows([]);
                        actionRef.current?.reloadAndRest?.();
                    }} type="primary">批量审批</Button>
                </FooterToolbar>
            )}
        </PageContainer>
    );
};

export default TableList;
