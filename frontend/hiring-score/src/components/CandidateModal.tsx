// frontend/hiring-score/src/components/CandidateModal.tsx
import {Candidate} from "@/types/types";
import {XMarkIcon} from "@heroicons/react/24/solid";

interface CandidateModalProps {
    candidate: Candidate;
    onClose: () => void;
    onSelect: (candidate: Candidate) => void;
    isSelected: boolean;
}

const CandidateModal: React.FC<CandidateModalProps> = ({candidate, onClose, onSelect, isSelected}) => {
    return (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
            <div className="bg-white rounded-lg w-full max-w-3xl max-h-[90vh] overflow-y-auto m-4">
                <div className="sticky top-0 bg-white p-6 border-b flex justify-between items-start">
                    <div>
                        <h2 className="text-2xl font-semibold">{candidate.name}</h2>
                        <p className="text-gray-600">{candidate.location}</p>
                    </div>
                    <div className="flex gap-4">
                        {isSelected ? (
                            <button
                                disabled
                                className="px-4 py-2 bg-gray-200 text-gray-600 rounded-md cursor-not-allowed"
                            >
                                Added
                            </button>
                        ) : (
                            <button
                                onClick={() => onSelect(candidate)}
                                className="px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600"
                            >
                                Add to Team
                            </button>
                        )}

                        <button onClick={onClose} className="text-gray-500 hover:text-gray-700">
                            <XMarkIcon className="w-6 h-6"/>
                        </button>
                    </div>
                </div>

                <div className="p-6 space-y-6">
                    <div className="flex items-center justify-between">
                        <h3 className="text-lg font-semibold">Score</h3>
                        <div className="flex items-center">
                            <span className="text-yellow-500">⚡</span>
                            <span className="ml-1 text-xl font-semibold">{candidate.score}</span>
                        </div>
                    </div>

                    <div>
                        <h3 className="text-lg font-semibold mb-3">Skills</h3>
                        <div className="flex flex-wrap gap-2">
                            {candidate.skills.map((skill, index) => (
                                <span
                                    key={index}
                                    className="px-3 py-1 bg-blue-100 text-blue-800 text-sm rounded-full"
                                >
                  {skill}
                </span>
                            ))}
                        </div>
                    </div>

                    <div>
                        <h3 className="text-lg font-semibold mb-3">Education</h3>
                        {candidate.education.degrees.map((degree, index) => (
                            <div key={index} className="mb-4">
                                <p className="font-medium">{degree.school}</p>
                                <p className="text-gray-600">{degree.degree} in {degree.subject}</p>
                                <p className="text-sm text-gray-500">{degree.endDate}</p>
                            </div>
                        ))}
                    </div>

                    <div>
                        <h3 className="text-lg font-semibold mb-3">Experience</h3>
                        {candidate.work_experiences.map((exp, index) => (
                            <div key={index} className="mb-4">
                                <p className="font-medium">{exp.company}</p>
                                <p className="text-sm text-gray-500">{exp.roleName}</p>
                            </div>
                        ))}
                    </div>
                </div>
            </div>
        </div>
    );
};

export default CandidateModal;