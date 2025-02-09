import React from "react";

interface ContainerProps {
    id: string;
    name: string;
    status: string;
    ip?: string;
}

const getStatusIcon = (status: string) => {
    return status === "running" ? "✅" : "❌";
};

const ContainerCard: React.FC<{ container: ContainerProps }> = ({ container }) => {
    return (
        <div className="container-card">
            <h3>{container.name}</h3>
            <p className="container-id">ID: {container.id}</p>
            <p className={`status ${container.status}`}>
                {getStatusIcon(container.status)} {container.status}
            </p>
            {container.ip && <p>IP: {container.ip}</p>}
        </div>
    );
};

export default ContainerCard;
