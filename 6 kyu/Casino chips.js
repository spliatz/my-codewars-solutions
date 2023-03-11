function solve(arr){
    const [a,b,c] = arr.sort((a,b)=> b - a);
    const [x,y] = [a, b+c];
    const z = y > x ? Math.floor((x-y)/2): 0;
    return y + z;
}