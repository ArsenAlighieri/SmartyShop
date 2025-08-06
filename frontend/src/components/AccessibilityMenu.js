import React, { useState } from 'react';
import './AccessibilityMenu.css';

const AccessibilityMenu = ({ screenReaderMode, setScreenReaderMode }) => {
  const [isOpen, setIsOpen] = useState(false);

  const handleModeChange = (e) => {
    setScreenReaderMode(e.target.value);
    setIsOpen(false); // Close menu on selection
  };

  return (
    <div className="accessibility-menu-container">
      <button 
        className="accessibility-toggle-button" 
        onClick={() => setIsOpen(!isOpen)}
        aria-label="Accessibility Settings"
      >
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"><path d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2zm0 18a8 8 0 1 1 8-8 8 8 0 0 1-8 8zm0-6a2 2 0 1 0-2-2 2 2 0 0 0 2 2zm-2-4a2 2 0 1 1 2 2 2 2 0 0 1-2-2z"/></svg>
      </button>
      {isOpen && (
        <div className="accessibility-dropdown">
          <fieldset>
            <legend>Screen Reader Options</legend>
            <div className="radio-option">
              <input
                type="radio"
                id="readOnHover"
                name="screenReader"
                value="hover"
                checked={screenReaderMode === 'hover'}
                onChange={handleModeChange}
              />
              <label htmlFor="readOnHover">Read on hover</label>
            </div>
            <div className="radio-option">
              <input
                type="radio"
                id="readAll"
                name="screenReader"
                value="all"
                checked={screenReaderMode === 'all'}
                onChange={handleModeChange}
              />
              <label htmlFor="readAll">Read the whole page</label>
            </div>
            <div className="radio-option">
              <input
                type="radio"
                id="readerOff"
                name="screenReader"
                value="off"
                checked={screenReaderMode === 'off'}
                onChange={handleModeChange}
              />
              <label htmlFor="readerOff">Turn off screen reader</label>
            </div>
          </fieldset>
        </div>
      )}
    </div>
  );
};

export default AccessibilityMenu;
