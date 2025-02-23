import {Candidate} from "@/types/types";
import {useState} from "react";
import CandidateModal from "@/components/CandidateModal";

interface CandidateCardProps {
    candidate: Candidate;
    onSelect: (candidate: Candidate) => void;
    isSelected: boolean;
}

const CandidateCard: React.FC<CandidateCardProps> = ({candidate, onSelect, isSelected}) => {
    const [showModal, setShowModal] = useState(false);

    return (
        <>
            <div className="border rounded-lg p-6 shadow-sm" onClick={() => setShowModal(true)}>
                <div className="flex justify-between items-start mb-4">
                    <div>
                        <h2 className="text-xl font-semibold">{candidate.name}</h2>
                        <p className="text-gray-600">{candidate.location}</p>
                    </div>
                    <div className="flex items-center">
                        <span className="text-yellow-500">⚡</span>
                        <span className="ml-1 font-semibold">{candidate.score}</span>
                    </div>
                </div>

                <div className="mt-4">
                    <div className="flex flex-wrap gap-2">
                        {candidate.skills.map((skill, index) => (
                            <span
                                key={index}
                                className="px-2 py-1 bg-blue-100 text-blue-800 text-sm rounded-full"
                            >
              {skill}
            </span>
                        ))}
                    </div>
                </div>

                <div className="mt-4 text-sm text-gray-600">
                    <p>{candidate.education.degrees[0]?.school}</p>
                    <p>{candidate.education.degrees[0]?.degree} in {candidate.education.degrees[0]?.subject}</p>
                </div>
            </div>
            {showModal && (
                <CandidateModal
                    candidate={candidate}
                    onClose={() => setShowModal(false)}
                    onSelect={onSelect}
                    isSelected={isSelected}
                />
            )}        </>
    );
}

export default CandidateCard

