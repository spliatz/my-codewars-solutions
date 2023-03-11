function moveZeros(arr) {
    const sortedArray = [];

    let catchNull = 0;

    for (let i = 0; i < arr.length; i++) {
        if (arr[i] === 0) {
            catchNull++;
            continue;
        }
        sortedArray.push(arr[i]);
    }

    for (let i = 0; i < catchNull; i++) {
        sortedArray.push(0);
    }

    return sortedArray;
}