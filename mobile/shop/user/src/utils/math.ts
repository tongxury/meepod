export const factorial = (num: number) => {
    if (num > 0) {
        return (num * factorial(num - 1));
    } else
        return (1);
}


export const cnm = (all: number, target: number) => {

    if (all < target) {
        return 0
    }

    return factorial(all) / (factorial(target) * factorial(all - target))
}

// export const arrange = (arr, len) => {
//     // let len = arr.length
//     let res = [] // 所有排列结果
//     /**
//      * 【全排列算法】
//      * 说明：arrange用来对arr中的元素进行排列组合，将排列好的各个结果存在新数组中
//      * @param tempArr：排列好的元素
//      * @param leftArr：待排列元素
//      */
//     let _arrange = (tempArr, leftArr) => {
//         if (tempArr.length === len) { // 这里就是递归结束的地方
//             // res.push(tempArr) // 得到全排列的每个元素都是数组
//             res.push(tempArr.join('')) // 得到全排列的每个元素都是字符串
//         } else {
//             leftArr.forEach((item, index) => {
//                 let temp = [].concat(leftArr)
//                 temp.splice(index, 1)
//                 // 此时，第一个参数是当前分离出的元素所在数组；第二个参数temp是传入的leftArr去掉第一个后的结果
//                 _arrange(tempArr.concat(item), temp) // 这里使用了递归
//             })
//         }
//     }
//     _arrange([], arr)
//     return res
// }



export const arrange = (arr, len) => {

    const xrt = [];

    if (len <= 0) {
        return xrt
    }

    function combine_increase(iArr, start, rt, count, NUM, arr_len) { // 从有arr_len个元素的数组中抽取NUM个元素的组合所有可能
        let i = 0;
        for (let i = start; i < arr_len + 1 - count; i++) {
            rt[count - 1] = i;
            if (count - 1 == 0) {
                let tmp = [];
                for (let j = NUM - 1; j >= 0; j--) {
                    tmp.push(arr[rt[j]]);
                }
                xrt.push(tmp);
            } else {
                combine_increase(iArr, i + 1, rt, count - 1, NUM, arr_len);
            }
        }
    }

    var rt=new Array(len)
    combine_increase(arr, 0, rt, len, len, arr.length);

    return xrt
}
// let arr=[1,2,3,4,5,6,7,8,9,10];
// var rt=new Array(4)
// combine_increase(arr, 0, rt, 6, 6, 10);
// console.log(xrt)