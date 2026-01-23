import md5 from 'js-md5'
import {btoa, atob, toByteArray} from 'react-native-quick-base64';

export const createMd5 = (src: string) => {
    return md5(src)
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


export const b642File = (b64Src: string, fileName: string) => {

    const parts = b64Src.split(',')
    const mime = parts[0].match(/:(.*?);/)[1]
    const u8arr = toByteArray(parts[1])

    const blob = new Blob([u8arr], {type: mime})

    // @ts-ignore
    blob.lastModifiedDate = new Date();  // 文件最后的修改日期
    // @ts-ignore
    blob.name = fileName;

    return new File([blob], fileName, {type: blob.type, lastModified: Date.now()});

}


export const sumReduce = (arr: number[]): number => {
    if (arr.length == 0) {
        return 0
    }

    return arr.reduce((a1, a2) => a1 + a2)
}
export const multiReduce = (arr: number[]): number => {
    if (arr.length == 0) {
        return 0
    }

    return arr.reduce((a1, a2) => a1 * a2)
}