
import React, { useState } from 'react';
import './ChatSidemenu.css';
import logo from '../assets/images/Adsız.png';

const ChatSidemenu = ({ onSearch, loading, onGeminiQuery, products }) => {
    const [query, setQuery] = useState('');
    const [geminiQuery, setGeminiQuery] = useState('');

    const handleSearch = (e) => {
        e.preventDefault();
        if (query.trim() && !loading) {
            onSearch(query);
        }
    };

    const handleGemini = (e) => {
        e.preventDefault();
        if (geminiQuery.trim() && products.length > 0 && !loading) {
            onGeminiQuery(geminiQuery, products);
            setGeminiQuery('');
        }
    };

    return (
        <div className="chat-sidemenu">
            <div className="logo-container">
                <img src={logo} alt="SmartyShop Logo" className="logo-img" />
                <h1 className="logo-text">SmartyShop</h1>
            </div>
            <p className="subtitle">Your AI-powered shopping assistant</p>

            <div className="search-section">
                <form onSubmit={handleSearch} className="input-group">
                    <input
                        type="text"
                        value={query}
                        onChange={(e) => setQuery(e.target.value)}
                        placeholder="e.g., wireless headphones..."
                        disabled={loading}
                    />
                    <button type="submit" disabled={loading}>
                        {loading ? 'Searching...' : 'Search Products'}
                    </button>
                </form>
            </div>

            {products.length > 0 && (
                <div className="gemini-section">
                    <h3>AI Shopping Assistant</h3>
                    <p>Ask for a comparison, summary, or anything about the found products.</p>
                    <form onSubmit={handleGemini} className="input-group">
                        <textarea
                            value={geminiQuery}
                            onChange={(e) => setGeminiQuery(e.target.value)}
                            placeholder="e.g., 'Which of these has the best battery life?'"
                            disabled={loading}
                        />
                        <button type="submit" disabled={loading || !geminiQuery.trim()}>
                            {loading ? 'Thinking...' : 'Ask Gemini'}
                        </button>
                    </form>
                </div>
            )}
        </div>
    );
};

export default ChatSidemenu;
