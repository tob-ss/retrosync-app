export function getHeaders(arg1:Array<Record<string, any>>) {
    console.log("this metadata is supposed to have things:", arg1)

    for (let map in arg1) {
        console.log("this is a save file:", arg1[map]["Header"])
    }

    let sortedArray = sortArray(arg1)

    console.log("This is the sorted array:", sortedArray)

    let headers: string[] = []

    for (let saveFile in sortedArray) {
        let header = sortedArray[saveFile]["Header"]
        console.log("we're going to add this header:", header)

        headers.push(header)
    }

    let uniqueHeaders = uniq(headers)

    console.log("this should be a unique list of headers which is also sorted:", uniqueHeaders)

    return uniqueHeaders
}

function sortArray(arrayMap:Array<Record<string, any>>): Array<Record<string, any>> {

    //let sortedArray: Array<Record<string, any>> = [] 

    arrayMap.sort((a, b) => a.TimeMod - b.TimeMod);

    arrayMap.reverse();

    return arrayMap
}

function uniq(a) {
    return Array.from(new Set(a));
}