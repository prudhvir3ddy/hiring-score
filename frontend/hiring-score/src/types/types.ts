interface Degree {
    degree: string;
    subject: string;
    school: string;
    isTop50: boolean;
    isTop25: boolean;
    endDate: string
}

export interface WorkExperience {
    company: string;
    roleName: string;
}

export interface Candidate {
    id: string;
    name: string;
    email: string;
    location: string;
    skills: string[];
    score: number;
    education: {
        highest_level: string
        degrees: Degree[];
    };
    work_experiences: WorkExperience[];
}

export interface PaginatedResponse {
    candidates: Candidate[];
    hasNextPage: boolean;
}