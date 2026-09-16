/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { FehlerKategorie } from '../models/FehlerKategorie';
import type { Laptop } from '../models/Laptop';
import type { LaptopInput } from '../models/LaptopInput';
import type { LogEintrag } from '../models/LogEintrag';
import type { LogEintragInput } from '../models/LogEintragInput';
import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';
export class DefaultService {
    /**
     * Liste aller Laptops abrufen
     * @param fehler
     * @returns Laptop Erfolgreich
     * @throws ApiError
     */
    public static getLaptops(
        fehler?: FehlerKategorie,
    ): CancelablePromise<Array<Laptop>> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/laptops',
            query: {
                'fehler': fehler,
            },
            errors: {
                400: `Ungültige Anfrageparameter`,
            },
        });
    }
    /**
     * Neuen Laptop zur Reparatur erfassen
     * @param requestBody
     * @returns Laptop Laptop angelegt
     * @throws ApiError
     */
    public static postLaptops(
        requestBody: LaptopInput,
    ): CancelablePromise<Laptop> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/laptops',
            body: requestBody,
            mediaType: 'application/json',
            errors: {
                400: `Ungültige Eingabedaten (z.B. fehlende Pflichtfelder)`,
            },
        });
    }
    /**
     * Einzelnen Laptop samt Logs abrufen
     * @param id
     * @returns Laptop Erfolgreich
     * @throws ApiError
     */
    public static getLaptops1(
        id: number,
    ): CancelablePromise<Laptop> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/laptops/{id}',
            path: {
                'id': id,
            },
            errors: {
                400: `Ungültiges ID-Format`,
                404: `Laptop nicht gefunden`,
            },
        });
    }
    /**
     * Laptop-Daten oder Fehlerstatus aktualisieren
     * @param id
     * @param requestBody
     * @returns Laptop Laptop erfolgreich aktualisiert
     * @throws ApiError
     */
    public static putLaptops(
        id: number,
        requestBody: LaptopInput,
    ): CancelablePromise<Laptop> {
        return __request(OpenAPI, {
            method: 'PUT',
            url: '/laptops/{id}',
            path: {
                'id': id,
            },
            body: requestBody,
            mediaType: 'application/json',
            errors: {
                400: `Ungültige Eingabedaten oder ID-Format`,
                404: `Laptop nicht gefunden`,
            },
        });
    }
    /**
     * Reparierten oder ausgemusterten Laptop löschen
     * @param id
     * @returns void
     * @throws ApiError
     */
    public static deleteLaptops(
        id: number,
    ): CancelablePromise<void> {
        return __request(OpenAPI, {
            method: 'DELETE',
            url: '/laptops/{id}',
            path: {
                'id': id,
            },
            errors: {
                400: `Ungültiges ID-Format`,
                404: `Laptop nicht gefunden`,
            },
        });
    }
    /**
     * IT-Log-Eintrag hinzufügen
     * @param id
     * @param requestBody
     * @returns LogEintrag Log-Eintrag erfolgreich hinzugefügt
     * @throws ApiError
     */
    public static postLaptopsLogs(
        id: number,
        requestBody: LogEintragInput,
    ): CancelablePromise<LogEintrag> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/laptops/{id}/logs',
            path: {
                'id': id,
            },
            body: requestBody,
            mediaType: 'application/json',
            errors: {
                400: `Ungültige Eingabedaten oder ID-Format`,
                404: `Laptop nicht gefunden`,
            },
        });
    }
    /**
     * Existierenden IT-Log-Eintrag aktualisieren
     * @param id
     * @param logId
     * @param requestBody
     * @returns LogEintrag Log-Eintrag erfolgreich aktualisiert
     * @throws ApiError
     */
    public static putLaptopsLogs(
        id: number,
        logId: number,
        requestBody: LogEintragInput,
    ): CancelablePromise<LogEintrag> {
        return __request(OpenAPI, {
            method: 'PUT',
            url: '/laptops/{id}/logs/{logId}',
            path: {
                'id': id,
                'logId': logId,
            },
            body: requestBody,
            mediaType: 'application/json',
            errors: {
                400: `Ungültige Eingabedaten oder ID-Format`,
                404: `Laptop oder Log-Eintrag nicht gefunden`,
            },
        });
    }
    /**
     * IT-Log-Eintrag löschen
     * @param id
     * @param logId
     * @returns void
     * @throws ApiError
     */
    public static deleteLaptopsLogs(
        id: number,
        logId: number,
    ): CancelablePromise<void> {
        return __request(OpenAPI, {
            method: 'DELETE',
            url: '/laptops/{id}/logs/{logId}',
            path: {
                'id': id,
                'logId': logId,
            },
            errors: {
                400: `Ungültiges ID-Format`,
                404: `Laptop oder Log-Eintrag nicht gefunden`,
            },
        });
    }
}
