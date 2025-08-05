import React from 'react';
import './GeminiResponseDisplay.css';

const GeminiResponseDisplay = ({ response }) => {
  if (!response) {
    return null;
  }

  return (
    <div className="gemini-response-display">
      <h3 className="gemini-response-title">Gemini's Insights</h3>
      <div className="gemini-response-content">
        <p>{response}</p>
      </div>
    </div>
  );
};

export default GeminiResponseDisplay;
