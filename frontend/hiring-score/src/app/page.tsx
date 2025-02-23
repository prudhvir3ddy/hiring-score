'use client';

import {useCallback, useEffect, useState} from "react";
import {Candidate, PaginatedResponse} from "@/types/types";
import SearchBar from "@/components/SearchBar";
import PaginationControls from "@/components/PaginationControls";
import {fetchCandidates} from "@/api/candidates";
import CandidateList from "@/components/CanditateList";
import TeamBuilder from "@/components/TeamBuilder";

export default function Home() {
    const [selectedCandidates, setSelectedCandidates] = useState<Candidate[]>([]);
    const [candidates, setCandidates] = useState<Candidate[]>([]);
    const [loading, setLoading] = useState(true);
    const [page, setPage] = useState(1);
    const [searchQuery, setSearchQuery] = useState("");
    const [hasNextPage, setHasNextPage] = useState(false);

    const handleCandidateSelect = useCallback((candidate: Candidate) => {
        setSelectedCandidates(prev => {
            if (prev.some(c => c.id === candidate.id)) return prev;
            return [...prev, candidate];
        });
    }, []);

    const handleForward = useCallback(() => {
        // Add your forward logic here
        console.log('Forwarding candidates:', selectedCandidates);
    }, [selectedCandidates]);

    const handleCandidateRemove = useCallback((id: string) => {
        setSelectedCandidates(prev => prev.filter(c => c.id !== id));
    }, []);

    const getCandidates = async (currentPage: number, query: string) => {
        try {
            setLoading(true);
            const data: PaginatedResponse = await fetchCandidates(currentPage, query);
            setCandidates(data.candidates);
            setHasNextPage(data.hasNextPage);
        } catch (error) {
            console.error('Error fetching candidates:', error);
        } finally {
            setLoading(false);
        }
    };

    // Handle search changes
    useEffect(() => {
        setPage(1); // Reset page only when search changes
        getCandidates(1, searchQuery);
    }, [searchQuery]);

    // Handle page changes
    useEffect(() => {
        getCandidates(page, searchQuery);
    }, [page]);

    const handleSearch = (value: string) => {
        setSearchQuery(value);
    };

    if (loading) {
        return <div className="flex justify-center items-center min-h-screen">Loading...</div>;
    }

    return (
        <div className="min-h-screen p-8">
            <div className="grid grid-cols-1 lg:grid-cols-4 gap-8">
                <div className="lg:col-span-3">
                    <div className="flex flex-col gap-6 mb-6">
                        <div className="flex justify-between items-center">
                            <h1 className="text-2xl font-bold">Candidates</h1>
                            <PaginationControls
                                page={page}
                                loading={loading}
                                hasNextPage={hasNextPage}
                                onPrevPage={() => setPage(p => p - 1)}
                                onNextPage={() => setPage(p => p + 1)}
                            />
                        </div>
                        <SearchBar value={searchQuery} onChange={handleSearch}/>
                    </div>
                    {loading ? (
                        <div className="flex justify-center items-center min-h-[200px]">Loading...</div>
                    ) : (
                        <CandidateList
                            candidates={candidates}
                            onSelect={handleCandidateSelect}
                            selectedCandidates={selectedCandidates}
                        />
                    )}
                </div>
                <div>
                    <TeamBuilder
                        selectedCandidates={selectedCandidates}
                        onRemoveCandidate={handleCandidateRemove}
                        onForward={handleForward}
                    />
                </div>
            </div>
        </div>
    );
}