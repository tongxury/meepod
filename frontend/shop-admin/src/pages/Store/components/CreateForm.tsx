import { addStore, fetchBanks, fetchLocs, updateStore } from '@/services';
import {
  ProFormDatePicker,
  ProFormGroup,
  ProFormSelect,
  ProFormText,
  ProFormTextArea,
} from '@ant-design/pro-components';
import {
  DrawerForm,
  ProFormCascader,
  ProFormUploadButton,
} from '@ant-design/pro-form';
import { Upload, message } from 'antd';
import { RcFile } from 'antd/lib/upload';
import React, { useState } from 'react';

const CreateForm: React.FC<any> = ({
  initialValues,
  onSubmit,
  children,
}: {
  initialValues?: any;
  onSubmit: () => void;
  children: JSX.Element;
}) => {
  const uploadProps = {
    maxCount: 1,
    accept: 'image/*',
    beforeUpload: (file: RcFile) => {
      const valid = file.size < 1024 * 1024;
      if (!valid) {
        message.error('文件大小不得超过1M, 请调整后重新上传');
      }
      return valid || Upload.LIST_IGNORE;
    },
  };

  const [banks, setBanks] = useState<any[]>();

  return (
    <DrawerForm
      // labelWidth="auto"
      grid={true}
      rowProps={{
        gutter: [16, 16],
      }}
      trigger={children}
      onFinish={async (values: any) => {
        const params = {
          ...values,
          bank_card_front: values.bank_card_front?.[0]?.response?.data?.[0],
          bank_card_back: values.bank_card_back?.[0]?.response?.data?.[0],
          id_card_front: values.id_card_front?.[0]?.response?.data?.[0],
          id_card_back: values.id_card_back?.[0]?.response?.data?.[0],
          id_card_handled: values.id_card_handled?.[0]?.response?.data?.[0],
          sales_card: values.sales_card?.[0]?.response?.data?.[0],
          store_front: values.store_front?.[0]?.response?.data?.[0],
          store_in_side: values.store_in_side?.[0]?.response?.data?.[0],
          loc: values.area_list?.join('-'),
          // bank_loc: values.bank_area?.join('-'),
        };

        if (params.id) {
          await updateStore([params.id], params);
        } else {
          await addStore(params);
        }

        onSubmit?.();
        return true;
      }}
      initialValues={{
        ...initialValues,
        store_name: initialValues?.name,
        phone: initialValues?.owner?.phone,
        area_list: initialValues?.loc?.split('-'),
        // bank_area: initialValues?.bank_loc?.split('-'),
        bank_card_front: initialValues?.bank_card_front
          ? [
            {
              url: initialValues?.bank_card_front,
              response: {
                data: [initialValues?.bank_card_front],
              },
            },
          ]
          : [],
        bank_card_back: initialValues?.bank_card_back
          ? [
            {
              url: initialValues?.bank_card_back,
              response: {
                data: [initialValues?.bank_card_back],
              },
            },
          ]
          : [],
        id_card_front: initialValues?.id_card_front
          ? [
            {
              url: initialValues?.id_card_front,
              response: {
                data: [initialValues?.id_card_front],
              },
            },
          ]
          : [],
        id_card_back: initialValues?.id_card_back
          ? [
            {
              url: initialValues?.id_card_back,
              response: {
                data: [initialValues?.id_card_back],
              },
            },
          ]
          : [],
        id_card_handled: initialValues?.id_card_handled
          ? [
            {
              url: initialValues?.id_card_handled,
              response: {
                data: [initialValues?.id_card_handled],
              },
            },
          ]
          : [],
        sales_card: initialValues?.sales_card
          ? [
            {
              url: initialValues?.sales_card,
              response: {
                data: [initialValues?.sales_card],
              },
            },
          ]
          : [],
        store_front: initialValues?.store_front
          ? [
            {
              url: initialValues?.store_front,
              response: {
                data: [initialValues?.store_front],
              },
            },
          ]
          : [],
        store_in_side: initialValues?.store_in_side
          ? [
            {
              url: initialValues?.store_in_side,
              response: {
                data: [initialValues?.store_in_side],
              },
            },
          ]
          : [],
      }}
    >
      <ProFormGroup>
        {initialValues?.id && (
          <ProFormText
            width="md"
            disabled={true}
            required
            name="id"
            label="店铺邀请码"
          />
        )}
        <ProFormText
          width="md"
          disabled={initialValues?.owner?.phone}
          required
          name="phone"
          label="店主手机号"
          placeholder="请输入店主手机号(必须与店主端APP登录手机号一致)"
        />
        <ProFormText
          width="md"
          required
          name="store_name"
          label="店铺名称"
          tooltip="最长为24个字"
          placeholder="请输入名称"
        />

        <ProFormText
          width="md"
          required
          name="username"
          label="店主姓名"
          placeholder="请输入店主姓名"
        />

        <ProFormText
          width="md"
          required
          name="email"
          label="店主邮箱"
          placeholder="请输入店主邮箱"
        />
      </ProFormGroup>
      <ProFormGroup>
        <ProFormUploadButton
          width="md"
          name="id_card_front"
          label="店主身份证正面"
          // listType={"text"}
          fieldProps={uploadProps}
          action="/api/v1/images"
        />
        <ProFormUploadButton
          width="md"
          name="id_card_back"
          label="店主身份证反面"
          fieldProps={uploadProps}
          action="/api/v1/images"
        />
        <ProFormUploadButton
          width="md"
          name="id_card_handled"
          label="店主手持身份证"
          fieldProps={uploadProps}
          action="/api/v1/images"
        />
        <ProFormText
          width="md"
          name="id_card_no"
          label="身份证号"
          placeholder="请输入店主身份证号"
        />
        <ProFormDatePicker
          width="md"
          name="id_card_from"
          label="身份证有效期开始"
          // format='YYYY-MM-DD'
          placeholder="请输入店主身份证有效期开始"
        />
        <ProFormDatePicker
          width="md"
          name="id_card_to"
          label="身份证有效期截止"
          placeholder="请输入店主身份证有效期截止"
        />
      </ProFormGroup>
      <ProFormGroup>
        <ProFormUploadButton
          width="md"
          name="sales_card"
          label="代销证（体彩或福彩）"
          fieldProps={uploadProps}
          action="/api/v1/images"
        />
        <ProFormUploadButton
          width="md"
          name="store_front"
          label="店铺门面照片"
          fieldProps={uploadProps}
          action="/api/v1/images"
        />
        <ProFormUploadButton
          width="md"
          name="store_in_side"
          label="店铺内景照片"
          fieldProps={uploadProps}
          action="/api/v1/images"
        />
      </ProFormGroup>

      <ProFormGroup>
        <ProFormText
          width="md"
          name="bank_account_name"
          label="银行卡开户名称"
          placeholder="请输入店主银行卡开户名称"
        />
        <ProFormText
          width="md"
          name="bank_name"
          label="银行名称"
          placeholder="请输入银行名称"
        />
        <ProFormSelect
          width="md"
          name="bank_branch"
          label="银行所属支行名称"
          placeholder="请输入所属支行"
          fieldProps={{
            showSearch: true,
            // onSearch: (value) => {
            //   fetchBanks({branchName: value}).then(rsp => {
            //     setBanks(rsp.data)
            //   })
            // },
            // options: banks?.map(c => ({label: c.branch_name, value: c.bank_id + "-" + c.branch_no}))
          }}
          request={async (params, props) =>
            fetchBanks({
              branchName: params.keyWords,
              branchNo: params.keyWords ? '' : initialValues?.bank_branch,
            }).then((rsp) => {
              console.log('rsprsprsprsprsp', rsp);

              const options = rsp.data?.map((c: any) => ({
                label: c.branch_name,
                value: c.bank_id + '-' + c.branch_no,
              }));
              return new Promise((resolve, reject) => resolve(options));
            })
          }
        // request={async () => [
        //   { label: '全部', value: 'all' },
        //   { label: '未解决', value: 'open' },
        //   { label: '已解决', value: 'closed' },
        //   { label: '解决中', value: 'processing' },
        // ]}
        />

        <ProFormText
          width="md"
          name="bank_account"
          label="银行卡号"
          placeholder="请输入店主银行卡号"
        />
        <ProFormText
          width="md"
          name="bank_phone"
          label="银行卡预留手机号"
          placeholder="请输入银行卡预留手机号"
        />

        <ProFormUploadButton
          // width="md"
          name="bank_card_front"
          label="银行卡正面"
          fieldProps={uploadProps}
          action="/api/v1/images"
        />
        <ProFormUploadButton
          width="md"
          name="bank_card_back"
          label="银行卡反面"
          fieldProps={uploadProps}
          action="/api/v1/images"
        />

        {/*<ProFormCascader*/}
        {/*  colProps={{ span: 24 }}*/}
        {/*  width="md"*/}
        {/*  required*/}
        {/*  request={() =>*/}
        {/*    fetchLocs().then((rsp) => {*/}
        {/*      return new Promise((resolve, reject) => resolve(rsp.data));*/}
        {/*    })*/}
        {/*  }*/}
        {/*  name="bank_area"*/}
        {/*  label="银行卡开户地址"*/}
        {/*  // initialValue={['zhejiang', 'hangzhou', 'xihu']}*/}
        {/*  addonAfter={''}*/}
        {/*/>*/}
      </ProFormGroup>

      <ProFormCascader
        colProps={{ span: 24 }}
        width="md"
        request={() =>
          fetchLocs().then((rsp) => {
            return new Promise((resolve, reject) => resolve(rsp.data));
          })
        }
        name="area_list"
        label="区域"
        // initialValue={['zhejiang', 'hangzhou', 'xihu']}
        addonAfter={''}
      />
      <ProFormTextArea
        colProps={{ span: 24 }}
        name="address"
        label="详细地址"
      />
    </DrawerForm>
  );
};

export default CreateForm;
