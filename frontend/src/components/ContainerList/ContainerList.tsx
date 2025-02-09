import React, { useEffect, useState } from 'react';
import api from "../../services/api";
import ContainerCard from "../../components/ContainerCard/ContainerCard";
import { Container } from "../../types/types";
import { FaServer, FaHistory } from 'react-icons/fa';

export const ContainerList: React.FC = () => {
    const [containers, setContainers] = useState<Container[]>([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        const fetchContainers = async () => {
            try {
                const data = (await api.getContainers()) as unknown as Container[];

                // Убираем дубликаты по ID
                const uniqueContainers = Array.from(new Map(data.map(c => [c.id, c])).values());

                setContainers(uniqueContainers);
                setError(null);
            } catch (err) {
                setError('Не удалось получить список контейнеров');
                console.error('Ошибка при загрузке контейнеров:', err);
            } finally {
                setLoading(false);
            }
        };

        fetchContainers();
    }, []);

    if (loading) return <p className="text-center text-white">Загрузка контейнеров...</p>;
    if (error) return <p className="text-center text-red-500">{error}</p>;

    // Разделение активных и завершенных контейнеров
    const activeContainers = containers.filter(c => c.status === 'running');
    const exitedContainers = containers.filter(c => c.status !== 'running');

    return (
        <div className="container mx-auto p-4 text-white">
            <h1 className="text-4xl font-bold text-center mb-6">Container Monitoring</h1>

            {/* Активные контейнеры */}
            <div className="mb-8">
                <h2 className="text-2xl font-semibold mb-4 flex items-center justify-center">
                    <FaServer className="mr-2 text-blue-400" /> Активные контейнеры ({activeContainers.length})
                </h2>
                <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
                    {activeContainers.length > 0 ? (
                        activeContainers.map(container => (
                            <ContainerCard key={container.id} container={container} />
                        ))
                    ) : (
                        <p className="text-gray-400">Нет активных контейнеров</p>
                    )}
                </div>
            </div>

            {/* История контейнеров */}
            {exitedContainers.length > 0 && (
                <div>
                    <h2 className="text-2xl font-semibold mb-4 flex items-center justify-center">
                        <FaHistory className="mr-2 text-yellow-400" /> История контейнеров ({exitedContainers.length})
                    </h2>
                    <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
                        {exitedContainers.map(container => (
                            <ContainerCard key={container.id} container={container} />
                        ))}
                    </div>
                </div>
            )}
        </div>
    );
};

export default ContainerList;
