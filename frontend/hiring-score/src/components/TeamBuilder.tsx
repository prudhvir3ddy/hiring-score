// frontend/hiring-score/src/components/TeamBuilder.tsx
import {Candidate} from "@/types/types";
import {XMarkIcon} from "@heroicons/react/24/solid";

interface TeamBuilderProps {
    selectedCandidates: Candidate[];
    onRemoveCandidate: (id: string) => void;
    onForward?: () => void;
}

const TeamBuilder: React.FC<TeamBuilderProps> = ({selectedCandidates, onRemoveCandidate, onForward}) => {
    const canForward = selectedCandidates.length >= 5;

    return (
        <div className="bg-white rounded-lg shadow flex flex-col sticky top-8 h-[calc(100vh-4rem)]">
            <div className="p-6 border-b">
                <div className="flex justify-between items-center">
                    <h2 className="text-xl font-bold">Team Builder</h2>
                    <span className="text-sm text-gray-600">
                        {selectedCandidates.length} selected
                    </span>
                </div>
            </div>

            <div className="flex-1 overflow-y-auto p-6">
                <div className="space-y-4">
                    {selectedCandidates.map((candidate) => (
                        <div key={candidate.id} className="bg-gray-50 rounded-lg p-4">
                            <div className="flex justify-between items-start">
                                <div className="flex-1">
                                    <div className="flex justify-between items-start mb-2">
                                        <h3 className="font-semibold">{candidate.name}</h3>
                                        <div className="flex items-center">
                                            <span className="text-yellow-500">⚡</span>
                                            <span className="ml-1 font-medium">{candidate.score}</span>
                                        </div>
                                    </div>
                                    <p className="text-sm text-gray-600 mb-2">{candidate.location}</p>
                                    <div className="flex flex-wrap gap-1 mb-2">
                                        {candidate.skills.slice(0, 3).map((skill, index) => (
                                            <span
                                                key={index}
                                                className="px-2 py-0.5 bg-blue-100 text-blue-800 text-xs rounded-full"
                                            >
                                            {skill}
                                        </span>
                                        ))}
                                        {candidate.skills.length > 3 && (
                                            <span className="text-xs text-gray-500">
                                            +{candidate.skills.length - 3} more
                                        </span>
                                        )}
                                    </div>
                                    <p className="text-xs text-gray-600">
                                        {candidate.education.degrees[0]?.degree} in {candidate.education.degrees[0]?.subject}
                                    </p>
                                </div>
                                <button
                                    onClick={() => onRemoveCandidate(candidate.id)}
                                    className="ml-2 text-gray-400 hover:text-gray-600"
                                    title="Remove from team"
                                >
                                    <XMarkIcon className="w-5 h-5"/>
                                </button>
                            </div>
                        </div>
                    ))}
                    {selectedCandidates.length === 0 && (
                        <p className="text-gray-500 text-center">No candidates selected</p>
                    )}
                </div>
            </div>
            <div className="p-6 border-t">
                <button
                    onClick={onForward}
                    disabled={!canForward}
                    className={`w-full py-2 px-4 rounded-md text-white text-center transition-colors
                        ${canForward
                        ? 'bg-green-600 hover:bg-green-700'
                        : 'bg-gray-300 cursor-not-allowed'}`}
                >
                    Forward to interview loop
                </button>
                {!canForward && (
                    <p className="text-sm text-gray-500 text-center mt-2">
                        Select {5 - selectedCandidates.length} more candidates
                    </p>
                )}
            </div>
        </div>
    );
};

export default TeamBuilder;