import React from 'react';
import './ProductGrid.css';
import formatPrice from '../utils/formatPrice';

const ProductGrid = ({ products }) => {
  if (!products || products.length === 0) {
    return null;
  }

  return (
    <div className="product-grid-container">
      {products.map((product, index) => (
        <div key={index} className="product-card">
          <a href={product.url} target="_blank" rel="noopener noreferrer" className="product-link">
            <div className="product-image-container">
              <img src={product.image_url} alt={product.title} className="product-image" />
            </div>
            <div className="product-info">
              <h3 className="product-name">{product.title}</h3>
              <p className="product-price">
                {product.site === 'Teknosa' ? formatPrice(product.price) : product.price}
              </p>
              {product.rating > 0 && <p className="product-rating">Rating: {product.rating}</p>}
            </div>
          </a>
        </div>
      ))}
    </div>
  );
};

export default ProductGrid;
