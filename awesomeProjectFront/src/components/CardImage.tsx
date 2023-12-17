import Card from 'react-bootstrap/Card';
import { useState, useEffect } from 'react';

interface CardImageProps {
    url: string;
    className?: string;
}

const CardImage = ({ url, className, ...props }: CardImageProps) => {
    const [src, setSrc] = useState(`${import.meta.env.BASE_URL}Default.jpg`);


    useEffect(() => {
        fetch(url)
            .then(response => {
                if (response.status >= 500 || response.headers.get("Server") == "GitHub.com") {
                    throw new Error(`Can't load image from ${url}`);
                }
                return response.blob();
            })
            .then(blob => setSrc(URL.createObjectURL(blob)))
            .catch(error => {
                console.error(error.message);
            });
    }, []);

    const handleError = () => {
        console.error(`Error loading image: ${url}`);
    };

    return <Card.Img src={src} className={className + ', img-fluid'} onError={handleError} {...props} />;
};

export default CardImage;