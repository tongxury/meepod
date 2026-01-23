import Guide from '@/components/Guide';
import {trim} from '@/utils/format';
import {PageContainer} from '@ant-design/pro-components';
import {useModel} from '@umijs/max';
import {StatisticCard} from '@ant-design/pro-components';
import {useRequest} from "umi";
import {fetchStats} from "@/services";

const HomePage = () => {

    const {data, loading} = useRequest(fetchStats)

    return (
        <PageContainer ghost>
            <StatisticCard.Group direction={'row'} loading={loading}>
                <StatisticCard
                    statistic={{
                        title: '店铺总量',
                        value: data?.store_count,
                    }}
                />
                <StatisticCard
                    statistic={{
                        title: '用户总量',
                        value: data?.user_count,
                    }}
                />
                <StatisticCard
                    statistic={{
                        title: '订单总量',
                        value: data?.order_count,
                    }}
                />
            </StatisticCard.Group>
        </PageContainer>
    );
};

export default HomePage;
