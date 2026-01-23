import {Button, Drawer, message, Modal} from 'antd';
import React, {PropsWithChildren} from 'react';
import {
    DrawerForm, ModalForm,
    ProForm, ProFormCascader,
    ProFormUploadButton,
} from "@ant-design/pro-form";
import {ProFormRadio, ProFormDatePicker, ProFormSelect, ProFormText, ProFormTextArea} from "@ant-design/pro-components";
import {addPaymentStore} from "@/services";

const CreateForm: React.FC<any> = ({initialValues, onSubmit, children}: {
    initialValues?: any,
    onSubmit?: () => void,
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
            onFinish={async (values: any) => {

                console.log('values', values)

                const params = {
                    ...values,
                }

                await addPaymentStore(params)
                onSubmit?.()

                return true
            }}
            initialValues={{
                ...initialValues,
            }}
        >
            <ProFormText
                width="md"
                required
                name="store_id"
                label="店铺邀请码"
                placeholder="店铺邀请码"
            />

        </ModalForm>
    );
};

export default CreateForm;
