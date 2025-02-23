// frontend/hiring-score/src/components/CanditateList.tsx
import {Candidate} from "@/types/types";
import CandidateCard from "./CandidateCard";

interface CandidateListProps {
    candidates: Candidate[];
    onSelect: (candidate: Candidate) => void;
    selectedCandidates: Candidate[];
}

const CandidateList: React.FC<CandidateListProps> = ({candidates, onSelect, selectedCandidates}) => {
    return (
        <div className="grid grid-cols-3 gap-6">
            {candidates.map((candidate) => (
                <CandidateCard
                    key={candidate.id}
                    candidate={candidate}
                    onSelect={onSelect}
                    isSelected={selectedCandidates?.some(c => c.id === candidate.id)}
                />
            ))}
        </div>
    );
};

export default CandidateList;