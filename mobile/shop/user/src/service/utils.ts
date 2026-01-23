export const updateList = (mutate, rsp, id,) => {
    if (rsp.data.code === 0) {
        mutate(oldData => {
            const newList = oldData.list.map(u => u.id !== id ? u : rsp.data.data);
            const newData = {...oldData}
            newData.list = newList
            return newData
        })
    }
}


export const updateListItemById = (mutate, newItem) => {
    mutate(oldData => {
        const newList = oldData.list.map(u => u.id !== newItem.id ? u : newItem);
        const newData = {...oldData}
        newData.list = newList
        return newData
    })
}

