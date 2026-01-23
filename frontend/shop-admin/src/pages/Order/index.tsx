import services from '@/services/demo';
import {
    ActionType,
    FooterToolbar,
    PageContainer,
    ProDescriptions,
    ProDescriptionsItemProps,
    ProTable,
} from '@ant-design/pro-components';
import {Avatar, Button, Divider, Drawer, Input, message, Tag} from 'antd';
import React, {useRef, useState} from 'react';
import {confirmStore, deleteStore, fetchOrders, fetchStores, fetchUsers} from "@/services";
import {PlusOutlined} from "@ant-design/icons";
import {Stack} from "@mui/material";

const TableList: React.FC<unknown> = () => {
    const actionRef = useRef<ActionType>();
    const [selectedRowsState, setSelectedRows] = useState<any[]>([]);
    const columns: ProDescriptionsItemProps<any>[] = [
        {
            title: 'ID',
            dataIndex: 'id',
            hideInForm: true,
        },
        {
            title: '店铺',
            dataIndex: 'store',
            render: (dom, entity, index, action, schema) => {
                return <Stack direction="row" spacing={1} alignItems={"center"}>
                    <Avatar src={entity?.store?.icon}/>
                    <div>{entity?.store?.name}</div>
                </Stack>
            },
            hideInForm: true,
            hideInSearch: true,
        },
        {
            title: '彩种',
            dataIndex: 'item',
            render: (dom, entity, index, action, schema) => {
                return <Stack direction="row" spacing={1} alignItems={"center"}>
                    <Avatar src={entity?.plan?.item?.icon}/>
                    <div>{entity?.plan?.item?.name}</div>
                </Stack>
            },
            hideInForm: true,
            hideInSearch: true,
        },
        {
            title: '用户',
            dataIndex: 'user',
            // render: (dom, entity, index, action, schema) => {
            //     return <Stack spacing={1} alignItems={"center"}>
            //         <div>{entity?.user?.nickname}</div>
            //         <div>{entity?.user?.phone}</div>
            //     </Stack>
            // },
            render: (dom, entity, index, action, schema) => {
                return <Stack direction="row" spacing={1} alignItems={"center"}>
                    <Avatar src={entity?.user?.icon}/>
                    <div>{entity?.user?.nickname}</div>
                </Stack>
            },
            hideInForm: true,
            hideInSearch: true,
        },
        {
            title: '金额',
            dataIndex: 'amount',
            valueType: 'text',
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
            title: '状态',
            dataIndex: 'status',
            render: (dom, entity, index, action, schema) => {
                return <Tag color={entity.status?.color}>{entity.status?.name}</Tag>
            },
            hideInForm: true,
            hideInSearch: true,
        },
        // {
        //     title: '操作',
        //     dataIndex: 'option',
        //     valueType: 'option',
        //     render: (_, record) => (
        //         <>
        //
        //         </>
        //     ),
        // },
    ];

    return (
        <PageContainer
            header={{title: '店铺',}}
        >
            <ProTable<any>
                headerTitle="店铺列表"
                actionRef={actionRef}
                rowKey="id"
                search={{
                    labelWidth: 120,
                }}
                toolBarRender={() => [
                    // <CreateForm>
                    //     <Button type="primary">
                    //         <PlusOutlined/>
                    //         创建店铺
                    //     </Button>
                    // </CreateForm>
                ]}
                request={async (params, sorter, filter) => {
                    const {data, success} = await fetchOrders({
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
