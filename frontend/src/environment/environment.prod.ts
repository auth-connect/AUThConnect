const env = (window as any).__env || {};

export const environment = {
  production: true,
  apiUrl: env.apiUrl || 'http://localhost:8000', // Fallback to default for development
};
