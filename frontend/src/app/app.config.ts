import { ApplicationConfig, inject, provideZoneChangeDetection } from '@angular/core';
import { provideRouter } from '@angular/router';

import { routes } from './app.routes';
import { HttpClient, provideHttpClient } from '@angular/common/http';
import { provideNgIconLoader } from '@ng-icons/core';

export const appConfig: ApplicationConfig = {
  providers: [
    provideZoneChangeDetection({ eventCoalescing: true }), 
    provideRouter(routes), 
    provideHttpClient(), 
    // provideNgIconLoader(name => {
    //   const http = inject(HttpClient);
    //   return http.get(`/assets/icons/${name}.svg`, { responseType: 'text' });
    // }),
  ]
};
