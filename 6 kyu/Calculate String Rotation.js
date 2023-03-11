function shiftedDiff(first,second){
    // ...
    const arr = [...second];


    for (let i = 0; i < arr.length; i++) {

        if (first === arr.join('')) return i;

        arr.push(arr.shift());
    }

    return -1;
}