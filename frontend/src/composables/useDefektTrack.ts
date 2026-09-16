import { ref } from 'vue';
import { DefaultService } from '../api/generated';
import type { Laptop, LaptopInput, LogEintragInput, FehlerKategorie } from '../api/generated';

export function useDefektTrack() {
  const laptops = ref<Laptop[]>([]);
  const selectedLaptop = ref<Laptop | null>(null);
  const activeFilter = ref<string>('Alle');
  const isLoading = ref<boolean>(false);
  const error = ref<string | null>(null);

  const fetchLaptops = async (filter?: string) => {
    isLoading.value = true;
    error.value = null;
    try {
      const category = filter && filter !== 'Alle' ? (filter as FehlerKategorie) : undefined;
      laptops.value = await DefaultService.getLaptops(category);
      if (selectedLaptop.value) {
        selectedLaptop.value = laptops.value.find(l => l.id === selectedLaptop.value?.id) || null;
      }
    } catch (e: any) {
      error.value = e.message || 'Fehler beim Laden der Laptops';
    } finally {
      isLoading.value = false;
    }
  };

  const createLaptop = async (input: LaptopInput) => {
    try {
      await DefaultService.postLaptops(input);
      await fetchLaptops(activeFilter.value);
    } catch (e: any) {
      error.value = e.message || 'Erfassung fehlgeschlagen';
    }
  };

  const removeLaptop = async (id: number) => {
    try {
      await DefaultService.deleteLaptops(id);
      if (selectedLaptop.value?.id === id) selectedLaptop.value = null;
      await fetchLaptops(activeFilter.value);
    } catch (e: any) {
      error.value = e.message || 'Löschen fehlgeschlagen';
    }
  };

  const addLog = async (laptopId: number, input: LogEintragInput) => {
    try {
      await DefaultService.postLaptopsLogs(laptopId, input);
      await fetchLaptops(activeFilter.value);
    } catch (e: any) {
      error.value = e.message || 'Log konnte nicht gespeichert werden';
    }
  };

  const removeLog = async (laptopId: number, logId: number) => {
    try {
      await DefaultService.deleteLaptopsLogs(laptopId, logId);
      await fetchLaptops(activeFilter.value);
    } catch (e: any) {
      error.value = e.message || 'Log konnte nicht gelöscht werden';
    }
  };

  return {
    laptops,
    selectedLaptop,
    activeFilter,
    isLoading,
    error,
    fetchLaptops,
    createLaptop,
    removeLaptop,
    addLog,
    removeLog,
  };
}