function count (str) {
    // The function code should be here

    if (str.length < 1) return {};
    const obj = {}
    const arr = [];

    for(let i = 0; i < str.length; i++) {
        arr.push(str[i])
    }

    arr.forEach((item,index) => {
        obj[item] = 0;
        let repeat = 0;
        for (let i = 0; i < arr.length; i++) {
            if (arr[i] === item) repeat++;
        }
        obj[item] = repeat;
    })

    return obj;
}