import React from 'react';
import './ProductGrid.css';
import { FaStar, FaStarHalfAlt, FaRegStar } from 'react-icons/fa';
import formatPrice from '../utils/formatPrice';

const renderStars = (rating) => {
    const fullStars = Math.floor(rating);
    const halfStar = rating % 1 >= 0.5;
    const emptyStars = 5 - fullStars - (halfStar ? 1 : 0);
    return (
        <div className="stars">
            {[...Array(fullStars)].map((_, i) => <FaStar key={`full-${i}`} className="star-icon" />)}
            {halfStar && <FaStarHalfAlt className="star-icon" />}
            {[...Array(emptyStars)].map((_, i) => <FaRegStar key={`empty-${i}`} className="star-icon" />)}
        </div>
    );
};

const ProductGrid = ({ products }) => {
    if (!products || products.length === 0) {
        return null; // Eğer ürün yoksa hiçbir şey render etme
    }

    return (
        <div className="product-grid">
            {products.map((product, index) => (
                <div className="product-card" key={product.id || `${product.url}-${index}`}>
                    <a href={product.url} target="_blank" rel="noopener noreferrer" className="product-image-container">
                        <img src={product.image_url || '/placeholder.png'} alt={product.title} className="product-image" />
                    </a>
                    <div className="product-info">
                        <h3 className="product-title">{product.title}</h3>
                        <p className="product-price">{formatPrice(product.price)}</p>
                        {product.rating && (
                            <div className="product-rating">
                                {renderStars(parseFloat(product.rating))}
                                <span>({product.rating})</span>
                            </div>
                        )}
                        <a href={product.url} target="_blank" rel="noopener noreferrer" className="product-link">
                            View Product
                        </a>
                    </div>
                </div>
            ))}
        </div>
    );
};

export default ProductGrid;