// frontend/hiring-score/src/components/PaginationControls.tsx
import { ChevronLeftIcon, ChevronRightIcon } from '@heroicons/react/24/solid';

interface PaginationControlsProps {
  page: number;
  loading: boolean;
  hasNextPage: boolean;
  onPrevPage: () => void;
  onNextPage: () => void;
}

const PaginationControls: React.FC<PaginationControlsProps> = ({
  page,
  loading,
  hasNextPage,
  onPrevPage,
  onNextPage
}) => {
  return (
    <div className="flex items-center gap-4">
      <button
        onClick={onPrevPage}
        disabled={page === 1 || loading}
        className="p-2 rounded-full hover:bg-gray-100 disabled:opacity-50 disabled:hover:bg-transparent transition-colors"
        aria-label="Previous page"
      >
        <ChevronLeftIcon className="w-5 h-5" />
      </button>
      <span className="text-sm font-medium">Page {page}</span>
      <button
        onClick={onNextPage}
        disabled={!hasNextPage || loading}
        className="p-2 rounded-full hover:bg-gray-100 disabled:opacity-50 disabled:hover:bg-transparent transition-colors"
        aria-label="Next page"
      >
        <ChevronRightIcon className="w-5 h-5" />
      </button>
    </div>
  );
};

export default PaginationControls;