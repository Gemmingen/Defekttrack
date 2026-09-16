import type { Laptop, LaptopInput, LogEintrag, LogEintragInput } from '../types/laptop';

const BASE_URL = '/laptops';

export const laptopApi = {
  async getAll(fehlerFilter?: string): Promise<Laptop[]> {
    const url = fehlerFilter && fehlerFilter !== 'Alle' 
      ? `${BASE_URL}?fehler=${encodeURIComponent(fehlerFilter)}` 
      : BASE_URL;
    const res = await fetch(url);
    if (!res.ok) throw new Error('Fehler beim Laden der Laptops');
    return res.json();
  },

  async getById(id: number): Promise<Laptop> {
    const res = await fetch(`${BASE_URL}/${id}`);
    if (!res.ok) throw new Error(`Laptop #${id} nicht gefunden`);
    return res.json();
  },

  async create(input: LaptopInput): Promise<Laptop> {
    const res = await fetch(BASE_URL, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(input),
    });
    if (!res.ok) throw new Error('Erfassung fehlgeschlagen');
    return res.json();
  },

  async update(id: number, input: LaptopInput): Promise<Laptop> {
    const res = await fetch(`${BASE_URL}/${id}`, {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(input),
    });
    if (!res.ok) throw new Error('Update fehlgeschlagen');
    return res.json();
  },

  async delete(id: number): Promise<void> {
    const res = await fetch(`${BASE_URL}/${id}`, { method: 'DELETE' });
    if (!res.ok) throw new Error('Löschen fehlgeschlagen');
  },

  async createLog(laptopId: number, input: LogEintragInput): Promise<LogEintrag> {
    const res = await fetch(`${BASE_URL}/${laptopId}/logs`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(input),
    });
    if (!res.ok) throw new Error('Log-Eintrag fehlgeschlagen');
    return res.json();
  },

  async deleteLog(laptopId: number, logId: number): Promise<void> {
    const res = await fetch(`${BASE_URL}/${laptopId}/logs/${logId}`, { method: 'DELETE' });
    if (!res.ok) throw new Error('Log-Löschung fehlgeschlagen');
  }
};