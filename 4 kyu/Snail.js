function snail (array) {
    let rotate = function(array) {
        let a = [];
        for (let i=0; i<array.length; i++)
            for (let j=0; j<array[i].length; j++) {
                let k = array[i].length - 1 - j;
                a[k] = a[k] || [];
                a[k][i] = array[i][j];
            }
        return a;
    };

    let r = [];
    while(array.length > 0){
        r = r.concat(array.shift());
        array = rotate(array);
    }
    return r;
}