export interface User {
  email: string;
  creditUser?: string;
  creditPassword?: string;
  cutOffTime?: string;
  minimalCredit?: number;
  debt: number;
}
