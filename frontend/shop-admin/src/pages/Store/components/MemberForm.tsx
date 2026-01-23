import {Button, Drawer, message, Modal} from 'antd';
import React, {PropsWithChildren} from 'react';
import {
    DrawerForm, ModalForm,
    ProForm, ProFormCascader,
    ProFormUploadButton,
} from "@ant-design/pro-form";
import {ProFormRadio, ProFormDatePicker, ProFormSelect, ProFormText, ProFormTextArea} from "@ant-design/pro-components";
import {addStore, fetchLocs, updateStore, updateStoreMember} from "@/services";
import moment from 'moment'
import {FormValueType} from "@/pages/Table/components/UpdateForm";

const MemberForm: React.FC<any> = ({initialValues, onSubmit, children}: {
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
                until: initialValues.until ? moment(initialValues.until, 'YYYY-MM-DD').unix()*1000 : undefined

            }}
        >
            <ProFormRadio.Group
                // width="md"
                name="level"
                label="会员等级"
                required
                radioType="button"
                fieldProps={{
                    buttonStyle: "solid"
                    // value: '',
                    // onChange: (e) => setType(e.target.value),
                }}
                options={[
                    {value: 'normal', label: '普通会员'},
                    {value: 'advanced', label: '高级会员'},
                ]}
            />

            <ProFormDatePicker
                name="until"
                label="会员截止日期"
                required
                transform={(value) => {
                    return {
                        until: moment(value).unix(),
                    };
                }}
            />
        </ModalForm>
    );
};

export default MemberForm;
