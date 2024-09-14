// Check if the global __env variable exists (it will only exist in production)
const env = (window as any).__env || {};

export const environment = {
  production: false,
  apiUrl: env.apiUrl || 'http://localhost:8000', // Fallback to default for development
};
