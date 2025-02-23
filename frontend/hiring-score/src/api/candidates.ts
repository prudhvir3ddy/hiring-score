// src/api/candidates.ts
export const fetchCandidates = async (page: number, query: string = '') => {
  const baseUrl = 'http://localhost:8080/api/candidates';
  const url = new URL(baseUrl);
  url.searchParams.append('page', page.toString());
  url.searchParams.append('page_count', '10');
  if (query) {
    url.searchParams.append('query', query);
  }

  const response = await fetch(url.toString());
  return await response.json();
};