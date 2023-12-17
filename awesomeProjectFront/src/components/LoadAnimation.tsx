import Spinner from 'react-bootstrap/Spinner';
import { FC } from 'react'

export const LoadAnimation: FC = () => (
    <div className='position-absolute top-50 start-50 translate-middled-flex justify-content-center'>
        <Spinner animation="border" role="status">
            <span className="visually-hidden">Загрузка...</span>
        </Spinner>
    </div>
)

export default LoadAnimation;