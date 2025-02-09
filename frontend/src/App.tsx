import React from 'react';
import ContainerList from './components/ContainerList/ContainerList';

const App: React.FC = () => {
    return (
        <div className="min-h-screen flex flex-col items-center justify-center">
            <ContainerList />
        </div>
    );
};

export default App;
