export const getBackendBaseUrl = (devBaseUrl: string): string => {
  return import.meta.env.PROD
    ? `http://${window.location.hostname}:8080`
    : devBaseUrl;
};
