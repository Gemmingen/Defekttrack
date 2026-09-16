import { ref } from 'vue';
import { laptopApi } from '../api/laptopApi';
import type { Laptop, LaptopInput, LogEintragInput } from '../types/laptop';

export function useDefektTrack() {
  const laptops = ref<Laptop[]>([]);
  const selectedLaptop = ref<Laptop | null>(null);
  const activeFilter = ref<string>('Alle');
  const isLoading = ref(false);
  const error = ref<string | null>(null);

  const fetchLaptops = async (filter?: string) => {
    isLoading.value = true;
    error.value = null;
    try {
      laptops.value = await laptopApi.getAll(filter);
    } catch (err: any) {
      error.value = err.message;
    } finally {
      isLoading.value = false;
    }
  };

  const createLaptop = async (input: LaptopInput) => {
    try {
      await laptopApi.create(input);
      await fetchLaptops(activeFilter.value);
    } catch (err: any) {
      error.value = err.message;
    }
  };

  const removeLaptop = async (id: number) => {
    try {
      await laptopApi.delete(id);
      if (selectedLaptop.value?.id === id) selectedLaptop.value = null;
      await fetchLaptops(activeFilter.value);
    } catch (err: any) {
      error.value = err.message;
    }
  };

  const addLog = async (laptopId: number, input: LogEintragInput) => {
    try {
      await laptopApi.createLog(laptopId, input);
      const updated = await laptopApi.getById(laptopId);
      selectedLaptop.value = updated;
      await fetchLaptops(activeFilter.value);
    } catch (err: any) {
      error.value = err.message;
    }
  };

  
  const removeLog = async (laptopId: number, logId: number) => {
    try {
      await laptopApi.deleteLog(laptopId, logId);
      const updated = await laptopApi.getById(laptopId);
      selectedLaptop.value = updated;
      await fetchLaptops(activeFilter.value);
    } catch (err: any) {
      error.value = err.message;
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
    removeLog
  };
}