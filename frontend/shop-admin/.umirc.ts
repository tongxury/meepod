import {defineConfig} from '@umijs/max';

export default defineConfig({
    antd: {},
    access: {},
    model: {},
    initialState: {},
    request: {},
    layout: {
        title: '@umijs/max',
    },
    routes: [
        {path: '/login', layout: false, component: './Login'},
        {path: '/', redirect: '/home',},
        {name: '首页', path: '/home', component: './Home',},
        {name: '店铺', path: '/store', component: './Store',},
        {name: '用户', path: '/user', component: './User',},
        {name: '订单', path: '/order', component: './Order',},
        {
            name: '支付', path: '/payment', routes: [
                {name: '商户', path: '/payment/store', component: './Payment/Store',},
            ]
        },
    ],
    npmClient: 'yarn',
    outputPath: '../../opt/projects/shop/dist/admin',
    proxy: {
        '/api/': {
            target: 'http://localhost:6677'
        }
    }
});

