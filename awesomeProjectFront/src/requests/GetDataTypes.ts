import { data_types, draft_forecast_app } from './MockData';
import { InterfaceDataTypeProps } from '../components/DataTypeCard';

export type Response = {
    draft_forecast_app: string | null;
    data_types: InterfaceDataTypeProps[];
}

export async function getDataTypes(filter?: string): Promise<Response> {
    let url = '/api/data_types'
    if (filter !== undefined) {
        url += `?data_type_name=${filter}`
    }
    return fetch(url)
        .then(response => {
            if (response.status >= 500 || response.headers.get("Server") == "GitHub.com") {
                return fromMock(filter)
            }
            return response.json() as Promise<Response>
        })
        .catch(_ => {
            return fromMock(filter)
        })
}

function fromMock(filter?: string): Response {
    let filteredDataTypes = Array.from(data_types.values())
    if (filter !== undefined) {
        let type = filter.toLowerCase()
        filteredDataTypes = filteredDataTypes.filter(
            (data_type) => data_type.data_type_name.toLowerCase().includes(type)
        )
    }
    return { draft_forecast_app, data_types: filteredDataTypes }
}