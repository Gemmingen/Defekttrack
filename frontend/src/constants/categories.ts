import type { SelectOption } from 'sit-onyx';
import { FehlerKategorie } from '../api/generated';

export const ALL_FILTER_VALUE = 'ALL' as const;
export type FilterCategory = FehlerKategorie | typeof ALL_FILTER_VALUE;

export const CATEGORY_LABELS: Record<FehlerKategorie, string> = {
  [FehlerKategorie.DISPLAY_SCHADEN]: 'Display Schaden',
  [FehlerKategorie.AKKU_DEFEKT]: 'Akku Defekt',
  [FehlerKategorie.TASTATUR_DEFEKT]: 'Tastatur Defekt',
  [FehlerKategorie.SOFTWARE_OS]: 'Software / OS',
  [FehlerKategorie.KEIN_FEHLER]: 'Kein Fehler',
};

export const CATEGORY_OPTIONS: SelectOption<FehlerKategorie>[] = Object.values(FehlerKategorie).map((value) => ({
  label: CATEGORY_LABELS[value],
  value,
}));

export const FILTER_OPTIONS: SelectOption<FilterCategory>[] = [
  { label: 'Alle Kategorien', value: ALL_FILTER_VALUE },
  ...CATEGORY_OPTIONS,
];