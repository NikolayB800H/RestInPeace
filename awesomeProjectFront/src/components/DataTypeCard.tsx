import { FC } from 'react'
import { Link } from 'react-router-dom';
import Card from 'react-bootstrap/Card';
import ListGroup from 'react-bootstrap/ListGroup';
import CardImage from './CardImage';

export interface InterfaceDataTypeProps {
    data_type_id: string
    image_path: string
    data_type_name: string
    precision: number
    description: string
    unit: string
    data_type_status: string
}

// const setPlaceholder = (event: any) => {
//     event.target.src = '/placeholder3.jpg';
// };

export const SmallDataTypeCard: FC<InterfaceDataTypeProps> = ({ data_type_id, image_path, data_type_name, precision, unit }) => (

    <Card className='w-100 mx-auto px-0 shadow text-center'>
        <div className="ratio ratio-16x9 overflow-hidden">
            <CardImage url={`http://${image_path}`} className='rounded object-fit-cover'/>
            {/* <Card.Img src={`http://${image_url}`} alt='картинка контейнера' onError={setPlaceholder} className='rounded object-fit-cover' /> */}
        </div>
        <Card.Body className='flex-grow-1'>
            <Card.Title>Прогноз {data_type_name} (в {unit})</Card.Title>
            <Card.Text>Погрешность ±{precision} {unit}</Card.Text>
        </Card.Body>
        <Link to={`/data_types/${data_type_id}`} className="btn btn-dark">Подробнее</Link>
    </Card>
)

export const BigDataTypeCard: FC<InterfaceDataTypeProps> = ({ image_path, data_type_name, precision, unit, description }) => (
    <Card className='shadow text-center text-md-start'>
        <div className='row'>
            <div className='col-12 col-md-8 overflow-hidden'>
                {/* <Card.Img src={`http://${image_url}`} onError={setPlaceholder}/> */}
                <CardImage url={`http://${image_path}`}/>
            </div>
            <Card.Body className='col-12 col-md-4 ps-md-0'>
                <ListGroup variant="flush">
                    <ListGroup.Item>
                        <Card.Title>Прогноз {data_type_name} (в {unit})</Card.Title>
                        <Card.Text>Погрешность ±{precision} {unit}</Card.Text>
                    </ListGroup.Item>
                    <ListGroup.Item>
                        <Card.Text>Описание: {description}</Card.Text>
                    </ListGroup.Item>
                </ListGroup>
            </Card.Body>
        </div>
    </Card>
);