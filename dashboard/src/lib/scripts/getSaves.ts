import { gql, GraphQLClient } from "graphql-request"

export async function getSaves():Promise<Array<Record<string, any>>> {
    const endpoint = `http://localhost:8080/query`

    const graphQLClient = new GraphQLClient(endpoint)

    const query = gql`
        query getIDs($userID: Int!) {
            getIDs(UserID: $userID) {
                IDs
            }
        }
    `

    const variables = {
        userID: 420,
    };

    

    let metadata = (async () =>  {
        const data = await graphQLClient.request(query, variables)
        //console.log(data)

        let metadata: Array<Record<string, any>> = []

        for (let value in data) {
            let idList = data[value]
            for (let listKey in idList) {
                let ids = idList[listKey]
                for (let item of ids) {
                    const saveInfo = await getSaveInfo(item, graphQLClient);
                    //console.log(metadata);
                    //for loop the saveInfo and append the value to metadata
                    for (let save in saveInfo) {
                        //console.log(saveInfo[save])
                        metadata.push(saveInfo[save])
                    }
                    
                }
            }
        }

        return metadata

    })();

    console.log("Returning from getsaves with metadata:", metadata)

    return metadata
}

async function getSaveInfo(saveID: number, graphQLClient: GraphQLClient): Promise<any>{
    //console.log("getting metadata for following saveid:" + saveID)
    const query = gql`
        query getLocalSave($ID: Int!) {
            getLocalSave(ID: $ID) {
                ID
                UserID
                Name
                Console
                Device
                TimeMod
                TimeString
                Path
                Thumbnail
                Header
            }
        }
    `

    const variables = {
        ID: saveID,
    };

    const data = (async () =>  {
        const data = await graphQLClient.request(query, variables)
        //console.log(data)

        return data
    })();

    return data
}

//getIDs(UserID: Int!): AllIDs!