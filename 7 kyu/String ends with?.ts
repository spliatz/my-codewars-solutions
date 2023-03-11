export function solution(str: string, ending: string): boolean {
    return str === str.slice(0, str.length - ending.length).concat(ending)
}