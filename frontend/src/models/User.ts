export interface User {
  email: string;
  creditUser?: string;
  creditPassword?: string;
  cutOffTime?: number;
  minimalCredit?: number;
  tariffKey?: string;
  debt: number;
}
