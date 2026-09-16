/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { FehlerKategorie } from './FehlerKategorie';
import type { LogEintrag } from './LogEintrag';
export type Laptop = {
    id: number;
    name: string;
    marke: string;
    os: string;
    fehler: FehlerKategorie;
    it_logs?: Array<LogEintrag>;
};

