import { data_types } from './MockData';
import { InterfaceDataTypeProps } from '../components/DataTypeCard';

const api = '/api/data_types/'

export async function getDataType(dataTypeId?: string): Promise<InterfaceDataTypeProps | undefined> {
    if (dataTypeId === undefined) {
        return undefined
    }
    let url = api + dataTypeId
    return fetch(url)
        .then(response => {
            if (response.status >= 500 || response.headers.get("Server") == "GitHub.com") {
                return data_types.get(dataTypeId)
            }
            return response.json() as Promise<InterfaceDataTypeProps | undefined>
        })
        .catch(_ => {
            return data_types.get(dataTypeId)
        })
}