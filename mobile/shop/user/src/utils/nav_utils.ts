export const toDetail = (navigation, biz_category, biz_id) => {

    switch (biz_category) {
        case 'order' :
            // @ts-ignore
            navigation.navigate('OrderDetail', {id: biz_id})
            break
        case 'groupShare':
            // @ts-ignore
            navigation.navigate('OrderGroupDetail', {id: biz_id, biz_category})
            break
    }
}
