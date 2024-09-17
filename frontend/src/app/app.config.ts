import { ApplicationConfig, inject, provideZoneChangeDetection } from '@angular/core';
import { provideRouter } from '@angular/router';

import { routes } from './app.routes';
import { HttpClient, provideHttpClient, withInterceptors } from '@angular/common/http';
import { provideNgIconLoader } from '@ng-icons/core';
import { authInterceptor } from './services/auth-service/auth.interceptor';

export const appConfig: ApplicationConfig = {
  providers: [
    provideZoneChangeDetection({ eventCoalescing: true }), 
    provideRouter(routes), 
    provideHttpClient(withInterceptors([authInterceptor])),
    // provideNgIconLoader(name => {
    //   const http = inject(HttpClient);
    //   return http.get(`/assets/icons/${name}.svg`, { responseType: 'text' });
    // }),
  ]
};
