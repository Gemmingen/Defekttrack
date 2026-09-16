export type FehlerKategorie = 'Hardware' | 'Display' | 'Akku' | 'Software' | 'Netzwerk' | 'Sonstiges';

export interface LogEintrag {
  id: number;
  bearbeiter: string;
  notiz: string;
  timestamp: string;
}

export interface LogEintragInput {
  bearbeiter: string;
  notiz: string;
}

export interface Laptop {
  id: number;
  marke: string;
  name: string;
  os: string;
  fehler: FehlerKategorie;
  it_logs?: LogEintrag[];
}

export type LaptopInput = Omit<Laptop, 'id' | 'it_logs'>;