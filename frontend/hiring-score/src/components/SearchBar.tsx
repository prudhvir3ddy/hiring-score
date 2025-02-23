import {useEffect, useState} from "react";

interface SearchBarProps {
    value: string;
    onChange: (value: string) => void;
    debounceTime?: number;
}

const SearchBar: React.FC<SearchBarProps> = ({
                                                 value,
                                                 onChange,
                                                 debounceTime = 500
                                             }) => {
    const [localValue, setLocalValue] = useState(value);

    useEffect(() => {
        const timer = setTimeout(() => {
            onChange(localValue);
        }, debounceTime);

        return () => clearTimeout(timer);
    }, [localValue, onChange, debounceTime]);

    return (
        <div className="mb-6">
            <input
                type="text"
                placeholder="Search for candidates, skills, work exp, college..."
                value={localValue}
                onChange={(e) => setLocalValue(e.target.value)}
                className="w-full p-3 border rounded-md"
            />
        </div>
    );
};

export default SearchBar;
