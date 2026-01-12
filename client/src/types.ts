export interface Grant {
  id: number;
  title: string;
  agency: string;
  deadline: string; // ISO Date string
  quantum: string;
  matchScore: number; // For your "Relevance Filtering" logic
}