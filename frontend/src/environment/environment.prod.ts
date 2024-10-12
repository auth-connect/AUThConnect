declare const window: any;

export const environment = {
  production: false,
  apiUrl: window.__env && window.__env.apiUrl ? window.__env.apiUrl : 'API_URL_NOT_SET',
};
