// 运行时配置

// 全局初始化数据配置，用于 Layout 用户信息和权限初始化
// 更多信息见文档：https://umijs.org/docs/api/runtime-config#getinitialstate
import { RequestConfig } from "@@/plugin-request/request";
import { message } from "antd";

export async function getInitialState(): Promise<{ name: string }> {
    return { name: '@umijs/max' };
}

export const layout = () => {
    return {
        logo: 'https://img.alicdn.com/tfs/TB1YHEpwUT1gK0jSZFhXXaAtVXa-28-27.svg',
        menu: {
            locale: false,
        },
        title: '福彩管理平台',
        layout: 'mix',
        siderWidth: 200,
    };
};

export const request: RequestConfig = {
    // timeout: 1000,
    // other axios options you want
    // headers: {'X-ACCOUNT': localStorage.getItem('X-ACCOUNT') || ''},
    errorConfig: {
        errorHandler() {
        },
        errorThrower() {
        }
    },
    requestInterceptors: [
        (url, options) => {
            // const {account} = useEthers()
            // console.log('requestInterceptors', 'xxxxxxx')
            // options.headers['X-ACCOUNT'] = account
            console.log(url, options)
            return { url, options }
        },
    ],
    responseInterceptors: [
        // 直接写一个 function，作为拦截器
        [(response) => {
            // 不再需要异步处理读取返回体内容，可直接在data中读出，部分字段可在 config 中找到
            const { data = {} as any, config } = response;
            if (data.code == 10401) {
                // location.href = '/login';
            } else if (data.message) {
                message.open({
                    type: "error",
                    content: data.message,
                }).then()
            }

            // do something
            return response
        },]

        // 一个二元组，第一个元素是 request 拦截器，第二个元素是错误处理
        // [(response) => {return response}, (error) => {return Promise.reject(error)}],
        // // 数组，省略错误处理
        // [(response) => {return response}]
    ]
};


// export function render(oldRender: any) {
//
//   if (!location.href.includes('login')) {
//     getAccount().then(result => {
//       if (result.data) {
//         localStorage.setItem('X-USER-EMAIL', result.data.email)
//         oldRender()
//       } else {
//         localStorage.removeItem('X-USER-EMAIL')
//         location.href = '/login';
//         oldRender()
//       }
//     })
//
//   } else {
//     oldRender()
//   }
// }