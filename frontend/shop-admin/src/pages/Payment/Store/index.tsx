import services from '@/services/demo';
import {
    ActionType,
    FooterToolbar,
    PageContainer,
    ProDescriptions,
    ProDescriptionsItemProps,
    ProTable,
} from '@ant-design/pro-components';
import {Avatar, Button, Divider, Drawer, Input, Popconfirm, Tag} from 'antd';
import React, {useRef, useState} from 'react';
import {confirmStore, deleteStore, fetchPaymentStores, fetchStores, updatePaymentStoreApplyXinsh} from "@/services";
import {PlusOutlined} from "@ant-design/icons";
import CreateForm from './CreateForm';
import JSONPretty from 'react-json-pretty';


const TableList: React.FC<unknown> = () => {
    const actionRef = useRef<ActionType>();


    const [selectedRowsState, setSelectedRows] = useState<any[]>([]);
    const columns: ProDescriptionsItemProps<any>[] = [
        {
            title: '店铺',
            dataIndex: 'store',
            render: (dom, entity, index, action, schema) => {
                return <div>{entity.store?.id}({entity.store?.name})</div>
            },
            hideInForm: true,
        },
        {
            title: '新生支付',
            dataIndex: 'xinsh',
            render: (dom, entity, index, action, schema) => {
                return <JSONPretty id="json-pretty" data={entity.xinsh}></JSONPretty>

                // return <div>{JSON.stringify(entity.xinsh)}</div>
            },
            hideInForm: true,
            hideInSearch: true,
        },
        {
            title: '支付宝',
            dataIndex: 'alipay',
            render: (dom, entity, index, action, schema) => {
                return <div>{JSON.stringify(entity.alipay)}</div>
            },
            hideInForm: true,
            hideInSearch: true,
        },
        {
            title: '创建时间',
            dataIndex: 'created_at',
            valueType: 'text',
            hideInForm: true,
            hideInSearch: true,
        },
        {
            title: '操作',
            dataIndex: 'option',
            valueType: 'option',
            render: (_, record) => (
                <>
                    <Popconfirm
                        title="开通支付"
                        description="确认为此商户开通新生支付吗?"
                        onConfirm={() => updatePaymentStoreApplyXinsh({store_id: record?.store?.id})}
                        onCancel={() => undefined}
                        okText="Yes"
                        cancelText="No"
                    >
                        <a>新生进件</a>
                    </Popconfirm>

                    {/*<CreateForm initialValues={record}>*/}
                    {/*    <a>编辑</a>*/}
                    {/*</CreateForm>*/}
                    {/*<Divider type="vertical"/>*/}
                    {/*<MemberForm id={record.id} initialValues={{level: record.member_level, until: record.member_until }}>*/}
                    {/*    <a>会员</a>*/}
                    {/*</MemberForm>*/}

                </>
            ),
        },
    ];

    return (
        <PageContainer
            header={{title: '商户',}}
        >
            <ProTable<any>
                headerTitle="商户列表"
                actionRef={actionRef}
                rowKey="store_id"
                search={{
                    span: 6
                    // labelWidth: 120,
                }}
                toolBarRender={() => [
                    <CreateForm  onSubmit={() => actionRef.current?.reloadAndRest?.()}>
                        <Button type="primary">
                            <PlusOutlined/>
                            添加商户
                        </Button>
                    </CreateForm>
                ]}
                request={async (params, sorter, filter) => {
                    const {data, success} = await fetchPaymentStores({
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
                columns={columns}
                rowSelection={{
                    onChange: (_, selectedRows) => setSelectedRows(selectedRows),
                }}
            />
            {selectedRowsState?.length > 0 && (
                <FooterToolbar
                    extra={
                        <div>
                            已选择{' '}
                            <a style={{fontWeight: 600}}>{selectedRowsState.length}</a>{' '}
                            项&nbsp;&nbsp;
                        </div>
                    }
                >
                    {/*<Button*/}
                    {/*    onClick={async () => {*/}

                    {/*        await deleteStore(selectedRowsState.map(t => t.id))*/}

                    {/*        setSelectedRows([]);*/}
                    {/*        actionRef.current?.reloadAndRest?.();*/}
                    {/*    }}*/}
                    {/*>*/}
                    {/*    批量删除*/}
                    {/*</Button>*/}
                    {/*<Button onClick={async () => {*/}
                    {/*    await confirmStore(selectedRowsState.map(t => t.id))*/}

                    {/*    setSelectedRows([]);*/}
                    {/*    actionRef.current?.reloadAndRest?.();*/}
                    {/*}} type="primary">批量审批</Button>*/}
                </FooterToolbar>
            )}
        </PageContainer>
    );
};

export default TableList;
