import { FC, useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import { BigDataTypeCard, InterfaceDataTypeProps } from '../components/DataTypeCard';
import Navbar from 'react-bootstrap/Navbar';
import { Link } from 'react-router-dom';
import Nav from 'react-bootstrap/Nav';
import LoadAnimation from '../components/LoadAnimation';
import { getDataType } from '../requests/GetDataType'

const DataTypeInfo: FC = () => {
    let { data_type_id } = useParams()
    const [dataType, setDataType] = useState<InterfaceDataTypeProps>()
    const [loaded, setLoaded] = useState<boolean>(false)

    useEffect(() => {
        getDataType(data_type_id)
            .then(data => {
                setDataType(data)
                setLoaded(true)
            })
            .catch((error) => {
                console.error("Error fetching data:", error);
            });
    }, []);

    return (
        <>
            <Navbar>
                <Nav>
                    <Link to="/data_types" className="nav-link p-0 text-dark" data-bs-theme="dark">
                        Виды данных
                    </Link>
                    <Nav.Item className='mx-1'>{">"}</Nav.Item>
                    <Nav.Item className="nav-link p-0 text-dark">
                        {`${dataType ? dataType.data_type_name : '/???/'}`}
                    </Nav.Item>
                </Nav>
            </Navbar>
            {loaded ? (
                dataType ? (
                    <BigDataTypeCard {...dataType} />
                ) : (
                    <h3 className='text-center'>Такого вида данных не существует</h3>
                )
            ) : (
                <LoadAnimation />
            )
            }
        </ >
    )
}

export { DataTypeInfo }