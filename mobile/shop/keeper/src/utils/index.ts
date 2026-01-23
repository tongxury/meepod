import md5 from 'js-md5'

export const createMd5 = (src: string) => {
    return md5(src + '1111JDo39Z112NmDk04Jmlzd7')
}


const dataURLtoBlob = (dataurl) => {
    var arr = dataurl.split(','),
        mime = arr[0].match(/:(.*?);/)[1],
        bstr = atob(arr[1]),
        n = bstr.length,
        u8arr = new Uint8Array(n);
    while (n--) {
        u8arr[n] = bstr.charCodeAt(n);
    }
    return new Blob([u8arr], {type: mime});
}

const blobToFile = (theBlob, fileName) => {
    theBlob.lastModifiedDate = new Date();  // 文件最后的修改日期
    theBlob.name = fileName;                // 文件名
    return new File([theBlob], fileName, {type: theBlob.type, lastModified: Date.now()});
}

export const base64ToFile = (base64Src: string, fileName: string) => {
    return blobToFile(dataURLtoBlob(base64Src), fileName)
}