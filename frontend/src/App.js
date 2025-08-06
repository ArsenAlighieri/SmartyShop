import React, { useState, useEffect } from 'react';
import AccessibilityMenu from './components/AccessibilityMenu';
import './App.css';
import ChatSidemenu from './components/ChatSidemenu';
import ProductGrid from './components/ProductGrid';
import LoadingSpinner from './components/LoadingSpinner';
import GeminiResponseDisplay from './components/GeminiResponseDisplay';

// Fisher-Yates (aka Knuth) Shuffle
const shuffleArray = (array) => {
  let currentIndex = array.length, randomIndex;
  while (currentIndex !== 0) {
    randomIndex = Math.floor(Math.random() * currentIndex);
    currentIndex--;
    [array[currentIndex], array[randomIndex]] = [array[randomIndex], array[currentIndex]];
  }
  return array;
};

function App() {
  const [products, setProducts] = useState([]);
  const [originalProducts, setOriginalProducts] = useState([]);
  const [geminiResponse, setGeminiResponse] = useState('');
  const [loading, setLoading] = useState(false);
  const [isChatActive, setIsChatActive] = useState(true);
  const [error, setError] = useState('');
  const [screenReaderMode, setScreenReaderMode] = useState('off'); // 'off', 'hover', 'all'

  // Speech synthesis setup
  useEffect(() => {
    const synth = window.speechSynthesis;
    if (!synth) {
      console.warn('Speech synthesis not supported by this browser.');
      return;
    }

    const speak = (text) => {
      if (synth.speaking) {
        synth.cancel();
      }
      const utterance = new SpeechSynthesisUtterance(text);
      utterance.lang = 'en-US';
      synth.speak(utterance);
    };

    const handleMouseOver = (e) => {
      if (screenReaderMode !== 'hover') return;
      const target = e.target;
      const text = target.innerText || target.alt || target.ariaLabel;
      if (text) {
        speak(text);
      }
    };

    if (screenReaderMode === 'all') {
      speak(document.body.innerText);
    } else {
      synth.cancel(); // Stop speaking if mode changes from 'all'
    }

    if (screenReaderMode === 'hover') {
      document.addEventListener('mouseover', handleMouseOver);
    } else {
      document.removeEventListener('mouseover', handleMouseOver);
    }

    return () => {
      synth.cancel();
      document.removeEventListener('mouseover', handleMouseOver);
    };
  }, [screenReaderMode]);

  const handleSearch = async (query) => {
    setLoading(true);
    setProducts([]);
    setGeminiResponse('');
    setError('');
    setIsChatActive(false);

    const sites = ['trendyol', 'teknosa', 'mediamarkt', 'amazon'];
    try {
      const fetchPromises = sites.map(site =>
        fetch(`http://localhost:8080/products?site=${site}&query=${encodeURIComponent(query)}`)
          .then(response => {
            if (!response.ok) {
              // Don't throw, just log and return empty for this source
              console.error(`Failed to fetch from ${site}: ${response.status}`);
              return [];
            }
            return response.json();
          })
          .catch(err => {
            console.error(`Error connecting to ${site}:`, err);
            return []; // Return empty array on network error
          })
      );

      const results = await Promise.all(fetchPromises);
      const allProducts = results.flat().filter(p => p && p.title && p.price); // Flatten and basic validation
      
      if (allProducts.length === 0) {
        setError('No products found for your query. Please try another search.');
      }

      setProducts(shuffleArray(allProducts));
      setOriginalProducts(allProducts); // Store original products with image_urls
      console.log("Products sent to Gemini:", allProducts); // Log products before sending to Gemini

    } catch (error) {
      console.error('Error during product search:', error);
      setError('An unexpected error occurred. Please try again.');
    }
    setLoading(false);
  };

  const handleGeminiQuery = async (query, productContext) => {
    setLoading(true);
    setGeminiResponse('');
    setError('');
    try {
      const response = await fetch('http://localhost:8080/gemini/query', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ query: query, products: productContext }),
      });
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      const data = await response.json();
      setGeminiResponse(data.answer); // Correctly access the 'answer' field
      if (data.products && data.products.length > 0) {
        const productsWithImages = data.products.map(geminiProduct => {
          const originalProduct = originalProducts.find(op => op.title === geminiProduct.title && op.url === geminiProduct.url);
          return {
            ...geminiProduct,
            image_url: originalProduct ? originalProduct.image_url : geminiProduct.image_url
          };
        });
        setProducts(productsWithImages); // Update products with Gemini's filtered list, now with image_urls
        console.log("Products from Gemini (with images):", productsWithImages); // Log Gemini's product list
      } else {
        setProducts([]); // Clear products if Gemini returns none
      }
    } catch (error) {
      console.error('Error querying Gemini:', error);
      setError('Failed to get a response from the AI assistant.');
    }
    setLoading(false);
  };

  return (
    <div className={`App ${isChatActive ? 'chat-centric' : 'grid-view'}`}>
      <AccessibilityMenu 
        screenReaderMode={screenReaderMode} 
        setScreenReaderMode={setScreenReaderMode} 
      />
      {loading && <LoadingSpinner />}
      <div className="chat-container">
        <ChatSidemenu 
          onSearch={handleSearch} 
          loading={loading} 
          onGeminiQuery={handleGeminiQuery} 
          products={products}
        />
      </div>
      <div className="product-grid-area">
        {error && <p className="error-message">{error}</p>}
        <GeminiResponseDisplay response={geminiResponse} />
        <ProductGrid products={products} />
      </div>
    </div>
  );
}

export default App;
