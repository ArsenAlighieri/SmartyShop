import React from 'react';
import ReactMarkdown from 'react-markdown';
import { Prism as SyntaxHighlighter } from 'react-syntax-highlighter';
import { tomorrow } from 'react-syntax-highlighter/dist/esm/styles/prism';
import remarkGfm from 'remark-gfm';
import './GeminiResponseDisplay.css';

const GeminiResponseDisplay = ({ response }) => {
    if (!response) {
        return null;
    }

    return (
        <div className="gemini-response-display">
            <ReactMarkdown
                children={response}
                remarkPlugins={[remarkGfm]} // GitHub Flavored Markdown desteği
                components={{
                    code({ node, inline, className, children, ...props }) {
                        const match = /language-(\w+)/.exec(className || '');
                        return !inline && match ? (
                            <SyntaxHighlighter
                                children={String(children).replace(/\n$/, '')}
                                style={tomorrow} // Kod bloğu tema
                                language={match[1]}
                                PreTag="div"
                                {...props}
                            />
                        ) : (
                            <code className={className} {...props}>
                                {children}
                            </code>
                        );
                    },
                }}
            />
        </div>
    );
};

export default GeminiResponseDisplay;