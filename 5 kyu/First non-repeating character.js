function firstNonRepeatingLetter(s) {

    const obj = {};
    const arr = s.split('');

    for (let i = 0; i < arr.length; i++) {
        if (obj[arr[i].toLowerCase()]) {
            obj[arr[i].toLowerCase()] = obj[arr[i].toLowerCase()] + 1;
            continue;
        }
        obj[arr[i].toLowerCase()] = 1;
    }

    console.log(obj);
    if (arr.filter(x => obj[x.toLowerCase()] === 1).length > 0) return arr.filter(x => obj[x.toLowerCase()] === 1)[0];
    return '';
}