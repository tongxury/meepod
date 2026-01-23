import {Button, Drawer, message, Modal} from 'antd';
import React, {PropsWithChildren} from 'react';
import {
    DrawerForm, ModalForm,
    ProForm, ProFormCascader, ProFormMoney,
    ProFormUploadButton,
} from "@ant-design/pro-form";
import {ProFormRadio, ProFormDatePicker, ProFormSelect, ProFormText, ProFormTextArea} from "@ant-design/pro-components";
import {addStore, addStorePayment, fetchLocs, updateStore, updateStoreMember} from "@/services";
import moment from 'moment'
import {FormValueType} from "@/pages/Table/components/UpdateForm";

const TopUpForm: React.FC<any> = ({initialValues, onSubmit, children}: {
    initialValues?: any,
    onSubmit: (values: FormValueType) => Promise<void>;
    children: JSX.Element
}) => {

    return (
        <ModalForm
            // labelWidth="auto"
            grid={true}
            rowProps={{gutter: [16, 16],}}
            trigger={
                children
            }
            onFinish={onSubmit}
            initialValues={{
                ...initialValues,
            }}
        >
            <ProFormMoney
                width="md"
                name="amount"
                label="充值金额"
                required
                fieldProps={{
                    // value: '',
                    // onChange: (e) => setType(e.target.value),
                }}

            />

        </ModalForm>
    );
};

export default TopUpForm;
