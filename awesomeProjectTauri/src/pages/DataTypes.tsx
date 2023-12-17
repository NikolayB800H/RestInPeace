import { useEffect, useState, FC } from 'react';
import { SmallDataTypeCard, InterfaceDataTypeProps } from '../components/DataTypeCard';
import LoadAnimation from '../components/LoadAnimation';
import Navbar from 'react-bootstrap/Navbar';
import Form from 'react-bootstrap/Form';
import Button from 'react-bootstrap/Button';
import { getDataTypes } from '../requests/GetDataTypes'

interface InterfaceSearchProps {
    setDataTypes: React.Dispatch<React.SetStateAction<InterfaceDataTypeProps[]>>
}

const Search: FC<InterfaceSearchProps> = ({ setDataTypes }) => {
    const [searchText, setSearchText] = useState<string>('');

    const handleSearch = (event: React.FormEvent<any>) => {
        event.preventDefault();
        getDataTypes(searchText)
            .then(data => {
                console.log(data)
                setDataTypes(data.data_types)
            })
    }
    return (
        <Navbar>
            <Form className="d-flex flex-row flex-grow-1 gap-2" onSubmit={handleSearch}>
                <Form.Control
                    type="text"
                    placeholder="Поиск"
                    className="form-control-sm flex-grow-1 shadow shadow-sm"
                    data-bs-theme="dark"
                    value={searchText}
                    onChange={(e) => setSearchText(e.target.value)}
                />
                <Button
                    variant="dark"
                    size="sm"
                    type="submit"
                    className="shadow">
                    🔍
                </Button>
            </Form>
        </Navbar>)
}

const DataTypes = () => {
    const [loaded, setLoaded] = useState<boolean>(false)
    const [dataTypes, setDataTypes] = useState<InterfaceDataTypeProps[]>([]);
    const [_, setDraftForecastApp] = useState<string | null>(null);

    useEffect(() => {
        getDataTypes()
            .then(data => {
                console.log(data)
                setDraftForecastApp(data.draft_forecast_app)
                setDataTypes(data.data_types)
                setLoaded(true)
            })
            .catch((error) => {
                console.error("Error fetching data:", error);
            });
    }, []);

    return (
        <>
            <Search setDataTypes={setDataTypes} />
            <div className='row row-cols-1 row-cols-sm-2 row-cols-md-3 row-cols-xl-4 px-1'>
                {loaded ? (
                    dataTypes.map((dataType) => (
                        <div className='d-flex p-2 justify-content-center' key={dataType.data_type_id}>
                            <SmallDataTypeCard  {...dataType} />
                        </div>
                    ))
                ) : (
                    <LoadAnimation />
                )}
            </div>
        </>
    )
}

export { DataTypes }