function rgb(r, g, b){
    // complete this function
    if(r < 0) {
        r = 0
    } else if (r > 255) {
        r = 255
    }
    if(g < 0) {
        g = 0
    } else if (g > 255) {
        g = 255
    }
    if(b < 0) {
        b = 0
    } else if (b > 255) {
        b = 255
    }


    let hexR = r.toString(16);
    if (hexR.length % 2 || hexR.length < 2) {
        hexR = '0' + hexR;
    }
    let hexG = g.toString(16);
    if (hexG.length % 2 || hexG.length < 2) {
        hexG = '0' + hexG;
    }
    let hexB = b.toString(16);
    if (b.length % 2 || hexB.length < 2) {
        hexB = '0' + hexB;
    }
    if(r === 0 && b === 0 && g === 0) {
        return '000000'
    }
    return(hexR.toUpperCase()+hexG.toUpperCase()+hexB.toUpperCase());
}
